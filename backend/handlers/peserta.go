package handlers

import (
	"fmt"
	"jogokariyan-backend/config"
	"jogokariyan-backend/models"
	"jogokariyan-backend/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CheckExisting checks if NoHP OR Email is already in use
func CheckExisting(c *fiber.Ctx) error {
	query := strings.TrimSpace(c.Query("query")) // Can be no_hp or email
	if query == "" {
		return utils.JSONError(c, 400, "Query parameter required")
	}

	var peserta models.Peserta
	result := config.DB.Where("no_hp = ? OR email = ?", query, query).First(&peserta)

	if result.Error == nil {
		return utils.JSONResponse(c, 409, "User already exists", fiber.Map{
			"exists": true,
			"msg":    "Nomor HP atau Email sudah terdaftar",
		})
	} else if result.Error == gorm.ErrRecordNotFound {
		return utils.JSONResponse(c, 200, "Available", fiber.Map{
			"exists": false,
		})
	}

	return utils.JSONError(c, 500, "Database error")
}

// Register creates a new participant
func Register(c *fiber.Ctx) error {
	var peserta models.Peserta

	// Parse form fields manually because we have file upload
	peserta.NamaLengkap = c.FormValue("nama_lengkap")
	peserta.Panggilan = c.FormValue("panggilan")
	peserta.JenisKelamin = c.FormValue("jenis_kelamin")
	peserta.AlamatDomisili = c.FormValue("alamat_domisili")
	peserta.NoHP = c.FormValue("no_hp")
	peserta.Email = c.FormValue("email")

	// Convert int
	if tahunLahirStr := c.FormValue("tahun_lahir"); tahunLahirStr != "" {
		fmt.Sscanf(tahunLahirStr, "%d", &peserta.TahunLahir)
	}

	// Handle File Upload
	file, err := c.FormFile("bukti_transfer")
	if err == nil {
		f, err := file.Open()
		if err == nil {
			defer f.Close()
			buffer := make([]byte, file.Size)
			f.Read(buffer)

			peserta.BuktiTransferData = buffer
			peserta.BuktiTransferContentType = file.Header.Get("Content-Type")
		}
	}

	// Create in DB
	peserta.ID = uuid.New() // Generate UUID manually if needed or let DB do it. Letting GORM do it is safer if configured.
	// But let's set it to be sure for the response

	if result := config.DB.Create(&peserta); result.Error != nil {
		return utils.JSONError(c, 500, result.Error.Error())
	}

	return utils.JSONResponse(c, 201, "Registration successful", peserta)
}

func FindPeserta(c *fiber.Ctx) error {
	query := strings.TrimSpace(c.Query("query"))
	var peserta models.Peserta

	err := config.DB.Where("no_hp = ? OR id::text = ?", query, query).First(&peserta).Error
	if err != nil {
		return utils.JSONError(c, 404, "Peserta not found")
	}

	return utils.JSONResponse(c, 200, "Found", peserta)
}

// GetPesertaList fetches participants for Admin with search
func GetPesertaList(c *fiber.Ctx) error {
	search := c.Query("search")
	var peserta []models.Peserta

	db := config.DB.Model(&models.Peserta{}).Order("created_at desc")

	if search != "" {
		likeQuery := "%" + search + "%"
		db = db.Where("nama_lengkap ILIKE ? OR panggilan ILIKE ? OR no_hp ILIKE ?", likeQuery, likeQuery, likeQuery)
	}

	if err := db.Find(&peserta).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to fetch participants")
	}

	return utils.JSONResponse(c, 200, "Success", peserta)
}

// UpdatePeserta updates participant details
func UpdatePeserta(c *fiber.Ctx) error {
	id := c.Params("id")
	var peserta models.Peserta

	if err := config.DB.First(&peserta, "id = ?", id).Error; err != nil {
		return utils.JSONError(c, 404, "Peserta not found")
	}

	type UpdateInput struct {
		NamaLengkap    string `json:"nama_lengkap"`
		Panggilan      string `json:"panggilan"`
		JenisKelamin   string `json:"jenis_kelamin"`
		AlamatDomisili string `json:"alamat_domisili"`
		NoHP           string `json:"no_hp"`
		Email          string `json:"email"`
		TahunLahir     int    `json:"tahun_lahir"`
	}

	var input UpdateInput
	if err := c.BodyParser(&input); err != nil {
		return utils.JSONError(c, 400, "Invalid input")
	}

	peserta.NamaLengkap = input.NamaLengkap
	peserta.Panggilan = input.Panggilan
	peserta.JenisKelamin = input.JenisKelamin
	peserta.AlamatDomisili = input.AlamatDomisili
	peserta.NoHP = input.NoHP
	peserta.Email = input.Email
	peserta.TahunLahir = input.TahunLahir

	if err := config.DB.Save(&peserta).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to update peserta")
	}

	return utils.JSONResponse(c, 200, "Peserta updated", peserta)
}

// DeletePeserta removes a participant
func DeletePeserta(c *fiber.Ctx) error {
	id := c.Params("id")
	var peserta models.Peserta

	if err := config.DB.First(&peserta, "id = ?", id).Error; err != nil {
		return utils.JSONError(c, 404, "Peserta not found")
	}

	if err := config.DB.Delete(&peserta).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to delete peserta")
	}

	return utils.JSONResponse(c, 200, "Peserta deleted", nil)
}

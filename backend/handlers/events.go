package handlers

import (
	"fmt"
	"jogokariyan-backend/config"
	"jogokariyan-backend/models"
	"jogokariyan-backend/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetActiveEvent(c *fiber.Ctx) error {
	var setting models.AppSettings
	// Find active event ID setting
	err := config.DB.Where("key = ?", "active_event_id").First(&setting).Error

	if err != nil {
		// Default message if no active event
		return utils.JSONResponse(c, 200, "No active event", fiber.Map{
			"message": "Silahkan datang ke Masjid Jogokariyan",
			"active":  false,
		})
	}

	var kegiatan models.Kegiatan
	if err := config.DB.Where("id = ?", setting.Value).First(&kegiatan).Error; err != nil {
		return utils.JSONError(c, 404, "Active event not found in database")
	}

	return utils.JSONResponse(c, 200, "Active event found", fiber.Map{
		"active": true,
		"event":  kegiatan,
	})
}

func CheckIn(c *fiber.Ctx) error {
	type CheckInRequest struct {
		PesertaID string `json:"peserta_id"`
		// Manually sent or inferred from active event
	}

	var req CheckInRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.JSONError(c, 400, "Invalid request")
	}

	// Get Active Event
	var setting models.AppSettings
	if err := config.DB.Where("key = ?", "active_event_id").First(&setting).Error; err != nil {
		return utils.JSONError(c, 400, "No active event configured")
	}

	eventID := setting.Value

	// Validate Peserta existence
	var peserta models.Peserta
	var errPeserta error

	// Check if input is UUID
	if _, err := uuid.Parse(req.PesertaID); err == nil {
		// Is UUID
		errPeserta = config.DB.Where("id = ?", req.PesertaID).First(&peserta).Error
	} else {
		// Is likely Phone Number
		errPeserta = config.DB.Where("no_hp = ?", req.PesertaID).First(&peserta).Error
	}

	if errPeserta != nil {
		return utils.JSONError(c, 404, "Data peserta tidak ditemukan (ID/No HP salah)")
	}

	// Check if already checked in
	var existingPresensi models.Presensi
	// USE peserta.ID (UUID) not req.PesertaID (which could be Phone Number)
	if err := config.DB.Where("peserta_id = ? AND kegiatan_id = ?", peserta.ID, eventID).First(&existingPresensi).Error; err == nil {
		return utils.JSONResponse(c, 200, "Berhasil Check-in! (Sudah pernah)", fiber.Map{
			"peserta": peserta,
			"new":     false,
		})
	}

	// Create Presensi
	newPresensi := models.Presensi{
		ID:           uuid.New(),
		PesertaID:    peserta.ID,
		KegiatanID:   uuid.MustParse(eventID),
		WaktuCheckIn: time.Now(),
	}

	if err := config.DB.Create(&newPresensi).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to record attendance")
	}

	return utils.JSONResponse(c, 200, fmt.Sprintf("Berhasil Check-in! Selamat datang %s", peserta.Panggilan), fiber.Map{
		"peserta": peserta,
		"new":     true,
	})
}

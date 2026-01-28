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

// GetKegiatanList fetches all activities
func GetKegiatanList(c *fiber.Ctx) error {
	var kegiatan []models.Kegiatan
	if err := config.DB.Order("waktu desc").Find(&kegiatan).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to fetch activities")
	}
	return utils.JSONResponse(c, 200, "Success", kegiatan)
}

// CreateKegiatan adds a new activity
func CreateKegiatan(c *fiber.Ctx) error {
	var kegiatan models.Kegiatan
	if err := c.BodyParser(&kegiatan); err != nil {
		fmt.Println("BodyParser Error:", err)
		return utils.JSONError(c, 400, "Invalid input: "+err.Error())
	}

	kegiatan.ID = uuid.New()
	// Ensure Waktu is parsed correctly if sent as string, but BodyParser handles ISO strings well usually.

	if err := config.DB.Create(&kegiatan).Error; err != nil {
		fmt.Println("DB Create Error:", err)
		return utils.JSONError(c, 500, "Failed to create activity: "+err.Error())
	}

	return utils.JSONResponse(c, 201, "Activity created", kegiatan)
}

// UpdateKegiatan edits an existing activity
func UpdateKegiatan(c *fiber.Ctx) error {
	id := c.Params("id")
	var kegiatan models.Kegiatan

	if err := config.DB.First(&kegiatan, "id = ?", id).Error; err != nil {
		return utils.JSONError(c, 404, "Activity not found")
	}

	type UpdateInput struct {
		NamaKegiatan string    `json:"nama_kegiatan"`
		Waktu        time.Time `json:"waktu"`
		Pemateri     string    `json:"pemateri"`
		Tempat       string    `json:"tempat"`
	}

	var input UpdateInput
	if err := c.BodyParser(&input); err != nil {
		fmt.Println("Update BodyParser Error:", err)
		return utils.JSONError(c, 400, "Invalid input: "+err.Error())
	}

	fmt.Printf("Updating Kegiatan ID: %s. New Time: %v\n", id, input.Waktu)

	kegiatan.NamaKegiatan = input.NamaKegiatan
	kegiatan.Waktu = input.Waktu
	kegiatan.Pemateri = input.Pemateri
	kegiatan.Tempat = input.Tempat

	if err := config.DB.Save(&kegiatan).Error; err != nil {
		fmt.Println("Update DB Error:", err)
		return utils.JSONError(c, 500, "Failed to update: "+err.Error())
	}

	return utils.JSONResponse(c, 200, "Activity updated", kegiatan)
}

// DeleteKegiatan removes an activity
func DeleteKegiatan(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Kegiatan{}, "id = ?", id).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to delete activity")
	}
	return utils.JSONResponse(c, 200, "Activity deleted", nil)
}

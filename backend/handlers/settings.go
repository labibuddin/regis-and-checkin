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

// GetSettings fetches all app settings
func GetSettings(c *fiber.Ctx) error {
	var settings []models.AppSettings
	if err := config.DB.Find(&settings).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to fetch settings")
	}
	return utils.JSONResponse(c, 200, "Success", settings)
}

// SetActiveEvent updates the active_event_id
func SetActiveEvent(c *fiber.Ctx) error {
	type Request struct {
		EventID string `json:"event_id"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return utils.JSONError(c, 400, "Invalid input")
	}

	setting := models.AppSettings{
		Key:   "active_event_id",
		Value: req.EventID,
	}

	if err := config.DB.Save(&setting).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to save setting")
	}

	return utils.JSONResponse(c, 200, "Active event updated", setting)
}

// UploadQRIS handles QRIS image upload to DB
func UploadQRIS(c *fiber.Ctx) error {
	file, err := c.FormFile("qris_image")
	if err != nil {
		return utils.JSONError(c, 400, "Image file is required")
	}

	// Read file content
	f, err := file.Open()
	if err != nil {
		return utils.JSONError(c, 500, "Failed to open file")
	}
	defer f.Close()

	// Convert to bytes
	buffer := make([]byte, file.Size)
	f.Read(buffer)

	// Save to DB
	image := models.AppImage{
		Key:         "qris_image",
		Data:        buffer,
		ContentType: file.Header.Get("Content-Type"),
	}

	if err := config.DB.Save(&image).Error; err != nil {
		fmt.Println("DB Save Error:", err)
		return utils.JSONError(c, 500, "Failed to save image to database: "+err.Error())
	}

	// Update Settings to point to the new endpoint pattern (just a flag or timestamp to force refresh)
	setting := models.AppSettings{
		Key:   "qris_image_path",
		Value: fmt.Sprintf("db_image?t=%d", time.Now().Unix()),
	}
	config.DB.Save(&setting)

	return utils.JSONResponse(c, 200, "QRIS uploaded to Database", setting)
}

// GetQRISImage serves the QRIS image from DB
func GetQRISImage(c *fiber.Ctx) error {
	var image models.AppImage
	if err := config.DB.Where("key = ?", "qris_image").First(&image).Error; err != nil {
		return c.Status(404).SendString("Image not found")
	}

	c.Set("Content-Type", image.ContentType)
	return c.Send(image.Data)
}

// UploadBackground handles Registration Background upload to DB
func UploadBackground(c *fiber.Ctx) error {
	file, err := c.FormFile("bg_image")
	if err != nil {
		return utils.JSONError(c, 400, "Image file is required")
	}

	f, err := file.Open()
	if err != nil {
		return utils.JSONError(c, 500, "Failed to open file")
	}
	defer f.Close()

	buffer := make([]byte, file.Size)
	f.Read(buffer)

	image := models.AppImage{
		Key:         "regis_bg_image",
		Data:        buffer,
		ContentType: file.Header.Get("Content-Type"),
	}

	if err := config.DB.Save(&image).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to save image to database")
	}

	setting := models.AppSettings{
		Key:   "regis_bg_path",
		Value: fmt.Sprintf("db_image?t=%d", time.Now().Unix()),
	}
	config.DB.Save(&setting)

	return utils.JSONResponse(c, 200, "Background uploaded to Database", setting)
}

// GetBackgroundImage serves the Registration Background from DB
// --- Background Gallery Handlers ---

// GetBackgrounds lists all background images (Admin)
func GetBackgrounds(c *fiber.Ctx) error {
	var images []models.BackgroundImage
	// Exclude Data field for list view to save bandwidth
	if err := config.DB.Select("id, content_type, is_active, created_at").Order("created_at desc").Find(&images).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to fetch backgrounds")
	}
	return utils.JSONResponse(c, 200, "Success", images)
}

// GetActiveBackgrounds lists active background IDs (Public)
func GetActiveBackgrounds(c *fiber.Ctx) error {
	var images []models.BackgroundImage
	if err := config.DB.Select("id").Where("is_active = ?", true).Order("created_at desc").Find(&images).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to fetch backgrounds")
	}
	// Return list of IDs
	ids := make([]string, len(images))
	for i, img := range images {
		ids[i] = img.ID.String()
	}
	return utils.JSONResponse(c, 200, "Success", ids)
}

// GetBackgroundImage serves the image data (Public)
func GetBackgroundImage(c *fiber.Ctx) error {
	id := c.Params("id")
	var image models.BackgroundImage
	if err := config.DB.Where("id = ?", id).First(&image).Error; err != nil {
		return c.Status(404).SendString("Image not found")
	}

	c.Set("Content-Type", image.ContentType)
	return c.Send(image.Data)
}

// UploadBackgroundGallery handles uploading new background (Admin)
func UploadBackgroundGallery(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return utils.JSONError(c, 400, "Image file is required")
	}

	f, err := file.Open()
	if err != nil {
		return utils.JSONError(c, 500, "Failed to open file")
	}
	defer f.Close()

	buffer := make([]byte, file.Size)
	f.Read(buffer)

	image := models.BackgroundImage{
		ID:          uuid.New(),
		Data:        buffer,
		ContentType: file.Header.Get("Content-Type"),
		IsActive:    true, // Default active
		CreatedAt:   time.Now(),
	}

	if err := config.DB.Create(&image).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to save image")
	}

	return utils.JSONResponse(c, 201, "Image uploaded", image.ID)
}

// ToggleBackgroundActive toggles active state (Admin)
func ToggleBackgroundActive(c *fiber.Ctx) error {
	id := c.Params("id")
	var image models.BackgroundImage
	if err := config.DB.First(&image, "id = ?", id).Error; err != nil {
		return utils.JSONError(c, 404, "Image not found")
	}

	image.IsActive = !image.IsActive
	config.DB.Save(&image)

	return utils.JSONResponse(c, 200, "Updated", image)
}

// DeleteBackground deletes an image (Admin)
func DeleteBackground(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.BackgroundImage{}, "id = ?", id).Error; err != nil {
		return utils.JSONError(c, 500, "Failed to delete")
	}
	return utils.JSONResponse(c, 200, "Deleted", nil)
}

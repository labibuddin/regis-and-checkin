package main

import (
	"log"

	"jogokariyan-backend/config"
	"jogokariyan-backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file (optional if using Docker env vars directly, but good for local dev)
	godotenv.Load()

	// Connect to Database
	config.ConnectDB()

	app := fiber.New()

	// Middleware
	app.Use(cors.New())

	// Static Files
	app.Static("/uploads", "./uploads")

	// Routes
	api := app.Group("/api")

	// Peserta
	api.Get("/check-existing", handlers.CheckExisting)
	api.Post("/register", handlers.Register)
	api.Get("/peserta", handlers.FindPeserta)

	// Admin
	api.Get("/admin/peserta", handlers.GetPesertaList)
	api.Put("/admin/peserta/:id", handlers.UpdatePeserta)
	api.Delete("/admin/peserta/:id", handlers.DeletePeserta)

	// Event & Check-in
	api.Get("/event/active", handlers.GetActiveEvent)
	api.Post("/check-in", handlers.CheckIn)

	// Admin Stats
	api.Get("/admin/stats", handlers.GetDashboardStats)

	// Kegiatan CRUD
	api.Get("/admin/kegiatan", handlers.GetKegiatanList)
	api.Post("/admin/kegiatan", handlers.CreateKegiatan)
	api.Put("/admin/kegiatan/:id", handlers.UpdateKegiatan)
	api.Delete("/admin/kegiatan/:id", handlers.DeleteKegiatan)

	// Settings
	api.Get("/admin/settings", handlers.GetSettings)
	api.Post("/admin/settings/active-event", handlers.SetActiveEvent)
	api.Post("/admin/settings/qris", handlers.UploadQRIS)
	api.Post("/admin/settings/background", handlers.UploadBackground)
	api.Get("/images/qris", handlers.GetQRISImage)
	api.Get("/images/qris", handlers.GetQRISImage)
	// api.Get("/images/background", handlers.GetBackgroundImage) // Deprecated

	// Background Gallery Routes
	api.Get("/backgrounds", handlers.GetActiveBackgrounds)
	api.Get("/backgrounds/:id", handlers.GetBackgroundImage)

	api.Get("/admin/backgrounds", handlers.GetBackgrounds)
	api.Post("/admin/backgrounds", handlers.UploadBackgroundGallery)
	api.Post("/admin/backgrounds/:id/toggle", handlers.ToggleBackgroundActive)
	api.Delete("/admin/backgrounds/:id", handlers.DeleteBackground)

	log.Fatal(app.Listen(":8080"))
}

package main

import (
	"log"

	"jogokariyan-backend/config"
	"jogokariyan-backend/handlers"
	"jogokariyan-backend/middleware"

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
	api.Get("/peserta", handlers.FindPeserta)
	api.Post("/register", handlers.Register)
	// Auth
	api.Post("/login", handlers.Login)

	// Admin
	admin := api.Group("/admin", middleware.Protected())

	admin.Get("/peserta", handlers.GetPesertaList)
	admin.Put("/peserta/:id", handlers.UpdatePeserta)
	admin.Delete("/peserta/:id", handlers.DeletePeserta)

	// Event & Check-in
	api.Get("/event/active", handlers.GetActiveEvent)
	api.Post("/check-in", handlers.CheckIn)

	// Admin Stats
	admin.Get("/stats", handlers.GetDashboardStats)

	// Kegiatan CRUD
	admin.Get("/kegiatan", handlers.GetKegiatanList)
	admin.Post("/kegiatan", handlers.CreateKegiatan)
	admin.Put("/kegiatan/:id", handlers.UpdateKegiatan)
	admin.Delete("/kegiatan/:id", handlers.DeleteKegiatan)

	// Settings
	admin.Get("/settings", handlers.GetSettings)
	admin.Post("/settings/active-event", handlers.SetActiveEvent)
	admin.Post("/settings/qris", handlers.UploadQRIS)
	admin.Post("/settings/background", handlers.UploadBackground)
	api.Get("/images/qris", handlers.GetQRISImage)
	api.Get("/images/qris", handlers.GetQRISImage)
	// api.Get("/images/background", handlers.GetBackgroundImage) // Deprecated

	// Background Gallery Routes
	api.Get("/backgrounds", handlers.GetActiveBackgrounds)
	api.Get("/backgrounds/:id", handlers.GetBackgroundImage)

	admin.Get("/backgrounds", handlers.GetBackgrounds)
	admin.Post("/backgrounds", handlers.UploadBackgroundGallery)
	admin.Post("/backgrounds/:id/toggle", handlers.ToggleBackgroundActive)
	admin.Delete("/backgrounds/:id", handlers.DeleteBackground)

	log.Fatal(app.Listen(":8080"))
}

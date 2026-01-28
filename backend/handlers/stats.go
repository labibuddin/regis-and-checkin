package handlers

import (
	"jogokariyan-backend/config"
	"jogokariyan-backend/models"
	"jogokariyan-backend/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type DashboardStats struct {
	TotalPeserta      int64 `json:"total_peserta"`
	TotalCheckinToday int64 `json:"total_checkin_today"`
	GenderStats       struct {
		LakiLaki  int64 `json:"laki_laki"`
		Perempuan int64 `json:"perempuan"`
	} `json:"gender_stats"`
	AgeStats struct {
		Remaja int64 `json:"remaja"` // < 20
		Dewasa int64 `json:"dewasa"` // 20 - 50
		Lansia int64 `json:"lansia"` // > 50
	} `json:"age_stats"`
	TopPeserta []struct {
		Nama       string `json:"nama"`
		TotalHadir int64  `json:"total_hadir"`
	} `json:"top_peserta"`
	PopularEvents []struct {
		NamaKegiatan string `json:"nama_kegiatan"`
		TotalHadir   int64  `json:"total_hadir"`
	} `json:"popular_events"`
	RegistrationTrend []struct {
		Date  string `json:"date"`
		Count int64  `json:"count"`
	} `json:"registration_trend"`
	ActiveEvent *struct {
		NamaKegiatan string `json:"nama_kegiatan"`
		Waktu        string `json:"waktu"`
		Tempat       string `json:"tempat"`
		Pemateri     string `json:"pemateri"`
	} `json:"active_event"`
}

func GetDashboardStats(c *fiber.Ctx) error {
	var stats DashboardStats

	// 1. Total Peserta
	config.DB.Model(&models.Peserta{}).Count(&stats.TotalPeserta)

	// 2. Total Checkin Today
	startOfDay := time.Now().Truncate(24 * time.Hour)
	config.DB.Model(&models.Presensi{}).Where("waktu_check_in >= ?", startOfDay).Count(&stats.TotalCheckinToday)

	// 3. Gender Stats
	config.DB.Model(&models.Peserta{}).Where("jenis_kelamin = ?", "Laki-laki").Count(&stats.GenderStats.LakiLaki)
	config.DB.Model(&models.Peserta{}).Where("jenis_kelamin = ?", "Perempuan").Count(&stats.GenderStats.Perempuan)

	// 4. Age Stats
	currentYear := time.Now().Year()
	// Remaja: < 20 (Born > currentYear - 20)
	config.DB.Model(&models.Peserta{}).Where("tahun_lahir > ?", currentYear-20).Count(&stats.AgeStats.Remaja)
	// Dewasa: 20 - 50 (Born between currentYear - 50 and currentYear - 20)
	config.DB.Model(&models.Peserta{}).Where("tahun_lahir <= ? AND tahun_lahir >= ?", currentYear-20, currentYear-50).Count(&stats.AgeStats.Dewasa)
	// Lansia: > 50 (Born < currentYear - 50)
	config.DB.Model(&models.Peserta{}).Where("tahun_lahir < ?", currentYear-50).Count(&stats.AgeStats.Lansia)

	// 5. Top Peserta (Most active) - Join Presensi
	type TopPesertaResult struct {
		Nama       string
		TotalHadir int64
	}
	config.DB.Table("presensi").
		Select("peserta.nama_lengkap as nama, count(presensi.id) as total_hadir").
		Joins("join peserta on peserta.id = presensi.peserta_id").
		Group("peserta.nama_lengkap").
		Order("total_hadir desc").
		Limit(5).
		Scan(&stats.TopPeserta)

	// 6. Popular Events
	config.DB.Table("presensi").
		Select("kegiatan.nama_kegiatan, count(presensi.id) as total_hadir").
		Joins("join kegiatan on kegiatan.id = presensi.kegiatan_id").
		Group("kegiatan.nama_kegiatan").
		Order("total_hadir desc").
		Limit(5).
		Scan(&stats.PopularEvents)

	// 7. Active Event
	var setting models.AppSettings
	if err := config.DB.Where("key = ?", "active_event_id").First(&setting).Error; err == nil {
		var kegiatan models.Kegiatan
		if err := config.DB.Where("id = ?", setting.Value).First(&kegiatan).Error; err == nil {
			stats.ActiveEvent = &struct {
				NamaKegiatan string `json:"nama_kegiatan"`
				Waktu        string `json:"waktu"`
				Tempat       string `json:"tempat"`
				Pemateri     string `json:"pemateri"`
			}{
				NamaKegiatan: kegiatan.NamaKegiatan,
				Waktu:        kegiatan.Waktu.Format("02 Jan 2006 15:04"),
				Tempat:       kegiatan.Tempat,
				Pemateri:     kegiatan.Pemateri,
			}
		}
	}

	// 8. Registration Trend (Last 7 Days)
	// Note: Providing data for last 7 days even if 0 needs a loop or complex query.
	// For simplicity, we just query available data in last 7 days and fill gaps in frontend or here if needed.
	// We'll fill gaps here.
	type TrendResult struct {
		Date  string
		Count int64
	}
	var results []TrendResult

	// Postgres specific date truncation
	config.DB.Model(&models.Peserta{}).
		Select("to_char(created_at, 'YYYY-MM-DD') as date, count(id) as count").
		Where("created_at >= ?", time.Now().AddDate(0, 0, -6)).
		Group("date").
		Order("date asc").
		Scan(&results)

	// Fill simple map
	trendMap := make(map[string]int64)
	for _, r := range results {
		trendMap[r.Date] = r.Count
	}

	// Generate last 7 days
	for i := 6; i >= 0; i-- {
		d := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		count := trendMap[d]
		stats.RegistrationTrend = append(stats.RegistrationTrend, struct {
			Date  string `json:"date"`
			Count int64  `json:"count"`
		}{Date: d, Count: count})
	}

	return utils.JSONResponse(c, 200, "Dashboard Stats", stats)
}

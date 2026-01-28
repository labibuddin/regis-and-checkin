package models

import (
	"time"

	"github.com/google/uuid"
)

type Peserta struct {
	ID                       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	NamaLengkap              string    `json:"nama_lengkap"`
	Panggilan                string    `json:"panggilan"`
	JenisKelamin             string    `json:"jenis_kelamin"`
	AlamatDomisili           string    `json:"alamat_domisili"`
	TahunLahir               int       `json:"tahun_lahir"`
	NoHP                     string    `gorm:"unique" json:"no_hp"`
	Email                    string    `gorm:"unique" json:"email"`
	BuktiTransferData        []byte    `gorm:"type:bytea" json:"-"`
	BuktiTransferContentType string    `json:"bukti_transfer_content_type"`
	CreatedAt                time.Time `json:"created_at"`
}

type Kegiatan struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	NamaKegiatan string    `json:"nama_kegiatan"`
	Waktu        time.Time `json:"waktu"`
	Pemateri     string    `json:"pemateri"`
	Tempat       string    `json:"tempat"`
}

type Presensi struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	PesertaID    uuid.UUID `json:"peserta_id"`
	Peserta      Peserta   `gorm:"foreignKey:PesertaID" json:"peserta"`
	KegiatanID   uuid.UUID `json:"kegiatan_id"`
	Kegiatan     Kegiatan  `gorm:"foreignKey:KegiatanID" json:"kegiatan"`
	WaktuCheckIn time.Time `json:"waktu_check_in"`
}

type AppSettings struct {
	Key   string `gorm:"primaryKey" json:"key"`
	Value string `json:"value"`
}

type AppImage struct {
	Key         string `gorm:"primaryKey" json:"key"`
	Data        []byte `gorm:"type:bytea" json:"-"` // Don't expose binary data in JSON
	ContentType string `json:"content_type"`
}

type BackgroundImage struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Data        []byte    `gorm:"type:bytea" json:"-"`
	ContentType string    `json:"content_type"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

// TableName overrides
func (Peserta) TableName() string {
	return "peserta"
}

func (Kegiatan) TableName() string {
	return "kegiatan"
}

func (Presensi) TableName() string {
	return "presensi"
}

func (AppSettings) TableName() string {
	return "app_settings"
}

func (BackgroundImage) TableName() string {
	return "background_images"
}

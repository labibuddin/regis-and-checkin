-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Table: Peserta
CREATE TABLE IF NOT EXISTS peserta (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nama_lengkap VARCHAR(255) NOT NULL,
    panggilan VARCHAR(100),
    jenis_kelamin VARCHAR(20) CHECK (jenis_kelamin IN ('Laki-laki', 'Perempuan')),
    alamat_domisili TEXT,
    tahun_lahir INT,
    no_hp VARCHAR(20) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_peserta_no_hp ON peserta(no_hp);
CREATE INDEX idx_peserta_email ON peserta(email);

-- Table: Kegiatan
CREATE TABLE IF NOT EXISTS kegiatan (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nama_kegiatan VARCHAR(255) NOT NULL,
    waktu TIMESTAMP NOT NULL,
    pemateri VARCHAR(255),
    tempat VARCHAR(255)
);

-- Table: Presensi
CREATE TABLE IF NOT EXISTS presensi (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    peserta_id UUID NOT NULL REFERENCES peserta(id) ON DELETE CASCADE,
    kegiatan_id UUID NOT NULL REFERENCES kegiatan(id) ON DELETE CASCADE,
    waktu_check_in TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_checkin UNIQUE (peserta_id, kegiatan_id)
);

-- Table: App Settings
CREATE TABLE IF NOT EXISTS app_settings (
    key VARCHAR(100) PRIMARY KEY,
    value TEXT
);

-- Insert default dummy event if needed (optional)
-- INSERT INTO kegiatan (nama_kegiatan, waktu, pemateri, tempat) VALUES ('Kajian Rutin Ahad Pagi', NOW(), 'Ustadz Jazir', 'Masjid Jogokariyan');

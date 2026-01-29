# Regis and Check In - Jgkryn

Aplikasi Pendaftaran dan Check-In Peserta untuk event Jogokariyan. Dibangun dengan Go Fiber (Backend), Vue.js + Vite (Frontend), dan PostgreSQL.

## 📋 Prerequisites

Sebelum menjalankan aplikasi, pastikan komputer Anda sudah terinstall:

- **Docker & Docker Compose**: Untuk menjalankan database PostgreSQL.
- **Go** (v1.23+): Untuk menjalankan backend.
- **Node.js** (v18+): Untuk menjalankan frontend.

## 🚀 Quick Start (Local)

Cara termudah untuk menjalankan seluruh aplikasi (Database, Backend, dan Frontend) adalah menggunakan helper script yang sudah disediakan.

1.  **Clone Repository** (jika belum):
    ```bash
    git clone https://github.com/labibuddin/regis-and-checkin.git
    cd regis-and-checkin
    ```

2.  **Setup Environment**:
    Pastikan file `.env` sudah ada di folder `backend/`. Jika belum, copy dari example atau buat baru.
    ```bash
    cp backend/.env.example backend/.env
    ```

3.  **Jalankan Aplikasi**:
    ```bash
    ./scripts/start-all.sh
    ```
    Script ini akan:
    - Menyalakan container database `db` via Docker Compose.
    - Menjalankan Backend Go di port `:8080`.
    - Menjalankan Frontend Vite di port `:1966`.

4.  **Akses Aplikasi**:
    - **Frontend**: [http://localhost:1966](http://localhost:1966)
    - **Backend API**: [http://localhost:8080](http://localhost:8080)
    
    *Tips: Tekan `Ctrl+C` di terminal untuk mematikan semua service.*

---

## 🌐 Public Access (Cloudflare Tunnel)

Ingin membagikan link ke client atau membuka akses dari internet? Gunakan script tunnel yang sudah terintegrasi.

**PENTING**: Pastikan aplikasi utama (`./scripts/start-all.sh`) **SEDANG BERJALAN** di terminal lain sebelum menjalankan tunnel.

1.  Buka terminal baru.
2.  Jalankan perintah berikut:
    ```bash
    ./scripts/start-tunnel.sh
    ```
    Script ini akan otomatis mendownload `cloudflared` (jika belum ada) dan membuat secure tunnel ke aplikasi lokal Anda.

3.  **Copy URL Publik**:
    Terminal akan menampilkan URL sementara, contohnya:
    `https://random-words.trycloudflare.com`
    
    Buka URL tersebut di browser HP atau laptop lain. Frontend dan Backend akan otomatis terhubung dengan lancar.

---

## 🛠️ Access via LAN (WiFi)

Aplikasi ini juga sudah dikonfigurasi untuk bisa diakses langsung via IP dalam satu jaringan WiFi.

1.  Jalankan aplikasi dengan `./scripts/start-all.sh`.
2.  Lihat output terminal bagian "Network".
3.  Buka URL IP address (misal `http://192.168.1.x:1966`) di browser perangkat lain.

---

## 📂 Project Structure

- `/backend`: Go Fiber App (API Service)
- `/frontend`: Vue 3 + Tailwind CSS App (UI)
- `/database`: SQL Init Scripts
- `/scripts`: Helper utilities for running the app

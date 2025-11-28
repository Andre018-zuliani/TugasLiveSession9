# 🚀 Tugas Live Session 9: Implementasi Unit Test Golang (session-9)

Proyek ini adalah implementasi CLI sederhana untuk manajemen data siswa (Student) menggunakan bahasa pemrograman Go (Golang) dengan fokus utama pada penerapan **Unit Testing** pada *Service Layer* dan mencapai target *Code Coverage*.

## 🎯 Target Tugas
Berdasarkan instruksi Latihan Live Session 9:

* Buatkan unit test lanjutan untuk `service/student.go`.
* Mencapai **Code Coverage 60% ke atas**.

## Struktur Proyek

Proyek ini menggunakan arsitektur berlapis (Layered Architecture):

| Folder | Deskripsi |
| :--- | :--- |
| `handler/` | Logika penanganan input dari CLI (`user.go`). |
| `service/` | **Logika Bisnis** utama (`student.go`) - **Fokus Unit Test**. |
| `repository/` | Logika akses data (menggunakan file JSON) (`student.go`). |
| `model/` | Definisi struktur data (`student.go`). |
| `utils/` | Fungsi-fungsi bantuan (errors, file I/O). |
| `data/` | Berisi `student.json` sebagai sumber data. |
| `main.go` | Titik masuk aplikasi (CLI menu). |

## Status Unit Test & Code Coverage

Saat ini, implementasi Unit Test sudah tersedia di `service/student_test.go` dengan menggunakan *in-memory mock repository*.

### Hasil Code Coverage

Berdasarkan laporan terakhir:

| File | Coverage (%) | Status |
| :--- | :--- | :--- |
| `service/student.go` | **18.0%** | 🔴 **Target Belum Tercapai** |
| File lain | 0.0% | |
| **Total Proyek** | Rendah (Perlu Unit Test pada fungsi `Create`, `Update`, `Delete` di Service Layer) | |

<br>

**Detail Coverage di `service/student.go` (18.0%):**
Fungsi yang sudah di-cover adalah `NewStudentService` dan bagian `GetByID` (termasuk kasus *found* dan *not found*).

### 🛠️ Perintah untuk Menjalankan Test

Untuk menjalankan semua unit test dan melihat coverage:

```bash
# Menjalankan semua test
go test ./...

# Menjalankan test dan melihat persentase coverage
go test -cover ./...

# Melihat laporan coverage dalam format HTML di browser
go test -coverprofile=cover.out ./...
go tool cover -html=cover.out

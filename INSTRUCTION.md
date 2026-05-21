# Instruksi Pengerjaan

Dokumen ini wajib dibaca sebelum mulai mengerjakan modul.

---

## 1. Setup Awal

```bash
# Clone repo
git clone https://github.com/24A-TI-ULBI/siakad-2A.git
cd siakad-2A

# Taruh file .env yang dibagikan via WA di root folder
# Jangan commit file .env ke GitHub

# Install dependencies
go mod tidy

# Jalankan aplikasi
go run main.go
# Server berjalan di http://localhost:8080
```

---

## 2. Workflow Git

Note: Sebagai contoh saja, sesuaikan dengan modul yang dikerjakan

```bash
# 1. Pastikan kamu berada di branch modulmu
git fetch origin
git checkout nama-modul
# Contoh: git checkout dosen

# 2. Sync dengan main sebelum mulai kerja
git pull origin main

# 3. Kerjakan modulmu

# 4. Stage hanya file milikmu, JANGAN git add .
git add controller/dosenController.go
git add model/dosen.go
git add url/dosenRoute.go
git add frontend/dosen/index.html
git add url/url.go

# 5. Commit
git commit -m "feat: tambah CRUD dosen dan jabatan"

# 6. Push ke branch-mu
git push origin nama-modul

# 7. Buat Pull Request di GitHub ke branch main
```

Tidak ada yang boleh push langsung ke `main`. Wajib lewat Pull Request.

---

## 3. File yang Dibuat

Setiap mahasiswa hanya membuat 4 file baru:

```
controller/[modul]Controller.go
model/[modul].go
url/[modul]Route.go
frontend/[modul]/index.html
```

Dan edit 1 baris di file ini:

```
url/url.go   tambah pemanggilan fungsi route modulmu
```

---

## 4. File yang Tidak Boleh Disentuh

```
main.go
config/
helper/
go.mod
go.sum
.env
```

File-file ini domain maintainer. Kalau ada yang perlu diubah, hubungi maintainer.

---

## 5. Environment Variables

| Variable | Keterangan |
|---|---|
| `MONGOSTRING` | Connection string MongoDB Atlas |
| `MONGODB_NAME` | Nama database (default: kampus) |
| `PORT` | Port server (diset otomatis oleh Alwaysdata) |
| `IP` | IP binding (diset otomatis oleh Alwaysdata) |

# Go Cinema API

## A. System Design Test

### 1️⃣ Flowchart Sistem

![Flowchart Sistem](assets/flowchart-bioskop.png)

### 2️⃣ Pemilihan Kursi, Restok Tiket & Refund

---

## B. Database Design Test

### 1️⃣ ERD / Database Schema

![ERD Database](assets/erd-bioskop.png)

---

## C. Skill Test

### 1️⃣ Clone Project

```bash
git clone https://github.com/username/go-cinema-api.git
cd go-cinema-api
```

### 2️⃣ Setup Dependencies

```bash
go mod tidy
```

### 3️⃣ Jalankan Project

Pastikan .env sudah dibuat sesuai .env.example:

```bash
go run main.go
```

Server akan berjalan di http://localhost:8080

Endpoint tersedia:

-   POST /login → login user dan dapatkan JWT
-   GET /showtimes → list jadwal tayang (butuh JWT)
-   POST /showtimes → tambah jadwal tayang (butuh JWT)
-   PUT /showtimes/:id → update jadwal tayang (butuh JWT)
-   DELETE /showtimes/:id → hapus jadwal tayang (butuh JWT)

## D. Postman Collection

Untuk memudahkan testing API, silakan download dan import file Postman Collection berikut:

[Download Postman Collection](assets/go-cinema-api.postman_collection.json)

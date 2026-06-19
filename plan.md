# Struktur Implementasi: CRUD & Search Data (Gin-Gonic + SQLite + GORM)

Struktur folder ter-restrukturisasi yang telah diterapkan:
```
go-sqlite-crud/
├── config/
│   └── db.go            # Konfigurasi & Inisialisasi Database
├── handlers/
│   ├── category.go      # Validasi payload & Response HTTP Kategori
│   └── product.go       # Validasi payload & Response HTTP & Query Param Produk
├── services/
│   ├── category.go      # Logika Database & Query Kategori
│   └── product.go       # Logika Database, Pencarian (Search/Filter) & Paginasi Produk
├── models/
│   ├── category.go      # Struct Entitas Kategori
│   └── product.go       # Struct Entitas & Schema Database Produk
├── routes/
│   └── routes.go        # Registrasi endpoint & Grouping Router Gin
├── go.mod
├── go.sum
├── .air.toml            # Konfigurasi Live Reloading (Air)
└── main.go              # Titik Masuk Utama Server (Bootstrap)
```

## Teknologi & Library Utama
1. **Gin Framework**: Untuk routing HTTP, parser request body, dan JSON response.
2. **GORM (SQLite Driver)**: Untuk ORM database SQLite, mempermudah relasi *Kategori - Produk*, mempermudah pencarian (Search), dan melakukan migrasi skema database secara otomatis.

---

## 1. Desain Model & Skema Database (`models/`)

### Kategori (`models/category.go`)
- `ID`: Primary key, auto-increment.
- `Name`: String, tidak boleh kosong (*not null*), unik.
- `CreatedAt` & `UpdatedAt`: Metadata waktu.
- `Products`: Relasi *has-many* (Kategori memiliki banyak Produk).

### Produk (`models/product.go`)
- `ID`: Primary key, auto-increment.
- `CategoryID`: Foreign key ke tabel Kategori.
- `Name`: String, tidak boleh kosong (*not null*).
- `Price`: Float/Decimal, harga produk (harus > 0).
- `Stock`: Integer, jumlah stok (default 0).
- `Description`: Teks deskripsi produk.
- `CreatedAt` & `UpdatedAt`: Metadata waktu.

---

## 2. Konfigurasi Database (`config/db.go`)
- Inisialisasi koneksi ke SQLite file `crud.db`.
- Melakukan auto-migrasi (`AutoMigrate`) untuk memetakan struct models ke dalam tabel database secara otomatis.

---

## 3. Logika Service (`services/`)

### Kategori (`services/category.go`)
- `CreateCategory`: Menyimpan kategori ke database.
- `GetAllCategories`: Mengambil seluruh kategori.
- `GetCategoryByID`: Mengambil detail satu kategori berserta produknya (Preload).
- `UpdateCategory`: Mengubah nama kategori.
- `DeleteCategory`: Menghapus kategori dari database.

### Produk (`services/product.go`)
- `CreateProduct`: Menyimpan produk baru setelah memvalidasi keberadaan `category_id`.
- `GetProducts`: Mengambil seluruh produk berdasarkan filter pencarian (`search`), kategori (`category_id`), harga minimum/maksimum (`min_price`/`max_price`), dan paginasi (`page`/`limit`).
- `GetProductByID`: Mengambil detail satu produk.
- `UpdateProduct`: Mengubah informasi produk.
- `DeleteProduct`: Menghapus produk.

---

## 4. Handler HTTP (`handlers/`)
Menerjemahkan input request HTTP (JSON body, Path parameter, Query parameter) untuk diproses oleh service, dan mengembalikan format respons standar (misal JSON error, status code HTTP).

---

## 5. Routing (`routes/routes.go`)
Mengelompokkan rute API di bawah group `/api/v1` untuk Kategori (`/api/v1/categories`) dan Produk (`/api/v1/products`) serta meneruskannya ke masing-masing Handler yang sesuai.

---

## 6. Entrypoint (`main.go`)
- Menginisialisasi koneksi database melalui `config.InitDB()`.
- Memanggil dan mengkonfigurasi Router melalui `routes.SetupRouter()`.
- Menjalankan server pada port default (`:8080`).

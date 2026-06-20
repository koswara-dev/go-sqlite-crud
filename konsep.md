# Penjelasan Penggunaan Pointer pada Go, Gin, dan GORM

Penggunaan pointer (`*`) sangat krusial dalam bahasa Go untuk mengelola memori secara efisien dan mengizinkan modifikasi data. Berdasarkan berkas [models/product.go](file:///c:/Projects/JuaraCoding/golang-batch-1/go-sqlite-crud/models/product.go), [repositories/category.go](file:///c:/Projects/JuaraCoding/golang-batch-1/go-sqlite-crud/repositories/category.go), dan [repositories/product.go](file:///c:/Projects/JuaraCoding/golang-batch-1/go-sqlite-crud/repositories/product.go), berikut adalah penjelasan mendalam tentang konsep dan alasannya:

---

## 1. Pointer pada Parameter Fungsi/Method (GORM Mutation)
Pada repository, method untuk membuat atau memperbarui data didefinisikan dengan parameter pointer:
```go
// repositories/category.go
Create(category *models.Category) error
Update(category *models.Category) error
```

### Mengapa menggunakan pointer `*models.Category`?
* **Mutasi Data Hasil Operasi (Side Effects)**: Saat GORM melakukan `db.Create(category)`, SQLite akan men-generate kolom auto-increment `ID` serta timestamp `CreatedAt` dan `UpdatedAt`. Karena parameter dikirim sebagai pointer, GORM dapat menuliskan kembali data ID dan timestamp baru tersebut langsung ke alamat memori variabel asli yang dikirim oleh pemanggil. Jika tidak dikirim sebagai pointer, nilai ID pada variabel asli pemanggil akan tetap `0`.
* **Efisiensi Memori (Pass by Reference)**: Secara default, Go menggunakan sifat *pass-by-value* (menyalin seluruh data). Jika struct dikirim tanpa pointer, Go akan menduplikasi struct tersebut di memori. Menggunakan pointer hanya akan menyalin alamat memorinya saja (biasanya hanya 8 byte), sangat efisien untuk objek yang besar.

---

## 2. Pointer pada Hubungan Relasi Struct (Models Association)
Di dalam [models/product.go](file:///c:/Projects/JuaraCoding/golang-batch-1/go-sqlite-crud/models/product.go), kita mendefinisikan relasi `Category` sebagai pointer:
```go
// models/product.go
type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CategoryID  uint      `json:"category_id" binding:"required"`
	Category    *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	// ...
}
```

### Mengapa menggunakan pointer `*Category` alih-alih `Category` langsung?
* **Representasi Nilai Kosong / Opsional (`nil`)**: Di Go, tipe data non-pointer memiliki *zero-value*. Jika kita menggunakan `Category` (bukan pointer), saat relasi kategori tidak di-load (tanpa `.Preload("Category")`), nilainya akan berupa struct kosong `Category{ID: 0, Name: ""}`. Ini akan muncul di response JSON.
* **Integrasi dengan Tag `omitempty`**: Dengan menggunakan pointer `*Category`, jika relasi kategori tidak di-preload, nilainya akan berupa `nil`. Dalam format JSON, nilai `nil` dikombinasikan dengan tag `omitempty` akan membuat field `"category"` **dihapus sepenuhnya** dari response API, membuat JSON menjadi jauh lebih rapi dan bersih.

---

## 3. Pointer Receiver pada Struct Method (Repository & Service)
Method pada repository ditulis dengan receiver berbentuk pointer:
```go
// repositories/product.go
func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}
```

### Mengapa menggunakan pointer receiver `r *productRepository`?
* **Konsistensi & Efisiensi**: Meskipun repository di atas bersifat *stateless* (tidak mengubah isi internal struct repo itu sendiri), menggunakan pointer receiver mencegah Go menyalin objek struct `productRepository` (termasuk pointer koneksi `*gorm.DB` di dalamnya) setiap kali method dipanggil.
* **Kemampuan Mutasi Status**: Jika di masa depan struct repository membutuhkan perubahan state internal (misalnya menyimpan cache, logger statis, dsb), pointer receiver mengizinkan perubahan tersebut bertahan di memori.

---

## 4. Pointer pada Gin Context Binding (Parsing Request JSON)
Pada file handler, kita menuliskan parsing JSON seperti ini:
```go
// handlers/category.go
var category models.Category
if err := c.ShouldBindJSON(&category); err != nil {
	// ...
}
```

### Mengapa kita mengirim `&category`?
* Operator ampersand (`&`) digunakan untuk mengambil alamat memori dari variabel `category` (menjadikannya pointer).
* Fungsi `ShouldBindJSON` milik Gin bertugas membaca request body JSON dari client, mem-parsing-nya, dan memasukkan nilai-nilai tersebut ke dalam variabel penampung. Agar fungsi internal Gin bisa mengisi nilai ke variabel `category` milik kita, Gin wajib mengetahui alamat memori aslinya melalui pointer (`&category`).

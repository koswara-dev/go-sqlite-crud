# Rencana Implementasi: Layer Repository (Refactoring Clean Architecture)

Rencana ini dibuat untuk menambahkan **Layer Repository** sebagai perantara data access antara database (GORM/SQLite) dengan Service layer. Selain itu, refactoring ini akan menerapkan **Dependency Injection** penuh di setiap layer.

Struktur folder ter-update yang akan diterapkan:
```
go-sqlite-crud/
├── config/
│   └── db.go            # Konfigurasi & Inisialisasi Database
├── repositories/        # (BARU) Mengisolasi semua query database
│   ├── category.go      
│   └── product.go       
├── services/            # Mengimplementasikan logika bisnis
│   ├── category.go      
│   └── product.go       
├── handlers/            # Melayani HTTP Bindings & Response format
│   ├── category.go      
│   └── product.go       
├── models/
│   ├── category.go      
│   └── product.go       
├── routes/
│   └── routes.go        # Router & Dependency Injection Wiring
├── go.mod
├── go.sum
└── main.go              # Server Entrypoint
```

---

## Desain Dependency Injection

Aliran data dan dependensi akan mengalir secara eksplisit sebagai berikut:
1. `main.go` menginisialisasi `config.DB` (`*gorm.DB`).
2. `config.DB` dioper ke dalam `routes.SetupRouter(db)`.
3. Di dalam router, instansiasi objek dilakukan secara berurutan:
   - **Repository** menerima dependensi `*gorm.DB`.
   - **Service** menerima dependensi **Repository Interface**.
   - **Handler** menerima dependensi **Service Interface**.
4. Rute API didaftarkan mengacu pada method dari masing-masing **Handler struct**.

---

## 1. Desain Repository (`repositories/`)

### Kategori Repository Interface
```go
type CategoryRepository interface {
    Create(category *models.Category) error
    FindAll() ([]models.Category, error)
    FindByID(id uint) (models.Category, error)
    Update(category *models.Category) error
    Delete(id uint) error
}
```

### Produk Repository Interface
```go
type ProductRepository interface {
    Create(product *models.Product) error
    FindAll(search string, categoryID int, minPrice, maxPrice float64, page, limit int) ([]models.Product, int64, error)
    FindByID(id uint) (models.Product, error)
    Update(product *models.Product) error
    Delete(id uint) error
}
```

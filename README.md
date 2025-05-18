
# 🧮 DRX Fullstack Test - Backend - Golang
---

## 📁 Struktur Direktori

```
.
├── main.go                 
├── handler.go              
├── usecase.go              
├── repository.go           
├── models.go               
├── discount.go             
├── database.go             
├── discount_test.go        
├── usecase_test.go         
├── repository_mock.go      
├── go.mod / go.sum         
├── Dockerfile              
└── .dockerignore / .gitignore
```

---

## 🛠️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/abdghn/drx-fs-test-be.git
cd drx-fs-test-be
```

### 2. Install Go Modules

Pastikan kamu sudah menginstall Go (versi 1.18+ direkomendasikan), lalu jalankan:

```bash
go mod download
```

### 3. Jalankan Aplikasi

```bash
go run main.go
```

Secara default, server akan berjalan di:  
👉 `http://localhost:8080`

---

## ✅ Testing Instructions

### 1. Jalankan Semua Unit Test

```bash
go test ./...
```

### 2. Jalankan Test Tertentu

```bash
go test -v discount_test.go
```

### 3. Jalankan Test dengan Coverage

```bash
go test -cover ./...
```

---

## 🐳 Jalankan dengan Docker (Optional)

### 1. Build Docker Image

```bash
docker build -t drx-fs-test-be .
```

### 2. Jalankan Container

```bash
docker run -p 8080:8080 drx-fs-test-be
```

---

## 📬 Contoh Endpoint API

### GET /products

#### Contoh Response:

```json
[
  {
    "id": 1,
    "name": "Macbook Air M3",
    "description": "Laptop ringan dan cepat",
    "originalPrice": 250,
    "finalPrice": 190
  }
]
```

### POST /products

#### Contoh Request

```json
{
  "name": "Macbook Air M3",
  "description": "Laptop ringan dan cepat",
  "originalPrice": 250,
  "discounts": [
    { "type": "fixed", "value": 20 },
    { "type": "percentage", "value": 10 },
    { "type": "conditional", "condition": 200, "value": 15 },
    {
      "type": "tiered",
      "tiers": [
        { "min": 0, "max": 99, "value": 5 },
        { "min": 100, "max": 199, "value": 10 },
        { "min": 200, "max": 9999, "value": 25 }
      ]
    },
    { "type": "cap", "maxDiscount": 60 }
  ]
}
```

#### Contoh Response:

```json
{
  "appliedDiscounts": [
    {
      "amount": 20,
      "type": "fixed"
    },
    {
      "amount": 23,
      "type": "percentage"
    },
    {
      "amount": 15,
      "type": "conditional"
    },
    {
      "amount": 25,
      "type": "tiered"
    },
    {
      "cappedAt": 60,
      "originalDiscountTotal": 83,
      "type": "cap"
    }
  ],
  "product": {
    "id": 1,
    "name": "Macbook Air M3",
    "description": "Laptop ringan dan cepat",
    "originalPrice": 250,
    "finalPrice": 190
  }
}
```

---

## 📄 License

MIT License

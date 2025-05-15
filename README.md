<!-- @format -->

# 🔐 Go CRUD API with JWT Authentication

A production-ready Golang REST API using Gin, GORM, bcrypt, and JWT that supports:

- ✅ User Registration & Login
- 🔐 JWT Authentication with Middleware
- 🔒 Password Encryption with bcrypt
- 🖼️ Image/File Upload
- 🧱 MVC Architecture

<br>

---

<br>

## 🛠️ Step-by-Step Installation

### 📦 1. Clone the Repository

    git clone https://github.com/prashantyadav008/Go-Crud-using-JWT-Token.git

    cd go-crud-msql

### 📁 2. Setup Environment File

Create a .env file in the root directory:

    cp example_env .env

### 🛢️ 3. Import SQL

    mysql -u root -p < go_cruddb.sql

### 🔧 4. Install Dependencies

    go mod tidy

### 🚀 5. Run the App

    go run cmd/app/main.go

Or with hot reloading (if you installed Air):

    export PATH=$PATH:/home/admin1/go/bin

    echo $(go env GOPATH)/bin

    air

<br>

---

<br>

## 🗃️ Project Folder Structure

```txt
go-crud-msql/
│
├── cmd/app/                 # Entry point (main.go)
├── internal/
│   ├── config/              # Configuration (env loader)
│   ├── database/            # DB connection and migrations
│   ├── handler/             # Route handler logic
│   ├── middleware/          # Custom middleware (e.g., JWT check)
│   ├── models/              # GORM models (User, etc.)
│   ├── repository/          # Database operations (CRUD logic)
│   ├── routes/              # Route grouping
│   ├── service/             # Business logic (e.g., JWT generation)
│   ├── uploads/             # Uploaded images storage
│   └── utils/               # Utility helpers (JWT, bcrypt, etc.)
│
├── tmp/                     # Temp folder (ignored by git)
├── go.mod                   # Module config
├── go.sum                   # Dependency lock file
├── README.md                # This file 📄
└── .gitignore               # Git ignore rules
```

<br>

---

<br>

## 🧩 Technology & Libraries

### 📦 1. Gin

A fast web framework for building APIs in Go.
URL routing, JSON binding, middleware handling.

**Example:**

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"msg": "pong"})
    })

### 🧰 2. GORM

ORM library to handle database operations.
Works with Go structs for models.

**Example:**

    db.Create(&User{Name: "Alice"})

### 🔑 3. bcrypt

Used for secure password encryption.
Compares hashed password securely.

**Example:**

    hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

### 🔐 4. JWT (github.com/golang-jwt/jwt/v5)

Used for generating/verifying JSON Web Tokens.
Protects APIs by verifying the token.

**Example:**

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

<br>

---

<br>

## ⚙️ Setup Commands

    go mod init go-crud-msql

    mkdir -p cmd/app
    mkdir -p internal/config
    mkdir -p internal/database
    mkdir -p internal/handler
    mkdir -p internal/middleware
    mkdir -p internal/models
    mkdir -p internal/repository
    mkdir -p internal/routes
    mkdir -p internal/service
    mkdir -p internal/uploads
    mkdir -p internal/utils

    go get github.com/joho/godotenv
    go get gorm.io/driver/mysql
    go get gorm.io/gorm
    go get github.com/golang-jwt/jwt/v5
    go get golang.org/x/crypto/bcrypt

## 🛠️ Helpful Commands

    # Install Air for auto reload (optional)
    go install github.com/cosmtrek/air@latest

    # Run project
    go run cmd/app/main.go

    # OR run with Air hot reload
    air

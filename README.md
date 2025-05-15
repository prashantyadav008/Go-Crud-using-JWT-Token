<!-- @format -->

# Go-Crud-using-JWT-Token

go mod tidy

export PATH=$PATH:/home/admin1/go/bin
echo $(go env GOPATH)/bin
air

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
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt/v5

<!-- gin -->

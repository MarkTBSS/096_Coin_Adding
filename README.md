```
go get github.com/go-playground/validator/v10
go get github.com/spf13/viper

go get gorm.io/gorm
go get gorm.io/driver/postgres

docker run -d -e POSTGRES_PASSWORD=Pass1234 -p 5434:5432 <image_name>

Create database name "isekaishopdb"
go run ./databases/migration/migration.go
go run ./databases/migration/migration_item/migration_item.go

go run .
localhost:8080/v1/player-coin
```

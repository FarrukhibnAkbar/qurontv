package utils

import(
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
)

var dsn = "host=localhost user=admin password=12345 dbname=qurontv port=5432 sslmode=disable TimeZone=Asia/Tashkent"
var DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
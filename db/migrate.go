package db

import (
	"go-fiber-hairstyle/internal/categories"
	"go-fiber-hairstyle/internal/hairstyles"
	"log"
)

// Migrate melakukan migrasi tabel-tabel yang diperlukan
func Migrate() {
	err := DB.AutoMigrate(
		&hairstyles.HairStyle{},
		&categories.Category{},
		&hairstyles.HairStyleCategory{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}
}

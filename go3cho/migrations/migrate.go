package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

func InitMigrate() {
	fmt.Println("1")
}

func Migrate(db *gorm.DB) {
	fmt.Println("Migrating...")
}

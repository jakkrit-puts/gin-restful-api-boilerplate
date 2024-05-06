package configs

import (
	"fmt"
	"os"

	"example.com/jakkrit/ginbackendapi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := os.Getenv("DB_DNS")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("----------------------------------")
		fmt.Println("Status: Database Connection Fail!")
		panic(err)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("Status: Database Connection Success.")
	fmt.Println("-------------------------------------")

	// Migration
	// db.Migrator().DropTable(&models.User{})

	db.AutoMigrate(&models.User{}, &models.Blog{})

	DB = db
}

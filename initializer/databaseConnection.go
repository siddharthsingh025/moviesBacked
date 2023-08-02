package initializer

import (
	"example/movieBackend/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDB() {
	var err error
	// host := os.Getenv("DB_HOST")
	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// databaseName := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, username, password, databaseName, port)

	dsn := "host=localhost user=postgres password=postgres dbname=moviesrecord port=3333 sslmode=disable TimeZone=Asia/Shanghai"
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// Migrate the schema
	Database.AutoMigrate(&models.Movie{})
	Database.AutoMigrate(&models.Director{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

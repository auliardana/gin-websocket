package databases

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"GO-RESTAPI-GIN/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USERNAME"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),

	// )
    database, err := gorm.Open(mysql.Open("root:password@tcp(db:3306)/go_restapi_gin"))
    // database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin"))
    // database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
	database.AutoMigrate(&models.Mahasiswa{})
	DB = database
}

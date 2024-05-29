package initializers

import (
	"camp-summer/internal/config"
	"camp-summer/internal/model/app"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s",
		config.Cfg.Database.Host,
		config.Cfg.Database.User,
		config.Cfg.Database.Password,
		config.Cfg.Database.Name,
		config.Cfg.Database.Port,
		config.Cfg.Database.Timezone,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", "", "")
	}

	log.Println("DB connected successfully", "", "")

	err = DB.AutoMigrate(&app.App{})
	if err != nil {
		log.Fatal("Failed to migrate to database", "", "")
	}

	log.Println("DB migrated successfully", "", "")
}

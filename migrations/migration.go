package migrations

import (
	"biography-api/config"
	"biography-api/models"
	"log"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&models.Experience{},
		&models.User{},
		&models.Biography{},
		&models.Awardee{},
		&models.Organization{},
		&models.Skill{},
		&models.TechnicalExperience{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	log.Println("Migration completed successfully")
}

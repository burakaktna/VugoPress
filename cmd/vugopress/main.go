// cmd/main.go
package main

import (
	"fmt"
	"github.com/burakaktna/VugoPress/internal/config"
	"github.com/burakaktna/VugoPress/internal/handlers"
	"github.com/burakaktna/VugoPress/internal/middleware"
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/pkg/utils"
	"github.com/jinzhu/gorm"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg := config.New()

	db, err := utils.ConnectToDatabase(cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Veritabanı bağlantısı kapatılamadı: %v", err)
		}
	}(db)

	// Migrate models
	migrateModels(db)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Handlers
	handlers.RegisterHandlers(app, db, middleware.JwtMiddleware(cfg.AppKey))

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.AppPort)))
}
func migrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.Article{}, &models.Contact{}, &models.Tag{}, &models.UsefulLink{}, &models.SocialMediaLink{})
}

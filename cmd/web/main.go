package main

import (
	_ "backend/docs"
	"backend/internal/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	viper.SetDefault("PORT", "3000")
	viper.SetDefault("PREFORK", false)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	app := fiber.New(
		fiber.Config{
			Prefork:           viper.GetBool("PREFORK"),
			EnablePrintRoutes: true,
		},
	)

	db := config.NewDatabase()
	v := config.NewValidator()

	config.Bootstrap(
		&config.BootstrapConfig{
			App:       app,
			DB:        db,
			Validator: v,
		},
	)

	// setup cors

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		},
	))

	app.Listen(":" + viper.GetString("PORT"))
}

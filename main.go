package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/oliver7100/fmf-rest/controllers"
	"github.com/oliver7100/fmf-rest/database"
	"github.com/oliver7100/fmf-rest/internal"
)

func main() {
	app := fiber.New()

	api := app.Group("api")

	imageStore, err := internal.NewCloudinaryStorage(
		&internal.CloudinaryStorageConfig{
			Cloud:  "zanzanzan",
			Key:    "748773632958652",
			Secret: "a5puHSHwEyy12RtXBz44fPr104s",
		},
	)

	if err != nil {
		panic(err)
	}

	dbConn, err := database.NewDatabaseConnection(
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("db.USERNAME"),
			os.Getenv("db.PASSWORD"),
			os.Getenv("db.HOSTNAME"),
			os.Getenv("db.PORT"),
			os.Getenv("db.DATABASE"),
		),
	)

	if err != nil {
		panic(err)
	}
	controllers.RegisterAuthController(api, dbConn)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	controllers.RegisterAdvertisementController(api, dbConn)
	controllers.RegisterUserController(api, dbConn)
	controllers.RegisterTaskController(api, dbConn, imageStore)

	panic(app.Listen(":3000"))
}

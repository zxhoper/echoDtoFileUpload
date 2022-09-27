package main

import (
	"fmt"

	_ "github.com/cloudinary/cloudinary-go"
	_ "github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello from My go app!")
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello from Cloudinary"})
	})

	e.Logger.Fatal(e.Start(":6000"))
}

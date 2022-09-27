package helper

import (
	"github.com/labstack/echo/v4"

	_ "github.com/cloudinary/cloudinary-go"
	_ "github.com/go-playground/validator/v10"
)

// Curl command to test: curl 127.0.0.1:6000
// result: {"data":"Hello from Cloudinary"}
func main() {
	//fmt.Println("Hello from My go app!")
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello from Cloudinary"})
	})

	e.Logger.Fatal(e.Start(":6000"))
}

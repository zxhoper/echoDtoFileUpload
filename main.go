package helper

import (
	"echoDtoFileUpload/controllers"
	"github.com/labstack/echo/v4"
)

// Curl command to test: curl 127.0.0.1:6000
// result: {"data":"Hello from Cloudinary"}
func main() {
	//fmt.Println("Hello from My go app!")
	e := echo.New()

	e.POST("/file", controllers.FileUpload)
	e.POST("/remote", controllers.RemoteUpload)

	e.Logger.Fatal(e.Start(":6000"))
}

package main

import "github.com/gin-gonic/gin"

// Endpoint routes to file names i.e. Key : Value
func setupStaticFiles(app *gin.Engine) {
	staticFiles := map[string]string{
		"/public/script.js":   "./src/public/script.js",
		"/assets/favicon.ico": "./src/assets/favicon.ico",
		"/src/app.html":       "./app/my-skeleton-app/app.html",
		"/":                   "./app/my-skeleton-app/src/login.html",
	}

	for urlPath, filePath := range staticFiles {
		app.StaticFile(urlPath, filePath)
	}
}

func main() {

	app := gin.Default()
	setupStaticFiles(app)
	app.Run(":8080")

}

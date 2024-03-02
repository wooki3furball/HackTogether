package main

import (
	"fmt"
	"log"
	"math/rand"

	"Toegether/mtd" // Local Module/relative_directory with package

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

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

const charsetAlpha = "abcdefghijklmnopqrstuvwxyz"
const charsetAlphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {

	app := gin.Default()
	setupStaticFiles(app)

	// Generate a random string
	randomStr := mtd.RandomString(10, charsetAlphaNum)
	// Seed the randomness with the generated string
	mtd.SeedRandomnessWithString(randomStr)
	randomDuration := mtd.GenerateRandomTimeInterval()
	fmt.Println("Random Time Interval:", randomDuration)

	// Generate a random number of endpoints, e.g., between 1 and 10
	numEndpoints := rand.Intn(10) + 1
	fmt.Printf("Creating %d random endpoints\n", numEndpoints)

	for i := 0; i < numEndpoints; i++ {
		// Generate a random path and message for each endpoint
		randomPath := "/" + mtd.RandomString(5, charsetAlpha) // Assume RandomString function exists and is correctly implemented
		randomMessage := "Message for " + randomPath

		// Register the endpoint using the factory
		err := mtd.RegisterEndpoint(app, randomPath, randomMessage)
		if err != nil {
			fmt.Println(err) // Log the error and continue
			continue
		}
	}

	store := sessions.NewCookieStore([]byte("secret"))
	app.Use(sessions.Sessions("mysession", store))

	caretaker := &mtd.Caretaker{}

	app.GET("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("name", "Gopher")
		session.Save()

		// Save session state
		originator := &mtd.SessionOriginator{Session: session}
		caretaker.AddMemento(originator.SaveToSessionMemento())
	})

	app.GET("/restore", func(c *gin.Context) {
		session := sessions.Default(c)
		// Save session state
		originator := &mtd.SessionOriginator{Session: session}

		// Assuming we want to restore the first saved state
		if len(caretaker.MementoList) > 0 {
			originator.RestoreFromSessionMemento(caretaker.GetMemento(0))
		}
	})

	err := mtd.RegisterEndpoint(app, "/hello", "Hello, World!")
	if err != nil {
		log.Fatal(err)
	}

	// Default Route Port
	app.Run(":8080")
}

// Flags for algorithm to reconfigure in shorter time periods for Project Demo

// Factory Functions for endpoint generation

// MTD Algorithm with a time interval, calls strategy to reconfigure code

// Place Query functions in main.go

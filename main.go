package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"

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

const charsetAlphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {

	app := gin.Default()
	setupStaticFiles(app)

	// Generate a random string
	randomStr := randomString(10, charsetAlphaNum)
	// Seed the randomness with the generated string
	seedRandomnessWithString(randomStr)
	randomDuration := generateRandomTimeInterval()
	fmt.Println("Random Time Interval:", randomDuration)

	// Default Route Port
	app.Run(":8080")
}

// randomString generates a random string of length n using a specified character set.
func randomString(n int, charset string) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// seedRandomnessWithString takes a string, hashes it, and uses it to seed the random number generator.
func seedRandomnessWithString(s string) {
	hash := sha256.Sum256([]byte(s))
	seed := int64(binary.BigEndian.Uint64(hash[:8]))
	rand.Seed(seed)
}

func generateRandomTimeInterval() time.Duration {
	// (inclusive)
	min, max := 30, 120
	randomSeconds := rand.Intn(max-min+1) + min

	// Convert the random number of seconds to a time.Duration and return
	return time.Duration(randomSeconds) * time.Second
}

// Flags for algorithm to reconfigure in shorter time periods for Project Demo

// Factory Functions for endpoint generation

// Memento Function // Prevent SQL Injection/Brute Force?

// MTD Algorithm with a time interval, calls strategy to reconfiggure code

package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"Toegether/mtd" // Local Module/relative_directory with package

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

/* All components needed for a Database URI. */
type QueryConfig struct {
	Name string
	Host string
	Port string
	User string
	Pass string
	SSL  string
	Cert string
}

/* Load Database Environment variables. */
func loadDBEnv() QueryConfig {
	return QueryConfig{
		Name: os.Getenv("DB_NAME"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		SSL:  os.Getenv("DB_SSL"),
		Cert: os.Getenv("CA_CERT"),
	}
}

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

	/* Endpoint for a web server user to login. */
	app.POST("/submit-login", func(c *gin.Context) {
		dbName := "exampledb"
		userName := c.PostForm("username") // Grab from the HTML frontend form
		password := c.PostForm("password")
		ip := c.ClientIP()
		attempt := FetchLoginAttempt(userName, ip)

		// Check if the current attempt is allowed
		if time.Since(attempt.LastAttempt) < ExponentialBackoffDelay(attempt) {
			// HTTP 429 Error
			c.String(http.StatusTooManyRequests, "Too many failed login attempts. Please try again later.")
			return
		}

		// Get user details from the database & log attempt
		user, userExists, err := FetchUser(userName)
		if err != nil {
			// Handle error (such as database connectivity issues)
			c.String(http.StatusInternalServerError, "Error fetching user")
			return
		}

		if !userExists {
			fmt.Printf("Login attempt for non-existing user: %s\n", userName)
			c.String(http.StatusUnauthorized, "Invalid credentials")
			return
		}

		// Use `user.PasswordHash` and `user.UserType` to validate and handle login logic
		if checkPasswordHash(password, user.PasswordHash) {
			UpdateLoginAttempt(userName, ip, true)
			// User authenticated, set session
			session := sessions.Default(c)
			session.Set("name", userName)
			if err := session.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
				return
			}

			// Construct a parameterized query
			query := `UPDATE logins SET "LoginSuccess" = true 
		  WHERE "UserName" = $1 AND "LoginDate" = (
			  SELECT MAX("LoginDate") 
			  FROM logins 
			  WHERE "UserName" = $1)`

			_, err := ExecuteQuery(query, dbName, userName)
			if err != nil {
				fmt.Printf("Error executing query: %v\n", err) // Detailed Query Error
				// log.Printf("Query: %s\n", query``)
				c.String(http.StatusInternalServerError, fmt.Sprintf("Query execution failed: %v", err1))
				return
			}
		}
	})

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

/* Generic Execute Query Function */
func ExecuteQuery(query string, dbName string, args ...interface{}) (string, error) {
	db := loadDBEnv()

	// Override DB_NAME if a radio button is selected in admin.html
	if dbName == "" {
		dbName = db.Name
	} else {
		db.Name = dbName
	}

	if db.Host == "" || db.Port == "" || db.User == "" || db.Name == "" || db.SSL == "" || db.Cert == "" {
		log.Fatal("One or more required DB environment variables are not set")
	}

	// Logging to GIN Terminal
	fmt.Printf("Accessing DB: %s \n", db.Name)

	// Build the connection string
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", db.User, db.Pass, db.Host, db.Port, db.Name, db.SSL)

	// Parse the original URI and append SSL options
	parsedURL, err := url.Parse(dbURI)
	if err != nil {
		log.Fatal("Failed to parse DATABASE_URL:", err)
	}

	// Append SSL root certificate to the query parameters if sslmode requires it
	if db.SSL != "disable" {
		q := parsedURL.Query()
		q.Set("sslrootcert", db.Cert)
		parsedURL.RawQuery = q.Encode()
	}

	// Connect to the database
	database, err := sql.Open("postgres", parsedURL.String())
	if err != nil {
		return "", err
	}
	defer database.Close()

	// Check the connection
	err = database.Ping()
	if err != nil {
		return "", err
	}

	// Execute the query with parameters
	rows, err := database.Query(query, args...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// Get column types
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return "", err
	}

	// Create a slice of interfaces to hold each value
	columns := make([]interface{}, len(columnTypes))
	columnPointers := make([]interface{}, len(columnTypes))
	for i := range columns {
		columnPointers[i] = &columns[i]
	}

	// Process query results
	var result string
	for rows.Next() {
		// Scan the result into our column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return "", err
		}

		// Convert each column to a string and format the result
		var rowResult string
		for i, col := range columns {
			var colStr string
			if col != nil {
				colStr = fmt.Sprintf("%v", col)
			}
			rowResult += fmt.Sprintf("%s: %s, ", columnTypes[i].Name(), colStr)
		}
		result += rowResult + "\n"
	}

	return result, nil
}

// Flags for algorithm to reconfigure in shorter time periods for Project Demo

// Factory Functions for endpoint generation

// MTD Algorithm with a time interval, calls strategy to reconfigure code

// Place Query functions in main.go

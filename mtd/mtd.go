package mtd

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// randomString generates a random string of length n using a specified character set.
func RandomString(n int, charset string) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// seedRandomnessWithString takes a string, hashes it, and uses it to seed the random number generator.
func SeedRandomnessWithString(s string) {
	hash := sha256.Sum256([]byte(s))
	seed := int64(binary.BigEndian.Uint64(hash[:8]))
	rand.Seed(seed)
}

func GenerateRandomTimeInterval() time.Duration {
	// (inclusive)
	min, max := 30, 120
	randomSeconds := rand.Intn(max-min+1) + min

	// Convert the random number of seconds to a time.Duration and return
	return time.Duration(randomSeconds) * time.Second
}

func EndpointFactory(path string, message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"path":    path,
			"message": message,
		})
	}
}

func RegisterEndpoint(router *gin.Engine, path string, message string) error {
	if _, exists := RegisteredPaths[path]; exists {
		return fmt.Errorf("duplicate endpoint: %s", path)
	}
	router.GET(path, EndpointFactory(path, message))
	RegisteredPaths[path] = true
	return nil
}

package configs

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func beforeAll() {
	if _, err := os.Stat(".env.test"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("exists .env.test")
}

func TestGetConfig(t *testing.T) {
	beforeAll()

	env := GetEnv()

	assert.Equal(t, env.Env, "test")
	assert.Equal(t, env.IsPork, false)
	assert.Equal(t, env.MajorVersion, 1)
	assert.Equal(t, env.Port, "3001")
	assert.Equal(t, env.Version, "1.0.0")

}

package configs

import (
	"log"
	"os"
	"strings"

	"github.com/fiber/utils"
	"github.com/joho/godotenv"
)

type secretEnv struct {
	Env          string
	Port         string
	IsPork       bool
	Version      string
	MajorVersion int
}

func GetEnv() secretEnv {

	var err error

	if utils.StringToBool(os.Getenv("IS_LOCAL")) {
		// use Docker
		switch os.Getenv("ENV") {
		case "dev":
			err = godotenv.Load(".env.dev")
		case "prod":
			err = godotenv.Load(".env.prod")
		case "debug":
			err = godotenv.Load(".env.debug")
		case "test":
			err = godotenv.Load(".env.test")
		default:
			err = godotenv.Load(".env.test")
		}
	}

	if err != nil {
		log.Fatalf("Error loading %s file => %s\n", os.Getenv("ENV"), err)
	}

	version, majorVersion := getVersion(os.Getenv("VERSION"))

	return secretEnv{
		Env:          os.Getenv("ENV"),
		Port:         os.Getenv("PORT"),
		IsPork:       utils.StringToBool(os.Getenv("IS_PORK")),
		Version:      version,
		MajorVersion: utils.StringToInt(majorVersion),
	}
}

func getVersion(version string) (string, string) {

	major := strings.Split(version, ".")
	return version, major[0]

}

package configs

import (
	"log"
	"os"
	"strings"

	"github.com/fiber/src/utils"
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
	switch os.Getenv("ENV") {
	case "dev":
		err = godotenv.Load(".env.dev")
	case "prod":
		err = godotenv.Load(".env.prod")
	case "debug":
		err = godotenv.Load(".env.debug")
	default:
		log.Fatalf("%s not exists", os.Getenv("ENV"))
	}

	if err != nil {
		log.Fatalf("Error loading %s file => %s", os.Getenv("ENV"), err)
	}

	version, majorVersion := getVersion(os.Getenv("VERSION"))

	return secretEnv{
		Env:          os.Getenv("ENV"),
		Port:         os.Getenv("PORT"),
		IsPork:       utils.StringToBool(os.Getenv("IsPork")),
		Version:      version,
		MajorVersion: utils.StringToInt(majorVersion),
	}
}

func getVersion(version string) (string, string) {

	major := strings.Split(v, ".")
	return version, major[0]

}

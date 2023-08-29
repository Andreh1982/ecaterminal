package environment

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

// Single is the singleton instance of the environment
type Single struct {
	ENVIRONMENT string // nolint: golint
	APP_VERSION string // nolint: golint
	AWS_REGION  string // nolint: golint
	AWS_PROFILE string // nolint: golint
	LOG_LEVEL   string // nolint: golint
	KUBE_CONFIG string // nolint: golint
	PROFILE_CFG string // nolint: golint
}

func init() {
	if os.Getenv("ENVIRONMENT") == "development" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Println("Error loading .env.local file")
		}
	}
	env := GetInstance()
	env.Setup()
}

func (e *Single) Setup() {
	e.ENVIRONMENT = os.Getenv("ENVIRONMENT")
	e.APP_VERSION = os.Getenv("APP_VERSION")
	e.AWS_REGION = getenv("AWS_REGION", "sa-east-1")
	e.AWS_PROFILE = getenv("AWS_PROFILE", "")
	e.LOG_LEVEL = getenv("LOG_LEVEL", "debug")
	e.KUBE_CONFIG = getenv("KUBE_CONFIG", "/home/atorres/.kube/config")
	e.PROFILE_CFG = getenv("PROFILE_CFG", "/home/atorres/.aws/credentials")
}

func (e *Single) IsDevelopment() bool {
	return e.ENVIRONMENT == "development"
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &Single{}
			singleInstance.Setup()
		} else {
			fmt.Println("Environment is already set")
		}
	}
	return singleInstance
}

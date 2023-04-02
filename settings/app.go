package settings

import (
	"os"
	"strconv"
	"sync"
)

type app struct {
	ServerPort int
	BasePath string
}

var singletonApp *app
var lock = &sync.Mutex{}

func InitializeApp() *app {
	if singletonApp != nil {
		return singletonApp
	}

	lock.Lock()
	defer lock.Unlock()

	var serverPort int

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))

	if err != nil {
		serverPort = 3006
	}

	singletonApp = &app{
		ServerPort: serverPort,
		BasePath: "/api/veterinary-employee",
	}

	return singletonApp
}

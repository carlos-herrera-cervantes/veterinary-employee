package settings

import (
	"os"
	"strconv"
)

type app struct {
	ServerPort int
}

func InitializeApp() app {
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))

	return app{
		ServerPort: serverPort,
	}
}

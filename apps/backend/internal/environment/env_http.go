package environment

import (
	"os"
	"strconv"
)

func HttpPort(defaultPort int) int {
	portAsStr, isSet := os.LookupEnv("HTTP_PORT")

	if !isSet {
		return defaultPort
	}

	port, err := strconv.Atoi(portAsStr)

	if err != nil {
		panic(err)
	}

	return port
}

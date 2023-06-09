package envservice

import "os"

func get(name string, def string) string {
	val := os.Getenv(name)

	if len(val) > 0 {
		return val
	} else {
		return def
	}
}

func HttpApiListen() string {
	return get("http_api_listen", ":8080")
}

func ClientRepoDataFilename() string {
	return get("data_clients", "./clients.data.yml")
}

func AccessLogFile() string {
	return get("log_access_filename", "./access.log")
}

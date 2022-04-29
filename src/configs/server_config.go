package configs

var ServerConfig = make(map[string]string)

func loadServerConfig() {
	ServerConfig["port"] = ":8000"
}

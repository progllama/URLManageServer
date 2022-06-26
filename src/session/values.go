package session

const SESSION_SECRET = "SESSION_SECRET"
const SESSION_TYPE = "SESSION_TYPE"

const (
	COOKIE     string = "COOKIE"
	REDIS      string = "REDIS"
	MEM_CACHED string = "MEM_CASHED"
	MONGO_DB   string = "MONGO_DB"
	GORM       string = "GORM"
	MEMORY     string = "MEMORY"
	POSTGRESQL string = "POSTGRESQL"
)

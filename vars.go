package crud

import (
	"fmt"
	"os"
)

var (
	dbUser      = os.Getenv("DB_USER")
	dbPassword  = os.Getenv("DB_PASSWORD")
	dbHost      = os.Getenv("DB_HOST")
	dbPort      = os.Getenv("DB_PORT")
	dbName      = os.Getenv("DB_NAME")
	dbCharset   = "utf8mb4"
	dbParseTime = "True"
	dbLoc       = "Local"
)
var dsn = fmt.Sprintf(
	"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
	dbUser,
	dbPassword,
	dbHost,
	dbPort,
	dbName,
	dbCharset,
	dbParseTime,
	dbLoc,
)

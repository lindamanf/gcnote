package users_db

import (
	"database/sql"
	"fmt"
	"log"
)

// const (
// 	mysqlUsersUsername = "mysql_users_username"
// 	mysqlUsersPassword = "mysql_users_password"
// 	mysqlUsersHost     = "mysql_users_host"
// 	mysqlUsersSchema   = "mysql_users_schema"
// )

var (
	Client *sql.DB

	// username = os.Getenv(mysqlUsersUsername)
	username = "gcnote"
	// password = os.Getenv(mysqlUsersPassword)
	password = "gcnote"
	// host     = os.Getenv(mysqlUsersHost)
	host = "gcnote_db_1"
	// schema   = os.Getenv(mysqlUsersSchema)
	schema = "gcnote"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}

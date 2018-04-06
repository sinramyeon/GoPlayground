package config

import (
	"hero0926-api-test/config/cfcookie"
	"hero0926-api-test/config/cflog"
	"hero0926-api-test/config/cfmongo"
	"hero0926-api-test/config/cfpostgres"
	"hero0926-api-test/config/cfredis"
	"hero0926-api-test/util"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx"
)

// Configuration ...
type Configuration struct {
	Postgres cfpostgres.Postgres
	Redis    cfredis.Redis
	Mongo    cfmongo.Mongo
	Secret   Secret
}

// Secret ...
// Secret Key for auth ...
type Secret struct {
	Key string
}

// Global variable
var (

	// ENV is Configuration instance.
	ENV          *Configuration
	SQL          *pgx.Conn
	RedisSession *cfredis.Redis
	CookiesNM    cfcookie.CookiesNM
	Log          *cflog.Log
)

// init() ...
// Init Config Files and Setup DB
func init() {
	err := newConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	SQL, err = cfpostgres.NewPostgres(ENV.Postgres)
	if err != nil {
		log.Fatalln(err)
	}

	Log = cflog.SavePrintLog()

	RedisSession = cfredis.NewPool(ENV.Redis.URL, 0)
	err = RedisSession.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	//cfmongo.Connect(ENV.Mongo)

	InfoF("Setup DB")
	SQL.Exec(InitSQL)
	InfoF("Create DB")

}

// newConfiguration ...
// make new configuration by config.json file
func newConfiguration() error {
	ENV = &Configuration{}

	var path string

	pwd, _ := os.Getwd()
	if strings.Contains(pwd, "config") {
		path = pwd + "/config/config.json"

	} else {
		os.Chdir("../config")
		pwd, _ = os.Getwd()
		path = pwd + "/config.json"
	}

	if !strings.Contains(path, "config/") {
		path = pwd + "/config/config.json"
	}
	return util.RequireJSON(path, ENV)
}

// Panic ...
func Panic(e error) {
	Log.Error.Panic(e)
}

// Errorf ...
func Errorf(format string, a ...interface{}) {
	Log.Error.Printf(format, a...)
}

// InfoF ...
func InfoF(format string, a ...interface{}) {
	Log.Info.Printf(format, a...)
}

// Trace ...
func Trace(format string, a ...interface{}) {
	Log.Trace.Printf(format, a...)
}

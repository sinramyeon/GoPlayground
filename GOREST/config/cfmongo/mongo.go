package cfmongo

import (
	"log"
	"strconv"

	mgo "gopkg.in/mgo.v2"
)

// Sessions ...
//var Sessions map[string]*mgo.Session

// Session ...
var Session *mgo.Session

// Mongo ...
type Mongo struct {
	URL  string
	Pool int
	DBs  map[string]string
}

// Connect ...
// Connect With Mongo
func Connect(m Mongo) {
	s, err := mgo.Dial(m.URL + "?maxPoolSize=" + strconv.Itoa(m.Pool))
	if err != nil {
		log.Fatal(err)
	}
	s.SetMode(mgo.Monotonic, true)
	Session = s
}

// GetSession ...
// Get Mongo Session
func GetSession() *mgo.Session {
	return Session.Copy()
}

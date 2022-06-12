package mongodb

import (
	"os"

	"gopkg.in/mgo.v2"
)

var (
	mongoUri    = os.Getenv("MONGO_URI")
	mongoUser   = os.Getenv("MONGO_USER")
	mongoPasswd = os.Getenv("MONGO_PASSWD")
)

// Connect to a MongoDB instance located by mongo-uri using the `mgo`
// driver.
func NewMongo(mongoUri string, failFast bool) (*mgo.Database, func(), error) {
	di, err := mgo.ParseURL(mongoUri)
	if err != nil {
		return nil, doNothing, err
	}

	di.FailFast = failFast
	session, err := mgo.DialWithInfo(di)

	if err != nil {
		return nil, doNothing, err
	}

	return session.DB(di.Database), session.Close, nil
}

func doNothing() {}

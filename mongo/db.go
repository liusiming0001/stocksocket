package mongo

import "gopkg.in/mgo.v2"

var (
	MongoSession *mgo.Session
)

func InitDb() {
	MongoSession, _ = mgo.Dial("127.0.0.1:27017")
}

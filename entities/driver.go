package entities

import mgo "gopkg.in/mgo.v2"

var (
	mgoSession   *mgo.Session
	databaseName = "youkdb"
)

// GetSession : todo
func GetSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err) // no, not really
		}
	}
	return mgoSession.Clone()
}

package entities

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Group : todo
type Group struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Title    string        `json:"title"`
	UsersID  []string      `json:"UsersID"`
	Creation time.Time     `json:"creation"`
}

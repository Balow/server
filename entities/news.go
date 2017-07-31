package entities

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// News : todo
type News struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	GroupID  bson.ObjectId `json:"groupId" bson:"groupId"`
	Title    string        `json:"title"`
	Category string        `json:"category"`
	Content  string        `json:"content"`
	Creation time.Time     `json:"creation"`
}

// GetNewsByID : todo
func GetNewsByID(id string) News {
	session := GetSession()
	c := session.DB("store").C("News")

	data := News{}
	c.FindId(bson.ObjectIdHex(id)).One(&data)

	return data
}

// AddNews : todo
func AddNews(news News) News {
	session := GetSession()
	c := session.DB("store").C("News")

	news.ID = bson.NewObjectId()
	news.Creation = time.Now()

	c.Insert(news)

	return news
}

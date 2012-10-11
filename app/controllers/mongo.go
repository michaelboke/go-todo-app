package controllers

import (
	"log"
	"labix.org/v2/mgo"
	"github.com/robfig/revel"
    "labix.org/v2/mgo/bson"
    "github.com/yanatan16/go-todo-app/conf"
)

var (
	sess *mgo.Session
)

type MongoPlugin struct {
	rev.EmptyPlugin
}

func (p MongoPlugin) OnAppStart() {
	log.Println("Open MongoDB")

	sess, err := mgo.Dial(conf.MongoConnectURL)
	if err != nil {
		log.Panicln("Error opening Database", err)
	}

    db = sess.DB(conf.MongoDatabase)

    // Insert a demo user
	_, err = db.C("user").Update(
        bson.M{
            "username": "demo",
        },
        bson.M{
            "username": "demo",
            "password": "demo",
            "name": "Demo User",
        },
    )
	if err != nil {
		log.Panicln("Error upserting Database", err)
	}
}

func (p MongoPlugin) BeforeRequest(c *rev.Controller) {
    newSess := sess.Copy()
    c.Args["mongo"] = newSess
    c.Args["db"] = newSess.DB(conf.MongoDatabase)
}

func (p MongoPlugin) AfterRequest(c *rev.Controller) {
	mongo := c.Args["mongo"]
    mongo.Close()
}

func init() {
	rev.RegisterPlugin(MongoPlugin{})
}

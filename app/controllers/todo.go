package controllers

import (
    "github.com/robfig/revel"
    "github.com/yanatan16/go-todo-app/models"
    "errors"
    "log"
    "labix.org/v2/mgo"
)

type Todo {
    *revel.Controllers
}

func init() {
    revel.InterceptFunc(mustUser, revel.BEFORE, &Todo{})
}

func (t *Todo) Index() revel.Result {
    username := t.RenderArgs["user"]
    db := t.Args["db"]

    list := &models.List{}
    err := db.C("todo").Find(
        bson.M{
            "username": username,
        },
    ).One(list)

    if err != nil {
        nerr = errors.New("Error executing Mongo Find " + err.Error())
        log.Println(nerr)
        return c.RenderError(nerr)
    }

    // If its not found, make it!
    if list.items == nil {
        list.items = map[string]Item{}
    }

    t.RenderArgs["list"] = list
    t.Render()
}

func (t *Todo) JsonReadList() revel.Result {
    
}

func (t *Todo) JsonReadItem(id string) revel.Result {
    
}

func (t *Todo) JsonUpdateItem(id string) revel.Result {
    
}

func (t *Todo) JsonCreateItem() revel.Result {
    
}

func (t *Todo) JsonDeleteItem(id string) revel.Result {
    
}
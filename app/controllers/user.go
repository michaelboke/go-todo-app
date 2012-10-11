package controllers

import (
    "github.com/robfig/revel"
    "github.com/yanatan16/go-todo-app/models"
    "labix.org/v2/mgo/bson"
)

func checkUser(c *revel.Controller) revel.Result {
    if user := connected(c); user != nil {
        c.RenderArgs["user"] = user
    }
    return nil
}

func mustUser(c *revel.Controller) revel.Result {
    if user := connected(c); user == nil {
        c.RenderArgs["user"] = user
    } else {
        c.Flash.Error("Please log in first")
        return c.Redirect(Application.Index)
    }
    return nil
}

func connected(c *revel.Controller) *models.User {
    if c.RenderArgs["user"] != nil {
        return c.RenderArgs["user"].(*models.User)
    }
    
    if username, ok := c.Session["user"]; !ok {
        return nil
    }
     
    user := &models.User{Username: username}
    err := c.Args["db"].C("users").Find(
        bson.M{
            "username": username,
        },
    ).One(user)

    if err != nil {
        log.Println("Error executing Mongo Find", err)
        return nil
    }

    return user
}


package controllers

import (
	"github.com/robfig/revel"
    "github.com/yanatan16/go-todo-app/models"
	"errors"
	"log"
)

func init() {
    revel.InterceptFunc(checkUser, revel.BEFORE, &Application{})
}

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	if _, ok := c.RenderArgs["user"]; ok {
		return c.Redirect(Todo.Index)
	}
	return c.Render()
}

func (c Application) Register() revel.Result {
	return c.Render()
}

func (c Application) SaveUser(user models.User, verifyPassword string) revel.Result {
	c.Validation.Required(verifyPassword).Key("verifyPassword")
	c.Validation.Required(verifyPassword == user.Password).Key("verifyPassword").
		Message("Password does not match")
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Application.Register)
	}

	if n, err := c.Args["db"].Find(bson.M{"username":user.Username}).Count(); 
		err != nil {

		nerr = errors.New("Error executing Mongo Count " + err.Error())
		log.Println(nerr)
		return c.RenderError(nerr)
	} else if n > 0 {
		c.Flash.Error("Username already exists!")
		return c.Redirect(Application.Register)
	}

	if err := c.Args["db"].Insert(user); err != nil {
		nerr = errors.New("Error executing Mongo Insert " + err.Error())
		log.Println(nerr)
		return c.RenderError(nerr)
	}

	c.Session["user"] = user.Username
	c.Flash.Success("Welcome, " + user.Name)
	return c.Redirect(Todo.Index)
}

func (c Application) Login(username, password string) revel.Result {
	c.Validation.Required(username).Key("username")
	c.Validation.Check(password,
		revel.Required{},
		rev.MinSize{5},
	).Key("password")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Application.Index)
	}

	user := &models.User{Username:user}

    user := &models.User{Username: username}
    err := c.Args["db"].C("users").Find(
        bson.M{
            "username": username,
        },
    ).One(user)

    if err != nil {
		nerr = errors.New("Error executing Mongo Find " + err.Error())
		log.Println(nerr)
		return c.RenderError(nerr)
    }

    if user.Password != "" && user.Password == password {
		c.Session["user"] = username
		c.Flash.Success("Welcome, " + username)
		return c.Redirect(Todo.Index)
    }

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed.")
	return c.Redirect(Application.Index)
}

func (c Application) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(Application.Index)
}


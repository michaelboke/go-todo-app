package controller

import (
	"github.com/yanatan16/go-todo-app/model"
	"net/url"
	"testing"
	"time"
)

func TestControllers(t *testing.T) {
	app := model.NewTodoApp()

	// Add 2 items
	ret, err := add(app, url.Values{"desc": {"do this"}})
	if err != nil {
		t.Error("Error adding first item to list!", err)
	} else if ret["num"] != 0 {
		t.Error("Return value num not equal to 0!", ret["num"])
	}

	ret, err = add(app, url.Values{"desc": {"do that"}})
	if err != nil {
		t.Error("Error adding second item to list!", err)
	} else if ret["num"] != 1 {
		t.Error("Return value num not equal to 1!", ret["num"])
	}

	i0 := app.List.Items[0]
	i1 := app.List.Items[1]

	if i0.Num != 0 {
		t.Error("First item number is not 0:", i0.Num)
	} else if i0.Desc != "do this" {
		t.Error("First item description is not correct:", i0.Desc)
	} else if i0.Done != false {
		t.Error("First item done is true!")
	}

	if i1.Num != 1 {
		t.Error("Second item number is not 1:", i1.Num)
	} else if i1.Desc != "do that" {
		t.Error("Second item description is not correct:", i1.Desc)
	} else if i1.Done != false {
		t.Error("Second item done is true!")
	}

	done(app, url.Values{"num": {"0"}, "done": {"true"}})

	// Wait for update
	<-time.After(10 * time.Millisecond)

	i0p := app.List.Items[0]
	if i0p.Done != true {
		t.Error("First item did not update done to true!")
	}

	// Check for a bad set
	done(app, url.Values{"num": {"1231321"}, "done": {"false"}})
	<-time.After(10 * time.Millisecond)
	if app.List.Items[1].Num != 1 {
		t.Error("Second item number changed with bad Set command:", app.List.Items[1].Num)
	}

}

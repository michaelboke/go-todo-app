package controller

import (
	"github.com/yanatan16/go-todo-app/model"
	"testing"
	"time"
)

func TestControllers(t *testing.T) {
	app := model.NewTodoApp(false)

	// Add 2 items
	ret, err := add(app, map[string]string{"desc": "do this"})
	if err != nil {
		t.Error("Error adding first item to list!", err)
	} else {
		if data, ok := ret.(map[string]interface{}); !ok {
			t.Error("Return value not correct type!", ret)
		} else if data["num"] != 0 {
			t.Error("Return value num not equal to 0!", data["num"])
		}
	}

	ret, err = add(app, map[string]string{"desc": "do that"})
	if err != nil {
		t.Error("Error adding second item to list!", err)
	} else {
		if data, ok := ret.(map[string]interface{}); !ok {
			t.Error("Return value not correct type!", ret)
		} else if data["num"] != 1 {
			t.Error("Return value num not equal to 1!", data["num"])
		}
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

	_, err = done(app, map[string]string{"num": "0", "done": "true"})
	if err != nil {
		t.Error("Error executing first done.", err)
	}

	// Wait for update
	<-time.After(10 * time.Millisecond)

	i0p := app.List.Items[0]
	if i0p.Done != true {
		t.Error("First item did not update done to true!")
	}

	// Check for a bad set
	_, err = done(app, map[string]string{"num": "1231321", "done": "false"})
	if err == nil {
		t.Error("No error when expected one for second done.")
	}

	<-time.After(10 * time.Millisecond)
	if app.List.Items[1].Num != 1 {
		t.Error("Second item number changed with bad Set command:", app.List.Items[1].Num)
	}

}

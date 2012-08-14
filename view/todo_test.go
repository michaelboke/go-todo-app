package view

import (
	"bytes"
	"github.com/yanatan16/go-todo-app/model"
	"github.com/yanatan16/go-todo-app/model/todo"
	"html/template"
	"os"
	"strings"
	"testing"
)

func TestApp(t *testing.T) {
	app := model.NewTodoApp(false)

	app.List.Items = []todo.Item{
		todo.Item{0, "do this", true},
		todo.Item{1, "do that", false},
	}

	// Test App is formatted as expected
	tmplt := `{{with .List}}Yo {{range .Items}}#{{.Num}}: {{.Desc}} / {{if .Done}}check!{{else}}no-check :({{end}} {{end}}{{end}}`

	expected := `Yo #0: do this / check! #1: do that / no-check :( `

	tmp, err := template.New("test").Parse(tmplt)
	if err != nil {
		t.Fatal("Error parsing test template:", err)
	}

	buf := bytes.NewBuffer(make([]byte, 1024))
	err = tmp.Execute(buf, app)
	if err != nil {
		t.Error("Error on execution!", err)
	}

	out := strings.TrimLeft(buf.String(), "\x00")
	if out != expected {
		t.Errorf("Output was not as expected! (tmplt: %q) != (expct: %q)", out, expected)
	}
}

func TestTodoTemplate(t *testing.T) {
	app := model.NewTodoApp(false)

	app.List.Items = []todo.Item{
		todo.Item{0, "do this", true},
		todo.Item{1, "do that", false},
	}

	// Test App is formatted as expected
	tmp, err := template.ParseFiles(templateRoot + "/todo.html")
	if err != nil {
		t.Fatal("Error parsing test template:", err)
	}

	err = tmp.Execute(os.Stdout, app)
	if err != nil {
		t.Error("Error on execution!", err)
	}
}

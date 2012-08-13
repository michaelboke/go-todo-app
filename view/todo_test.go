package view

import (
	"bytes"
	"github.com/yanatan16/go-todo-app/model"
	"github.com/yanatan16/go-todo-app/model/todo"
	"html/template"
	"strings"
	"testing"
)

func TestTodo(t *testing.T) {
	app := model.NewTodoApp()

	app.List.Items = []todo.Item{
		todo.Item{0, "do this", true},
		todo.Item{1, "do that", false},
	}

	tmplt := `{{with .List}}Yo {{range .Items}}#{{.Num}}: {{.Desc}} / {{if .Done}}check!{{else}}no-check :({{end}} {{end}}{{end}}`

	expected := `Yo #0: do this / check! #1: do that / no-check :( `

	tmp, err := template.New("test").Parse(tmplt)
	if err != nil {
		t.Fatal("Error parsing test template:", err)
	}

	buf := bytes.NewBuffer(make([]byte, 1024))
	tmp.Execute(buf, app)

	out := strings.TrimLeft(buf.String(), "\x00")
	if out != expected {
		t.Errorf("Output was not as expected! (tmplt: %q) != (expct: %q)", out, expected)
	}

}

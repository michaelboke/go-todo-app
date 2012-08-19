package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
	"io"
	"net/http"
	"testing"
	"time"
)

const (
	port int    = 16123
	path string = "/test"
)

type ExampleController map[string]string

func (e ExampleController) Create(attr string) (interface{}, error) {
	return e, nil
}
func (e ExampleController) Read(id string) (interface{}, error) {
	return e, nil
}
func (e ExampleController) Update(id, attr string) (interface{}, error) {
	return e, nil
}
func (e ExampleController) Delete(id string) error {
	return nil
}

func ExampleInitController(svr *web.Server, path string) {
	m := map[string]string{"hello": "world"}
	ctrl := ExampleController(m)

	BindController(svr, path, ctrl)
}

func init() {
	svr := web.NewServer()
	ExampleInitController(svr, path)
	go svr.Run(fmt.Sprintf("0.0.0.0:%d", port))
	<-time.After(10 * time.Millisecond)
}

func checkRespBody(t *testing.T, r io.Reader) {
	m := map[string]string{}
	j := json.NewDecoder(r)
	err := j.Decode(&m)
	if err != nil {
		t.Error("Response Body was not decodable into json", err)
	} else if m["hello"] != "world" {
		t.Error("Response Body did not contain correct map", m)
	}
}

func TestCreate(t *testing.T) {
	url := fmt.Sprintf("http://127.0.0.1:%d%s", port, path)
	client := &http.Client{}

	cbody := bytes.NewBuffer([]byte(`{}`))
	creq, err := http.NewRequest("POST", url, cbody)
	if err != nil {
		t.Fatal("Error creating the create request", err)
	}
	cres, err := client.Do(creq)
	if err != nil {
		t.Fatal("Error performing create request", err)
	}
	defer cres.Body.Close()
	checkRespBody(t, cres.Body)
}

func TestRead(t *testing.T) {
	url := fmt.Sprintf("http://127.0.0.1:%d%s/id", port, path)
	client := &http.Client{}

	creq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal("Error creating the create request", err)
	}
	cres, err := client.Do(creq)
	if err != nil {
		t.Fatal("Error performing create request", err)
	}
	defer cres.Body.Close()
	checkRespBody(t, cres.Body)
}

func TestUpdate(t *testing.T) {
	url := fmt.Sprintf("http://127.0.0.1:%d%s/id", port, path)
	client := &http.Client{}

	creq, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		t.Fatal("Error creating the create request", err)
	}
	cres, err := client.Do(creq)
	if err != nil {
		t.Fatal("Error performing create request", err)
	}
	defer cres.Body.Close()
	checkRespBody(t, cres.Body)
}

func TestDelete(t *testing.T) {
	url := fmt.Sprintf("http://127.0.0.1:%d%s/id", port, path)
	client := &http.Client{}

	creq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal("Error creating the create request", err)
	}
	cres, err := client.Do(creq)
	if err != nil {
		t.Fatal("Error performing create request", err)
	}
	defer cres.Body.Close()
}

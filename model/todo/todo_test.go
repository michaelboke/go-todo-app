package todo

import (
	"encoding/json"
	"testing"
	"reflect"
)

func checkErrors(t *testing.T, i interface{}, err error) Item {
	if err != nil {
		t.Fatal("Error creating item:", i, err)
	}

	it, ok := i.(Item)
	if !ok {
		t.Fatal("Error casting item:", i)
	}

	return it
}

func TestJsonEncoding(t *testing.T) {
	it := Item{}
	_, err := json.Marshal(it)
	if err != nil {
		t.Error("Item is not json serializable!")
	}
}

func TestList(t *testing.T) {
	l := NewList()

	// Add 2 items
	iti, err := l.Create(`{"desc":"do this"}`)
	it1 := checkErrors(t, iti, err)
	if it1.Desc != "do this" {
		t.Error("First item does not have correct description.")
	} else if it1.Done {
		t.Error("First item is set to done!")
	}

	iti, err = l.Create(`{"desc":"do that"}`)
	it2 := checkErrors(t, iti, err)
	if it2.Desc != "do that" {
		t.Error("Second item does not have correct description.")
	} else if it2.Done {
		t.Error("Second item is set to done!")
	}

	if it1.Id == it2.Id {
		t.Error("Items 1 and 2 have the same id!")
	}

	iri1, err := l.Read(it1.Id)
	ir1 := checkErrors(t, iri1, err)
	if !reflect.DeepEqual(ir1, it1) {
		t.Error("Read value for item 1 is not right:", it1, ir1)
	}

	iui1, err := l.Update(it1.Id, `{"done":true}`)
	checkErrors(t, iui1, err)

	err = l.Delete(it2.Id)
	if err != nil {
		t.Error("Error deleting item 2", err)
	}

	_, err = l.Read(it2.Id)
	if err == nil {
		t.Error("Should of had an error reading item 2 after delete!")
	}
}

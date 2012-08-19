package todo

import (
	"encoding/json"
	"errors"
	"fmt"
)

// A Todo Item with a description, number and done boolean
type Item struct {
	// The index in the Todo List
	Id string
	// A description of the item
	Desc string
	// Whether the item is done or not
	Done bool
}

type ItemCreation struct {
	Desc string
}

// A Todo List contain a number of items
// It is thread-safe, because all changes use channels
type List struct {
	exec    chan func()
	items   map[string]Item
	counter int
}

// Create a new Todo List with no items.
// Returns an empty list.
func NewList() *List {
	l := &List{
		exec:    make(chan func(), 10),
		items:   make(map[string]Item),
		counter: 0,
	}

	// Start listening for updates
	go l.listen()

	return l
}

// Listen for updates on the channels and update the list
// when necessary
func (l *List) listen() {
	for f := range l.exec {
		f()
	}
}

// Add a new todo item with a description
// Arguments: 
//	desc - Description of the new todo item
func (l *List) createItem(desc string) Item {

	its := make(chan Item)
	l.exec <- func() {
		id := l.counter
		l.counter += 1

		it := Item{
			Id:   fmt.Sprintf("%d", id),
			Desc: desc,
			Done: false,
		}

		l.items[it.Id] = it
		its <- it
	}

	return <-its
}

func (l *List) getItem(id string) (it Item, err error) {
	its := make(chan Item)
	errs := make(chan error)

	l.exec <- func() {
		it, ok := l.items[id]
		if !ok {
			errs <- errors.New("id not valid")
		} else {
			its <- it
		}
	}

	select {
	case it = <-its:
		err = nil
	case err = <-errs:
		it = Item{}
	}
	return it, err
}

func (l *List) getAll() map[string]Item {
	its := make(chan map[string]Item)
	l.exec <- func() {
		its <- l.items
	}
	return <- its
}

func (l *List) updateItem(id string, it Item) (Item, error) {
	suc := make(chan bool)
	errs := make(chan error)

	l.exec <- func() {
		_, ok := l.items[id]
		if !ok {
			errs <- errors.New("id not valid")
		} else {
			l.items[id] = it
			suc <- true
		}
	}

	var err error
	select {
	case <-suc:
	case err = <-errs:
	}

	return it, err
}

func (l *List) deleteItem(id string) {
	l.exec <- func() {
		delete(l.items, id)
	}
}

// Create an item
// attr is a json-formatted string of attributes
// Return a json-formattable object of all model attributes
func (l *List) Create(attr string) (interface{}, error) {
	itreq := ItemCreation{}
	err := json.Unmarshal([]byte(attr), &itreq)
	if err != nil {
		return nil, err
	}

	desc := itreq.Desc
	if desc == "" {
		return nil, errors.New("Create request requires 'desc' attribute.")
	}

	it := l.createItem(desc)

	return it, nil
}

// Read an item
// ID may be empty string
// Return a json-formattable object of all model attributes
func (l *List) Read(id string) (interface{}, error) {
	if id == "" {
		return l.getAll(), nil
	}
	
	return l.getItem(id)
}

// Update a model object based on parameters. 
// ID is required and will be non-empty
// attr is a json-formatted string of attributes
// Return a json-formattable object of updated model attributes
// If no attributes other than the updated ones changed, it is acceptable to return nil
func (l *List) Update(id string, attr string) (interface{}, error) {
	it := Item{}
	err := json.Unmarshal([]byte(attr), &it)
	if err != nil {
		return nil, err
	}

	return l.updateItem(id, it)
}

// Delete a model object.
// ID is required and will be non-empty
func (l *List) Delete(id string) error {
	l.deleteItem(id)
	return nil
}

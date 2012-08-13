package todo

import (
	"log"
)

// A Todo Item with a description, number and done boolean
type Item struct {
	// The index in the Todo List
	Num int
	// A description of the item
	Desc string
	// Whether the item is done or not
	Done bool
}

// An ItemRequest is how one adds items to a list.
type ItemRequest struct {
	// Description of the new item.
	Desc string
	// Once the request is processed, the corresponding num will
	// be sent back over the channel Num (if not nil)
	// Note that the listener will not wait for syncing on this channel.
	// It should be bufferred if the receiver is not waiting for it.
	Num chan<- int
}

// A Todo List contain a number of items
// It is thread-safe, because all changes use channels
type List struct {
	// To add a new item, send the items description over the Add channel
	Add chan ItemRequest
	// To set a previously known item (with changes), send it over the Set channel
	Set chan Item

	// Do not modify this list directly. That would violate thread-safety.
	// Use this for read access only.
	Items []Item
}

// Create a new Todo List with no items.
// Returns an empty list.
func NewList() *List {
	l := &List{
		Items: make([]Item, 0),
		Add:   make(chan ItemRequest, 5),
		Set:   make(chan Item, 5),
	}

	// Start listening for updates
	go l.listen()

	return l
}

// Listen for updates on the channels and update the list
// when necessary
func (l *List) listen() {
	for {
		select {
		case req := <-l.Add:
			i := l.addItem(req.Desc)
			log.Printf("Added item: num %d, desc: %s", i, req.Desc)
			if req.Num != nil {
				// Perform a send over Num or break
				select {
				case req.Num <- i:
				default:
					log.Printf("Warning: Return Number channel was not ready to recieve. Not sending...")
				}
			}

		case it := <-l.Set:
			if it.Num < len(l.Items) {
				l.Items[it.Num] = it
				log.Printf("Set Item at num %d with done: %t", it.Num, it.Done)
			} else {
				log.Printf("Warning: Set Item received out of bouds num: %d", it.Num)
			}
		}
	}
}

// Add a new todo item with a description
// Arguments: 
//	desc - Description of the new todo item
func (l *List) addItem(desc string) int {
	i := Item{
		Num:  len(l.Items),
		Desc: desc,
		Done: false,
	}
	l.Items = append(l.Items, i)

	return i.Num
}

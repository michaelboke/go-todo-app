package todo

import (
	"testing"
	"time"
)

func TestList(t *testing.T) {
	l := NewList()
	key := make(chan int)

	// Add 2 items
	l.Add <- ItemRequest{"do this", nil}
	l.Add <- ItemRequest{"do that", key}

	select {
	case n := <-key:
		if n != 1 {
			t.Errorf("Adding item did not return n == 1: %d", n)
		}
	case <-time.After(time.Second / 2):
		t.Error("Adding item did not respond with key after timeout!")
	}

	i0 := l.Items[0]
	i1 := l.Items[1]

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

	i0.Done = true
	l.Set <- i0

	// Wait for update
	<-time.After(10 * time.Millisecond)

	i0p := l.Items[0]
	if i0p.Done != true {
		t.Error("First item did not update done to true!")
	}

	// Check for a bad set
	i1.Num = 1231321
	l.Set <- i1
	<-time.After(10 * time.Millisecond)
	if l.Items[1].Num != 1 {
		t.Error("Second item number changed with bad Set command:", l.Items[1].Num)
	}

	// Add an item with a channel return but don't listen on it
	l.Add <- ItemRequest{"now what", key}
	<-time.After(10 * time.Millisecond)
	if len(l.Items) != 3 {
		t.Error("Item was not added after Item request with no waiting receiver for key chan")
	}
}

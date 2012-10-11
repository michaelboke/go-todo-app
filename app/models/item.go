package models

import (
    "github.com/robfig/revel"
    "fmt"
)


// A Todo Item with a description, number and done boolean
type Item struct {
    // The index in the Todo List
    Id string `json:"id",bson:"id"`
    // A description of the item
    Desc string `json:"desc",bson:"desc"`
    // Whether the item is done or not
    Done bool `json:"done",bson:"done"`
}

func (i *Item) String() string {
    return fmt.Sprintf("Item:{Id:%s,Desc:%s,Done:%v", i.Id, i.Desc, i.Done)
}

func (i *Item) Validate(v *revel.Validation) {
    v.Check(i.Id,
        revel.Required{},
        revel.MaxSize(64),
        revel.MinSize{32},
    ).Key("item.Id")

    v.Check(i.Desc,
        revel.Required{},
        revel.MinSize(1),
    ).Key("item.Desc")
}
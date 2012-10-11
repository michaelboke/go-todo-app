package models

import (
	"github.com/robfig/revel"
	"strings"
)

type List struct {
	Items map[string] Item `json:"items",bson:"items"`
}

func (l *List) String() string {
	strs := make([]string,0,len(l))
	for _, item := range l {
		strs = append(strs, item.String())
	}
	return strings.Join(strs, ",")
}

func (l *List) Validate(v *revel.Validation) {
	if !l {
		l = make(map[string]Item)
	}
}

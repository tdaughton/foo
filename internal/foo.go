package server

import (
	"github.com/google/uuid"
	"sync"
)

type Foo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type FooManager struct {
	mu          sync.Mutex
	foosById    map[string]*Foo
	initialized bool
}

func (c *FooManager) init() {
	c.foosById = make(map[string]*Foo)
	c.initialized = true
}

func (c *FooManager) insert(name string) Foo {
	c.mu.Lock()
	if c.initialized == false {
		c.init()
	}
	defer c.mu.Unlock()
	newId := uuid.New().String()
	foo := Foo{newId, name}
	c.foosById[foo.Id] = &foo
	return foo
}

func (c *FooManager) retrieve(id string) (Foo, bool) {
	c.mu.Lock()
	if c.initialized == false {
		c.init()
	}
	defer c.mu.Unlock()
	foo := c.foosById[id]
	if foo == nil {
		return Foo{}, false
	}
	return *foo, true
}

func (c *FooManager) delete(id string) bool {
	c.mu.Lock()
	if c.initialized == false {
		c.init()
	}
	defer c.mu.Unlock()
	foo := c.foosById[id]
	if foo == nil {
		return false
	} else {
		delete(c.foosById, id)
		return true
	}
}

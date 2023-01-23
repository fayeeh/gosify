package main

import "time"

type Item struct {
	Name   string
	Value  string
	Expire time.Time
}

type Cache struct {
	Items []*Item
}

func (c *Cache) Add(name string, value string, expire time.Time) {
	item := &Item{
		Name:   name,
		Value:  value,
		Expire: expire,
	}

	idx := indexOf(c.Items, func(elem *Item) bool {
		return elem.Name == name
	})

	if idx == -1 {
		c.Items = append(c.Items, item)
	} else {
		c.Items[idx] = item
	}
}

func (c *Cache) Get(name string) *Item {
	for _, i := range c.Items {
		if i.Name == name {
			return i
		}
	}
	return nil
}

func (c *Cache) Remove(name string) {
	idx := indexOf(c.Items, func(elem *Item) bool {
		return elem.Name == name
	})

	if idx > -1 {
		c.Items = append(c.Items[:idx], c.Items[idx+1:]...)
	}
}

func (c *Cache) Start() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for range ticker.C {
			c.deleteExpired()
		}
	}()
}

func (c *Cache) deleteExpired() {
	for _, elem := range c.Items {
		if time.Now().After(elem.Expire) && !elem.Expire.IsZero() {
			c.Remove(elem.Name)
		}
	}
}

func indexOf[T comparable](arr []T, callback func(T) bool) int {
	for idx, elem := range arr {
		if callback(elem) {
			return idx
		}
	}
	return -1
}

package models

type Item struct {
	Name  string
	Count int64
	Price int64
}

func (c *Item) IsAvailable() bool {
	return c.Count > 0
}

func (c *Item) Add() error {
	c.Count++
	return nil
}

func (c *Item) Remove() {
	c.Count--
}

func (c *Item) GetName() string {
	return c.Name
}

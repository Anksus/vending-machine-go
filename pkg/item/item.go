package item

import "errors"

type Item struct {
	Name  string
	Count int
	Price int
}

func (c *Item) IsAvailable() bool {
	return c.Count > 0
}

func (c *Item) Add() error {
	c.Count++
	return nil
}

func (c *Item) Remove() error {
	if c.Count == 0 {
		return errors.New("item is not available")
	}
	c.Count--
	return nil
}

func (c *Item) GetName() string {
	return c.Name
}

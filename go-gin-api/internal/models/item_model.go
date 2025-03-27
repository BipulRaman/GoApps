package models

import "fmt"

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (e *Item) Validate() error {
	if e.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}

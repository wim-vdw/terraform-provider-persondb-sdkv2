package myprovider

import "errors"

type Client struct {
	CustomDatabase string
	Persons        map[string]string
}

func (c *Client) AddPerson(nameID string, lastName, firstName string) error {
	if _, exists := c.Persons[nameID]; exists {
		return errors.New("a person with the same name_id already exists")
	}
	c.Persons[nameID] = lastName + " " + firstName
	return nil
}

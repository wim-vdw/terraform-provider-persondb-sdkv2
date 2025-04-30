package myprovider

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type Person struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

type Client struct {
	CustomDatabase string
	Persons        map[string]Person
}

var databaseLock = sync.Mutex{}

func (c *Client) AddPerson(nameID string, lastName, firstName string) error {
	if _, exists := c.Persons[nameID]; exists {
		return errors.New("a person with the same name_id already exists")
	}
	databaseLock.Lock()
	defer databaseLock.Unlock()
	c.Persons[nameID] = Person{
		LastName:  lastName,
		FirstName: firstName,
	}
	return nil
}

func (c *Client) GetPerson(nameID string) (Person, error) {
	if person, exists := c.Persons[nameID]; exists {
		return person, nil
	}
	return Person{}, errors.New("person not found")
}

func (c *Client) LoadDB() error {
	databaseLock.Lock()
	defer databaseLock.Unlock()
	data, err := os.ReadFile(c.CustomDatabase)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c.Persons)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SaveDB() error {
	databaseLock.Lock()
	defer databaseLock.Unlock()
	data, err := json.MarshalIndent(c.Persons, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(c.CustomDatabase, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

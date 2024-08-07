package user

import (
	"encoding/json"
	"fmt"
)

type Client struct {
	ID         uint64
	FirstName  string
	SecondName string
}

func (c *Client) String() string {
	encoded, err := json.Marshal(c)
	if err != nil {
		encoded = []byte("")
	}

	return fmt.Sprintf("Client: %s", encoded)
}

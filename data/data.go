package data

import (
	"fmt"
	"encoding/json"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err;
	}

	fmt.Println(string(b));

	return nil;
}
package data

import (
	"io/ioutil"
	"encoding/json"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items);
	if err != nil {
		return err;
	}

	err = ioutil.WriteFile(filename, b, 0644);
	if err != nil {
		return err;
	}


	return nil;
}
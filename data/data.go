package data

import (
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type Item struct {
	Text string
	Priority int
	position int
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

func ReadItems(filename string)([]Item, error) {
	b, err := ioutil.ReadFile(filename);
	if err != nil {
		return []Item{}, err;
	}

	var items []Item;
	err = json.Unmarshal(b, &items);
	if err != nil {
		return []Item{}, err;
	}

	for i, _ := range items {
		items[i].position = i + 1;
	}

	return items, nil; 
}

func (i *Item) SetPriority(pri int) {
	switch pri {
		case 1: 
			i.Priority = 1;
		case 3:
			i.Priority = 3;
		default:
			i.Priority = 2;
	}
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)";
	}
	if i.Priority == 3 {
		return "(3)";
	}

	return " ";
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + ".";
}
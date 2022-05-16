package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Data struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Zip     string `json:"zip"`
	ID      string `json:"id"`
}

type info struct {
	name string
	address string
	zip string
}

func (d *Data) Validate() bool {
	return d.Name != "" && d.Address != "" && d.zipCheck()
}

// I'm from Canada, we don't use Zip Codes; from a quick search these seem to be the bounds however there may be more to it
func (d *Data) zipCheck() bool {
	if len(d.Zip) == 5 {
		zip, err := strconv.Atoi(d.Zip)
		if err != nil {
			return false
		}
		if zip >= 1 && zip <= 99950 {
			return true
		}
	}
	return false
}

func main() {
	jsonFile, err := os.Open("../data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data []Data
	json.Unmarshal(byteValue, &data)
	invalidIDs := validate(data)
	for _, id := range invalidIDs {
		fmt.Println(id)
	}
}

func validate(data []Data) []string {
	var invalidIDs []string
	dupes := map[info][]string{}
	for _, d := range data {
		if !d.Validate() {
			invalidIDs = append(invalidIDs, d.ID)
		} else {
			dupes[info{d.Name, d.Address, d.Zip}] = append(dupes[info{d.Name, d.Address, d.Zip}], d.ID)
		}
	}
	for _, v := range dupes {
		if len(v) > 1 {
			invalidIDs = append(invalidIDs, v...)
		}
	}
	return invalidIDs
}
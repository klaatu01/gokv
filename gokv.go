package gokv

import (
	json "encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"log"
)

var jsonpath string
var keyvals []KeyVal

type KeyVal struct {
	Key       string    `json:"Key"`
	Value     string    `json:"Value"`
	Timestamp time.Time `json:"Time"`
}

func SetPath(path string) {
	jsonpath = path
}

func loadKeyVals(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		_, err = file.WriteString("[]\n")
		defer file.Close()
	}
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error Reading File")
		os.Exit(0)
	}
	json.Unmarshal(raw, &keyvals)
}

func removeIndex(slice []KeyVal, i int) []KeyVal {
	slice[len(slice)-1], slice[i] = slice[i], slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func GetAll() []KeyVal {
	loadKeyVals(jsonpath)
	return keyvals
}

func Remove(key string) {
	loadKeyVals(jsonpath)
	index := -1
	for i, keyval := range keyvals {
		if keyval.Key == key {
			index = i
			break
		}
	}
	keyvals = removeIndex(keyvals, index)
	Commit(jsonpath)
}

func Get(key string) string {
	loadKeyVals(jsonpath)
	for _, keyval := range keyvals {
		if keyval.Key == key {
			return keyval.Value
		}
	}
	return ""
}

func SaveAs(path string) {
	data, err := json.Marshal(keyvals)
	err = ioutil.WriteFile(path, data, os.FileMode(0777))
	if err != nil {
		fmt.Println("Couldn't save")
	}
}

func Commit(path string) {
	data, err := json.Marshal(keyvals)
	err = ioutil.WriteFile(path, data, os.FileMode(0777))
	if err != nil {
		fmt.Println("Couldn't save")
	}
}

func Set(key string, val string) {
	loadKeyVals(jsonpath)
	kv := KeyVal{Key: key, Value: val, Timestamp: time.Now()}
	var placed bool = false
	for i, keyval := range keyvals {
		if keyval.Key == kv.Key {
			keyvals[i].Key = kv.Key
			keyvals[i].Value = kv.Value
			keyvals[i].Timestamp = kv.Timestamp
			placed = true
			break
		}
	}
	if placed != true {
		keyvals = append(keyvals, kv)
	}
	Commit(jsonpath)
}

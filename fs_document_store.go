package main

import (
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"code.google.com/p/go-uuid/uuid"
	"path/filepath"
	//"os"
)

type FSDocumentStore struct {
	Root string
}

func (this *FSDocumentStore) Insert(doc interface{}) (interface{}, error) {
	jsonbytes, err := json.Marshal(doc)

	if err != nil {
		return nil, err
	}

	id := uuid.NewRandom().String()

	//fmt.Println(id)
	//fmt.Println(string(jsonbytes))

	ioutil.WriteFile(filepath.Join(this.Root,  id + ".json"), jsonbytes, 0666)

	out := map[string]interface{}{}
	err = json.Unmarshal(jsonbytes, &out)

	out["_id"] = id

	return out, err
}
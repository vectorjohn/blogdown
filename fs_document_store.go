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

func (this *FSDocumentStore) Insert(doc interface{}) (Document, error) {
	jsonbytes, err := json.Marshal(doc)

	if err != nil {
		return nil, err
	}

	id := uuid.NewRandom().String()

	//fmt.Println(id)
	//fmt.Println(string(jsonbytes))

	ioutil.WriteFile(filepath.Join(this.Root,  id + ".json"), jsonbytes, 0666)

	out := Document{}
	err = json.Unmarshal(jsonbytes, &out)
	if err != nil {
		return nil, err
	}
	out["_id"] = id

	return out, nil
}

func (this *FSDocumentStore) Find(id string) (Document, error) {
	jsonbytes, err := ioutil.ReadFile(filepath.Join(this.Root, id + ".json"))

	if err != nil {
		return nil, err
	}

	out := Document{}
	err = json.Unmarshal(jsonbytes, &out)
	if err != nil {
		return nil, err
	}
	out["_id"] = id

	return out, nil
}
package main

import (
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"code.google.com/p/go-uuid/uuid"
	"path/filepath"
	"os"
	"strings"
)

type FSDocumentStore struct {
	Root string
}

func (this *FSDocumentStore) FindAll() (*Collection, error) {
	dir, err := os.Open(this.Root)
	if err != nil {
		return nil, err
	}


	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err;
	}

	var docs Collection = make(Collection, 0, 100)

	for _, file := range files {
		fnparts := strings.Split(file.Name(), ".")
		if len(fnparts) < 2 || fnparts[1] != "json" {
			continue
		}

		id := fnparts[0]
		doc, err := this.FindId(id)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	return &docs, nil
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

func (this *FSDocumentStore) FindId(id string) (Document, error) {
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


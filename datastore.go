package main

type Document map[string]interface{}

type DocumentStore interface {
	Insert(doc interface{}) (Document, error)
	Update(id string, doc interface{}) error
	Find(id string) (Document, error)
}


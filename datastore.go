package main

type DocumentStore interface {
	Insert(interface{}) (interface{}, error)
	Update(id string, update interface{}) error
}

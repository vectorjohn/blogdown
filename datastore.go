package main

type Document map[string]interface{}

type Collection []Document

type DocumentStore interface {
	FindAll() (*Collection, error)
	Insert(doc interface{}) (Document, error)
	Update(id string, doc interface{}) error
	FindId(id string) (Document, error)
}


func (this *Collection) Filter(filter func(Document, int) bool) *Collection {
	var results Collection = make(Collection, 0, 100)

	for i, v := range *this {
		if filter(v, i) {
			results = append(results, v)
		}
	}

	return &results
}
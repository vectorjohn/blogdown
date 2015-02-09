package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type ServerConfig struct {
	PagePath string `json:"page_path"`
	WebRoot  string `json:"web_root"`
	Index    string `json:"index"`
	Port     int    `json:"port"`
}

type TestData struct {
	Foo string
	Bar int
	baz string
}

var config *ServerConfig = &ServerConfig{"pages", "www", "index", 4003}

func main() {
	/*
		data, err := loadFile(config.Index)
		if err != nil {
			fmt.Println( err.Error() )
			return
		}

		output := blackfriday.MarkdownCommon(data)
		fmt.Println(string(output))
	*/

	if 1==1 {
		docstore := &FSDocumentStore{Root: "data"}
		/*
		var out Document
		out, _ = docstore.Insert(&TestData{"Hello world", 1234, "bla bla bla"})
		out, _ = docstore.Insert(&Document{"foo":"bar", "baz": 11})
		out, _ = docstore.Find("19291aee-3208-470a-9686-3bcc2386ec91")
		fmt.Println(out)
		*/
		all, err := docstore.FindAll()
		fmt.Println("ALL: ", all)
		all = all.Filter(func(doc Document, i int) bool {
			return doc["foo"] == "bar"
		})
		fmt.Println(err)
		fmt.Println("SOME: ", all)
		return
	}

	http.HandleFunc("/admin", configHandler(config, admin_handler))

	indexServer := pageServer("index")
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			indexServer(resp, req)
			return
		}

		pageServer(req.URL.Path)(resp, req)
	})

	listenOn := "localhost:" + strconv.Itoa(config.Port)
	fmt.Println("Listening on: " + listenOn)
	http.ListenAndServe(listenOn, nil)
}

func configHandler(conf ServerConfig, handler func(resp http.ResponseWriter, req *http.Request)) {
	handler(conf, resp, req)
}

func pageServer(page string) func(http.ResponseWriter, *http.Request) {
	data, err := loadFile(page)
	if err != nil {
		if page == "404" {
			return internalServerError
		}

		return pageServer("404")
	}

	data = blackfriday.MarkdownCommon(data)
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Write(data)
	}
}

func internalServerError(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("The server had a big problem."))
}

func loadFile(name string) ([]byte, error) {
	file, err := os.Open(config.PagePath + "/" + name + ".md")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

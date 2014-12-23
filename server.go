package main

import (
	"github.com/russross/blackfriday"
	"fmt"
	//"net/http"
	"os"
	"io/ioutil"
)

type ServerConfig struct {
	PagePath	string	`json:"page_path"`
	WebRoot	string	`json:"web_root"`
	Index	string	`json:"index"`
	Port	int	`json:"port"`
}

var config *ServerConfig = &ServerConfig{"pages", "www", "index", 4003}

func main() {
	data, err := loadFile(config.Index)
	if err != nil {
		fmt.Println( err.Error() )
		return
	}

	output := blackfriday.MarkdownCommon(data)
	fmt.Println(string(output))
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
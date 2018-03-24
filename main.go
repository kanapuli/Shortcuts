package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	toml "github.com/pelletier/go-toml"
	browser "github.com/skratchdot/open-golang/open"
)

func init() {
	_, err := os.Stat("/home/user/.z/config.toml")
	if os.IsNotExist(err) {
		//Config file does not exists
		os.Mkdir("/home/user/.z", 0777)
		os.Create("/home/user/.z/config.toml")
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		//the binary name is arg0. hence min arg length should be 1
		fmt.Println("What should I search?")
		os.Exit(1)
	}
	//read the config toml
	content, err := ioutil.ReadFile("/home/user/.z/config.toml")
	if err != nil {
		log.Println("No config file: ", err)
	}
	config, _ := toml.Load(string(content))
	q := fmt.Sprintf("urls.%v", args[0])
	urls := config.Get(q).([]interface{})
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(urlString interface{}) {
			defer wg.Done()
			browser.Run(urlString.(string))
		}(url)
	}
	wg.Wait()
}

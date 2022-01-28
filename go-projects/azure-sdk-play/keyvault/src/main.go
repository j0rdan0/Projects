package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	file, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		log.Panic(err)
	}
	_ = json.Unmarshal([]byte(file), &config)

}

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		getSecret("az-token")
	} else {
		// can also query different secrets, but main purpose was to fetch auth bearer for Azure REST API
		getSecret(args[0])
		// NOTE: doesn`t check for existance of the key prior to searching
	}

}

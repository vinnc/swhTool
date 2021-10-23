/** ConfigReader for ShopWare**/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"strconv"
)

const envFileName = ".env"

func ReadConfig() map[string]string {
	config := make(map[string]string)
	err := godotenv.Load(envFileName)
	if err != nil {
		fatal("Error loading %s file", envFileName)
	}

	config["mode"] = os.Getenv("APP_ENV")
	config["database"] = os.Getenv("DATABASE_URL")
	readComposerData()
	return config
}

func readComposerData() {
	jsonFile, err := os.Open("composer.lock")

	if err != nil {
		fatal(err.Error())
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)


	_packages := result["packages"]

	_packages.
	for i := 0; i < len(_packages); i++ {
		fmt.Println("sadf")
	}
	/*for _, _package := range _packages {
		fmt.Println(_package)
		os.Exit(1)
	}*/

	defer jsonFile.Close()
}

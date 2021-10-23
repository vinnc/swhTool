package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var programName = ""

func getAvailableCommands() []string {
	return []string {"list", "info", "dump"}
}

func main() {
	if len(os.Args) <= 2 {
		//showAllAvailableCommands()
		//os.Exit(1)
	}

	//_command := os.Args[1]
	_command := "info"

	_config  := ReadConfig()
	for _type,  _configItem := range _config {
		debug(_configItem)
		debug(_type)
	}

	os.Exit(1)

	switch _command {
	case "list": showAllAvailableCommands()
	case "info": showInfo()

	default:
		fmt.Println("please enter a command")
	}



	os.Exit(1)

}

func init() {
	var err error
	programName, err := os.Executable()
	//programName, err = os.Executable()
	if  err != nil {
		programName = os.Args[0]
	}
	programName = filepath.Base(programName)
	fmt.Println(programName)
	flag.Parse()
}

func showAllAvailableCommands() {
	fmt.Println("Available commands:")
	for _ , _command := range getAvailableCommands() {
		fmt.Println(_command)
	}
}

func showInfo()  {
	//_dirs := readCurrentDirectories()
	/*for _, _dir := range _dirs {

	}*/
	fmt.Printf("Shopware %s", "6.4")
}

func stderr(f string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, programName+": "+fmt.Sprintf(f, args...)+"\n")
}

func fatal(f string, args ...interface{}) {
	stderr(f, args...)
	os.Exit(1)
}

func warning(f string, args ...interface{}) {
	stderr(f, args...)
}

func debug(f string)  {
	warning(f)
}

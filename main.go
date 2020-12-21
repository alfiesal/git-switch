package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "gsu",
		Usage: "Switch git users quickly. Switches locally by default",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list all available users",
				Action:  listAction,
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a user to the list",
				Action: addAction,
			},
			{
				Name:    "reset",
				Aliases: []string{"a"},
				Usage:   "clear profile list",
				Action: resetAction,
			},
			{
				Name:    "switch",
				Aliases: []string{"s"},
				Usage:   "switch git users. Switches locally by default",
				Action: switchAction,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func store(name string, email string) {

	config := readConfig("git-users.json")
	config.Users = append(config.Users, User{name, email})

	file, err := os.OpenFile("git-users.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.Marshal(config)
	fmt.Println(string(json))
	_, err = file.WriteAt(json, 0)
	if err != nil {
		log.Println(err)
	}
}

func readConfig(configPath string) Config {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	json.Unmarshal(body, &config)

	return config
}

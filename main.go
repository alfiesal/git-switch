package main

import (
	"encoding/json"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Name:  "Git Switch",
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
				Action:  addAction,
			},
			{
				Name:    "reset",
				Aliases: []string{"a"},
				Usage:   "clear profile list",
				Action:  resetAction,
			},
			{
				Name:    "switch",
				Aliases: []string{"s"},
				Usage:   "switch git users. Switches locally by default",
				Action:  switchAction,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func store(name string, email string) {
	config := read()
	config.Users = append(config.Users, User{name, email})

	file, err := os.OpenFile(configFilePath(), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.Marshal(config)

	_, err = file.WriteAt(json, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func read() Config {
	file, err := os.OpenFile(configFilePath(), os.O_RDONLY|os.O_CREATE, 0644)
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

func configFilePath() string {
	homeDir, _ := os.UserHomeDir()

	return homeDir + "/git-users.json"
}

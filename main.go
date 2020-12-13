package main

import (
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
				Action: func(c *cli.Context) error {
					fmt.Println("run list")
					list()

					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a user to the list",
				Action: func(c *cli.Context) error {
					prompt := promptui.Prompt{
						Label:    "Username",
					}

					username, err := prompt.Run()

					if err != nil {
						log.Fatal(err)
					}

					prompt = promptui.Prompt{
						Label:    "Email",
					}

					email, err := prompt.Run()

					if err != nil {
						log.Fatal(err)
					}

					add(username, email)

					return nil
				},
			},
			{
				Name:    "reset",
				Aliases: []string{"a"},
				Usage:   "clear profile list",
				Action: func(c *cli.Context) error {
					os.Remove("git-users.json")
					return nil
				},
			},
			{
				Name:    "switch",
				Aliases: []string{"s"},
				Usage:   "Switch git users. Switches locally by default",
				Action: func(c *cli.Context) error {

					var users []User
					config := readConfig("git-users.json")

					var items []string

					for _, user := range config.Users {
						users = append(users, User{user.Name, user.Email})
						items = append(items, user.Name + " " + user.Email)
					}

					prompt := promptui.Select{
						Label: "Select User",
						Items: items,
					}

					i, _, err := prompt.Run()

					if err != nil {
						fmt.Printf("Prompt failed %v\n", err)
						return nil
					}

					selectedUser := users[i]

					cmd := exec.Command("git","config", "user.name", selectedUser.Name)
					err = cmd.Run()

					if err != nil {
						log.Fatal(err)
					}

					cmd = exec.Command("git","config", "user.email", selectedUser.Email)
					err = cmd.Run()

					if err != nil {
						log.Fatal(err)
					}

					fmt.Printf("Profile has been set\n")

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


func list() {
	config := readConfig("git-users.json")

	for _, user := range config.Users {
		fmt.Println(user.Name + "\t <" + user.Email + ">")
	}
}

func add(name string, email string) {

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

type User struct {
	Name  string
	Email string
}
type Config struct {
	Users []User
}

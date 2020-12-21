package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
)

func listAction(c *cli.Context) error {
	config := readConfig("git-users.json")

	for _, user := range config.Users {
		fmt.Println(user.Name + "\t <" + user.Email + ">")
	}

	return nil
}

func addAction(c *cli.Context) error {
	prompt := promptui.Prompt{
		Label: "Username",
	}

	username, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	prompt = promptui.Prompt{
		Label: "Email",
	}

	email, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	store(username, email)

	return nil
}

func resetAction(c *cli.Context) error {
	os.Remove("git-users.json")
	return nil
}

func switchAction(c *cli.Context) error {

	var users []User
	config := readConfig("git-users.json")

	var items []string

	for _, user := range config.Users {
		users = append(users, User{user.Name, user.Email})
		items = append(items, user.Name+" "+user.Email)
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

	cmd := exec.Command("git", "config", "user.name", selectedUser.Name)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("git", "config", "user.email", selectedUser.Email)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Profile has been set\n")

	return nil
}

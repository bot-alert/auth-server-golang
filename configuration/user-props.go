package configuration

import (
	"gopkg.in/yaml.v2"
	"os"
)

var Users UserList

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type UserList struct {
	Users []User `yaml:"user-list"`
}

func init() {
	data, err := os.ReadFile("configuration/props.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &Users)
	if err != nil {
		panic(err)
	}
}

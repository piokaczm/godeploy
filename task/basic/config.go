package basic

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	rsync = "rsync"
	scp   = "scp"
)

type Config struct {
	Mode         string                 `yaml:"mode"`
	Name         string                 `yaml:"binary_name"`
	Strategy     string                 `yaml:"strategy"`
	Goos         string                 `yaml:"goos"`
	Goarch       string                 `yaml:"goarch"`
	Test         bool                   `yaml:"test"`
	Vendor       bool                   `yaml:"vendor"`
	Environments map[string]Environment `yaml:"environments"`
	Slack        Slack                  `yaml:"slack"`
	CurrentEnv   string
}

type Environment struct {
	Hosts []string
	User  string
	Name  string
}

// load yml with config, prepare tasks and slack config based on it
func loadConfig(path string, currentEnv string) (*Config, *Slack, *Environment, []*Task) {
	boldRed := color.New(color.FgRed, color.Bold)

	currentUser := os.Getenv("USER")
	c := Config{CurrentEnv: currentEnv}

	// load yml
	data, err := ioutil.ReadFile(path) // maybe use abs here
	checkErr(err)
	err = yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		boldRed.Println("Unmarshalling configuration file failed!")
		panic(err)
	}

	c.chooseStrategy()
	slack := c.loadSlack(currentUser)
	env := c.loadEnvironment(currentUser)
	// prepare tasks
	return &c, slack, env, []*Task{}
}

// return pointer to Environment struct with current env variables
func (config Config) loadEnv(currentUser string) *Environment {
	return &Environment{
		Hosts: c.getHosts(),
		User:  currentUser,
		Name:  c.CurrentEnv,
	}
}

func (config Config) loadSlack(currentUser string) *Slack {
	return &config.Slack{
		User: config.currentUser,
		Env:  config.CurrentEnv,
	}
}

// retrieve hosts from config, panic if no hosts found.
func (c *Config) getHosts() []string {
	hosts := c.Environments[c.CurrentEnv]["hosts"]
	if len(hosts) == 0 {
		err := errors.New("No hosts found in config file!")
		panic(err)
	}
	return hosts
}

// if no strategy specified fallback to rsync
func (c *Config) chooseStrategy() {
	if blank(c.Strategy) {
		c.Strategy = rsync
	}
}

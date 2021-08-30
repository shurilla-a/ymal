package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigYmal struct {
	Host          int    `yaml:"host"`
	Port          int    `yaml:"port"`
	Login         string `yaml:"login"`
	Password      string `yaml:"passwd"`
	QueueName     string `yaml:"QueueName"`
	QueueMessages int    `yaml:"QueueMessages"`
	QueueCount    string `yaml:"queueCount"`
}

func inConfigParsing(Configfile string) (*ConfigYmal, error) {
	file, err := ioutil.ReadFile(Configfile)
	if err != nil {
		return nil, err
	}
	c := &ConfigYmal{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", Configfile, err)
	}
	return c, nil
}

func main() {
	c, err := inConfigParsing("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", c)
}

//func (c *ConfigYmal ) Parse(data []byte) error {
//	return yaml.Unmarshal(data,c)
//}
//func main()  {
//data ,err := ioutil.ReadFile("config.yml")
//if err != nil {
//	log.Fatal(err)
//}
//	var config ConfigYmal
//if err := config.Parse(data);err !=nil {
//	log.Fatal(err)
//}
//fmt.Printf("%+v\n", config)
//}

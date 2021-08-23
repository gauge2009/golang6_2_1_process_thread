package main

import (
	"Common"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	path_root := Common.GetCurrentAbPath() // // go run 与 go build  统一使用go run 制指定的
	fmt.Println("GetCurrentAbPath = ", path_root)
	path_yaml := path_root + "\\test.yaml"

	// resultMap := make(map[string]interface{})
	conf := new(Common.Yaml)
	yamlFile, err := ioutil.ReadFile(path_yaml)

	// conf := new(module.Yaml1)
	// yamlFile, err := ioutil.ReadFile("test.yaml")

	// conf := new(module.Yaml2)
	//  yamlFile, err := ioutil.ReadFile("test1.yaml")

	log.Println("yamlFile:", yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)
	// log.Println("conf", resultMap)
}

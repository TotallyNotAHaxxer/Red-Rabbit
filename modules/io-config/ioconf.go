package IO_Configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type YAML_CONF struct {
	Clear_When_Command_Ran bool `yaml:"Clear_On_Command"`
}

func Open(filename string) bool {
	yfile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]YAML_CONF)
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	for _, l := range data {
		if !l.Clear_When_Command_Ran {
			return false // false means it will not clear
		}
	}
	return true // true meaning it will clear after a command is ran
}

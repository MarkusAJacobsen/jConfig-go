package JSONConfig

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strings"
)

const jsonMime = ".json"

type Config struct {
	fileName string
}



func (c *Config) CreateConfig(fileName string) (err error){
	fileName = fileName + jsonMime

	f, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer f.Close()

	stat, _ := f.Stat()
	c.fileName = stat.Name()

	return
}

func (c *Config) Open(fileName string) (err error){
	if !strings.Contains(fileName, jsonMime) {
		err = errors.New("File extension is not of type json")
		return
	}

	f, err := os.Open(fileName)
	if err != nil {
		return
	}

	stat, _ := f.Stat()
	c.fileName = stat.Name()

	return
}

func (c *Config) Write(v interface{}) (err error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return
	}

	f, err := os.OpenFile(c.fileName, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	f.Write(bytes)

	return
}

func (c *Config) Get(v interface{}) (err error) {
	f, err := os.Open(c.fileName)
	if err != nil {
		return
	}
	defer f.Close()

	cb, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	json.Unmarshal(cb, &v)

	return
}


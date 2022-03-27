package json

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jsmzr/bootstrap-config/config"
	"github.com/jsmzr/bootstrap-config/util"
)

type JsonConfig struct{}

type JsonContainer struct {
	dict map[string]interface{}
}

func (c *JsonConfig) Load(filename string) (config.Configer, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var dict map[string]interface{}
	err = json.Unmarshal(data, &dict)
	if err != nil {
		return nil, err
	}
	return &JsonContainer{
		dict: util.FlatStringMap(&dict),
	}, nil
}

func (c *JsonContainer) Get(key string) (interface{}, bool) {
	value, ok := c.dict[key]
	return value, ok
}

func (c *JsonContainer) Resolve(preifx string, p interface{}) error {
	return util.ResolveStruct(&c.dict, preifx, p)
}

func init() {
	config.Register("json", &JsonConfig{})
}

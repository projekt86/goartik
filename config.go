package goartik

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const filePath = "/config/%s.json"

type config struct {
	Pins map[string]uint
}

func newConfig(moduleID ArtikModuleID) (config, error) {
	c := config{}
	id := strings.ToLower(moduleID.String())
	// read json config file
	f, err := ioutil.ReadFile(getRelativePath(fmt.Sprintf(filePath, id)))
	if err != nil {
		return c, err
	}
	// unmarshal json
	err = json.Unmarshal(f, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

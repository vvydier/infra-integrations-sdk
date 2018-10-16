package args

import (
	"encoding/json"
	"flag"
	"fmt"
)

// JSON type, to be used from the arguments structs.
// This argument type will parse a serialized JSON string into a map which
type JSON struct {
	value interface{}
}

// NewJSON returns new JSON
func NewJSON(value interface{}) *JSON {
	return &JSON{value}
}

// Set sets its content
func (i *JSON) Set(s string) error {
	if err := json.Unmarshal([]byte(s), &(i.value)); err != nil {
		return fmt.Errorf("Bad JSON, %v", err)
	}
	return nil
}

// Get provides internal value
func (i *JSON) Get() interface{} { return i.value }

// String returns string representation
func (i *JSON) String() string {
	s, _ := json.Marshal(&(i.value))
	return string(s)
}

func jsonVar(p *JSON, name string, value string, usage string) {
	flag.CommandLine.Var(p, name, usage)
}

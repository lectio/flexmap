package flexmap

import (
	"fmt"
)

// TextKeyMap is an implementation of FlexMap for string keys and untyped values
type textKeyMap map[string]interface{}

// MakeTextKeyMap returns a new, empty TextKeyMap
func MakeTextKeyMap() TextKeyMap {
	return make(textKeyMap)
}

// MakeCustomTextKeyMap returns a new TextKeyMap after calling the initializer function to fill contents
func MakeCustomTextKeyMap(initFn func(v interface{}) error) (TextKeyMap, error) {
	m := make(textKeyMap)
	err := initFn(&m)
	return m, err
}

// Map returns the underlying map data structure
func (m textKeyMap) Map() (interface{}, error) {
	return m, nil
}

// MapValue returns the value of key
func (m textKeyMap) MapValue(key interface{}) (interface{}, bool, error) {
	value, ok := m[key.(string)]
	return value, ok, nil
}

// TextKeyValue returns the generic value of a key
func (m textKeyMap) TextKeyValue(key string) (interface{}, bool) {
	value, ok := m[key]
	return value, ok
}

// TextKeyTextValue returns the value of key as text
func (m textKeyMap) TextKeyTextValue(key string) (string, bool, error) {
	value, ok := m[key]
	switch v := value.(type) {
	case string:
		return v, ok, nil
	default:
		return "", ok, fmt.Errorf("key %q value %v is not a string", key, value)
	}
}

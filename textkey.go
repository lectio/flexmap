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

// MakeCustomTextKeyMap returns a new TextKeyMap after calling the initializer function to fill contents of v using Unmarshal or similar
func MakeCustomTextKeyMap(initFn func(v interface{}) error) (TextKeyMap, error) {
	m := make(textKeyMap)
	err := initFn(&m)
	return m, err
}

// MakeTextKeyMapWithDefaults returns a new TextKeyMap by preloading the array of keys and values
func MakeTextKeyMapWithDefaults(rangeFn func() (startIndex int, endIndex int, iteratorFn func(index int) (Item, bool))) (TextKeyMap, error) {
	m := make(textKeyMap)
	start, end, fn := rangeFn()
	for i := start; i <= end; i++ {
		item, keepGoing := fn(i)
		m[item.Key().(string)] = item.Value()
		if !keepGoing {
			break
		}
	}
	return m, nil
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

// ForEachKey iterates and executes function fn for each key value pair
func (m textKeyMap) ForEachKey(fn func(key interface{}, value interface{}) bool) {
	for key, value := range m {
		keepGoing := fn(key, value)
		if !keepGoing {
			break
		}
	}
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

// ForEachTextKey iterates and executes function fn for each key value pair
func (m textKeyMap) ForEachTextKey(fn func(key string, value interface{}) bool) {
	for key, value := range m {
		keepGoing := fn(key, value)
		if !keepGoing {
			break
		}
	}
}

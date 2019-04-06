package flexmap

// GenericMap is an implementation of FlexMap for untyped keys and values
type genericMap map[interface{}]interface{}

// MakeGenericMap returns a new, empty GenericMap
func MakeGenericMap() Map {
	return make(genericMap)
}

// Map returns the underlying map data structure
func (m genericMap) Map() (interface{}, error) {
	return m, nil
}

// MapValue returns the value of key
func (m genericMap) MapValue(key interface{}) (interface{}, bool, error) {
	value, ok := m[key]
	return value, ok, nil
}

package flexmap

// GenericMap is an implementation of FlexMap for untyped keys and values
type GenericMap map[interface{}]interface{}

// Map returns the underlying map data structure
func (m GenericMap) Map() (interface{}, error) {
	return m, nil
}

// MapValue returns the value of key in the GenericMap
func (m GenericMap) MapValue(key interface{}) (interface{}, bool, error) {
	value, ok := m[key]
	return value, ok, nil
}

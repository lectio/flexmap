package flexmap

// FlexMap is a flexible map structure which allows generic keys and values
type FlexMap interface {
	Map() (interface{}, error)
	MapValue(key interface{}) (interface{}, bool, error)
}

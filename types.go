package flexmap

// Map is a flexible map structure which allows generic keys and values
type Map interface {
	Map() (interface{}, error)
	MapValue(key interface{}) (interface{}, bool, error)
}

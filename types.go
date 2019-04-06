package flexmap

// Map is a flexible map structure which allows generic keys and values
type Map interface {
	Map() (interface{}, error)
	MapValue(key interface{}) (interface{}, bool, error)
}

// TextKeyMap is a flexible map structure which allows text keys and generic values
type TextKeyMap interface {
	Map
	TextKeyValue(key string) (interface{}, bool)
	TextKeyTextValue(key string) (string, bool, error)
}

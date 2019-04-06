package flexmap

// Item is a key value pair
type Item interface {
	Key() interface{}
	Value() interface{}
}

// Map is a flexible map structure which allows generic keys and values
type Map interface {
	Map() (interface{}, error)
	MapValue(key interface{}) (interface{}, bool, error)
	ForEachKey(fn func(key interface{}, value interface{}) bool)
}

// TextKeyMap is a flexible map structure which allows text keys and generic values
type TextKeyMap interface {
	Map
	TextKeyValue(key string) (interface{}, bool)
	TextKeyTextValue(key string) (string, bool, error)
	ForEachTextKey(fn func(key string, value interface{}) bool)
}

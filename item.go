package flexmap

// Item is a key value pair
type Item interface {
	Key() interface{}
	Value() interface{}
}

type item struct {
	key   interface{}
	value interface{}
}

func (i item) Key() interface{} {
	return i.key
}

func (i item) Value() interface{} {
	return i.value
}

// MakeItem creates a new generic key value pair
func MakeItem(key, value interface{}) Item {
	i := new(item)
	i.key = key
	i.value = value
	return i
}

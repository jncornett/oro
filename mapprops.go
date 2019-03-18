package oro

// MapProps ...
type MapProps map[Key]interface{}

var _ Props = MapProps(nil)

// Get ...
func (props MapProps) Get(k Key) interface{} {
	return props[k]
}

// Len ...
func (props MapProps) Len() int {
	return len(props)
}

// Keys ...
func (props MapProps) Keys() (keys []Key) {
	if len(props) > 0 {
		keys = make([]Key, 0, len(keys))
	}
	for k := range props {
		keys = append(keys, k)
	}
	return
}

// Dispatch ...
func (props MapProps) Dispatch(Action) {
	panic("not implemented")
}

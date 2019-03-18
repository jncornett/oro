package oro

// Props ...
type Props interface {
	Get(Key) interface{}
	Len() int
	Keys() []Key
}

// WithProp ...
func WithProp(props Props, prop Prop) Props {
	if props.Get(prop.Key) == prop.Value {
		return props
	}
	out := make(MapProps, props.Len()+1)
	for _, key := range props.Keys() {
		if key == prop.Key {
			out[key] = prop.Value
		} else {
			out[key] = props.Get(key)
		}
	}
	return out
}

// Prop ...
type Prop struct {
	Key   Key
	Value interface{}
}

// IntProp ...
func IntProp(k Key, v int64) Prop {
	return Prop{Key: k, Value: v}
}

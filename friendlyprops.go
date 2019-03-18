package oro

// FriendlyProps ...
type FriendlyProps struct {
	Props
}

// Wrap ...
func Wrap(props Props) *FriendlyProps {
	return &FriendlyProps{props}
}

func (props *FriendlyProps) String(k Key) string {
	s, _ := props.Props.Get(k).(string)
	return s
}

// Bool ...
func (props *FriendlyProps) Bool(k Key) bool {
	b, _ := props.Props.Get(k).(bool)
	return b
}

// Int ...
func (props *FriendlyProps) Int(k Key) int64 {
	i, _ := props.Props.Get(k).(int64)
	return i
}

// Float ...
func (props *FriendlyProps) Float(k Key) float64 {
	f, _ := props.Props.Get(k).(float64)
	return f
}

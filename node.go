package oro

import "strconv"

// Kind ...
type Kind int

const (
	// KindInvalid ...
	KindInvalid Kind = iota
	// KindElement ...
	KindElement
	// KindText ...
	KindText
	// KindFragment ...
	KindFragment
)

func (k Kind) String() string {
	switch k {
	case KindInvalid:
		return "Invalid"
	case KindElement:
		return "Element"
	case KindText:
		return "Text"
	case KindFragment:
		return "Fragment"
	default:
		return "Kind(" + strconv.Itoa(int(k)) + ")"
	}
}

// Event ...
type Event interface{}

// Callback ...
type Callback func(Event)

// Attributes ...
type Attributes map[Key]interface{}

// Callbacks ...
type Callbacks map[Key]Callback

// Node ...
type Node struct {
	Kind       Kind
	Value      string
	Callbacks  Callbacks
	Attributes Attributes
	Children   []*Node
}

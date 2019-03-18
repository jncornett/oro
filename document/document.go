package document

import (
	"errors"
	"syscall/js"
)

var errNoSuchElement = errors.New("no such element")

// GetElementByID ...
func GetElementByID(id string) Element {
	v := js.Global().Get("document").Call("getElementById", js.ValueOf(id))
	if Nil(v) {
		return Element{}
	}
	return Element{v}
}

// CreateElement ...
func CreateElement(tagName string) Element {
	v := js.Global().Get("document").Call("createElement", js.ValueOf(tagName))
	return Element{v}
}

// CreateTextNode ...
func CreateTextNode(data string) Element {
	v := js.Global().Get("document").Call("createTextNode", js.ValueOf(data))
	return Element{v}
}

// Nil ...
func Nil(v js.Value) bool {
	switch v.Type() {
	case js.TypeNull, js.TypeUndefined:
		return true
	}
	return false
}

// Element ...
type Element struct {
	RawValue js.Value
}

// Nil ...
func (el *Element) Nil() bool {
	if el == nil {
		return true
	}
	return Nil(el.RawValue)
}

// Children ...
func (el *Element) Children() []Element {
	if el.Nil() {
		return nil
	}
	var out []Element
	children := el.RawValue.Get("childNodes")
	for i := 0; i < children.Length(); i++ {
		out = append(out, Element{children.Index(i)})
	}
	return out
}

// Tag ...
func (el *Element) Tag() string {
	if el.Nil() {
		return ""
	}
	return el.RawValue.Get("nodeName").String()
}

// Text ...
func (el *Element) Text() bool {
	if el.Nil() {
		return false
	}
	return "#text" == el.RawValue.Get("nodeName").String()
}

// Append ...
func (el *Element) Append(child Element) {
	if el.Nil() {
		return
	}
	el.RawValue.Call("appendChild", child.RawValue)
}

// Replace ...
func (el *Element) Replace(newChild, oldChild Element) {
	if el.Nil() {
		return
	}
	el.RawValue.Call("replaceChild", newChild.RawValue, oldChild.RawValue)
}

// Remove ...
func (el *Element) Remove(child Element) {
	if el.Nil() {
		return
	}
	el.RawValue.Call("removeChild", child.RawValue)
}

// Set ...
func (el *Element) Set(key string, v interface{}) {
	if el.Nil() {
		return
	}
	el.RawValue.Set(key, v)
}

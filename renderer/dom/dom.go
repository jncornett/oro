package dom

import (
	"syscall/js"

	"github.com/jncornett/oro"
	"github.com/jncornett/oro/document"
)

// Renderer ...
type Renderer struct {
	DOMContainer document.Element
}

// Render ...
func (r *Renderer) Render(root *oro.Node) {
	resolveDOM(r.DOMContainer, r.DOMContainer.Children(), []*oro.Node{root})
}

func resolveDOM(parent document.Element, children []document.Element, maybeNodes []*oro.Node) {
	nodes := clean(maybeNodes)

	var (
		deleteMode     bool
		updateBoundary int
	)
	if len(nodes) < len(children) {
		deleteMode = true
		updateBoundary = len(nodes)
	} else {
		updateBoundary = len(children)
	}

	for i := 0; i < updateBoundary; i++ {
		// FIXME handle case where child.Nil() is true.
		prevChild := children[i]
		node := nodes[i]
		var nextChild document.Element
		switch node.Kind {
		case oro.KindText:
			if prevChild.Text() {
				prevChild.Set("data", node.Value)
			} else {
				nextChild = document.CreateTextNode(node.Value)
				parent.Replace(nextChild, prevChild)
			}
		case oro.KindElement:
			if prevChild.Tag() == node.Value {
				setAttributes(prevChild, node.Attributes)
				setCallbacks(prevChild, node.Callbacks)
				resolveDOM(prevChild, prevChild.Children(), node.Children)
			} else {
				nextChild = document.CreateElement(node.Value)
				setAttributes(nextChild, node.Attributes)
				setCallbacks(nextChild, node.Callbacks)
				parent.Replace(nextChild, prevChild)
				resolveDOM(nextChild, nil, node.Children)
			}
		default:
			// unhandled node.Kind
			continue
		}
	}

	if deleteMode {
		for i := updateBoundary; i < len(children); i++ {
			parent.Remove(children[i])
		}
	} else {
		// create mode
		for i := updateBoundary; i < len(nodes); i++ {
			node := nodes[i]
			switch node.Kind {
			case oro.KindText:
				child := document.CreateTextNode(node.Value)
				parent.Append(child)
			case oro.KindElement:
				child := document.CreateElement(node.Value)
				setAttributes(child, node.Attributes)
				setCallbacks(child, node.Callbacks)
				resolveDOM(child, nil, node.Children)
				parent.Append(child)
			default:
				// unhandled node.Kind
				continue
			}
		}
	}
}

func setAttributes(el document.Element, attrs oro.Attributes) {
	for name := range attrs {
		switch name {
		default:
			panic("TODO handle other types of attributes")
		}
	}
}

func setCallbacks(el document.Element, callbacks oro.Callbacks) {
	for name, cb := range callbacks {
		cb := cb
		switch name {
		case "onClick":
			el.Set("onclick", js.NewEventCallback(0, func(js.Value) {
				cb(nil)
			}))
		}
	}
}

func clean(nodes []*oro.Node) []oro.Node {
	if len(nodes) == 0 {
		return nil
	}
	out := make([]oro.Node, 0, len(nodes))
	for _, node := range nodes {
		if node == nil {
			continue
		}
		if node.Kind == oro.KindFragment {
			out = append(out, clean(node.Children)...)
			continue
		}
		out = append(out, *node)
	}
	return out
}

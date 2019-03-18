package oro

func textNode(value string) *Node {
	return &Node{
		Kind:  KindText,
		Value: value,
	}
}

func toNodes(stringsOrNodes ...interface{}) []*Node {
	nodes := make([]*Node, 0, len(stringsOrNodes))
	for _, v := range stringsOrNodes {
		if v == nil {
			nodes = append(nodes, nil)
		} else {
			switch t := v.(type) {
			case *Node:
				nodes = append(nodes, t)
			case string:
				nodes = append(nodes, textNode(t))
			default:
				// FIXME this should be a warning
				nodes = append(nodes, nil)
			}
		}
	}
	return nodes
}

func toAttributes(props Props) Attributes {
	if props == nil {
		return nil
	}
	keys := props.Keys()
	if len(keys) == 0 {
		return nil
	}
	attrs := make(Attributes, len(keys))
	for _, key := range keys {
		attrs[key] = props.Get(key)
	}
	return attrs
}

// Element ...
func Element(name string, props Props, callbacks Callbacks, stringsOrNodes ...interface{}) *Node {
	return &Node{
		Kind:       KindElement,
		Value:      name,
		Attributes: toAttributes(props),
		Callbacks:  callbacks,
		Children:   toNodes(stringsOrNodes...),
	}
}

// Fragment ...
func Fragment(stringsOrNodes ...interface{}) *Node {
	return &Node{
		Kind:     KindFragment,
		Children: toNodes(stringsOrNodes...),
	}
}

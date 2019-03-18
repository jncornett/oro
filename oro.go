package oro

// Key ...
type Key string

// Component ...
type Component func(State) *Node

// Renderer ...
type Renderer interface {
	Render(*Node)
}

// Action ...
type Action struct {
	Name Key
	Args []interface{}
}

// Run ...
func Run(store *Store, renderer Renderer, c Component) {
	for props := range store.Subscribe() {
		node := c(State{Props: props, Dispatch: store.Dispatch})
		renderer.Render(node)
	}
}

package oro

// State ...
type State struct {
	Props
	Dispatch func(Action)
}

package main

import (
	"strconv"

	"github.com/jncornett/oro/oroquickstart"

	"github.com/jncornett/oro"
)

// Reducer ...
func Reducer(props oro.Props, action oro.Action) oro.Props {
	switch action.Name {
	case "increment":
		curCount := oro.Wrap(props).Int("count")
		return oro.WithProp(props, oro.IntProp("count", curCount+1))
	default:
		panic("TODO: invalid action")
	}
}

// App ...
func App(state oro.State) *oro.Node {
	count := state.Props.Get("count").(int64)
	return oro.Fragment(
		oro.Element("h1", nil, nil, "Hello, wasm"),
		oro.Element("p", nil, nil, "Count: ", strconv.FormatInt(count, 10)),
		oro.Element("button", nil, oro.Callbacks{
			"onClick": func(oro.Event) {
				state.Dispatch(oro.Action{Name: "increment"})
			},
		}, "Increment counter"),
	)
}

func main() {
	oroquickstart.Start(
		"root",
		oro.MapProps{"count": int64(0)},
		Reducer,
		App,
	)
}

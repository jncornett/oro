package main

import (
	"strconv"

	"github.com/jncornett/oro"
	"github.com/jncornett/oro/document"
	"github.com/jncornett/oro/renderer/dom"
)

func main() {
	done := make(chan struct{})
	var (
		renderer = &dom.Renderer{DOMContainer: document.GetElementByID("root")}
		store    = &oro.Store{
			InitialProps: oro.MapProps{"count": int64(0)},
			Reducer: func(props oro.Props, action oro.Action) oro.Props {
				switch action.Name {
				case "increment":
					curCount := oro.Wrap(props).Int("count")
					return oro.WithProp(props, oro.IntProp("count", curCount+1))
				default:
					panic("TODO: invalid action")
				}
			},
		}
		app = func(state oro.State) *oro.Node {
			count := state.Props.Get("count").(int64)
			return oro.Fragment(
				oro.Element("h1", nil, nil, "Hello, wasm"),
				oro.Element("p", nil, nil, "Count: ", strconv.FormatInt(count, 10)),
				oro.Element("button", nil, oro.Callbacks{
					"onClick": func(oro.Event) {
						store.Dispatch(oro.Action{Name: "increment"})
					},
				}, "Increment counter"),
			)
		}
	)
	go oro.Run(store, renderer, app)
	<-done
}

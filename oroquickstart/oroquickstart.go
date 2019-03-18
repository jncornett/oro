package oroquickstart

import (
	"github.com/jncornett/oro"
	"github.com/jncornett/oro/document"
	"github.com/jncornett/oro/renderer/dom"
)

// Start ...
func Start(
	elementID string,
	initialState oro.Props,
	reducer oro.Reducer,
	rootComponent oro.Component,
) {
	done := make(chan struct{})
	renderer := &dom.Renderer{DOMContainer: document.GetElementByID("root")}
	store := &oro.Store{
		InitialProps: initialState,
		Reducer:      reducer,
	}
	go oro.Run(store, renderer, rootComponent)
	<-done
}

package oro

import "sync"

// Reducer ...
type Reducer func(Props, Action) Props

// Store ...
type Store struct {
	InitialProps Props
	Reducer      Reducer
	mux          sync.Mutex
	lastProps    Props
	subscribers  map[int64]chan Props
}

func (store *Store) init() {
	if store.subscribers == nil {
		store.subscribers = make(map[int64]chan Props)
	}
}

func (store *Store) notify(props Props) {
	store.mux.Lock()
	defer store.mux.Unlock()
	store.init()
	store.lastProps = props
	for _, sub := range store.subscribers {
		sub <- props
	}
}

// Subscribe ...
func (store *Store) Subscribe() <-chan Props {
	props := make(chan Props, 1)
	store.mux.Lock()
	defer store.mux.Unlock()
	store.init()
	if store.lastProps == nil {
		props <- store.InitialProps
	} else {
		props <- store.lastProps
	}
	key := int64(len(store.subscribers))
	store.subscribers[key] = props
	return props
}

// Props ...
func (store *Store) Props() Props {
	store.mux.Lock()
	defer store.mux.Unlock()
	if store.lastProps == nil {
		return store.InitialProps
	}
	return store.lastProps
}

// Dispatch ...
func (store *Store) Dispatch(action Action) {
	nextProps := store.Reducer(store.Props(), action)
	store.notify(nextProps)
}

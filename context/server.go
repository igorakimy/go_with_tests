package context

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, store.Fetch())
	}
}

type Store interface {
	Fetch() string
}

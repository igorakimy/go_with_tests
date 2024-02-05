package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {

	player := strings.TrimPrefix(r.URL.String(), "/players/")

	if player == "Pepper" {
		_, _ = fmt.Fprint(w, "20")
		return
	} else {
		_, _ = fmt.Fprint(w, "10")
		return
	}
}

package cowsay

import (
	"fmt"
	"net/http"
	"zeitfiles/templates"
)

// Handler function
func Handler(w http.ResponseWriter, r *http.Request) {
	content, err := templates.FileContent()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, content)
}

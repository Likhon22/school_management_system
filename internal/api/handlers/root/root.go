package root

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from test!")
}

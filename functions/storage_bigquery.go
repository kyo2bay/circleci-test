// Package functions provides a set of Cloud Functions samples.
package functions

import (
	"fmt"
	"net/http"
)

// HelloGet is an HTTP Cloud Function.
func HelloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

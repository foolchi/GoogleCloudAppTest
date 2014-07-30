package hello

import (
	"appengine"
	"appengine/user"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	fmt.Fprintf(w, "Error 1")
	if u == nil {
		fmt.Fprintf(w, "Error 2")
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			fmt.Fprintf(w, "Error 3")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Error 4")
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprint(w, "你好, %v!", u)
}

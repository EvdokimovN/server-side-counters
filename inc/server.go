package inc

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (I inc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	// Should be safe, since we will always get at least empty path
	lastSegment := pathSegments[len(pathSegments)-1]
	switch lastSegment {
	case "peek":
		m := make(map[string][]int)
		m["nums"] = make([]int, I.Size())
		for i := range I.nums {
			c := I.nums[i]
			// Pop value and send it back to channel
			// to prevent blocking
			n := <-c
			c <- n
			m["nums"][i] = n
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(m)
	case "switch":
		m, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			return
		}
		id, err := strconv.Atoi(m.Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			return
		}
		I.Switch(id)
		w.WriteHeader(http.StatusAccepted)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

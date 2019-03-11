package api

import (
	"encoding/json"
	"net/http"

	"../../db"
	keyMgmt "../key_management"
)

var apikeyQueryParam = "apikey"

// LatestGet ...
// @Summary List Posts
// @Description get posts
// @Produce  json
// @Param apikey query string true "apikey"
// @Param offset query int true "offset"
// @Success 200 {array} models.Posts
// @Router /latest [get]
func LatestGet(w http.ResponseWriter, r *http.Request) {
	apikey := r.URL.Query().Get(apikeyQueryParam)
	user, err := keyMgmt.GetValueFromKey(apikey)
	if err != nil || user == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Probablty you are not logged in"))
	} else {
		posts, err := db.LatestPosts()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Error or No Posts"))
		} else if len(posts) < 1 {
			w.WriteHeader(http.StatusPartialContent)
			w.Write([]byte("No Content in network"))
		} else {
			w.WriteHeader(http.StatusOK)
			postsMarshall, err := json.Marshal(posts)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Error. Marshal Error"))
			}
			w.Write(postsMarshall)
		}

	}
}

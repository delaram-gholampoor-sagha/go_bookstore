// we need data that is unmarshaled data
// we will recieve some request from the user and it will be json we need to unmarshal it to be able to use it out controller

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}

}

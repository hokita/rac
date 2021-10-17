package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/hokita/ac/testapi/domain"
)

// IndexHandler struct
type IndexHandler struct{}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u1 := &domain.User{ID: 1, Name: "hokita"}
	u2 := &domain.User{ID: 2, Name: "hideee"}

	b, err := json.Marshal([]*domain.User{u1, u2})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(b))
}

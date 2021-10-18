package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hokita/rac/testapi/domain"
)

// IndexUsersHandler struct
type IndexUsersHandler struct{}

func (h *IndexUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u1 := &domain.User{ID: 1, Name: "hokita"}
	u2 := &domain.User{ID: 2, Name: "hideee"}

	b, err := json.Marshal([]*domain.User{u1, u2})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(b))
}

// ShowUserHandler struct
type ShowUserHandler struct{}

func (h *ShowUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad request")
		return
	}

	u1 := &domain.User{ID: 1, Name: "hokita"}
	u2 := &domain.User{ID: 2, Name: "hideee"}
	users := []*domain.User{u1, u2}

	var user *domain.User
	for _, u := range users {
		if u.ID == id {
			user = u
		}
	}
	if user == nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "user not found")
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(b))
}

// CreateUserHandler struct
type CreateUserHandler struct{}

func (h *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var inputUser struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&inputUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad request")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, inputUser.Name)
}

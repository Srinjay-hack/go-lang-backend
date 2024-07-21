package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"srinjay.com/m/types"
	"srinjay.com/m/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {

	return &Handler{store: store}

}

func (h *Handler) RegistryRoutes(router *mux.Router) {

	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//getJSOnPayload
	var payload types.RegistryUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//check if the user exist;

	// if it doesnt we create the user;

}

package user

import (
	"encoding/json"
	"go/types"
	"net/http"
	"github.com/gorilla/mux"
	"srinjay.com/m/utils"
)

type Handler struct {
}

func NewHandler() *Handler {

	return &Handler{}

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
	if err := utils.ParseJSON(r.Body,payload); err != nil{

	}

	//check if the user exist;

	// if it doesnt we create the user;

}

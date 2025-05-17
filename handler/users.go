package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()

	user, err := h.UseCase.ListUsers(ctx)
	if err != nil {
		WriteError(w, 500, err)
	}
	WriteData(w, 200, user)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()
	userID, err := strconv.ParseInt(params.ByName("user_id"), 10, 64)
	if err != nil {
		WriteError(w, 400, err)
	}
	user, err := h.UseCase.GetUser(ctx, userID)
	if err != nil {
		WriteError(w, 500, err)
	}
	WriteData(w, 200, user)
}

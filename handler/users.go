package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/ryanbaskara/learning-go/entity"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()

	var userReq entity.CreateUserRequest
	if err := UnmarshalRequestBody(r, &userReq); err != nil {
		WriteError(w, 400, err)
		return
	}

	user, err := h.UseCase.CreateUser(ctx, &userReq)
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	WriteData(w, 200, user)
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()

	user, err := h.UseCase.ListUsers(ctx)
	if err != nil {
		WriteError(w, 500, err)
		return
	}
	WriteData(w, 200, user)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()
	userID, err := strconv.ParseInt(params.ByName("user_id"), 10, 64)
	if err != nil {
		WriteError(w, 400, err)
		return
	}
	user, err := h.UseCase.GetUser(ctx, userID)
	if err != nil {
		WriteError(w, 500, err)
		return
	}
	WriteData(w, 200, user)
}

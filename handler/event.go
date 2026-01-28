package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ryanbaskara/learning-go/entity"
	errlib "github.com/ryanbaskara/learning-go/errors"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()

	var req entity.CreateEventRequest
	if err := UnmarshalRequestBody(r, &req); err != nil {
		newErr := errlib.NewErrorDetail(err.Error(), http.StatusBadRequest)
		WriteErrorDetail(w, newErr)
		return
	}

	var err error
	if req.UserID == "" {
		err = errlib.NewErrorDetail("User ID can't empty", http.StatusBadRequest)
	}
	if req.Type == "" {
		err = errlib.NewErrorDetail("Type can't empty", http.StatusBadRequest)
	}
	if err != nil {
		WriteErrorDetail(w, err)
		return
	}

	event, err := h.eventUseCase.CreateEvent(ctx, &req)
	if err != nil {
		WriteErrorDetail(w, err)
		return
	}

	WriteData(w, http.StatusOK, event)
}

func (h *Handler) ListEvents(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()
	queries := r.URL.Query()
	userId := queries.Get("user_id")

	if userId == "" {
		newErr := errlib.NewErrorDetail("User ID can't empty", http.StatusBadRequest)
		WriteErrorDetail(w, newErr)
		return
	}

	req := &entity.ListEventRequest{
		UserID: userId,
	}

	summary, err := h.eventUseCase.ListEventByUserID(ctx, req)
	if err != nil {
		WriteErrorDetail(w, err)
		return
	}
	WriteData(w, http.StatusOK, summary)

}

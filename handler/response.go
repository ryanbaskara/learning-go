package handler

import (
	"encoding/json"
	"net/http"
)

type Meta map[string]interface{}

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Meta    Meta        `json:"meta,omitempty"`
}

func Write(w http.ResponseWriter, httpStatus int, data interface{}, message string, meta ...Meta) {
	if len(meta) == 0 {
		meta = append(meta, Meta{})
	}
	m := meta[0]
	m["http_status"] = httpStatus
	resp := Response{
		Meta:    m,
		Data:    data,
		Message: message,
	}
	b, _ := json.Marshal(resp)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(b) // #nosec
}

func WriteMessage(w http.ResponseWriter, httpStatus int, message string) {
	Write(w, httpStatus, nil, message)
}

func WriteError(w http.ResponseWriter, httpStatus int, err error) {
	WriteMessage(w, httpStatus, err.Error())
}

func WriteData(w http.ResponseWriter, httpStatus int, data interface{}, meta ...Meta) {
	Write(w, httpStatus, data, "", meta...)
}

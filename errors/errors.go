package errors

import "net/http"

type Error struct {
	Status  int16  `json:"status"`
	Message string `json:"message"`
}

func NotFound() Error {
	return Error{
		Status:  http.StatusNotFound,
		Message: "Element was not found.",
	}
}

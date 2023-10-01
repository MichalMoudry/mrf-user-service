package transport

import (
	"net/http"
	"user-service/internal/transport/util"
)

// Method for handling requests for deleting users.
func (handler *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	err := handler.Services.UserService.DeleteUser(r.Context(), "userId")
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, nil)
}

package transport

import (
	"net/http"
	"user-service/internal/transport/util"
)

// Method for handling requests for deleting users.
func (handler *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	handler.Services.UserService.DeleteUser(r.Context(), "test-user")
	util.WriteResponse(w, http.StatusOK, nil)
}

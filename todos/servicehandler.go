package todos

//
//    THIS IS AUTOGENERATED BY sysl
//

import (
	"encoding/json"
	"net/http"

	"github.com/anz-bank/syslgen-examples/restlib"
)

// Handler interface for Todos
type Handler interface {
	GetCommentsHandler(w http.ResponseWriter, r *http.Request)
	GetPostsHandler(w http.ResponseWriter, r *http.Request)
	GetTodosIDHandler(w http.ResponseWriter, r *http.Request)
	PostCommentsHandler(w http.ResponseWriter, r *http.Request)
}

// ServiceHandler for Todos API
type ServiceHandler struct {
	serviceInterface ServiceInterface
}

// NewServiceHandler for Todos
func NewServiceHandler(serviceInterface ServiceInterface) *ServiceHandler {
	return &ServiceHandler{serviceInterface}
}

// GetCommentsHandler ...
func (s *ServiceHandler) GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	PostID := restlib.GetQueryParam(r, "postId")
	httpStatus, headerMap, Posts := s.serviceInterface.GetComments(PostID)
	restlib.SetHeaders(w, headerMap)
	restlib.SendHTTPResponse(w, httpStatus, Posts)
}

// GetPostsHandler ...
func (s *ServiceHandler) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	httpStatus, headerMap, Posts := s.serviceInterface.GetPosts()
	restlib.SetHeaders(w, headerMap)
	restlib.SendHTTPResponse(w, httpStatus, Posts)
}

// GetTodosIDHandler ...
func (s *ServiceHandler) GetTodosIDHandler(w http.ResponseWriter, r *http.Request) {
	ID := restlib.GetURLParam(r, "id")
	httpStatus, headerMap, Todo := s.serviceInterface.GetTodosID(ID)
	restlib.SetHeaders(w, headerMap)
	restlib.SendHTTPResponse(w, httpStatus, Todo)
}

// PostCommentsHandler ...
func (s *ServiceHandler) PostCommentsHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newPost Post

	decodeErr := decoder.Decode(&newPost)
	if decodeErr != nil {
		errResp := s.serviceInterface.GetErrorResponse(http.StatusBadRequest, "Error reading request body", decodeErr)
		restlib.SendHTTPResponse(w, http.StatusBadRequest, errResp)
		return
	}

	httpStatus, headerMap, Post := s.serviceInterface.PostComments(newPost)
	restlib.SetHeaders(w, headerMap)
	restlib.SendHTTPResponse(w, httpStatus, Post)
}

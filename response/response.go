// response/response.go
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse1 interface {
	OK(c *gin.Context, data any) (int, any)
	Accepted(c *gin.Context, data any) (int, any)
	Created(c *gin.Context, data any) (int, any)
	Updated(c *gin.Context) (int, any)
	Deleted(c *gin.Context) (int, any)
	NoContent(c *gin.Context) (int, any)
	BindFailed(c *gin.Context, err error) (int, any)
	ParamFailed(c *gin.Context, err error) (int, any)
	Error(c *gin.Context, who string, err error) (int, any)
	ValidateFailed(c *gin.Context, who string, err error) (int, any)
}

type DefaultResponse struct{}

func (r DefaultResponse) OK(c *gin.Context, data any) (int, any) {
	return http.StatusOK, gin.H{"status": "success", "data": data}
}

func (r DefaultResponse) Accepted(c *gin.Context, data any) (int, any) {
	return http.StatusAccepted, gin.H{"status": "accepted", "data": data}
}

func (r DefaultResponse) Created(c *gin.Context, data any) (int, any) {
	return http.StatusCreated, gin.H{"status": "created", "data": data}
}

func (r DefaultResponse) Updated(c *gin.Context) (int, any) {
	return http.StatusOK, gin.H{"status": "updated"}
}

func (r DefaultResponse) Deleted(c *gin.Context) (int, any) {
	return http.StatusOK, gin.H{"status": "deleted"}
}

func (r DefaultResponse) NoContent(c *gin.Context) (int, any) {
	return http.StatusNoContent, nil
}

func (r DefaultResponse) BindFailed(c *gin.Context, err error) (int, any) {
	return http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()}
}

func (r DefaultResponse) ParamFailed(c *gin.Context, err error) (int, any) {
	return http.StatusUnprocessableEntity, gin.H{"status": "error", "message": err.Error()}
}

func (r DefaultResponse) Error(c *gin.Context, who string, err error) (int, any) {
	return http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error(), "where": who}
}

func (r DefaultResponse) ValidateFailed(c *gin.Context, who string, err error) (int, any) {
	return http.StatusUnprocessableEntity, gin.H{"status": "error", "message": err.Error(), "where": who}
}

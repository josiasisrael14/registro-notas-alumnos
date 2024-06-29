package login

import (
	"notas/domain/login"
	"notas/model"
	"notas/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const _headerAuthorization = "Authorization"

var jwtKey = []byte("clave_secreta_del_jwt")

type handler struct {
	useCase  login.UseCase
	response response.ApiResponse1
}

func newHandler(useCase login.UseCase, response response.ApiResponse1) handler {
	return handler{useCase: useCase, response: response}
}

func (h handler) authenticate(c *gin.Context) {
	var login model.Login

	if err := c.BindJSON(&login); err != nil {
		c.JSON(h.response.BindFailed(c, err))
		return
	}

	m, err := h.useCase.Login(c.Request.Context(), login)

	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.Login()", err))
		return
	}

	token, err := generarToken(m.IdTeacher)
	if err != nil {
		c.JSON(h.response.Error(c, "h.useCase.generarToken()", err))
		return
	}

	m.Token = token

	c.Header(_headerAuthorization, "Bearer "+token)

	c.JSON(h.response.OK(c, m))

}

func generarToken(idUser string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &model.Claims{
		UserID: idUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

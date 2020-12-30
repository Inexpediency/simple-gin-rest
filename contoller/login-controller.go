package controller

import (
	"Inexpediency/simple-gin-rest/dto"
	"Inexpediency/simple-gin-rest/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController interface {
	Login(c *gin.Context)
}

type loginController struct {
	loginService service.LoginService
	jwtService service.JWTService
}

func NewLoginController() LoginController {
	return &loginController{
		loginService: service.NewLoginService(),
		jwtService: service.NewJWTService(),
	}
}

// Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Router /auth/token [post]
func (controller *loginController) Login(c *gin.Context) {
	var credentials dto.AuthCredentials

	err := c.BindJSON(&credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		generatedToken, err := controller.jwtService.GenerateToken(credentials.Username, true)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": generatedToken})
		return
	}

	c.AbortWithStatus(http.StatusUnauthorized)
}

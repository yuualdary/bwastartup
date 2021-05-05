package middleware

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(AuthService auth.Service, UserService user.Service) gin.HandlerFunc{

return func(c *gin.Context){

	AuthHeader := c.GetHeader("Authorization")

	if !strings.Contains(AuthHeader,"Bearer"){
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	//bearer tokentokentoken split bearer dengan spasi

	TokenString := ""

	ArrayToken := strings.Split(AuthHeader, " ")

	if len(ArrayToken) == 2{
		TokenString = ArrayToken[1]

	}
	//check apakah token valid
	Token, err := auth.NewService().ValidateToken(TokenString)

	if err != nil {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
		}
	//ngeclaim token 
	Claim, IsTokenOk := Token.Claims.(jwt.MapClaims)
 
	if !IsTokenOk || !Token.Valid {

		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return

	}

	UserID := int(Claim["user_id"].(float64))
	
	User, err := UserService.GetUserById(UserID)
	
	if err != nil {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
		}

		c.Set("CurrentUser", User)

	}

}
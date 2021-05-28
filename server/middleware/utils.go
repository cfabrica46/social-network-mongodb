package middleware

import (
	"errors"
	"time"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/cfabrica46/social-network-mongodb/server/token"
	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) (user database.User, err error) {
	var tokenValue database.Token

	if err = c.ShouldBindHeader(&tokenValue); err != nil {
		err = errors.New("internal Error")
		return
	}

	check := database.CheckIfTokenIsInBlackList(tokenValue.Content)
	if check {
		err = errors.New("el Token no es válido")
		return
	}

	user, err = token.ExtractUserFromClaims(tokenValue.Content)

	if err != nil {
		err = errors.New("internal Error")
		return
	}

	user.Token = tokenValue.Content

	deadline, err := time.Parse(time.ANSIC, user.Deadline)

	if err != nil {
		err = errors.New("internal Error")
		return
	}

	checkTime := time.Now().Local().After(deadline)

	if !checkTime {
		err = errors.New("el Token no es válido")
		return
	}
	return
}

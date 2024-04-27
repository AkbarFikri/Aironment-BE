package helper

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/pkg/model"

)

func GetUserLoginData(c *gin.Context) model.UserTokenData {
	getUser, _ := c.Get("user")
	user := getUser.(model.UserTokenData)

	return user
}

func GenerateRandomInt(n int) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var letterRunes = []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}

	dump := string(b)
	dump2, _ := strconv.ParseInt(dump, 0, 16)

	return dump2
}

package helpers

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/gin-gonic/gin"
)

type UrlHelper struct{}

func (u *UrlHelper) GenerateUrl(c *gin.Context, linkId string) string {

	url := c.Request.Host

	return fmt.Sprintf("%s/service/download/%s", url, linkId)

}

func (u *UrlHelper) GenerateId() string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()

	return str

}

package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagString"
	"time"
)

var defaultTotpSecret = ""

func main() {
	secretStr := flagString.NewDefault("secret", "E24K3WEYMZQZ74YC", "the default secret for fix secret api").BindCmd()
	portStr := flagString.NewDefault("port", "8081", "default port usage for totp generation").BindCmd()
	flag.Parse()
	defaultTotpSecret = secretStr.Value()
	fmt.Println("Fix secret key: " + defaultTotpSecret)

	router := gin.Default()

	router.GET("/api/totp/fix", func(context *gin.Context) {
		totpVal, err := totp.GenerateCode(defaultTotpSecret, time.Now())
		if err != nil {
			panic(err)
		}
		context.JSON(200, gin.H{
			"totp": totpVal,
		})
	})

	router.POST("/api/totp/fix/:key", func(context *gin.Context) {
		key := context.Param("key")
		fmt.Println("Set fix secret key: " + defaultTotpSecret)
		if key == "" {
			context.JSON(400, gin.H{
				"result": "fail",
			})
			return
		}

		defaultTotpSecret = key
		context.JSON(200, gin.H{
			"result": "success",
		})
	})

	router.GET("/api/totp/dynamic/:key/now", func(context *gin.Context) {
		key := context.Param("key")
		totpVal, err := totp.GenerateCode(key, time.Now())
		if err != nil {
			panic(err)
		}
		context.JSON(200, gin.H{
			"totp": totpVal,
		})
	})
	err := router.Run(":" + portStr.Value())
	if err != nil {
		panic(err)
	}

}

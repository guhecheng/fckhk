package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
	"os"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func main()  {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		fmt.Println("hello world")
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Posts",
		})
	})

	router.Static("/upload", "./upload")

	router.GET("/qrCode/getQrcode", func (c *gin.Context) {
		qrcodeContent := c.Query("qrcode_content")
		fmt.Println(qrcodeContent)
		cs, _ := code128.Encode(qrcodeContent)
		fmt.Println(cs)
		path := "./upload/qr.png"
		file, _ := os.Create(path)
		fmt.Println(file)
		defer file.Close()

		qrCode, _ := barcode.Scale(cs, 350, 70)
		fmt.Println(qrCode)
		png.Encode(file, qrCode)

		data := map[string]interface{}{
			"code": 0,
			"qrCode_url" : "/upload/qr.png",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	router.Run(":8080")
}

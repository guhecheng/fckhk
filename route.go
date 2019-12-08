package main

import (
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
	"os"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func main()  {
	router := gin.Default()
	router.LoadHTMLGlob("views/**/*")
	router.GET("/index/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/qrCode/getQrcode", func (c *gin.Context) {
		qrcodeContent := c.Query("qrcode_content")
		cs, _ := code128.Encode(qrcodeContent)
		file, _ := os.Create("qr.png")
		defer file.Close()

		qrCode, _ := barcode.Scale(cs, 350, 70)
		png.Encode(file, qrCode)
	})

	router.Run(":80")
}

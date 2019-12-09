package main

import (
	"fmt"
	"github.com/boombuler/barcode/qr"
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
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl", gin.H{
			"title": "FckHk",
		})
	})

	router.Static("/upload", "./upload")
	router.Static("/static", "./static")

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

	router.GET("/qrCode/getCode", func (c *gin.Context) {
		qrcodeContent := c.Query("content")
		fmt.Println(qrcodeContent)
		qrCode, _ := qr.Encode(qrcodeContent, qr.M, qr.Auto)
		fmt.Println(qrCode)
		path := "./upload/qr_code.png"
		file, _ := os.Create(path)
		fmt.Println(file)
		defer file.Close()

		qrCode, _ = barcode.Scale(qrCode, 256, 256)
		fmt.Println(qrCode)
		png.Encode(file, qrCode)

		data := map[string]interface{}{
			"code": 0,
			"qrCode_url" : "/upload/qr_code.png",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	router.Run(":8080")
}

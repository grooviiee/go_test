package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver"
)

// reference -> https://opendart.fss.or.kr/guide/detail.do?apiGrpCd=DS001&apiId=2019001
const apikey string = "1436a1b25f6279f51ed0daf3719a04f1f0f2a333"
const query string = "https://opendart.fss.or.kr/api/list.json?crtfc_key=" + apikey + "&corp_code=aasdd&bgn_de=20200117&end_de=20200117&corp_cls=Y&page_no=1&page_count=10"
const queryCorpCode string = "https://opendart.fss.or.kr/api/corpCode.xml?crtfc_key=" + apikey

//GET	https://opendart.fss.or.kr/api/list.json

func setupRouter(data string) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, data)
	})
	return r
}

func main() {

	resp, err := http.Get(queryCorpCode)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("data.zip", data, 0644)

	err = archiver.Unarchive("test.zip", "Unarchive_output")
	if err != nil {
		panic(err)
	}
	err = archiver.Walk("test.zip", func(f archiver.File) error {
		fmt.Println(f.Name(), f.Size())
		return nil
	})
	if err != nil {
		panic(err)
	}
}
	r := setupRouter(string(data))
	r.Run(":8080")
}

// package main

// import "fmt"

// func main() {
// 	fmt.Print("Hello World")
// }
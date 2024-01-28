package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func fetchHTML(url string) (string,error) {

	resp, err := http.Get(url)

	if err != nil {
		
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	 if err != nil {
		return "", err
	}

	html := string(body)

	return html,nil
}

func getPage(c *gin.Context){
	url, ok := c.GetQuery("url")


	if ok == false{
		c.String(http.StatusBadRequest, "URL malformed")
		return 
	}


	html,err := fetchHTML(url)

	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error fetching HTML: %s", err))
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}


func main(){
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("fetch",getPage)
	router.Run("localhost:8080")
}
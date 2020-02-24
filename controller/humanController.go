package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
const WEBSERVERPORT = ":8180"

func GetHandler(c *gin.Context)  {
	c.File("view/humanForm.html")
	c.Writer.WriteHeader(http.StatusOK)
}
func PostHandler(c *gin.Context) {
	fname := c.PostForm("fname")
	lname:=c.PostForm("lname")
	age,err:=strconv.Atoi(c.PostForm("Age"))
	if err!=nil{
		gin.Logger()
		return
	}
	c.JSON(200, gin.H{
		"fname": fname,
		"lname": lname,
		"age"  : age,
	})

}

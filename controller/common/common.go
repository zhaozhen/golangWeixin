package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
	"io"
)

func Fileupload(c *gin.Context){
	//得到上传的文件
	//c.Request.MultipartForm()

	//d:=c.Request.MultipartForm()
	file, header, err := c.Request.FormFile("file") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	//文件的名称
	filename := header.Filename

	fmt.Println(file, err, filename)
	//创建文件
	out, err := os.Create("static/"+filename)
	//注意此处的 static/uploadfile/ 不是/static/uploadfile/
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return
	}

	//c.String(http.StatusCreated, "upload successful")


	filepath := "http://127.0.0.1:8023/api/file/" + filename
	c.JSON(http.StatusOK, gin.H{
		"filePath": filepath,
	})

}

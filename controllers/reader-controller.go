package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"soundcase/worker"
)

func ReadFile(c *gin.Context) {
	dirName := c.Query("dir")

	fileNames := worker.ReadFilesFromDirectory(dirName)

	c.JSON(http.StatusOK, gin.H{
		"fileNames": fileNames,
	})
	return
}

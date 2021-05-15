package storage

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/RatelData/ratel-drive-core/common/util/config"
	"github.com/gin-gonic/gin"
)

type FileInfo struct {
	FileName string `json:"file_name"`
	IsDir    bool   `json:"is_dir"`
}

// Files godoc
// @tags files
// @summary Retrieve files information
// @description get files by specified path
// @accept  json
// @produce json
// @param   path path string true "the path that you want to list the files"
// @success 200 {object} misc.JSONResult{data=[]FileInfo}
// @failure 400 {object} error.ErrorResult{error=string}
// @router /api/storage/files [get]
func QueryFilesHandler(c *gin.Context) {
	rootDir := config.GetStorageConfig().StorageRootDir
	path := c.Query("path")

	files, err := ioutil.ReadDir(rootDir + "/" + path)
	if err != nil {
		log.Panic(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var fiArray []FileInfo
	for _, fi := range files {
		fiArray = append(fiArray, FileInfo{fi.Name(), fi.IsDir()})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": fiArray,
	})
}

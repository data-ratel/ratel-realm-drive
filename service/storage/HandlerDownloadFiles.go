package storage

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RatelData/ratel-drive-core/common/util"
	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver"
)

type DownloadParams struct {
	FilePaths []string `json:"file_paths"`
}

// Files godoc
// @tags files
// @summary Download a single file
// @description Download a single file by the specified file path
// @accept  json
// @produce octet-stream
// @param   file path string true "the file that you want to download"
// @success 200 {file} binary
// @failure 400 {object} types.ErrorResult{error=string}
// @router /api/storage/download [get]
func DownloadSingleFileHandler(c *gin.Context) {
	rootDir := util.GetStorageConfig().StorageRootDir
	path := c.Query("file")

	targetFilePath := fmt.Sprintf("%s/%s", rootDir, path)
	if util.IsPathExists(targetFilePath) {
		c.File(targetFilePath)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file is not existed",
		})
	}
}

// Files godoc
// @tags files
// @summary Download multiple files
// @description Download files by the specified file paths, will be zipped
// @accept  json
// @produce octet-stream
// @param   files body DownloadParams true "the files that you want to download"
// @success 200 {file} binary
// @failure 400 {object} types.ErrorResult{error=string}
// @failure 500 {object} types.ErrorResult{error=string}
// @router /api/storage/download [post]
func DownloadMultiFilesHandler(c *gin.Context) {
	var params DownloadParams
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Panicln(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(params.FilePaths) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no files to download",
		})
		return
	}

	// if download multiple files or directories
	// zip them to a temporary file
	// serve this zipped file
	tempDir := util.GetStorageConfig().TempDir
	targetFilePath := fmt.Sprintf("%s/archive-%d.zip", tempDir, time.Now().Unix())
	defer os.Remove(targetFilePath)

	rootDir := util.GetStorageConfig().StorageRootDir
	var sourceFilesPaths []string
	for _, path := range params.FilePaths {
		sourceFilesPaths = append(sourceFilesPaths, rootDir+"/"+path)
	}

	err := archiver.Archive(sourceFilesPaths, targetFilePath)
	if err != nil {
		log.Panicln(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal issue",
		})
		return
	}

	if util.IsPathExists(targetFilePath) {
		c.File(targetFilePath)
	} else {
		log.Panicln("[WARN] Something wrong while creating zipped file for downloading")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file is not existed",
		})
	}
}

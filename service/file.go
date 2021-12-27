package service

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/littlelittlelittleboy/scriptkilling/constants"
	uuid "github.com/satori/go.uuid"
)

type FileService struct{}

func (service *FileService) Save(file multipart.File, header *multipart.FileHeader, scriptName string) (string, constants.UploadFileStatus) {
	fileExt := path.Ext(header.Filename)
	fileName := strings.ReplaceAll(uuid.NewV4().String(), "-", "_")

	fileBasePath := fmt.Sprintf("%s/%s", constants.FILE_BASE_PATH, strings.ReplaceAll(scriptName, "/", "_"))
	if _, err := os.Stat(fileBasePath); err != nil {
		err := os.MkdirAll(fileBasePath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	filePath := fmt.Sprintf("%s/%s%s", fileBasePath, fileName, fileExt)
	if _, err := os.Stat(filePath); err == nil {
		return "", constants.UPLOAD_FILE_EXIST
	}

	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
		return "", constants.UPLOAD_FILE_SAVE_ERROR
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		return "", constants.UPLOAD_FILE_SAVE_ERROR
	}

	return filePath, constants.UPLOAD_FILE_SUCCESS
}

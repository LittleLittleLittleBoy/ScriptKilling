package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/scriptkilling/common/response"
	"github.com/littlelittlelittleboy/scriptkilling/common/vo"
	"github.com/littlelittlelittleboy/scriptkilling/constants"
)

type FileUploadControler struct{}

func (controller *FileUploadControler) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		response.GenerErrorResponse(c, nil,
			int(constants.UPLOAD_FILE_GET_FROM_REQUEST_ERROR),
			"can not parse file from request")
		return
	}

	scriptName := c.GetHeader(constants.HEADER_SCRIPT_NAME)

	fileService := ServiceInstance.FileService

	filePath, status := fileService.Save(file, header, scriptName)
	if status != constants.UPLOAD_FILE_SUCCESS {
		response.GenerErrorResponse(c, nil,
			int(constants.UPLOAD_FILE_SAVE_ERROR),
			"save to disk error")
		return
	}

	responseData := &vo.UploadResponse{
		Path: filePath,
	}

	response.GenerResponse(c, responseData)
}

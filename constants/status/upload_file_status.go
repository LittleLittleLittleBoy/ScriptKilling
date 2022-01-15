package status

type UploadFileStatus int

const (
	UPLOAD_FILE_SUCCESS                UploadFileStatus = 0
	UPLOAD_FILE_GET_FROM_REQUEST_ERROR UploadFileStatus = -1
	UPLOAD_FILE_SAVE_ERROR             UploadFileStatus = -2
	UPLOAD_FILE_EXIST                  UploadFileStatus = -3
)

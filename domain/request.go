package domain

type StartRequest struct {
}

type UploadRequest struct {
}

type ProcessRequest struct {
	TaskID string      `json:"task" validate:"required"`
	Tool   string      `json:"tool" validate:"required"`
	Files  FileRequest `json:"files" validate:"required"`
}
type DownloadRequest struct {
}

type FileRequest struct {
	ServerFileName string `json:"server_filename" validate:"required"`
	Filename       string `json:"filename" validate:"required"`
}

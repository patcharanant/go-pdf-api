package domain

type StartResponse struct {
}

type UploadResponse struct {
}

type ProcessResponse struct {
	DownloadFileName string `json:"download_filename"`
	FileSize         string `json:"file_size"`
	OutputFileSize   string `json:"output_filesize"`
	OutputFileNumber string `json:"output_filenumber"`
	OutputExtensions string `json:"output_extensions"`
	Timer            string `json:"timer"`
	Status           string `json:"status"`
}

type DownloadResponse struct {
}

type ResponseError struct {
	Message string `json:"message"`
}

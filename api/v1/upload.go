package v1

// FileInfo 上传的文件信息
type FileInfo struct {
	FileName string `json:"fileName" dc:"文件名"`  // 文件名
	FileSize uint64 `json:"fileSize" dc:"文件大小"` // 文件大小
	FileUrl  string `json:"fileUrl" dc:"文件Url"` // 文件Url
	FileType string `json:"fileType" dc:"文件类型"` // 文件类型
}

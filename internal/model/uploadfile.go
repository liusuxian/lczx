package model

// UploadFileInfo 上传的文件信息
type UploadFileInfo struct {
	FileName      string `json:"fileName" dc:"文件名"`          // 文件名
	FileSize      int64  `json:"fileSize" dc:"文件大小"`         // 文件大小
	OriginFileUrl string `json:"originFileUrl" dc:"原始文件Url"` // 原始文件Url
	PdfFileUrl    string `json:"pdfFileUrl" dc:"pdf文件Url"`   // pdf文件Url
	FileType      string `json:"fileType" dc:"文件类型"`         // 文件类型
}

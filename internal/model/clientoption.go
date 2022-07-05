package model

// ClientOption 客户端选项
type ClientOption struct {
	Name  string `json:"name" dc:"选项显示名"` // 选项显示名
	Value string `json:"value" dc:"选项值"`  // 选项值
}

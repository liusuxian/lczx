package service

type sAuthRule struct{}

var (
	insAuthRule = sAuthRule{}
)

// AuthRule 验证规则服务
func AuthRule() *sAuthRule {
	return &insAuthRule
}

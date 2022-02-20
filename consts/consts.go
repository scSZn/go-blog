package consts

type TraceKey string

// 环境信息
const (
	EnvDev  = "dev"
	EnvProd = "prod"
)

// 日志相关信息
const (
	LogTraceKey   = TraceKey("trace")
	LogTimeLayout = "2006-01-02 15:04:05.000-07:00"
)

// 删除状态
const (
	DelStatus   = 1 // 删除标识
	NoDelStatus = 0 // 未删除
)

// 状态枚举值
const (
	StatusDraft   = 1 // 草稿（博客特有）
	StatusEnable  = 2 // 可用（博客和标签都可用）
	StatusDisable = 3 // 禁用（博客和标签都可用）
	StatusDeleted = 4 // 删除（博客和标签都可用）
)

// TokenHeaderKey 登录Token
const TokenHeaderKey = "token"

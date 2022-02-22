package consts

type TraceKey string

// 格式化日期字符串
const TimeFormatLayout = "2006-01-02 15:04:05"

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
	StatusDraft        = 0  // 草稿（博客特有）
	StatusWaitingAudit = 10 // 待审核（博客特有）
	StatusEnable       = 20 // 可用（博客和标签都可用）
	StatusReject       = 30 // 禁用（博客特有）
	StatusDisable      = 40 // 禁用（博客和标签都可用）
	StatusDeleted      = 50 // 删除（博客和标签都可用）
)

// TokenHeaderKey 登录Token
const TokenHeaderKey = "token"

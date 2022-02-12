package consts

type TraceKey string

const (
	EnvDev  = "dev"
	EnvProd = "prod"

	// 存放在上下文信息中的trace key
	LogTraceKey   = TraceKey("trace")
	LogTimeLayout = "2006-01-02 15:04:05.000-07:00"
)

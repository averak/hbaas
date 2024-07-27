package interceptor

import (
	"connectrpc.com/connect"
)

func New() []connect.Interceptor {
	// 上から順に実行される。
	// ただし、connect.UnaryFunc の戻り値に対して処理する interceptor は下から順番に実行される。
	return []connect.Interceptor{
		NewTraceInterceptor(),
		NewAccessLogInterceptor(),
		NewFixPreconditionParamsInterceptor(),
	}
}
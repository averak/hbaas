package error_response

import (
	"testing"

	"connectrpc.com/connect"
	"github.com/averak/hbaas/protobuf/api/api_errors"
	"github.com/averak/hbaas/testutils/testconnect"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args[CODE CodeType] struct {
		code     CODE
		severity api_errors.ErrorSeverity
		message  string
	}
	type testCase[CODE CodeType] struct {
		name string
		args args[CODE]
		then func(t *testing.T, got Error)
	}
	tests := []testCase[api_errors.ErrorCode_Common]{
		{
			name: "[ケース1] COMMON_UNKNOWN",
			args: args[api_errors.ErrorCode_Common]{
				code: api_errors.ErrorCode_COMMON_UNKNOWN,
			},
			then: func(t *testing.T, got Error) {
				assert.Equal(t, connect.CodeUnknown, got.connectErr.Code())

				wantDetail := &api_errors.ErrorDetail{
					ErrorCode:         int64(api_errors.ErrorCode_COMMON_UNKNOWN),
					ErrorHandlingType: api_errors.ErrorHandlingType_ERROR_HANDLING_TYPE_TEMPORARY,
				}
				assert.EqualExportedValues(t, wantDetail, testconnect.GetErrorDetail(got.connectErr))
			},
		},
		{
			name: "[ケース2] METHOD_COMMON_INVALID_SESSION",
			args: args[api_errors.ErrorCode_Common]{
				code: api_errors.ErrorCode_COMMON_INVALID_SESSION,
			},
			then: func(t *testing.T, got Error) {
				assert.Equal(t, connect.CodeUnauthenticated, got.connectErr.Code())

				wantDetail := &api_errors.ErrorDetail{
					ErrorCode:         int64(api_errors.ErrorCode_COMMON_INVALID_SESSION),
					ErrorHandlingType: api_errors.ErrorHandlingType_ERROR_HANDLING_TYPE_RECOVERABLE,
				}
				assert.EqualExportedValues(t, wantDetail, testconnect.GetErrorDetail(got.connectErr))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := New(tt.args.code, tt.args.severity, tt.args.message)
			tt.then(t, err)
		})
	}
}

package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

type RetryableObj struct {
	fn        func() error
	histories []time.Time
}

func NewRetryableObj(fn func(n int) error) *RetryableObj {
	obj := &RetryableObj{
		histories: make([]time.Time, 0),
	}
	obj.fn = func() error {
		err := fn(len(obj.histories))
		obj.histories = append(obj.histories, time.Now())
		return err
	}
	return obj
}

func TestDo(t *testing.T) {
	var testErr = errors.New("test error")

	type args struct {
		opts []Option
	}
	type when struct {
		args args
		obj  *RetryableObj
	}
	tests := []struct {
		name string
		when when
		then func(t *testing.T, obj *RetryableObj, begin time.Time, err error)
	}{
		{
			name: "デフォルト設定 => 100ms の固定バックオフを最大 5 回行う",
			when: when{
				args: args{},
				obj: NewRetryableObj(func(n int) error {
					return testErr
				}),
			},
			then: func(t *testing.T, obj *RetryableObj, begin time.Time, err error) {
				assert.ErrorIs(t, err, testErr)

				wantHistories := []time.Time{
					begin,
					begin.Add(100 * time.Millisecond),
					begin.Add(200 * time.Millisecond),
					begin.Add(300 * time.Millisecond),
				}
				if diff := cmp.Diff(wantHistories, obj.histories, cmpopts.EquateApproxTime(10*time.Millisecond)); diff != "" {
					t.Errorf("(-want, +got)\n%s", diff)
				}
			},
		},
		{
			name: "最大リトライ回数到達前に成功する => エラーを返さない",
			when: when{
				args: args{
					opts: []Option{
						WithMaxRetries(10),
					},
				},
				obj: NewRetryableObj(func(n int) error {
					if n >= 2 {
						return nil
					}
					return testErr
				}),
			},
			then: func(t *testing.T, obj *RetryableObj, begin time.Time, err error) {
				assert.NoError(t, err)

				wantHistories := []time.Time{
					begin,
					begin.Add(100 * time.Millisecond),
					begin.Add(200 * time.Millisecond),
				}
				if diff := cmp.Diff(wantHistories, obj.histories, cmpopts.EquateApproxTime(10*time.Millisecond)); diff != "" {
					t.Errorf("(-want, +got)\n%s", diff)
				}
			},
		},
		{
			name: "WithMaxRetries => 最大リトライ回数を設定できる",
			when: when{
				args: args{
					opts: []Option{
						WithMaxRetries(2),
					},
				},
				obj: NewRetryableObj(func(n int) error {
					return testErr
				}),
			},
			then: func(t *testing.T, obj *RetryableObj, begin time.Time, err error) {
				assert.ErrorIs(t, err, testErr)
				assert.Len(t, obj.histories, 3)
			},
		},
		{
			name: "WithConstantBackOff => 固定バックオフのリトライ間隔を設定できる",
			when: when{
				args: args{
					opts: []Option{
						WithConstantBackOff(50 * time.Millisecond),
					},
				},
				obj: NewRetryableObj(func(n int) error {
					return testErr
				}),
			},
			then: func(t *testing.T, obj *RetryableObj, begin time.Time, err error) {
				assert.ErrorIs(t, err, testErr)

				wantHistories := []time.Time{
					begin,
					begin.Add(50 * time.Millisecond),
					begin.Add(100 * time.Millisecond),
					begin.Add(150 * time.Millisecond),
				}
				if diff := cmp.Diff(wantHistories, obj.histories, cmpopts.EquateApproxTime(10*time.Millisecond)); diff != "" {
					t.Errorf("(-want, +got)\n%s", diff)
				}
			},
		},
		{
			name: "WithExponentialBackoff => 指数バックオフでリトライできる",
			when: when{
				args: args{
					opts: []Option{
						WithExponentialBackoff(50*time.Millisecond, 300*time.Millisecond),
					},
				},
				obj: NewRetryableObj(func(n int) error {
					return testErr
				}),
			},
			then: func(t *testing.T, obj *RetryableObj, begin time.Time, err error) {
				assert.ErrorIs(t, err, testErr)

				wantHistories := []time.Time{
					begin,
					// delay 50ms
					begin.Add(50 * time.Millisecond),
					// delay 50ms + 100ms
					begin.Add(150 * time.Millisecond),
					// delay 50ms + 100ms + 200ms
					begin.Add(350 * time.Millisecond),
					// delay 50ms + 100ms + 200ms + 300ms
					begin.Add(650 * time.Millisecond),
				}
				if diff := cmp.Diff(wantHistories, obj.histories, cmpopts.EquateApproxTime(10*time.Millisecond)); diff != "" {
					t.Errorf("(-want, +got)\n%s", diff)
				}
			},
		},
		{
			name: "WithExponentialBackoff と WithMaxRetries を組み合わせる => 最大リトライ回数が WithMaxRetries で指定した値になる",
			when: when{
				args: args{
					opts: []Option{
						WithExponentialBackoff(50*time.Millisecond, 300*time.Millisecond),
						WithMaxRetries(2),
					},
				},
				obj: NewRetryableObj(func(n int) error {
					return testErr
				}),
			},
			then: func(t *testing.T, obj *RetryableObj, begin time.Time, err error) {
				assert.ErrorIs(t, err, testErr)

				wantHistories := []time.Time{
					begin,
					// delay 50ms
					begin.Add(50 * time.Millisecond),
					// delay 50ms + 100ms
					begin.Add(150 * time.Millisecond),
					// リトライ間隔が max に到達していないが、WithMaxRetries(2) が指定されたため、ここで終了する。
				}
				if diff := cmp.Diff(wantHistories, obj.histories, cmpopts.EquateApproxTime(10*time.Millisecond)); diff != "" {
					t.Errorf("(-want, +got)\n%s", diff)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			begin := time.Now()
			err := Do(context.Background(), tt.when.obj.fn, tt.when.args.opts...)
			tt.then(t, tt.when.obj, begin, err)
		})
	}
}

func TestWithConstantBackOff(t *testing.T) {
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		then func(t *testing.T, got Option)
	}{
		{
			name: "指定した遅延時間で固定の遅延時間を返す",
			args: args{
				d: 1 * time.Second,
			},
			then: func(t *testing.T, got Option) {
				opts := &Options{}
				got(opts)

				assert.Equal(t, opts.backoff(0), 1*time.Second)
				assert.Equal(t, opts.backoff(1), 1*time.Second)
				assert.Equal(t, opts.backoff(2), 1*time.Second)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WithConstantBackOff(tt.args.d)
			tt.then(t, got)
		})
	}
}

func TestWithExponentialBackoff(t *testing.T) {
	type args struct {
		base time.Duration
		max  time.Duration
	}
	tests := []struct {
		name string
		args args
		then func(t *testing.T, got Option)
	}{
		{
			name: "N回目のリトライ間隔 > 最大遅延時間の場合 => N回リトライする",
			args: args{
				base: 1 * time.Second,
				max:  10 * time.Second,
			},
			then: func(t *testing.T, got Option) {
				opts := &Options{}
				got(opts)

				assert.Equal(t, 5, opts.maxRetries)
				assert.Equal(t, opts.backoff(0), 1*time.Second)
				assert.Equal(t, opts.backoff(1), 2*time.Second)
				assert.Equal(t, opts.backoff(2), 4*time.Second)
				assert.Equal(t, opts.backoff(3), 8*time.Second)
				assert.Equal(t, opts.backoff(4), 10*time.Second)
			},
		},
		{
			name: "N回目のリトライ間隔 == 最大遅延時間の場合 => N回リトライする",
			args: args{
				base: 1 * time.Second,
				max:  8 * time.Second,
			},
			then: func(t *testing.T, got Option) {
				opts := &Options{}
				got(opts)

				assert.Equal(t, 4, opts.maxRetries)
				assert.Equal(t, opts.backoff(0), 1*time.Second)
				assert.Equal(t, opts.backoff(1), 2*time.Second)
				assert.Equal(t, opts.backoff(2), 4*time.Second)
				assert.Equal(t, opts.backoff(3), 8*time.Second)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WithExponentialBackoff(tt.args.base, tt.args.max)
			tt.then(t, got)
		})
	}
}

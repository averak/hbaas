package retry

import (
	"context"
	"math"
	"math/rand/v2"
	"time"
)

type (
	Option  func(*Options)
	Options struct {
		maxRetries int
		backoff    func(n int) time.Duration
		jitter     func() time.Duration
	}
)

// Do は、fn が成功するまで繰り返し実行します。
// デフォルトだと、100ms の固定バックオフで最大 3 回リトライします。
func Do(ctx context.Context, fn func() error, opts ...Option) error {
	options := &Options{
		maxRetries: 3,
		backoff: func(n int) time.Duration {
			return 100 * time.Millisecond
		},
		jitter: func() time.Duration {
			return 0
		},
	}
	for _, opt := range opts {
		opt(options)
	}

	var err error
	for i := range options.maxRetries + 1 {
		err = fn()
		if err == nil {
			return nil
		}

		if i != options.maxRetries {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(options.backoff(i) + options.jitter()):
			}
		}
	}
	return err
}

// WithMaxRetries は、最大リトライ回数を設定するオプションです。
func WithMaxRetries(n int) Option {
	return func(o *Options) {
		o.maxRetries = n
	}
}

// WithJitter は、リトライ間隔をランダムに変動させるオプションです。
// これにより、なんらかの状況によって多数のクライアントが同期され、再試行が同時に実行される状況を避けることができます。
func WithJitter(max time.Duration) Option {
	return func(o *Options) {
		o.jitter = func() time.Duration {
			return time.Duration(rand.IntN(int(max)))
		}
	}
}

// WithConstantBackOff は、定数バックオフを行うオプションです。
func WithConstantBackOff(d time.Duration) Option {
	return func(o *Options) {
		o.backoff = func(n int) time.Duration {
			return d
		}
	}
}

// WithExponentialBackoff は、指数バックオフを行うオプションです。
// min((2^n * base), max) でリトライ間隔を計算し、リトライ間隔が max に到達した時点でリトライを打ち切ります。
func WithExponentialBackoff(base time.Duration, max time.Duration) Option {
	return func(o *Options) {
		o.maxRetries = int(math.Floor(math.Log2(float64(max/base)))) + 1
		if base*(1<<uint(o.maxRetries-1)) < max {
			o.maxRetries++
		}
		o.backoff = func(n int) time.Duration {
			delay := base * (1 << uint(n))
			if delay > max {
				return max
			}
			return delay
		}
	}
}

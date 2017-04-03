package benchmarks

import (
	"io/ioutil"
	"testing"
	"time"

	logkit "github.com/go-kit/kit/log"
	logzap "go.uber.org/zap"
)

func BenchmarkZap(b *testing.B) {
	logger := logzap.NewNop()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Test",
				logzap.Int("attempt", 1),
				logzap.String("url", "string"),
				logzap.Duration("backoff", time.Second),
			)
		}
	})
}

func BenchmarkZapSugared(b *testing.B) {
	logger := logzap.NewNop().Sugar()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("attempt", 1, "url", "string", "backoff", time.Second)
		}
	})
}

func BenchmarkGoKit(b *testing.B) {
	logger := logkit.NewJSONLogger(logkit.NewSyncWriter(ioutil.Discard))

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Log("attempt", 1, "url", "string", "backoff", time.Second)
		}
	})
}

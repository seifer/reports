package benchmarks

import (
	"os"
	"testing"
	"time"

	hpwriter "github.com/seifer/go-hpwriter"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// var w = zapcore.AddSync(&zaptest.Discarder{})
var w = zapcore.AddSync(os.Stderr)
var wShardedBuffers = zapcore.AddSync(hpwriter.NewThroughShardedBuffers(w, 32))

func BenchmarkZapFree(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.DebugLevel,
	))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info(getMessage(0),
				zap.Int("int", 1),
				zap.Int64("int64", 2),
				zap.Float64("float", 3.0),
				zap.String("string", "four!"),
				zap.Bool("bool", true),
				zap.Time("time", time.Unix(0, 0)),
				zap.Error(errExample),
				zap.Duration("duration", time.Second),
				zap.Object("user-defined type", _jane),
				zap.String("another string", "done!"),
			)
		}
	})
}

func BenchmarkZapLock(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(w),
		zap.DebugLevel,
	))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info(getMessage(0),
				zap.Int("int", 1),
				zap.Int64("int64", 2),
				zap.Float64("float", 3.0),
				zap.String("string", "four!"),
				zap.Bool("bool", true),
				zap.Time("time", time.Unix(0, 0)),
				zap.Error(errExample),
				zap.Duration("duration", time.Second),
				zap.Object("user-defined type", _jane),
				zap.String("another string", "done!"),
			)
		}
	})
}

func BenchmarkZapShardedBuffers(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		wShardedBuffers,
		zap.DebugLevel,
	))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info(getMessage(0),
				zap.Int("int", 1),
				zap.Int64("int64", 2),
				zap.Float64("float", 3.0),
				zap.String("string", "four!"),
				zap.Bool("bool", true),
				zap.Time("time", time.Unix(0, 0)),
				zap.Error(errExample),
				zap.Duration("duration", time.Second),
				zap.Object("user-defined type", _jane),
				zap.String("another string", "done!"),
			)
		}
	})
}

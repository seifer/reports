package benchmarks

import (
	"os"
	"testing"
	"time"

	hpwriter "github.com/seifer/go-hpwriter"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var w, wLock, wChannel, wShardedBuffers zapcore.WriteSyncer

func init() {
	var err1 error

	w, err1 = os.OpenFile("free", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if err1 != nil {
		panic(err1)
	}

	file1, err := os.OpenFile("lock", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	wLock = zapcore.Lock(file1)

	file2, err := os.OpenFile("channel", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	wChannel = zapcore.AddSync(hpwriter.NewThroughChannel(file2))

	file3, err := os.OpenFile("buffers", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	wShardedBuffers = zapcore.AddSync(hpwriter.NewThroughShardedBuffers(file3, 32))
}

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
		wLock,
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

func BenchmarkZapChannel(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		wChannel,
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

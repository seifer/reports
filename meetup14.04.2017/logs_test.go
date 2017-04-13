package benchmarks

import (
	"testing"

	"go.uber.org/zap"
)

func BenchmarkAddingFields(b *testing.B) {
	b.Run("Zap", func(b *testing.B) {
		logger := newZapLogger(zap.DebugLevel)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0), fakeFields()...)
			}
		})
	})

	b.Run("Zap.Sugar", func(b *testing.B) {
		logger := newZapLogger(zap.DebugLevel).Sugar()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Infow(getMessage(0), fakeSugarFields()...)
			}
		})
	})

	b.Run("apex/log", func(b *testing.B) {
		logger := newApexLog()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.WithFields(fakeApexFields()).Info(getMessage(0))
			}
		})
	})

	b.Run("go-kit/kit/log", func(b *testing.B) {
		logger := newKitLog()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Log(fakeSugarFields()...)
			}
		})
	})

	b.Run("inconshreveable/log15", func(b *testing.B) {
		logger := newLog15()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0), fakeSugarFields()...)
			}
		})
	})

	b.Run("sirupsen/logrus", func(b *testing.B) {
		logger := newLogrus()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.WithFields(fakeLogrusFields()).Info(getMessage(0))
			}
		})
	})

	b.Run("go.pedge.io/lion", func(b *testing.B) {
		logger := newLion()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.WithFields(fakeLogrusFields()).Infof(getMessage(0))
			}
		})
	})
}

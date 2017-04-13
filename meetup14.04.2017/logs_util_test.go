package benchmarks

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	kitlog "github.com/go-kit/kit/log"
	"github.com/inconshreveable/log15"
	lion "github.com/peter-edge/lion-go"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

var (
	_messages  = fakeMessages(1000)
	errExample = errors.New("fail")
)

func fakeMessages(n int) []string {
	messages := make([]string, n)
	for i := range messages {
		messages[i] = fmt.Sprintf("Test logging, but use a somewhat realistic message length. (#%v)", i)
	}
	return messages
}

func getMessage(iter int) string {
	return _messages[iter%1000]
}

type user struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (u user) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", u.Name)
	enc.AddString("email", u.Email)
	enc.AddInt64("created_at", u.CreatedAt.UnixNano())
	return nil
}

var _jane = user{
	Name:      "Jane Doe",
	Email:     "jane@test.com",
	CreatedAt: time.Date(1980, 1, 1, 12, 0, 0, 0, time.UTC),
}

func newZapLogger(lvl zapcore.Level) *zap.Logger {
	// use the canned production encoder configuration
	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	return zap.New(zapcore.NewCore(
		enc,
		&zaptest.Discarder{},
		lvl,
	))
}

func fakeFields() []zapcore.Field {
	return []zapcore.Field{
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
	}
}

func fakeSugarFields() []interface{} {
	return []interface{}{
		"int", 1,
		"int64", 2,
		"float", 3.0,
		"string", "four!",
		"bool", true,
		"time", time.Unix(0, 0),
		"error", errExample,
		"duration", time.Second,
		"user-defined type", _jane,
		"another string", "done!",
	}
}

func newApexLog() *log.Logger {
	return &log.Logger{
		Handler: json.New(ioutil.Discard),
		Level:   log.DebugLevel,
	}
}

func fakeApexFields() log.Fields {
	return log.Fields{
		"int":               1,
		"int64":             int64(1),
		"float":             3.0,
		"string":            "four!",
		"bool":              true,
		"time":              time.Unix(0, 0),
		"error":             errExample.Error(),
		"duration":          time.Second,
		"user-defined type": _jane,
		"another string":    "done!",
	}
}

func newKitLog() kitlog.Logger {
	return kitlog.NewJSONLogger(kitlog.NewSyncWriter(ioutil.Discard))
}

func newLog15() log15.Logger {
	logger := log15.New()
	logger.SetHandler(log15.StreamHandler(ioutil.Discard, log15.JsonFormat()))
	return logger
}

func newLogrus() *logrus.Logger {
	return &logrus.Logger{
		Out:       ioutil.Discard,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}

func fakeLogrusFields() logrus.Fields {
	return logrus.Fields{
		"int":               1,
		"int64":             int64(1),
		"float":             3.0,
		"string":            "four!",
		"bool":              true,
		"time":              time.Unix(0, 0),
		"error":             errExample.Error(),
		"duration":          time.Second,
		"user-defined type": _jane,
		"another string":    "done!",
	}
}

func newLion() lion.Logger {
	return lion.NewLogger(lion.NewJSONWritePusher(ioutil.Discard))
}

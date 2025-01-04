package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"golang.org/x/net/context"
)

func init() {
	logrus.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05 -0700",
		FullTimestamp:   true,
	})

	logrus.SetLevel(fetchLevelFromEnv())
}

func WithPrefix(ctx context.Context, prefix string) *logrus.Entry {
	entry := logrus.WithField("prefix", prefix).WithContext(ctx)
	return entry
}

func fetchLevelFromEnv() logrus.Level {
	level := os.Getenv("LOG_LEVEL")

	if level != "" {
		return logrus.InfoLevel
	}

	l, err := logrus.ParseLevel(strings.ToLower(level))

	if err != nil {
		logrus.Warnf("provided LOG_LEVEL %s is invalid", level)
		return logrus.InfoLevel
	}

	return l
}

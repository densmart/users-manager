package pkg

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

func InitLogger() {
	logger := &lumberjack.Logger{
		Filename:   filepath.ToSlash(viper.GetString("logger.file")),
		MaxSize:    viper.GetInt("logger.max_size"),      // MB
		MaxBackups: viper.GetInt("logger.backups_count"), // count
		MaxAge:     viper.GetInt("logger.max_age"),       // days
		Compress:   true,
	}
	multiWriter := io.MultiWriter(os.Stderr, logger)

	logrus.SetReportCaller(true)
	logFormatter := &logrus.TextFormatter{
		ForceColors: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			fname := path.Base(f.File)
			fpath := path.Dir(f.File)
			now := time.Now()

			return "-", fmt.Sprintf(" [%s] %s/%s:%d",
				now.Format("02.01.2006 15:04:05"), fpath, fname, f.Line)
		},
	}
	level := logrus.DebugLevel
	switch viper.GetString("logger.level") {
	case "panic":
		level = logrus.PanicLevel
		break
	case "fatal":
		level = logrus.FatalLevel
		break
	case "error":
		level = logrus.ErrorLevel
		break
	case "warning":
		level = logrus.WarnLevel
		break
	case "info":
		level = logrus.InfoLevel
		break
	case "trace":
		level = logrus.TraceLevel
		break
	}
	logrus.SetFormatter(logFormatter)
	logrus.SetLevel(level)
	logrus.SetOutput(multiWriter)
}

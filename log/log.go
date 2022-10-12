package log

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"sort"
	"strings"
	"time"
)


func Info(args ...interface{}) {
	Log.Info(args...)
}

func Debug(args ...interface{}) {
	Log.Debug(args...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func diyTextFormatter() *logrus.TextFormatter{
	return &logrus.TextFormatter{
		// No color log is required
		DisableColors:   false,
		// Define timestamp format
		TimestampFormat: "06/01/02 15:04:05.999999999",
		// Cancel fields sorting
		DisableSorting: false,
		SortingFunc: func(keys []string) {
			sort.Slice(keys, func(i, j int) bool {
				switch keys[i] {
				case "level":
					return true
				case "module":
					if keys[j] == "level"{
						return false
					}
					return true
				case "method":
					if keys[j] == "module" || keys[j] == "level" {
						return false
					}
					return true
				case "time":
					if keys[j] == "module" || keys[j] == "level" || keys[j] == "method" {
						return false
					}
					return true
				case "msg":
					if keys[j] == "module" || keys[j] == "level" || keys[j] == "method" || keys[j] == "time" {
						return false
					}
					return true
				}
				// If it is other fields, the main field will be automatically placed behind
				switch keys[j] {
					case "level":
						return false
					case "time":
						return false
					case "module":
						return false
					case "method":
						return false
					case "msg":
						return false
					default:
						return strings.Compare(keys[i], keys[j]) == -1
				}
			})
		},
	}
}


// Enable log persistence
// @param life   Log save time
// @param split  After this time, the log will be automatically split into separate files
func EnableLogStorage(logPath string,life,split time.Duration)  {
	baseLogPath := path.Join(logPath)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),      // Generate a soft chain and point to the latest log file
		rotatelogs.WithMaxAge(life),
		rotatelogs.WithRotationTime(split),
	)
	if err != nil {
		panic(err)
	}
	logHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, nil)

	Log.AddHook(logHook)

	for _,logModule := range logModules{
		logModule.AddHook(logHook) //gei
	}
}

//Reinitialize at the specified log level
func InitLogger(logLevel logrus.Level) {
	Log = logrus.New()
	defaultLogLevel = logLevel
	Log.SetLevel(logLevel)
	Log.SetFormatter(diyTextFormatter()) //
}

type nilWriter struct {
}

func (nw *nilWriter) Write(data []byte) (n int, err error) {
	return 0, nil
}

//When building logs, the module log level will be automatically set
//Log persistence will not be started automatically. It will not be persistent until enablelogstorage is called
func BuildLog(funcName string, module LogModule) *logrus.Entry {
	moduleName := string(module)
	if !CheckModuleExists(module) {
		panic("log module "+module+" not have registerd")
	}
	logEntry := logModules[moduleName].WithField("module", strings.ToLower(moduleName))
	//logEntry := logrus.NewEntry(logModules[moduleName]).WithField("module", strings.ToLower(moduleName))
	if funcName != "" {
		logEntry = logEntry.WithField("method", strings.ToLower(funcName))
	}
	return logEntry
}
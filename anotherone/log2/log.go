package log2

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// 颜色
const (
	Black = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan //青色
	Gray
)

var levelMap = map[logrus.Level]string{
	logrus.TraceLevel: "[TRA]",
	logrus.DebugLevel: "[DEB]",
	logrus.InfoLevel:  "[INF]",
	logrus.WarnLevel:  "[WAR]",
	logrus.ErrorLevel: "[ERR]",
	logrus.PanicLevel: "[PAN]",
	logrus.FatalLevel: "[FAT]",
}

type LogFormatter struct {
	opt Options
}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.InfoLevel:
		levelColor = Gray
	case logrus.WarnLevel:
		levelColor = Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = Red
	default:
		levelColor = Blue
	}
	b := &bytes.Buffer{}
	if entry.Buffer != nil {
		b = entry.Buffer
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		dir, err2 := os.Getwd()
		if err2 != nil {
			panic(err2)
		}
		Path, err := filepath.Rel(dir, path.Dir(entry.Caller.File))
		if err != nil {
			panic(err)
		}
		fileVal := fmt.Sprintf(".\\%s:%d", Path+"\\"+path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[4%dm%s\x1b[0m %-40s %s %s\n", timestamp, levelColor, levelMap[entry.Level], fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[4%dm%s\x1b[0m %s\n", timestamp, levelColor, levelMap[entry.Level], entry.Message)
	}
	return b.Bytes(), nil
}

type LogFormatter2 struct {
	Options Options
}

func getCallerInfo(skip int) (info string) {

	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		info = "runtime.Caller() failed"
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // Base函数返回路径的最后一个元素
	return fmt.Sprintf("FuncName:%s, file:%s, line:%d ", funcName, fileName, lineNo)
}

func getFuncName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		panic("runtime.Caller() failed")
	}
	funcName := runtime.FuncForPC(pc).Name()
	return funcName
}

func GetJump(skip int) string {
	_, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		panic("runtime.Caller() failed")
	}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	Path, err := filepath.Rel(dir, file)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_ = path.Base(file) // Base函数返回路径的最后一个元素
	return fmt.Sprintf(".\\%s:%d", Path, lineNo)
}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter2) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = Gray
	case logrus.WarnLevel:
		levelColor = Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = Red
	default:
		levelColor = Blue
	}
	b := &bytes.Buffer{}
	if entry.Buffer != nil {
		b = entry.Buffer
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		//funcVal := entry.Caller.Function

		dir, err2 := os.Getwd()
		if err2 != nil {
			panic(err2)
		}
		Path, err := filepath.Rel(dir, path.Dir(entry.Caller.File))
		if err != nil {
			panic(err)
		}
		_ = fmt.Sprintf(".\\%s:%d", Path+"\\"+path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[1;4%dm%s\x1b[0m ", timestamp, levelColor, levelMap[entry.Level])
		if t.Options.WithJump {
			fmt.Fprintf(b, "%-40s|", GetJump(9))
		}
		if t.Options.WithFunc {
			split := strings.Split(getFuncName(9), ".")
			before := strings.Join(split[:len(split)-1], ".")
			funName := split[len(split)-1]
			fmt.Fprintf(b, " %-40s| ", before+".\x1b[36m"+funName+"\x1b[0m")
		}
		fmt.Fprintf(b, "\"%s\"\n", entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm%s\x1b[0m %s\n", timestamp, levelColor, levelMap[entry.Level], entry.Message)
	}
	return b.Bytes(), nil
}

var (
	myLogger    *logrus.Logger
	debugLogger *logrus.Logger
	infoLogger  *logrus.Logger
	warnLogger  *logrus.Logger
	errorLogger *logrus.Logger
)

func init() {
	myLogger, _ = New(StdConf())
	debugLogger, _ = New(FileConf("./src/lg/logs/debug.log"))
	infoLogger, _ = New(FileConf("./src/lg/logs/info.log"))
	warnLogger, _ = New(FileConf("./src/lg/logs/warn.log"))
	errorLogger, _ = New(FileConf("./src/lg/logs/error.log"))
}

//func init01() {
//	myLogger = NewLogger()
//	debugLogger = NewFileLogger("./src/lg/logs/debug.log")
//	infoLogger = NewFileLogger("./src/lg/logs/info.log")
//	warnLogger = NewFileLogger("./src/lg/logs/warn.log")
//	errorLogger = NewFileLogger("./src/lg/logs/error.log")
//}

func Error(v ...any) {
	myLogger.Errorln(v...)
	errorLogger.Errorln(v...)
}

func Debug(v ...any) {
	myLogger.Debugln(v...)
	debugLogger.Debugln(v...)
}

func Warn(v ...any) {
	myLogger.Warnln(v...)
	warnLogger.Warnln(v...)
}

func Info(v ...any) {
	myLogger.Infoln(v...)
	infoLogger.Infoln(v...)
}

func Ping() {
	myLogger.Println("log: Pong!")
}

//func NewLogger() *logrus.Logger {
//	mLog := logrus.New() //新建一个实例
//	file, err := os.OpenFile("./src/lg/logs/All.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
//	if err != nil {
//		panic(err)
//	}
//	mLog.SetOutput(io.MultiWriter(os.Stdout, file)) //设置输出类型
//	mLog.SetReportCaller(true)                      //开启返回函数名和行号
//	mLog.SetFormatter(&LogFormatter{})              //设置自己定义的Formatter
//	mLog.SetLevel(logrus.TraceLevel)                //设置最低的Level
//	return mLog
//}

//func NewFileLogger(filePath string) *logrus.Logger {
//	fLog := logrus.New() //新建一个实例
//	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
//	if err != nil {
//		panic(err)
//	}
//	fLog.SetOutput(io.MultiWriter(file)) //设置输出类型
//	fLog.SetReportCaller(true)           //开启返回函数名和行号
//	fLog.SetFormatter(&LogFormatter{})   //设置自己定义的Formatter
//	fLog.SetLevel(logrus.TraceLevel)     //设置最低的Level
//	return fLog
//}

type Options struct {
	OutputFiles []string
	WithStdout  bool
	WithFunc    bool
	WithJump    bool
	Level       logrus.Level
	Colors
}
type Colors struct {
	cInfo    int
	cDebug   int
	cWarn    int
	cError   int
	bgcInfo  int
	bgcDebug int
	bgcWarn  int
	bgcError int
}

func (c Colors) SetColor(level logrus.Level, color Colors) {
	switch level {
	case logrus.DebugLevel:

	}
}

func New(opt Options) (*logrus.Logger, error) {
	logger := logrus.New()
	if opt.OutputFiles != nil {
		var writers []io.Writer
		if opt.WithStdout {
			writers = append(writers, os.Stdout)
		}
		for _, writer := range opt.OutputFiles {
			file, err := os.OpenFile(writer, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				return nil, errors.New("invalid filepath")
			}
			writers = append(writers, file)
		}
		logger.SetOutput(io.MultiWriter(writers...)) //设置输出类型
	}

	logger.SetReportCaller(true)                      //开启返回函数名和行号
	logger.SetFormatter(&LogFormatter2{Options: opt}) //设置自己定义的Formatter
	logger.SetLevel(opt.Level)                        //设置最低的Level
	return logger, nil
}

func StdConf() Options {
	return Options{
		OutputFiles: nil,
		WithStdout:  true,
		WithFunc:    true,
		WithJump:    true,
		Level:       logrus.TraceLevel,
	}
}

func FileConf(filepath ...string) Options {
	return Options{
		OutputFiles: filepath,
		WithStdout:  false,
		WithFunc:    true,
		WithJump:    true,
		Level:       logrus.TraceLevel,
	}
}

type Option func(*Options)

func NewByOptions(options ...Option) *Options {
	opt := &Options{}
	for _, option := range options {
		option(opt)
	}
	return opt
}

func oWithStdout() Option {
	return func(o *Options) {
		o.WithStdout = true
	}
}

func oWithFile() {

}

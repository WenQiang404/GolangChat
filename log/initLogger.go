package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Logger struct {
	*zap.Logger
}

func NewLogger() *Logger {

	coreList := make([]zapcore.Core, 0)

	//==============================输出终端==============================
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel))

	//==============================输出文件(INFO)==============================
	infoBlogLogger := &lumberjack.Logger{
		Filename:   "log/info.log",
		MaxSize:    50,
		MaxBackups: 30,
		MaxAge:     90, //日志最大保存天数
		LocalTime:  true,
		Compress:   false,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(infoBlogLogger), zapcore.InfoLevel))

	//==============================输出文件(ERROR)==============================
	errBlogLogger := &lumberjack.Logger{
		Filename:   "log/err.log",
		MaxSize:    50,
		MaxBackups: 30,
		MaxAge:     90, //日志最大保存天数
		LocalTime:  true,
		Compress:   false,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(errBlogLogger), zapcore.ErrorLevel))

	//==============================输出文件(Fatal)==============================
	fatalBlogLogger := &lumberjack.Logger{
		Filename:   "log/fatal.log",
		MaxSize:    50,
		MaxBackups: 30,
		MaxAge:     90, //日志最大保存天数
		LocalTime:  true,
		Compress:   false,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(fatalBlogLogger), zapcore.ErrorLevel))
	//==============================把所有的output,都汇聚成coreList，一次性初始化==============================
	core := zapcore.NewTee(coreList...)

	//==============================设置是否打印堆栈==============================
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{
		Logger: zapLogger,
	}

}

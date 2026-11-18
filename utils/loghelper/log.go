package loghelper

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/zeromicro/go-zero/core/logx"
)

func GetColorLevel(level uint32) string {
	levelMap := map[uint32]string{
		logx.DebugLevel:  "\x1b[97;32m", //绿色
		logx.InfoLevel:   "\x1b[97;0m",  // 白色
		logx.ErrorLevel:  "\x1b[97;33m", // 黄色
		logx.SevereLevel: "\x1b[97;31m", // 红色
	}
	//背景色。。。。
	//levelMap := map[string]string{
	//	"Fatal": "\x1b[97;41m", // 红色
	//	"Panic": "\x1b[97;41m", // 红色
	//	"Error": "\x1b[97;41m", // 红色
	//	"Warn":  "\x1b[97;43m", // 黄色
	//	"Info":  "\x1b[97;0m",  // 白色
	//	"Debug": "\x1b[97;42m", //绿色
	//}
	return levelMap[level]
}

func InitLog(level uint32) {
	//logx.SetWriter() //todo:如果想输出正确的caller，需要修改writer
	logx.WithCallerSkip(1) //todo:通过获取一个不输出的caller，然后，我们自己再增加字段！
}

//	func Printf(level uint32, messageFormat string, args ...any) {
//		if logger.Level >= level {
//			logger.SetPrefix(GetColorLevel(level) + logManager.levels[level] + "[" + logger.scope + "][" + time.Now().Format("2006-01-02 15:04:05.000") + "]")
//			format := messageFormat + "\x1b[0m"
//			switch level {
//			case logx.DebugLevel:
//				logx.Debug(msg)
//			case logx.InfoLevel:
//				logx.Info(msg)
//			case logx.ErrorLevel:
//				logx.Error(msg)
//			case logx.SevereLevel:
//				logx.Servere(msg)
//			default:
//			}
//		}
//	}

// ★ 调用层级:
//
//	skip = 0: getRealCaller 自身
//	skip = 1: 调用 getRealCaller 的函数（如 Info）
//	skip = 2: 调用 Info 的业务代码 ← 这是我们要的
func getRealCaller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func enableWindowsANSI() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setConsoleMode := kernel32.NewProc("SetConsoleMode")
	handle := syscall.Stdout
	var mode uint32
	getConsoleMode := kernel32.NewProc("GetConsoleMode")
	getConsoleMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&mode)))
	// ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004
	setConsoleMode.Call(uintptr(handle), uintptr(mode|0x0004))
}
func SetColorLogLevel(level uint32, msg string) string {
	if runtime.GOOS == "windows" {
		// Windows 10+ 可通过启用 VT 支持 ANSI
		enableWindowsANSI()
	}
	return GetColorLevel(level) + msg + "\x1b[0m"
}

func Debug(message string) {
	logx.Debug(SetColorLogLevel(logx.DebugLevel, message))
}
func Debugf(messageFormat string, args ...any) {
	logx.Debugf(SetColorLogLevel(logx.DebugLevel, messageFormat), args...)
}

func Info(message string) {
	//logx.Infow(SetColorLogLevel(logx.InfoLevel, message), logx.Field("caller", getRealCaller(4)))
	logx.Info(SetColorLogLevel(logx.InfoLevel, message))
}

func Infof(messageFormat string, args ...any) {
	logx.Info(fmt.Sprintf(messageFormat, args...))
}

func Error(message string) {
	logx.Error(SetColorLogLevel(logx.ErrorLevel, message))
}

func Errorf(messageFormat string, args ...any) {
	logx.Errorf(SetColorLogLevel(logx.ErrorLevel, messageFormat), args...)
}

func Severe(message string) {
	logx.Severe(SetColorLogLevel(logx.SevereLevel, message))
}
func Severef(messageFormat string, args ...any) {
	logx.Severef(SetColorLogLevel(logx.SevereLevel, messageFormat), args...)
}

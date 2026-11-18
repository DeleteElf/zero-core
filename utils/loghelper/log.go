package loghelper

import (
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
	logx.Info(SetColorLogLevel(logx.InfoLevel, message))
}

func Infof(messageFormat string, args ...any) {
	logx.Infof(SetColorLogLevel(logx.InfoLevel, messageFormat), args...)
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

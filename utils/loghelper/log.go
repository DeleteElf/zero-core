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
		logx.ErrorLevel:  "\x1b[97;31m", // 红色
		logx.SevereLevel: "\x1b[97;33m", // 黄色
	}
	//颜色      前景色    背景色
	//──────────────────────────
	//黑色       30        40
	//红色       31        41
	//绿色       32        42
	//黄色       33        43
	//蓝色       34        44
	//品红       35        45
	//青色       36        46
	//白色       37        47
	return levelMap[level]
}

var logger logx.Logger

func InitLog() {
	logger = logx.WithCallerSkip(1) //我们包装了一层，让它跳过一层即可
	if runtime.GOOS == "windows" {
		// Windows 10+ 可通过启用 VT 支持 ANSI
		enableWindowsANSI()
	}
	Info("日志系统启动！！！")
}

//	 ★ 调用层级（这个代码在这里保留，仅作为技术存储）:
//		skip = 0: getRealCaller 自身
//		skip = 1: 调用 getRealCaller 的函数（如 Info）
//		skip = 2: 调用 Info 的业务代码 ← 这是我们要的
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
	return GetColorLevel(level) + msg + "\x1b[0m"
}

func Debug(message string) {
	logger.Debug(SetColorLogLevel(logx.DebugLevel, message))
}
func Debugf(messageFormat string, args ...any) {
	logger.Debugf(SetColorLogLevel(logx.DebugLevel, messageFormat), args...)
}

func Info(message string) {
	//logx.Infow(SetColorLogLevel(logx.InfoLevel, message), logx.Field("caller", getRealCaller(4)))
	logger.Info(SetColorLogLevel(logx.InfoLevel, message))
}

func Infof(messageFormat string, args ...any) {
	//Info(fmt.Sprintf(messageFormat, args...)) //我们现在不用重度优化，为了保证caller级别的正确性，这里也只做一级调用
	logger.Infof(SetColorLogLevel(logx.InfoLevel, messageFormat), args...)
}

func Error(message string) {
	logger.Error(SetColorLogLevel(logx.ErrorLevel, message))
}

func Errorf(messageFormat string, args ...any) {
	logger.Errorf(SetColorLogLevel(logx.ErrorLevel, messageFormat), args...)
}

func Slow(message string) {
	logger.Slow(SetColorLogLevel(logx.InfoLevel, message))
}
func Slowf(messageFormat string, args ...any) {
	logger.Slowf(SetColorLogLevel(logx.InfoLevel, messageFormat), args...)
}

//func Severe(message string) {
//	logger.Severe(SetColorLogLevel(logx.SevereLevel, message))
//}
//func Severef(messageFormat string, args ...any) {
//	logger.Severef(SetColorLogLevel(logx.SevereLevel, messageFormat), args...)
//}

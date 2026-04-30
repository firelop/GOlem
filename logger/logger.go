package logger

import (
	"fmt"
	"log"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

type Channel string

const (
	Protocol Channel = "Protocol"
	NBT      Channel = "NBT"
	Server   Channel = "Server"
	Network  Channel = "Network"
)

var DebugChannels map[Channel]bool = make(map[Channel]bool)

func Debug(channel Channel, v ...any) {
	if DebugChannels[channel] {
		prefix := fmt.Sprintf("[DEBUG] [%s]", channel)
		log.Println(append([]any{prefix}, v...)...)
	}
}

func Debugf(channel Channel, format string, v ...any) {
	if DebugChannels[channel] {
		prefix := fmt.Sprintf("[DEBUG] [%s] ", channel)
		log.Printf(prefix+format, v...)
	}
}

func Info(v ...any) {
	log.Println(append([]any{"[INFO]"}, v...)...)
}

func Infof(format string, v ...any) {
	log.Printf("[INFO] "+format, v...)
}

func Warn(v ...any) {
	log.Println(append([]any{"[WARN]"}, v...)...)
}

func Warnf(format string, v ...any) {
	log.Printf("[WARN] "+format, v...)
}

func Error(v ...any) {
	log.Println(append([]any{"[ERROR]"}, v...)...)
}

func Errorf(format string, v ...any) {
	log.Printf("[ERROR] "+format, v...)
}

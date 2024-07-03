package config

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

type LogFormatter struct {
	log.TextFormatter
}

func (f *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	prettyCaller := func(frame *runtime.Frame) string {
		_, fileName := filepath.Split(frame.File)
		return fmt.Sprintf("%s:%d ", fileName, frame.Line)
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	b.WriteString(fmt.Sprintf("[%s]:%s ", entry.Time.Format(f.TimestampFormat), strings.ToUpper(entry.Level.String())))
	if entry.HasCaller() {
		b.WriteString(prettyCaller(entry.Caller))
	}
	b.WriteString(entry.Message)
	b.WriteByte('\n')
	return b.Bytes(), nil
}

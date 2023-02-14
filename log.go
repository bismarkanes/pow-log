/*
Copyright © 2023 Bismark <bismark.john.anes@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package log

import (
	"io"
	"runtime"
	"strconv"
	"time"

  "github.com/bismarkanes/pow-log/logengine"
)

const (
	openTag           = "["
	closeTag          = "]"
	fileLinumDelim    = ":"
	delim             = " " // delimiter or separator
	prefixInfo        = delim + openTag + "info" + closeTag + delim
	prefixWarn        = delim + openTag + "warn" + closeTag + delim
	prefixError       = delim + openTag + "error" + closeTag + delim
	defaultFormatTime = "2006-01-02T15:04:05"
)

// Logger for logging things
type Logger interface {
	Info(formatMsg string, args ...any)
	Warn(formatMsg string, args ...any)
	Error(formatMsg string, args ...any)
	Label(formatMsg, label string, args ...any)
	Errorn(formatMsg string, args ...any)
	Infot(formatMsg, tag string, args ...any)
}

type logger struct {
	useLogConsole bool
	useLogCustom  bool
	loggerCustom  logengine.LoggerEngine // custom log engine
	loggerConsole logengine.LoggerEngine // console log engine
	formatTime    string
	logCustomOpen func() (io.WriteCloser, error)
}

// linum use line number
func (l *logger) log(formatMsg string, linum bool, args ...any) {
	var msg string = formatMsg

	if linum {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			msg = formatMsg + delim + file + fileLinumDelim + strconv.Itoa(line)
		}
	}

	if l.useLogCustom {
		h, err := l.logCustomOpen()
		if err != nil {
			return
		}
		defer h.Close()

		l.loggerCustom.SetOutput(h)
		l.loggerCustom.Printf(msg, args...)
	}

	if l.useLogConsole {
		l.loggerConsole.Printf(msg, args...)
	}

	return
}

func (l *logger) timeNow() string {
	return time.Now().Format(l.formatTime)
}

func (l *logger) Info(msg string, args ...any) {
	l.log(l.timeNow()+prefixInfo+msg, false, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.log(l.timeNow()+prefixWarn+msg, false, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.log(l.timeNow()+prefixError+msg, false, args...)
}

func (l *logger) Label(msg, label string, args ...any) {
	l.log(l.timeNow()+delim+openTag+label+closeTag+delim+msg, false, args...)
}

func (l *logger) Errorn(msg string, args ...any) {
	l.log(l.timeNow()+prefixError+msg, true, args...)
}

func (l *logger) Infot(msg, tag string, args ...any) {
	l.log(l.timeNow()+prefixInfo+openTag+tag+closeTag+delim+msg, false, args...)
}

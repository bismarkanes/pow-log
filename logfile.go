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
	"os"

	"github.com/bismarkanes/pow-log/logengine"
)

type loggerFile struct {
	logger
	filename string
}

func (lf *loggerFile) open() (io.WriteCloser, error) {
	return os.OpenFile(lf.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

// NewCustomLogFile custom logger engine log to local file
func NewCustomLogFile(filename string, useLogConsole, useLogFile bool, newEngine logengine.NewLoggerEngine) Logger {
	var engineFile logengine.LoggerEngine
	var engineConsole logengine.LoggerEngine

	if useLogFile {
		engineFile = newEngine(nil)
	}

	if useLogConsole {
		engineConsole = newEngine(os.Stdout)
	}

	lf := loggerFile{}

	lf.filename = filename
	lf.useLogConsole = useLogConsole
	lf.useLogCustom = useLogFile
	lf.loggerCustom = engineFile
	lf.loggerConsole = engineConsole
	lf.formatTime = defaultFormatTime
	lf.logCustomOpen = lf.open

	return &lf
}

// NewLogFile default logger to local file
func NewLogFile(filename string, isLogConsole, isLogFile bool) Logger {
	return NewCustomLogFile(
		filename,
		isLogConsole,
		isLogFile,
		logengine.NewBuiltInLoggerEngine,
	)
}

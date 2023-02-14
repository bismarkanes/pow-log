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
	"os"

	"github.com/bismarkanes/pow-log/logengine"
)

type loggerConsole struct {
	logger
}

// NewCustomLogConsole custom logger engine log to local file
func NewCustomLogConsole(newEngine logengine.NewLoggerEngine) Logger {
	lc := loggerConsole{}

	lc.useLogConsole = true
	lc.useLogCustom = false
	lc.loggerCustom = nil
	lc.loggerConsole = newEngine(os.Stdout)
	lc.formatTime = defaultFormatTime
	lc.logCustomOpen = nil

	return &lc
}

// NewLogConsole default
func NewLogConsole() Logger {
	return NewCustomLogConsole(logengine.NewBuiltInLoggerEngine)
}

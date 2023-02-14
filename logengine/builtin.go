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

// sample for customizing your own log engine

package logengine

import (
	"io"
	"log"
)

// built in logger engine (module log)
type builtInLoggerEnginer struct {
	log *log.Logger
}

func NewBuiltInLoggerEngine(wc io.WriteCloser) LoggerEngine {
	loggerStd := log.New(wc, "", 0)

	cle := builtInLoggerEnginer{
		log: loggerStd,
	}

	return &cle
}

func (bie *builtInLoggerEnginer) SetOutput(wc io.WriteCloser) {
	bie.log.SetOutput(wc)
}

func (bie *builtInLoggerEnginer) Printf(msg string, args ...any) {
	bie.log.Printf(msg, args...)
}

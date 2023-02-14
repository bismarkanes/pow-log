package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogFile(t *testing.T) {
	mockDatas := []struct {
		title            string
		actualFilename   string
		expectedFilename string
		content          string
		useFile          bool
		useConsole       bool
		fileExist        bool
	}{
		{
			title:            "Test NewLogFile 1",
			actualFilename:   "testing1.log",
			expectedFilename: "testing1.log",
			content:          "TEST 1",
			useFile:          true,
			useConsole:       true,
			fileExist:        true,
		},
		{
			title:            "Test NewLogFile 2",
			actualFilename:   "testing1.log",
			expectedFilename: "testing1.log",
			content:          "TEST 2",
			useFile:          false,
			useConsole:       true,
			fileExist:        false,
		},
		{
			title:            "Test NewLogFile 3",
			actualFilename:   "testing1.log",
			expectedFilename: "testing1.log",
			content:          "TEST 3",
			useFile:          true,
			useConsole:       false,
			fileExist:        true,
		},
		{
			title:            "Test NewLogFile 4",
			actualFilename:   "testing1.log",
			expectedFilename: "testing1.log",
			content:          "TEST 4",
			useFile:          false,
			useConsole:       false,
			fileExist:        false,
		},
	}

	for _, mockData := range mockDatas {
		// delete log file if exist
		if _, err := os.Stat(mockData.actualFilename); err == nil {
			err = os.Remove(mockData.actualFilename)
			assert.Nil(t, err, mockData.title)
		}

		log := NewLogFile(mockData.actualFilename, mockData.useConsole, mockData.useFile)

		log.Info("This is info %s", mockData.content)
		log.Warn("This is warn %s", mockData.content)
		log.Error("This is error %s", mockData.content)
		log.Label("This is label %s", "BISMARK", mockData.content)
		log.Errorn("This is error %s", mockData.content)
		log.Infot("This is info tag %s", "ORDERS", mockData.content)

		// check file
		f, err := os.Open(mockData.expectedFilename)

		if mockData.fileExist {
			assert.Nil(t, err, mockData.title)
			f.Close()
		} else {
			assert.Error(t, err, mockData.title)
		}
	}
}

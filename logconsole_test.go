package log

import (
	"testing"
)

func TestNewLogConsole(t *testing.T) {
	mockDatas := []struct {
		title   string
		content string
	}{
		{
			title:   "Test NewLogConsole",
			content: "HELLO",
		},
	}

	for _, mockData := range mockDatas {
		log := NewLogConsole()

		log.Info("This is info %s", mockData.content)
		log.Warn("This is warn %s", mockData.content)
		log.Error("This is error %s", mockData.content)
		log.Label("This is label %s", "BISMARK", mockData.content)
		log.Errorn("This is error %s", mockData.content)
	}
}

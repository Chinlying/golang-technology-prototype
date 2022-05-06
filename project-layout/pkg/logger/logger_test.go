package logger

import "testing"

func TestAll(t *testing.T) {
	logger.Infof("starting service")
	SetServiceName("testing service")
	logger.Errorf("Failed to start service")
}

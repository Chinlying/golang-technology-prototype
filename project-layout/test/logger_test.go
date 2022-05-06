package test

import (
	serviceLog "growing-into-an-excellent-golang-developer/project-layout/pkg/logger"
	"testing"
)

func TestAll(t *testing.T) {
	serviceName := "testing service"
	serviceLog.SetLogLevel("debug")
	logger := serviceLog.GetDefaultLogger()

	logger.Infof("starting service")
	serviceLog.SetServiceName("testing service")
	logger.Errorf("Failed to start service %s", serviceName)
}

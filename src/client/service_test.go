package ser_test

import (
	"testing"
	"service"
)

func TestPackage(t *testing.T) {
	t.Log(service.HelloWorld())
}
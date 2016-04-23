package kafkalog_logrus

import (
	"bytes"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHook(t *testing.T) {
	var buffer bytes.Buffer
	formater := &logrus.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	}
	lgr := logrus.Logger{
		Out:       &buffer,
		Formatter: formater,
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	hook, err := NewKafkalogHook("name", 3600, "/tmp")
	assert.Nil(t, err)
	lgr.Hooks.Add(hook)

	lgr.Errorf("Log this")
	assert.Equal(t, buffer.String(), "level=error msg=\"Log this\" \n")
	// TODO assert file /tmp/DATE_3600_UTC-name.szn was created
}

func TestHookError(t *testing.T) {
	var buffer bytes.Buffer
	formater := &logrus.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	}
	lgr := logrus.Logger{
		Out:       &buffer,
		Formatter: formater,
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	hook, err := NewKafkalogHook("name", 3600, "/neexistujiciadresarzpusobipanic")
	assert.Nil(t, err)
	lgr.Hooks.Add(hook)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Function did not panic")
		}

		called := false
		onerr := func(err error) {
			called = true
		}
		hook.SetOnError(onerr)
		lgr.Errorf("Log this")
		if !called {
			t.Error("On error func was not called")
		}
	}()
	lgr.Errorf("Log this") // panic
}

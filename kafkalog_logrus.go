package kafkalog_logrus

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	kafkalog "github.com/sejvlond/go-kafkalog/common"
	kafkalog_writer "github.com/sejvlond/go-kafkalog/writer"
)

type KafkalogHookOnError func(error)

type KafkalogHook struct {
	writer  kafkalog_writer.KafkalogWriter
	onError KafkalogHookOnError
}

func NewKafkalogHook(name string, interval uint, dir string) (
	hook *KafkalogHook, err error) {
	dw, err := kafkalog_writer.NewRotate(name, interval, dir,
		kafkalog.COMPRESS_SNAPPY)
	if err != nil {
		return
	}
	hook = &KafkalogHook{
		writer: dw,
		onError: func(err error) {
			fmt.Printf("KAFKALOG HOOK ERROR '%v'\n", err)
			panic(err)
		},
	}
	return
}

func (hook *KafkalogHook) SetOnError(fn KafkalogHookOnError) {
	hook.onError = fn
}

func (hook *KafkalogHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	bytes := []byte(line)
	if err != nil {
		hook.onError(err)
		return fmt.Errorf("Unable to read entry: '%v'", err)
	}
	_, err = hook.writer.Write(bytes, bytes, 0)
	if err != nil {
		hook.onError(err)
		return fmt.Errorf("Unable to write to kafkalog: '%v'", err)
	}
	return nil
}

func (hook *KafkalogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

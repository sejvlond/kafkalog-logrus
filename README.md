# Kafkalog Logrus

kafkalog hook for logrus https://github.com/Sirupsen/logrus

## Usage

```
import (
    "github.com/Sirupsen/logrus"
    "github.com/sejvlond/kafkalog-logrus"
)

func main() {
    lgr := logrus.New()
    kafkalog_hook, err := kafkalog_logrus.NewKafkalogHook("componentName", 3600, "/tmp")
    if err != nil {
        panic(err)
    }
    lgr.Hooks.Add(kafkalog_hook)

    lgr.Infof("Hello world")
}
```

By default it panics on error, it can be overwritten by KafkalogHookOnError func
```
type KafkalogHookOnError func(error)

onError := func(err error) {}
kafkalog_hook.SetOnError(onError)
```

package main

import (
	"fmt"
	"go8/configuration"
)

func main() {
	conf, err := configuration.Load("config.json")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Port: ", conf.Port)
		fmt.Println("Jaeger url: ", conf.JaegerUrl)
		fmt.Println("Sentry url: ", conf.SentryUrl)
		fmt.Println("Kafka broker: ", conf.KafkaBroker)
		fmt.Println("App id: ", conf.AppId)
		fmt.Println("App key: ", conf.AppKey)
	}

}

package configuration

import (
	"errors"
	"flag"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Port        string
	DbUrl       string
	JaegerUrl   string
	SentryUrl   string
	KafkaBroker string
	AppId       string
	AppKey      string
}

func Load() (*Configuration, error) {
	var conf Configuration

	flag.StringVar(&conf.Port, "port", "", "port")
	flag.StringVar(&conf.DbUrl, "db_url", "", "database url")
	flag.StringVar(&conf.JaegerUrl, "jaeger_url", "", "jaeger url")
	flag.StringVar(&conf.SentryUrl, "sentry_url", "", "sentry url")
	flag.StringVar(&conf.KafkaBroker, "kafka_broker", "", "kafka broker")
	flag.StringVar(&conf.AppId, "some_app_id", "", "application id")
	flag.StringVar(&conf.AppKey, "some_app_key", "", "application key")
	flag.Parse()

	err := godotenv.Load()
	if err == nil {
		if conf.Port == "" {
			conf.Port = os.Getenv("PORT")
		}
		if conf.DbUrl == "" {
			conf.DbUrl = os.Getenv("db_url")
		}
		if conf.JaegerUrl == "" {
			conf.JaegerUrl = os.Getenv("jaeger_url")
		}
		if conf.SentryUrl == "" {
			conf.SentryUrl = os.Getenv("sentry_url")
		}
		if conf.KafkaBroker == "" {
			conf.KafkaBroker = os.Getenv("kafka_broker")
		}
		if conf.AppId == "" {
			conf.AppId = os.Getenv("some_app_id")
		}
		if conf.AppKey == "" {
			conf.AppKey = os.Getenv("some_app_key")
		}
	}
	err = conf.ValidateConf()

	return &conf, err

}

func (conf Configuration) ValidateConf() error {
	var errString string

	valid, _ := regexp.MatchString(`\d{4}`, conf.Port)
	if !valid {
		errString += "Port invalid "
	}

	if !strings.HasPrefix(conf.DbUrl, "postgres://db-user:db-password@") {
		errString += "database url invalid "
	}

	valid, _ = regexp.MatchString(`http://jaeger:\d{4}`, conf.JaegerUrl)
	if !valid {
		errString += "jaeger url invalid "
	}

	valid, _ = regexp.MatchString(`http://sentry:\d{4}`, conf.SentryUrl)
	if !valid {
		errString += "Sentry url invalid "
	}

	if errString != "" {
		return errors.New(errString)
	} else {
		return nil
	}
}

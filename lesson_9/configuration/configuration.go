package configuration

import (
	"encoding/json"
	"errors"
	"flag"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Port        string `json:"port" yaml:"port"`
	DbUrl       string `json:"db_url" yaml:"db_url"`
	JaegerUrl   string `json:"jaeger_url" yaml:"jaeger_url"`
	SentryUrl   string `json:"sentry_url" yaml:"sentry_url"`
	KafkaBroker string `json:"kafka_broker" yaml:"kafka_broker"`
	AppId       string `json:"some_app_id" yaml:"some_app_id"`
	AppKey      string `json:"some_app_key" yaml:"some_app_key"`
}

func Load(fileName ...string) (*Configuration, error) {
	var conf Configuration

	if len(fileName) > 0 {
		fName := fileName[0]
		f, err := os.Open(fName)
		if err != nil {
			return &conf, err
		}

		fExtension := filepath.Ext(fName)

		if fExtension == ".json" {
			if err := json.NewDecoder(f).Decode(&conf); err != nil {
				return &conf, err
			}
		} else if fExtension == ".yaml" {
			if err := yaml.NewDecoder(f).Decode(&conf); err != nil {
				return &conf, err
			}
		} else {
			return &conf, errors.New("cant read config from file")
		}
	} else {
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
	}

	err := conf.ValidateConf()

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

	valid, _ = regexp.MatchString(`http://jaeger:\d{5}`, conf.JaegerUrl)
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

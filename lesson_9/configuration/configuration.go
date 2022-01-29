package configuration

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
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
		err := godotenv.Load()
		if err != nil {
			fmt.Println("There is no .env file with config, you must set config with flags")
		} else {
			conf.Port = os.Getenv("PORT")
			conf.DbUrl = os.Getenv("db_url")
			conf.JaegerUrl = os.Getenv("jaeger_url")
			conf.SentryUrl = os.Getenv("sentry_url")
			conf.KafkaBroker = os.Getenv("kafka_broker")
			conf.AppId = os.Getenv("some_app_id")
			conf.AppKey = os.Getenv("some_app_key")
		}
	}
	flag.StringVar(&conf.Port, "port", conf.Port, "port")
	flag.StringVar(&conf.DbUrl, "db_url", conf.DbUrl, "database url")
	flag.StringVar(&conf.JaegerUrl, "jaeger_url", conf.JaegerUrl, "jaeger url")
	flag.StringVar(&conf.SentryUrl, "sentry_url", conf.SentryUrl, "sentry url")
	flag.StringVar(&conf.KafkaBroker, "kafka_broker", conf.KafkaBroker, "kafka broker")
	flag.StringVar(&conf.AppId, "some_app_id", conf.AppId, "application id")
	flag.StringVar(&conf.AppKey, "some_app_key", conf.AppKey, "application key")
	flag.Parse()

	err := conf.ValidateConf()

	return &conf, err
}

func (conf Configuration) ValidateConf() error {

	valid, _ := regexp.MatchString(`\d{4}`, conf.Port)
	if !valid {
		return errors.New("port invalid ")
	}

	if !strings.HasPrefix(conf.DbUrl, "postgres://db-user:db-password@") {
		return errors.New("database url invalid ")
	}

	valid, _ = regexp.MatchString(`http://jaeger:\d{5}`, conf.JaegerUrl)
	if !valid {
		return errors.New("jaeger url invalid ")
	}

	valid, _ = regexp.MatchString(`http://sentry:\d{4}`, conf.SentryUrl)
	if !valid {
		return errors.New("sentry url invalid ")
	}
	return nil
}

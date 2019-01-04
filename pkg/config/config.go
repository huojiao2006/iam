// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// OpenPitrix IAM service config package.
package config

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/koding/multiconfig"
	"gopkg.in/yaml.v2"

	"openpitrix.io/iam/pkg/internal/jsonutil"
	"openpitrix.io/logger"
)

const EnvPrefix = "OPENPITRIX_IAM"

//
// Environment:
//
// OPENPITRIX_IAM_DB_TYPE
// OPENPITRIX_IAM_DB_HOST
// OPENPITRIX_IAM_DB_PORT
// OPENPITRIX_IAM_DB_USER
// OPENPITRIX_IAM_DB_PASSWORD
// OPENPITRIX_IAM_DB_DATABASE
//
// OPENPITRIX_IAM_HOST
// OPENPITRIX_IAM_PORT
//
// OPENPITRIX_IAM_TLS_ENABLED
// OPENPITRIX_IAM_TLS_CERT_FILE
// OPENPITRIX_IAM_TLS_KEY_FILE
//
// OPENPITRIX_IAM_LOG_LEVEL
//

type Config struct {
	DB DBConfig

	Host        string `default:"openpitrix-iam2-service"`
	Port        int    `default:"9119"`
	TlsEnabled  bool   `default:"false"`
	TlsCertFile string `default:"server.cert"`
	TlsKeyFile  string `default:"server.key"`

	LogLevel string `default:"info"` // debug, info, warn, error, fatal
}

type DBConfig struct {
	Type     string `default:"mysql"`
	Host     string `default:"openpitrix-db"`
	Port     int    `default:"3306"`
	User     string `default:"root"`
	Password string `default:"password"`
	Database string `default:"iam2"`
}

func (m *Config) Clone() *Config {
	q := *m
	return &q
}

func (m *DBConfig) GetHostUrl() string {
	if m.Type == "sqlite3" {
		return m.Database
	}
	if m.Type == "mysql" {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/", m.User, m.Password, m.Host, m.Port)
	}
	return m.Database
}

func (m *DBConfig) GetUrl() string {
	if m.Type == "sqlite3" {
		return m.Database
	}
	if m.Type == "mysql" {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.User, m.Password, m.Host, m.Port, m.Database)
	}
	return m.Database
}
func (m *DBConfig) GetUrlWithParseTime() string {
	if m.Type == "sqlite3" {
		return m.Database
	}
	if m.Type == "mysql" {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", m.User, m.Password, m.Host, m.Port, m.Database)
	}
	return m.Database
}

func Default() *Config {
	conf := new(Config)

	loader := newWithPath("")
	if err := loader.Load(conf); err != nil {
		logger.Criticalf(nil, "%v", err)
		panic(err)
	}
	return conf
}

func Load(path string) (*Config, error) {
	conf := new(Config)

	loader := newWithPath(path)
	if err := loader.Load(conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func MustLoad(path string) *Config {
	conf := new(Config)

	loader := newWithPath(path)
	if err := loader.Load(conf); err != nil {
		logger.Criticalf(nil, "%s: %v", path, err)
		panic(err)
	}

	return conf
}

func (p *Config) JSONString() string {
	return string(jsonutil.Encode(p))
}

func (p *Config) TOMLString() string {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(p); err != nil {
		logger.Criticalf(nil, "%v", err)
		panic(err)
	}
	return buf.String()
}

func (p *Config) YAMLString() string {
	data, err := yaml.Marshal(p)
	if err != nil {
		logger.Criticalf(nil, "%v", err)
		panic(err)
	}
	return string(data)
}

func newWithPath(path string) *multiconfig.DefaultLoader {
	loaders := []multiconfig.Loader{}

	// Read default values defined via tag fields "default"
	loaders = append(loaders, &multiconfig.TagLoader{})

	// Choose what while is passed
	if strings.HasSuffix(path, ".toml") {
		loaders = append(loaders, &multiconfig.TOMLLoader{Path: path})
	}

	if strings.HasSuffix(path, ".json") {
		loaders = append(loaders, &multiconfig.JSONLoader{Path: path})
	}

	if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
		loaders = append(loaders, &multiconfig.YAMLLoader{Path: path})
	}

	loaders = append(loaders, &multiconfig.EnvironmentLoader{
		Prefix:    EnvPrefix,
		CamelCase: true,
	})

	loader := multiconfig.MultiLoader(loaders...)

	d := &multiconfig.DefaultLoader{}
	d.Loader = loader
	d.Validator = multiconfig.MultiValidator(&multiconfig.RequiredValidator{})
	return d
}

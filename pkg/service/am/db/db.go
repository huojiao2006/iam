// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"openpitrix.io/iam/pkg/service/am/config"
	"openpitrix.io/logger"
)

type Database struct {
	cfg *config.Config
	*gorm.DB
}

type Options struct {
	SqlInitTable []string
	SqlInitData  []string
}

func OpenDatabase(cfg *config.Config, opt *Options) (*Database, error) {
	cfg = cfg.Clone()

	logger.Infof(nil, "DB config: begin")
	logger.Infof(nil, "\tType: %s", cfg.DB.Type)
	logger.Infof(nil, "\tHost: %s", cfg.DB.Host)
	logger.Infof(nil, "\tPort: %d", cfg.DB.Port)
	logger.Infof(nil, "\tUser: %s", cfg.DB.User)
	logger.Infof(nil, "\tDatabase: %s", cfg.DB.Database)
	logger.Infof(nil, "DB config: end")

	var p = &Database{cfg: cfg}
	var err error

	p.DB, err = gorm.Open(cfg.DB.Type, cfg.DB.GetUrl())
	if err != nil {
		return nil, err
	}

	p.DB.SingularTable(true)

	// init hook
	if opt != nil && len(opt.SqlInitTable) > 0 {
		for _, sql := range opt.SqlInitTable {
			if err := p.DB.Exec(sql).Error; err != nil {
				logger.Warnf(nil, "%+v", err)
			}
		}
	}
	if opt != nil && len(opt.SqlInitData) > 0 {
		for _, sql := range opt.SqlInitData {
			if err := p.DB.Exec(sql).Error; err != nil {
				logger.Warnf(nil, "%+v", err)
			}
		}
	}

	return p, nil
}

func (p *Database) Close() error {
	return p.DB.Close()
}

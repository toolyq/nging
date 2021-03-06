package upload

import (
	"io"

	uploadClient "github.com/webx-top/client/upload"

	"github.com/admpub/color"
	"github.com/admpub/log"
	modelFile "github.com/admpub/nging/application/model/file"
)

type (
	DBSaver func(fileM *modelFile.File, result *uploadClient.Result, reader io.Reader) error
)

var (
	dbSavers       = map[string]DBSaver{}
	DefaultDBSaver = func(fileM *modelFile.File, result *uploadClient.Result, reader io.Reader) error {
		return nil
	}
)

func DBSaverRegister(key string, dbsaver DBSaver) {
	dbSavers[key] = dbsaver
	log.Info(color.YellowString(`dbsaver.register:`), key)
}

func DBSaverUnregister(keys ...string) {
	for _, key := range keys {
		_, ok := dbSavers[key]
		if ok {
			delete(dbSavers, key)
		}
	}
}

func DBSaverGet(key string, defaults ...string) DBSaver {
	if dbsaver, ok := dbSavers[key]; ok {
		return dbsaver
	}
	if len(defaults) > 0 {
		return DBSaverGet(defaults[0])
	}
	return DefaultDBSaver
}

package container

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var dbWriterOnce sync.Once
var dbWriter *gorm.DB

//SetDbWriter creates a dbWriter used for writing into the db
func SetDbWriter(dbCon *gorm.DB) {
	dbWriterOnce.Do(func() {
		dbWriter = dbCon
	})
}

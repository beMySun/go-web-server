package dao

import "github.com/jinzhu/gorm"

// DAO definenion
type DAO struct {
	dbMap map[string]*gorm.DB
}

// NewDAO returns the DAO struct
func NewDAO() *DAO {
	dsns := map[string]string{
		"todo":     "root:123456@/todoLists",
		"database": "root:123456@/database",
	}

	// todo_lists_db
	// todo_tab

	dbMap := make(map[string]*gorm.DB)

	for key, dsn := range dsns {
		db, err := gorm.Open("mysql", dsn)
		if err != nil {
			panic("db " + key + " connection error: " + err.Error())
		}
		dbMap[key] = db.Debug()
	}

	dao := &DAO{
		dbMap: dbMap,
	}
	return dao
}

// SelectDB connect DBS
func (d *DAO) SelectDB(key string) *gorm.DB {
	if db, ok := d.dbMap[key]; ok {
		return db
	}
	return nil
}

// Close DBs
func (d *DAO) Close() {
	for _, db := range d.dbMap {
		db.Close()
	}
}

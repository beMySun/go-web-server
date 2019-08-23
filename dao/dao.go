package dao

import "github.com/jinzhu/gorm"

type DAO struct {
	db *gorm.DB
}

// NewDAO returns the DAO struct
func NewDAO() *DAO {
	db, err := gorm.Open("mysql", "root:123456@/database")
	if err != nil {
		panic("db connection error")
	}
	dao := &DAO{
		db: db.Debug(),
	}
	return dao
}

func (d *DAO) Close() {
	d.db.Close()
}

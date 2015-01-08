package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var (
	db gorm.DB
)

type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

func IntiDB() {
	var (
		driver, spec string
		found        bool
	)

	// Read configuration.
	if driver, found = revel.Config.String("db.driver"); !found {
		revel.ERROR.Fatal("No db.driver found.")
	}
	if spec, found = revel.Config.String("db.spec"); !found {
		revel.ERROR.Fatal("No db.spec found.")
	}

	// Open a connection.
	var err error
	db, err = gorm.Open(driver, spec)
	if err != nil {
		revel.ERROR.Fatal(err)
	}
}

// Begin a transaction
func (c *GormController) Begin() revel.Result {
	c.Txn = db.Begin()
	return nil
}

// Rollback if it's still going (must have panicked).
func (c *GormController) Rollback() revel.Result {
	if c.Txn != nil {
		c.Txn.Rollback()
	}
	return nil
}

// Commit the transaction.
func (c *GormController) Commit() revel.Result {
	if c.Txn != nil {
		c.Txn.Commit()
		c.Txn = nil
	}
	return nil
}

func init() {
	revel.InterceptMethod((*GormController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GormController).Commit, revel.AFTER)
	revel.InterceptMethod((*GormController).Rollback, revel.FINALLY)
}

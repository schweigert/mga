package main

import (
	"github.com/schweigert/mga/libraries/db"
	"github.com/schweigert/mga/model"
)

func migrate() {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	dbc.AutoMigrate(
		&model.Account{},
		&model.Character{},
		&model.Item{},
	)
}

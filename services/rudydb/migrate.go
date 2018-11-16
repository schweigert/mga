package main

import (
	"github.com/schweigert/mga/model"
	"github.com/schweigert/mga/services/rudydb/internal/db"
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

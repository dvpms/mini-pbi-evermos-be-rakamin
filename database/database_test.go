package database_test

import (
	"mini-project-pbi/database"
	"testing"
)

func TestConnectDatabaseAndMigrate(t *testing.T) {
	db := database.ConnectDatabase()
	if db == nil {
		t.Fatal("Expected non-nil db instance")
	}

	// Verify tables exist
	tables := []string{
		"users",
		"tokos",
		"alamats",
		"categories",
		"produks",
		"foto_produks",
		"log_produks",
		"transaksis",
		"detail_transaksis",
	}

	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			t.Errorf("Expected table %s to exist in database", table)
		} else {
			t.Logf("Table %s verified exists", table)
		}
	}
}

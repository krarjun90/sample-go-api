package migrations

//go:generate go-bindata -pkg $GOPACKAGE -o generated_migration_file.go -prefix "sql/" ./sql

import (
	"fmt"
	"github.com/krarjun90/sample-go-api/config"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/mattes/migrate/source/go-bindata"
)

func Migrate() {
	m := setupMigrate()
	err := m.Up()
	if err != nil {
		panic(fmt.Sprintf("migration up failed, error: %s", err.Error()))
	}
}

func Rollback() {
	m := setupMigrate()
	err := m.Down()
	if err != nil {
		panic(fmt.Sprintf("migration down failed, error: %s", err.Error()))
	}
}

func setupMigrate() *migrate.Migrate {
	var err error
	defer func(){
		if err != nil {
			panic(fmt.Sprintf("migration setup failed, error: %s", err.Error()))
		}
	}()

	s := bindata.Resource(AssetNames(),
		func(name string) ([]byte, error) {
			return Asset(name)
		})
	d, err := bindata.WithInstance(s)
	if err != nil {
		return nil
	}
	m, err := migrate.NewWithSourceInstance("go-bindata", d, config.DatabaseUrl())
	return m
}
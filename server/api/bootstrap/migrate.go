package bootstrap

import (
	"github.com/RomainC75/todo2/config"
	"github.com/RomainC75/todo2/data/database"
	"github.com/RomainC75/todo2/data/migration"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}

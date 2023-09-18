package cmds

import (
	"log"

	"github.com/fabiante/persurl/config"
	"github.com/fabiante/persurl/db"
	"github.com/fabiante/persurl/db/migrations"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate application's database",
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		database, _, err := db.SetupPostgresDB(config.DbDSN(), config.DbMaxConnections())
		if err != nil {
			log.Fatalf("setting up database failed: %s", err)
		}

		err = migrations.RunSQLite(database)
		if err != nil {
			log.Fatalf("migrating database failed: %s", err)
		}
	}

	Root.AddCommand(cmd)
}

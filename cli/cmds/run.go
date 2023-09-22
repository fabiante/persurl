package cmds

import (
	"log"

	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/app"
	"github.com/fabiante/persurl/config"
	"github.com/fabiante/persurl/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the application",
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		_, database, err := db.SetupPostgresDB(config.DbDSN(), config.DbMaxConnections())
		if err != nil {
			log.Fatalf("setting up database failed: %s", err)
		}

		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: database,
		}))
		if err != nil {
			log.Fatalf("setting up gorm database failed: %s", err)
		}

		engine := gin.Default()
		service := app.NewAdminService(gormDB)
		server := api.NewServer(service, service)
		api.SetupRouting(engine, server)
		if err := engine.Run(":8060"); err != nil {
			log.Fatalf("running api failed: %s", err)
		}
	}

	Root.AddCommand(cmd)
}

package cmds

import (
	"log"

	"github.com/fabiante/persurl/api"
	"github.com/fabiante/persurl/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the application",
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		dataDir := envDataDir()

		dbFile := envDbFile(dataDir)

		_, database, err := db.SetupDB(dbFile)
		if err != nil {
			log.Fatalf("setting up database failed: %s", err)
		}

		engine := gin.Default()
		service := db.NewDatabase(database)
		server := api.NewServer(service)
		api.SetupRouting(engine, server)
		if err := engine.Run(":8060"); err != nil {
			log.Fatalf("running api failed: %s", err)
		}
	}

	Root.AddCommand(cmd)
}
package cmd

import (
	"github.com/sinameshkini/microkit/pkg/clients/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"lexilift/internal/repository/entities"
)

var (
	migrateCmd = &cobra.Command{
		Use:     "migrate",
		Aliases: []string{"m"},
		Short:   "Drop and Migrate all database tables",
		Run:     migrate,
	}
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate(cmd *cobra.Command, args []string) {
	//conf := initConfig()

	db, err := database.NewDBWithDsn("host=localhost user=admin password=admin dbname=lexilift port=5432 sslmode=disable", true)
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}

	if err = database.Drop(db.DB(), append(entities.AllTables, "word_tag")); err != nil {
		logrus.Errorln(err.Error())
		return
	}

	if err = database.Migrate(db.DB(), entities.AllTables); err != nil {
		logrus.Errorln(err.Error())
		return
	}
}

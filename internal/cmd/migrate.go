package cmd

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/pkg"

	"github.com/spf13/cobra"
)

var (
	genGormCmd = &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			err := global.GetDB().Migrator().AutoMigrate(
				pkg.Recruitment{},
				pkg.Application{},
				pkg.Interview{},
				pkg.Comment{},
			)
			if err != nil {
				panic(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(genGormCmd)
}

package cmd

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/models"
	"github.com/spf13/cobra"
)

var (
	genGormCmd = &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			err := global.GetDB().Migrator().AutoMigrate(
				//models.ApplicationEntity{},
				//models.CandidateEntity{},
				//models.CommentEntity{},
				//models.InterviewEntity{},
				//models.MemberEntity{},
				//models.RecruitmentEntity{},
				models.RecruitmentEntity{},
				models.ApplicationEntity{},
				models.InterviewEntity{},
				models.CommentEntity{},
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

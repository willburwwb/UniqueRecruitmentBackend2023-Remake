package request

import "UniqueRecruitmentBackend/internal/models"

type CreateCommentRequest struct {
	ApplicationID string            `json:"applicationId"`
	MemberID      string            `json:"memberId"`
	Content       string            `json:"content"`
	Evaluation    models.Evaluation `json:"evaluation"`
}

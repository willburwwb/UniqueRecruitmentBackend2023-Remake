package request

type CreateCommentRequest struct {
	ApplicationID string `json:"applicationId"`
	MemberID      string `json:"memberId"`
	Content       string `json:"content"`
	Evaluation    int    `json:"evaluation"`
}

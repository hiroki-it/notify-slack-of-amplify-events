package inputs

type DetailInput struct {
	AppId         string
	BranchName    string
	JobId         string
	JobStatusType string
}

func NewDetailInput(appId string, branchName string, jobId string, jobStatusType string) *DetailInput {

	return &DetailInput{
		AppId:         appId,
		BranchName:    branchName,
		JobId:         jobId,
		JobStatusType: jobStatusType,
	}
}

package inputs

type EventPostInput struct {
	AppId         string
	BranchName    string
	JobId         string
	JobStatusType string
}

func NewEventPostInput(appId string, branchName string, jobId string, jobStatusType string) *EventPostInput {

	return &EventPostInput{
		AppId:         appId,
		BranchName:    branchName,
		JobId:         jobId,
		JobStatusType: jobStatusType,
	}
}

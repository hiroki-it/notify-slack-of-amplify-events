package inputs

type EventPostInput struct {
	AppId         string
	BranchName    string
	JobId         string
	JobStatusType int
}

func NewEventPostInput(appId string, branchName string, jobId string, jobStatusType int) *EventPostInput {

	return &EventPostInput{
		AppId:         appId,
		BranchName:    branchName,
		JobId:         jobId,
		JobStatusType: jobStatusType,
	}
}

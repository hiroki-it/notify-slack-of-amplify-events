package inputs

type EventPostInput struct {
	AppId      string
	BranchName string
	JobId      string
	JobStatus  string
}

func NewEventPostInput(appId string, branchName string, jobId string, JobStatus string) *EventPostInput {

	return &EventPostInput{
		AppId:      appId,
		BranchName: branchName,
		JobId:      jobId,
		JobStatus:  JobStatus,
	}
}

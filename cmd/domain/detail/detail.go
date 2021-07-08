package detail

type Detail struct {
	appId         *AppId
	branchName    *BranchName
	jobId         *JobId
	jobStatusType JobStatusType
}

// NewDetail コンストラクタ
func NewDetail(appId *AppId, branchName *BranchName, jobId *JobId, jobStatusType JobStatusType) *Detail {

	return &Detail{
		appId:         appId,
		branchName:    branchName,
		jobId:         jobId,
		jobStatusType: jobStatusType,
	}
}

// AppId AppIdを返却します.
func (d *Detail) AppId() *AppId {
	return d.appId
}

// BranchName BranchNameを返却します.
func (d *Detail) BranchName() *BranchName {
	return d.branchName
}

// JobId JobIdを返却します.
func (d *Detail) JobId() *JobId {
	return d.jobId
}

// JobStatusType JobStatusTypeを返却します.
func (d *Detail) JobStatusType() JobStatusType {
	return d.jobStatusType
}

package entities

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/ids"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/values"
)

type Detail struct {
	appId         *ids.AppId
	branchName    *values.BranchName
	jobId         *ids.JobId
	jobStatusType values.JobStatusType
}

// NewDetail コンストラクタ
func NewDetail(appId *ids.AppId, branchName *values.BranchName, jobId *ids.JobId, jobStatusType values.JobStatusType) *Detail {

	return &Detail{
		appId:         appId,
		branchName:    branchName,
		jobId:         jobId,
		jobStatusType: jobStatusType,
	}
}

// AppId AppIdを返却します.
func (d *Detail) AppId() *ids.AppId {
	return d.appId
}

// BranchName BranchNameを返却します.
func (d *Detail) BranchName() *values.BranchName {
	return d.branchName
}

// JobId JobIdを返却します.
func (d *Detail) JobId() *ids.JobId {
	return d.jobId
}

// JobStatusType JobStatusTypeを返却します.
func (d *Detail) JobStatusType() values.JobStatusType {
	return d.jobStatusType
}

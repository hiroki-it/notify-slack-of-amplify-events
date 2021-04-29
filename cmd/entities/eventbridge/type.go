package eventbridge

/**/
type Event struct {
	Version    string   `json:"version"`
	Id         string   `json:"id"`
	DetailType string   `json:"detail-type"`
	Source     string   `json:"source"`
	Account    string   `json:"account"`
	Time       string   `json:"time"`
	Region     string   `json:"region"`
	Resources  []string `json:"resources"`
	Detail     struct {
		AppId      string `json:"appId"`
		BranchName string `json:"branchName"`
		JobId      string `json:"jobId"`
		JobStatus  string `json:"jobStatus"`
	} `json:"detail"`
}

package honorpush

const (
	grantType = "grant_type"
)

const (
	PostRetryTimes = 3 //重试次数
	MaxTimeToLive  = 3600 * 24
)

const (
	AuthHost = "https://iam.developer.honor.com"
	SendHost = "https://push-api.cloud.honor.com"

	AuthURL = "/auth/token"
	SendURL = "/api/v1/%d/sendMessage"
)

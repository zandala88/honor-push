package honorpush

const (
	grantType = "client_credentials"
)

const (
	maxTimeToLive = 3600 * 24
)

const (
	authHost = "https://iam.developer.honor.com"
	sendHost = "https://push-api.cloud.honor.com"

	authURL = "/auth/token"
	sendURL = "/api/v1/%d/sendMessage"
)

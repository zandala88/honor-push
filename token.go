package honorpush

import (
	"time"
)

type HonorToken struct {
	CreateTime  int64  `json:"create_time"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

var tokenInstance *HonorToken

func init() {
	tokenInstance = &HonorToken{
		AccessToken: "",
		CreateTime:  0,
	}
}

// GetToken 获取AccessToken值
func GetToken(clientId, clientSecret string) (*HonorToken, error) {
	nowMilliSecond := time.Now().UnixNano() / 1e6
	if (nowMilliSecond-tokenInstance.CreateTime) < maxTimeToLive*1000 && tokenInstance.AccessToken != "" {
		return tokenInstance, nil
	}
	request, err := post(authHost+authURL).
		header("Content-Type", "application/x-www-form-urlencoded").
		jsonBody(honorTokenReq{
			GrantType:    grantType,
			ClientId:     clientId,
			ClientSecret: clientSecret,
		})
	if err != nil {
		return nil, err
	}

	resp := &HonorToken{}
	if err := request.toJSON(&resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type honorTokenReq struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

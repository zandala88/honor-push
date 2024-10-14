package honorpush

import "fmt"

type HonorPush struct {
	AppId         int
	ClientId      string `json:"client_id"`
	ClientSecret  string `json:"client_secret"`
	Authorization string `json:"authorization"`
}

func NewClient(appId int, clientId, clientSecret string) (*HonorPush, error) {
	token, err := GetToken(clientId, clientSecret)
	if err != nil {
		return nil, err
	}
	return &HonorPush{
		AppId:         appId,
		ClientId:      clientId,
		ClientSecret:  clientSecret,
		Authorization: token.AccessToken,
	}, nil
}

func (h *HonorPush) Send(message *Message) (*HonorSendResp, error) {
	tokenInstance, err := GetToken(h.ClientId, h.ClientSecret)
	if err != nil {
		return nil, err
	}
	request, err := post(sendHost+fmt.Sprintf(sendURL, h.AppId)).
		header("Content-Type", "application/json").
		header("Authorization", tokenInstance.AccessToken).
		jsonBody(message)
	if err != nil {
		return nil, err
	}

	resp := &HonorSendResp{}
	if err := request.toJSON(&resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type HonorSendResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		SendResult   bool          `json:"sendResult"`
		RequestId    string        `json:"requestId"`
		FailTokens   []interface{} `json:"failTokens"`
		ExpireTokens []interface{} `json:"expireTokens"`
	} `json:"data"`
}

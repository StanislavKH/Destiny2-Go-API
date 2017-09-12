package ratking

import (
	"encoding/json"
	"errors"
)

type PlatformService struct {
	client *Client
}

const SuccessStatus = "Success"

type PlatformResponse struct {
	Response        *json.RawMessage `json: Response`
	ErrorCode       *int             `json: ErrorCode`
	ThrottleSeconds *float64         `json: ThrottleSeconds`
	ErrorStatus     *string          `json: ErrorStatus`
	Message         *string          `json: Message`
	MessageData     *interface{}     `json: MessageData, omitempty`
}

func (ps *PlatformService) PlatformRequest(requestType, urlStr string) ([]byte, error) {
	incomingData := PlatformResponse{}

	req, err := ps.client.NewRequest(requestType, urlStr, nil)
	if err != nil {
		return nil, err
	}

	r := new(PlatformResponse)
	response, err := ps.client.Do(req, r)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &incomingData)
	if err != nil {
		return nil, err
	}
	if incomingData.ErrorStatus != nil && *incomingData.ErrorStatus != SuccessStatus {
		return []byte{}, errors.New(*incomingData.Message)
	}
	return []byte(*incomingData.Response), nil
}

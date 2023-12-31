package web_session

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type WebSession struct {
	IP                  string
	ApiToken            string
	maxRequestPerSecond int64
}

func NewWebSession(IP string, ApiToken string, IPThrottling int64, TokenThrottling int64) (WebSession, error) {
	res := WebSession{
		IP:       IP,
		ApiToken: ApiToken,
	}
	err := res.IsValid()
	if err != nil {
		return WebSession{}, err
	}
	if res.ApiToken != "" {
		res.maxRequestPerSecond = TokenThrottling
	} else {
		res.maxRequestPerSecond = IPThrottling
	}
	return res, nil
}

func (h *WebSession) IsValid() error {
	if h.IP == "" {
		return errors.New("invalid IP address")
	}
	return nil
}

func (h *WebSession) GetSessionId() string {
	//The API Token precedes the IP address.
	if h.ApiToken != "" {
		return fmt.Sprintf("%x", sha256.Sum256([]byte(h.ApiToken)))
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(h.IP)))

}

func (h *WebSession) GetRequestCounterId() string {
	return h.GetSessionId() + "_counter"
}

func (h *WebSession) GetRequestsLimit() int64 {
	return h.maxRequestPerSecond
}

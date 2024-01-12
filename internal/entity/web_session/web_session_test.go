package web_session

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/aluferraz/go-expert-rate-limiter/internal/value_objects"
	"github.com/aluferraz/go-expert-rate-limiter/pkg/event_dispatcher"
	"github.com/stretchr/testify/suite"
	"testing"
)

type WebSessionTestSuite struct {
	suite.Suite
	ctx           context.Context
	evtDispatcher *event_dispatcher.MockDispatcher
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(WebSessionTestSuite))
}

func (suite *WebSessionTestSuite) SetupSuite() {

}
func (suite *WebSessionTestSuite) TestWebSessionIP() {

	session, err := NewWebSession("10.0.0.1", "", value_objects.NewRequestLimit(10, 15))
	suite.NoError(err)
	suite.Equal(session.GetSessionId(), fmt.Sprintf("%x", sha256.Sum256([]byte(session.IP))))
	suite.Equal(session.GetRequestCounterId(), fmt.Sprintf("%s%s", session.GetSessionId(), CounterSuffix))

}
func (suite *WebSessionTestSuite) TestWebSessionAPIKEY() {

	session, err := NewWebSession("10.0.0.1", "LUCAO", value_objects.NewRequestLimit(10, 15))
	suite.NoError(err)
	suite.Equal(session.GetSessionId(), fmt.Sprintf("%x", sha256.Sum256([]byte(session.ApiToken))))
	suite.Equal(session.GetRequestCounterId(), fmt.Sprintf("%s%s", session.GetSessionId(), CounterSuffix))

}

package rate_limit

import (
	"context"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RateLimitTestSuite struct {
	suite.Suite
	ctx context.Context
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RateLimitTestSuite))
}

func (suite *RateLimitTestSuite) SetupSuite() {
	suite.ctx = context.Background()
}
func (suite *RateLimitTestSuite) TestRateLimit() {
	db, mock := redismock.NewClientMock()
	print(db, mock)
	/* //TODO: Implement test, here is an inspiration to you!
	       uc := NewUseCase(
	   		&suite.ctx,
	   		InputDTO{},
	   		event_dispatcher.NewDummyEvent(),
	   		suite.evtDispatcher,
	   	)
	   	output, err := uc.Execute()
	   	suite.NoError(err)
	   	suite.NotEmpty(output.ExecutionId)
	   	suite.evtDispatcher.AssertNumberOfCalls(suite.T(), "Dispatch", 1)
	*/
}

package web_session
import (
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
    suite.ctx = context.Background()
	ed := event_dispatcher.NewMockDispatcher()
	ed.On("Dispatch", mock.Anything).Return(nil)
	suite.evtDispatcher = ed
}
func (suite *WebSessionTestSuite) TestWebSession() {
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

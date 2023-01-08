package async_test

import (
	"math"
	"sync"
	"testing"

	"github.com/kuckchuck96/async-exec-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type executorTestSuite struct {
	suite.Suite
	wg           *sync.WaitGroup
	input        int
	output       chan float64
	testFunction func(int) float64
}

func (suite *executorTestSuite) SetupTest() {
	suite.wg = new(sync.WaitGroup)
	suite.input = 2
	suite.output = make(chan float64, 1)
	suite.testFunction = func(i int) float64 {
		return math.Pow10(i)
	}
}

func (suite *executorTestSuite) TestAsyncExecutor() {
	async.Executor(suite.testFunction, suite.input, suite.output, suite.wg)
	suite.wg.Wait()

	result, ok := <-suite.output

	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), result, float64(100))
}

func TestExecutorTestSuite(t *testing.T) {
	suite.Run(t, new(executorTestSuite))
}

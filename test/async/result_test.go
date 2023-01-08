package async_test

import (
	"math"
	"sync"
	"testing"
	"time"

	"github.com/kuckchuck96/async-exec-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type resultTestSuite struct {
	suite.Suite
	wg           *sync.WaitGroup
	duration     time.Duration
	input        int
	output       chan float64
	sleep        time.Duration
	testFunction func(int) float64
}

func (suite *resultTestSuite) SetupTest() {
	suite.wg = new(sync.WaitGroup)
	suite.duration = 5 * time.Second
	suite.input = 2
	suite.output = make(chan float64, 1)
	suite.sleep = 3 * time.Second
	suite.testFunction = func(i int) float64 {
		time.Sleep(suite.sleep)
		return math.Pow10(i)
	}
}

func (suite *resultTestSuite) TestResult() {
	async.Executor(suite.testFunction, suite.input, suite.output, suite.wg)

	result, err := async.Result(suite.output, suite.duration)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), result, float64(100))
}

func (suite *resultTestSuite) TestResult_Timeout() {
	suite.sleep = 10 * time.Second

	async.Executor(suite.testFunction, suite.input, suite.output, suite.wg)

	result, err := async.Result(suite.output, suite.duration)

	assert.Error(suite.T(), err)
	assert.Zero(suite.T(), result)
}

func TestResultTestSuite(t *testing.T) {
	suite.Run(t, new(resultTestSuite))
}

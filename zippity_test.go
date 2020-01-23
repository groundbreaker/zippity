package zippity

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ZippityTestSuite struct {
	suite.Suite
}

func (suite *ZippityTestSuite) TestVersion() {
	suite.NotEmpty(Version(), "it returns the SemVer for this library")
}

func TestZippityTestSuite(t *testing.T) {
	suite.Run(t, new(ZippityTestSuite))
}

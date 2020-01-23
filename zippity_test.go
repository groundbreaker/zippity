package zippity

import (
	"archive/zip"
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ZippityTestSuite struct {
	suite.Suite
	Zipfile *Zipfile
	File1   *File
	File2   *File
	File3   *File
}

func (suite *ZippityTestSuite) SetupTest() {
	suite.Zipfile = New()
	suite.File1 = ReadFile("venture.pdf", "venture.pdf")
	suite.File2 = ReadFile("fine.pdf", "fine.pdf")
	suite.File3 = ReadFile("yezzy.pdf", "yezzy.pdf")
}

func (suite *ZippityTestSuite) TestVersion() {
	suite.NotEmpty(Version(), "it returns the SemVer for this library")
}

func (suite *ZippityTestSuite) TestNew() {
	suite.IsType(&Zipfile{}, New())
	suite.IsType(&bytes.Buffer{}, suite.Zipfile.Body)
	suite.IsType(&zip.Writer{}, suite.Zipfile.Client)
}

func (suite *ZippityTestSuite) TestReadFile() {
	emptyFile := &File{}
	file := ReadFile("venture.pdf", "venture.pdf")
	suite.Equal([]byte(nil), emptyFile.Body)
	suite.IsType(&File{}, file)
	suite.NotEmpty(file.Body)
}

func (suite *ZippityTestSuite) TestAdd() {
	suite.Equal(&bytes.Buffer{}, suite.Zipfile.Body)
	zf := suite.Zipfile.Add(suite.File1)
	suite.NotEqual(&bytes.Buffer{}, zf.Body)
}

func (suite *ZippityTestSuite) TestAddChained() {
	suite.Equal(&bytes.Buffer{}, suite.Zipfile.Body)
	zf := suite.Zipfile.Add(suite.File1).Add(suite.File2).Add(suite.File3)
	suite.NotEqual(&bytes.Buffer{}, zf.Body)
}

func (suite *ZippityTestSuite) TestDone() {
	suite.NotEqual([]byte(nil), suite.Zipfile.Add(suite.File1).Done())
}

func (suite *ZippityTestSuite) TestDoneChained() {
	suite.NotEqual([]byte(nil), suite.Zipfile.Add(suite.File1).Add(suite.File2).Add(suite.File3).Done())
}

func (suite *ZippityTestSuite) TestSave() {
	suite.Zipfile.Add(suite.File1).Save("test.zip")
	suite.True(fileExists("test.zip"))
}

func (suite *ZippityTestSuite) TestSaveChained() {
	suite.Zipfile.Add(suite.File1).Add(suite.File2).Add(suite.File3).Save("test.zip")
	suite.True(fileExists("test.zip"))
}

func TestZippityTestSuite(t *testing.T) {
	suite.Run(t, new(ZippityTestSuite))
}

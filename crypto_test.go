package gocrypto

import (
	//"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

var aesKey = "astaxie12798akljzmknm.ahkjkljl;k"
var input = "test skcloud crypto"
var md5Output = "37d5bf6f759c49527e78f8db7c051276"
var base64Output = "dGVzdCBza2Nsb3VkIGNyeXB0bw=="
var aesOutput = "XGFffDlOwueH4DCVD+5Um/S/OA=="

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestMD5(c *C) {
	out := GetMD5(input)
	c.Assert(out, Equals, md5Output)
	c.Assert(CheckMD5(input, out), Equals, true)
	c.Assert(CheckMD5(input, input), Equals, false)

}

func (s *MySuite) TestBase64(c *C) {
	out := Base64Encode(input)
	c.Assert(out, Equals, base64Output)

	in, e := Base64Decode(base64Output)
	c.Assert(e, Equals, nil)
	c.Assert(in, Equals, input)
}

func (s *MySuite) TestGetRandomKey(c *C) {
	out1 := GetRandomKey()
	out2 := GetRandomKey()
	c.Assert(len(out1), Equals, 40)
	c.Assert(len(out2), Equals, 40)
	c.Assert(out1, Not(Equals), out2)

}

func (s *MySuite) TestAES(c *C) {
	out1, err1 := AESEncode(input, aesKey)
	c.Assert(err1, Equals, nil)
	c.Assert(out1, Equals, aesOutput)

	out2, err2 := AESDecode(aesOutput, aesKey)
	c.Assert(err2, Equals, nil)
	c.Assert(out2, Equals, input)
}

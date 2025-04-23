package cmd /* Copyright © 2022 Ken Pepple <ken@pepple.io> */

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_AlphaNumeric(t *testing.T) {
	Convey("test alphanumeric", t, func() {
		type test struct {
			testname, s string
			want        bool
		}
		tests := []test{
			{testname: "all empty", s: "", want: false},
			{testname: "happy path", s: "u375jggk7cjyvbe4fbgwxrunqmtgbn", want: true},
			{testname: "invalid space only", s: " ", want: false},
		}
		for _, tc := range tests {
			SoMsg(tc.testname, isAlphaNumeric(tc.s), ShouldEqual, tc.want)
		}
	})
}

func Test_CheckPushoverKeys(t *testing.T) {
	Convey("test pushover key check", t, func() {
		type test struct {
			testname, appkey, userkey string
			want                      bool
		}
		tests := []test{
			{testname: "all empty", appkey: "", userkey: "", want: false},
			{testname: "userkey empty", appkey: "a35k5uvmtse8cjmexo3r93dt5oai1g", userkey: "", want: false},
			{testname: "appkey empty", appkey: "", userkey: "a35k5uvmtse8cjmexo3r93dt5oai1g", want: false},
			{testname: "happy path", appkey: "u375jggk7cjyvbe4fbgwxrunqmtgbn", userkey: "a35k5uvmtse8cjmexo3r93dt5oai1g", want: true},
			{testname: "invalid space only appkey", appkey: " ", userkey: "a35k5uvmtse8cjmexo3r93dt5oai1g", want: false},
		}
		for _, tc := range tests {
			SoMsg(tc.testname, checkPushoverKeys(tc.userkey, tc.appkey), ShouldEqual, tc.want)
		}
	})
}

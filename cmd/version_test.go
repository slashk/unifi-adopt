package cmd /* Copyright © 2022 Ken Pepple <kpepple@weedmaps.com> */

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_VersionCmd(t *testing.T) {
	Convey("Ensure version respects variables", t, func() {
		v := "1.2.1"
		d := "2022-10-03T16:59:16Z"
		c := "6dd8734"
		expected := "1.2.1, commit 6dd8734, built at 2022-10-03T16:59:16Z"
		So(PrintVersion(v, c, d), ShouldEqual, expected)
	})
}

func Test_Empty(t *testing.T) {
	Convey("Ensure version prints empty defaults", t, func() {
		expected := "dev, commit none, built at unknown"
		So(PrintVersion(version, commit, date), ShouldEqual, expected)
	})
}

func Test_ExecuteVersion(t *testing.T) {
	Convey("test executable version output", t, func() {
		actual := new(bytes.Buffer)
		rootCmd.SetOut(actual)
		rootCmd.SetErr(actual)
		rootCmd.SetArgs([]string{"version"})
		rootCmd.Execute()

		expected := "dev, commit none, built at unknown\n"

		So(actual.String(), ShouldEqual, expected)
	})
}

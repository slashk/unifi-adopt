package cmd /* Copyright © 2022 Ken Pepple <ken@pepple.io> */

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_ParseInfoHappy(t *testing.T) {
	Convey("test parseInfo for happy path", t, func() {
		s := `Model:       UAP-AC-Pro-Gen2
		Version:     6.2.44.14098
		MAC Address: 74:83:c2:b0:73:45
		IP Address:  10.0.1.63
		Hostname:    ac-pro-1
		Uptime:      1733214 seconds
		Status:      Connected (http://10.0.3.202:30002/inform)`
		parseString, _ := parseInfo(s)
		So(parseString["Status"], ShouldEqual, "Connected (http://10.0.3.202:30002/inform)")
		So(parseString["Model"], ShouldEqual, "UAP-AC-Pro-Gen2")
		So(parseString["IP Address"], ShouldEqual, "10.0.1.63")
		So(parseString["Hostname"], ShouldEqual, "ac-pro-1")
		So(parseString["Version"], ShouldEqual, "6.2.44.14098")
		So(parseString["MAC Address"], ShouldEqual, "74:83:c2:b0:73:45")
		So(parseString["Uptime"], ShouldEqual, "1733214 seconds")
	})
}

func Test_ParseInfoEmpty(t *testing.T) {
	Convey("test parseInfo for empty info results", t, func() {
		s := ``
		_, err := parseInfo(s)
		So(err, ShouldBeError)
	})
	Convey("test parseInfo for partial info results", t, func() {
		s := `Model:       UAP-AC-Pro-Gen2
		Version:     6.2.44.14098
		MAC Address: 74:83:c2:b0:73:45`
		parseString, err := parseInfo(s)
		So(err, ShouldBeError)
		So(parseString["Model"], ShouldEqual, "UAP-AC-Pro-Gen2")
		So(parseString["Version"], ShouldEqual, "6.2.44.14098")
		So(parseString["MAC Address"], ShouldEqual, "74:83:c2:b0:73:45")
	})
}

func Test_ParseInfoIsolated(t *testing.T) {
	Convey("test parseInfo for happy path", t, func() {
		s := `Model:       UAP-AC-Pro-Gen2
		Version:     6.2.44.14098
		MAC Address: 74:83:c2:b0:73:45
		IP Address:  10.0.1.63
		Hostname:    ac-pro-1
		Uptime:      1733214 seconds
		Status:      Isolated`
		parseString, _ := parseInfo(s)
		So(parseString["Status"], ShouldEqual, "Isolated")
		So(parseString["Model"], ShouldEqual, "UAP-AC-Pro-Gen2")
	})
}

func Test_ParseWAP(t *testing.T) {
	Convey("test parsewap for happy path", t, func() {
		s := "10.0.1.56,10.0.1.63,10.0.1.61,10.0.1.55,10.0.1.51"
		parseIPs, err := parseWAP(s)
		So(len(parseIPs), ShouldEqual, 5)
		So(err, ShouldBeNil)
	})
	Convey("test parsewap for empty", t, func() {
		s := ""
		parseIPs, err := parseWAP(s)
		So(err, ShouldBeError)
		So(parseIPs, ShouldBeNil)
	})
}

func TestParseInfoValidInput(t *testing.T) {
	s := `Model: Router
Version: 1.2.3
MAC Address: 00:1A:2B:3C:4D:5E
IP Address: 192.168.1.1
Hostname: router.local
Uptime: 1 day
Status: Connected`
	expected := map[string]string{
		"Model":       "Router",
		"Version":     "1.2.3",
		"MAC Address": "00:1A:2B:3C:4D:5E",
		"IP Address":  "192.168.1.1",
		"Hostname":    "router.local",
		"Uptime":      "1 day",
		"Status":      "Connected",
	}
	result, err := parseInfo(s)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	for k, v := range expected {
		if result[k] != v {
			t.Errorf("Expected %s to be %s but got %s", k, v, result[k])
		}
	}
}

func TestParseInfoInvalidInput(t *testing.T) {
	s := `Model: Router
Version: 1.2.3
MAC Address: 00:1A:2B:3C:4D:5E
IP Address: 192.168.1.1
Hostname: router.local`
	_, err := parseInfo(s)
	if err == nil || err.Error() != "missing info configs" {
		t.Errorf("Expected error 'missing info configs' but got %v", err)
	}
}

func TestParseInfoExtraData(t *testing.T) {
	s := `Model: Router
Version: 1.2.3
MAC Address: 00:1A:2B:3C:4D:5E
IP Address: 192.168.1.1
Hostname: router.local
Uptime: 1 day
Status: Connected
Extra: Data`
	expected := map[string]string{
		"Model":       "Router",
		"Version":     "1.2.3",
		"MAC Address": "00:1A:2B:3C:4D:5E",
		"IP Address":  "192.168.1.1",
		"Hostname":    "router.local",
		"Uptime":      "1 day",
		"Status":      "Connected",
	}
	result, err := parseInfo(s)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	for k, v := range expected {
		if result[k] != v {
			t.Errorf("Expected %s to be %s but got %s", k, v, result[k])
		}
	}
}

func TestParseInfoEmptyInput(t *testing.T) {
	s := ""
	_, err := parseInfo(s)
	if err == nil || err.Error() != "missing info configs" {
		t.Errorf("Expected error 'missing info configs' but got %v", err)
	}
}

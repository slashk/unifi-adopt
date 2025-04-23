package cmd /* Copyright © 2022 Ken Pepple <ken@pepple.io> */

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/helloyi/go-sshclient"
)

const infoCommand = "mca-cli-op info"
const informCommand = "mca-cli-op set-inform http://%s/inform; echo \"exit:\" $?" // %s allow for inform IP

// checkConnected queries the WAP for it's connection status through the info command
func checkConnected(ip, username, certfile string) (bool, error) {
	if debug {
		fmt.Println("Checking ", ip)
	}
	if debug {
		fmt.Printf("ssh -i %s %s@%v\n", certfile, username, ip)
	}
	client, err := sshclient.DialWithKey(ip+":22", username, certfile)
	if err != nil {
		fmt.Println("dial error: ", err)
		return false, err
	}
	defer client.Close()
	out, err := client.Cmd(infoCommand).SmartOutput()
	if err != nil {
		// the 'out' is stderr output
		fmt.Println("ssh command error:", err, "\n", out)
		return false, err
	}
	// the 'out' is stdout output
	ws, _ := parseInfo(string(out))
	if debug {
		fmt.Println(ws)
	}
	if status, exists := ws["Status"]; exists {
		if status != "Connected (http://10.0.3.202:30002/inform)" {
			return false, errors.New(ws["Hostname"] + " " + ws["Status"])
		} else {
			fmt.Printf("%s (%s) is %s\n", ws["Hostname"], ws["Model"], ws["Status"])
		}
	}
	if debug {
		fmt.Println("completed check")
	}
	return true, nil
}

// setInform sets the WAPs config to the informURL string provided
// % ssh -i ~/.ssh/unifi kenpepple@10.0.1.63 'mca-cli-op set-inform http://10.0.3.202:30002/inform; echo "exit:" $?'
//
// Adoption request sent to 'http://10.0.3.202:30002/inform'.  Use the controller to complete the adopt process.
//
// exit: 0
func setInform(ip, username, certfile, informURL string) (bool, error) {
	client, err := sshclient.DialWithKey(ip+":22", username, certfile)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()
	i := fmt.Sprintf(informCommand, informURL)
	out, err := client.Cmd(i).SmartOutput()
	if err != nil {
		// the 'out' is stderr output
		fmt.Println(out)
	}
	// the 'out' is stdout output
	// fmt.Println(string(out))
	return true, nil
}

// parseInfo returns the Status component of Unifi's info command
// ac-pro-1-BZ.6.2.44# mca-cli-op info
// Model:       UAP-AC-Pro-Gen2
// Version:     6.2.44.14098
// MAC Address: 74:83:c2:b0:73:45
// IP Address:  10.0.1.63
// Hostname:    ac-pro-1
// Uptime:      1733214 seconds
// Status:      Connected (http://10.0.3.202:30002/inform)
func parseInfo(s string) (map[string]string, error) {
	s += "\n"
	ws := make(map[string]string)
	// search string for Status line and return it? or parse for Connected with correct inform URL ?
	stat := []string{"Model", "Version", "MAC Address", "IP Address", "Hostname", "Uptime", "Status"}
	for x := range stat {
		// use regex page at https://regex101.com/
		r, _ := regexp.Compile(stat[x] + ":(.+)\n")
		match := r.FindAllStringSubmatch(s, 1)
		if match != nil {
			ws[stat[x]] = strings.TrimSpace(match[0][1])
		} else {
			return ws, errors.New("missing info configs")
		}
	}
	// fmt.Println(ws)
	return ws, nil
}

func parseWAP(s string) ([]string, error) {
	if s == "" {
		return nil, errors.New("improper or empty WAP address/name formatting")
	}
	x := strings.Split(s, ",")
	return x, nil
}

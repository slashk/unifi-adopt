/*  Copyright © 2022 Ken Pepple <kpepple@weedmaps.com>  */

package cmd

import (
	"regexp"

	"github.com/gregdel/pushover"
)

const ddURL = "https://app.datadoghq.com/billing/usage"

func sendPush(m, PushoverRecp, PushoverToken string) error {
	app := pushover.New(PushoverToken)
	recipient := pushover.NewRecipient(PushoverRecp)
	message := &pushover.Message{
		Message:  m,
		Title:    "Datadog Charges",
		Priority: pushover.PriorityNormal,
		URL:      ddURL,
		URLTitle: "Datadog Usage Link",
	}
	_, err := app.SendMessage(message, recipient)
	return err
}

func checkPushoverKeys(PushoverRecp, PushoverToken string) bool {
	if isAlphaNumeric(PushoverRecp) && isAlphaNumeric(PushoverToken) {
		return true
	}
	return false
}

func isAlphaNumeric(s string) bool {
	if s == "" {
		return false
	}
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(s)
}

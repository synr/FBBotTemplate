// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"bytes"
)

var mess = &Messenger{}

func main() {
	port := os.Getenv("PORT")
	log.Println("Server start in port:", port)
	mess.VerifyToken = os.Getenv("TOKEN")
	mess.AccessToken = os.Getenv("TOKEN")
	log.Println("Bot start in token:", mess.VerifyToken)
	mess.MessageReceived = MessageReceived
	http.HandleFunc("/webhook", mess.Handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func HttpPost_JANDI(body, connectColor, title string) error {
	log.Print("已經進來 JANDI POST")
	url := "https://wh.jandi.com/connect-api/webhook/11691684/46e7f45fd4f68a021afbd844aed66430"
	jsonStr := `{
		"body":"` + body + `",
		"connectColor":"` + connectColor + `",
		"connectInfo" : [{
				"title" : "` + title + `",
				"description" : "這是來自 LINE BOT 的通風報信",
				"imageUrl": "https://line.me/R/ti/p/@pyv6283b"
		}]
	}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		log.Print(err)
		return err
	}

	// Content-Type 設定
	req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)		
		return err
	}
	defer resp.Body.Close()

	log.Print(err)
	return err
}

func HttpPost_IFTTT(body string) error {
	//https://internal-api.ifttt.com/maker
	log.Print("已經進來 IFTTT POST")
	url := "https://maker.ifttt.com/trigger/linebot/with/key/WJCRNxQhGJuzPd-sUDext"
	jsonStr := `{
		"value1":"` + body + `",
		"value2":"這是 LINE BOT 的同步通知",
		"value3": "由 Heroku 的 GO 語言寫成"
	}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		log.Print(err)
		return err
	}

	// Content-Type 設定
	req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)		
		return err
	}
	defer resp.Body.Close()

	log.Print(err)
	return err
}

//MessageReceived :Callback to handle when message received.
func MessageReceived(event Event, opts MessageOpts, msg ReceivedMessage) {
	// log.Println("event:", event, " opt:", opts, " msg:", msg)
	profile, err := mess.GetProfile(opts.Sender.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	HttpPost_JANDI(opts.Sender.ID + " 從 FB 說：" + msg.Text, "yellow" , "FB 對話同步")
	HttpPost_IFTTT(opts.Sender.ID + " 從 FB 說：" + msg.Text)

	//resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("Hello   , %s %s, %s", profile.FirstName, profile.LastName, msg.Text))
	resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf(msg.Text))
	log.Print("profile.FirstName = " + profile.FirstName)
	log.Print("profile.LastName = " + profile.LastName)
	log.Print("opts.Sender.ID = " +o pts.Sender.ID)
	log.Print("msg.Text = " + msg.Text)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}

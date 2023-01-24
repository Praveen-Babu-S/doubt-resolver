package notify

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/backend-ids/utils"
)

var AdminMail = "admin@ids.in"

// This function returns mail body
func Mailinfo(receiver string, sub string, msg string) string {
	mailContent := fmt.Sprintf(`{
        "personalizations": [
            {
                "to": [
                    {
                        "email": `+`"%s"`+`
                    }
                ],
                "subject": `+`"%s"`+`
            }
        ],
        "from": {
            "email": `+`"%s"`+`
        },
        "content": [
            {
                "type": "text/plain",
                "value":`+`"%s"`+`
            }
        ]
    }`, receiver, sub, AdminMail, msg)
	return mailContent
}

// This function is used to send mail telling that account is created
func Notify(mail string) {
	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"
	payload := strings.NewReader(mail)
	// fmt.Println(mail)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "6d4d4c82c2mshda87a108b5a2271p19106cjsn908cefcbc7e3")
	req.Header.Add("X-RapidAPI-Host", "rapidprod-sendgrid-v1.p.rapidapi.com")
	res, err := http.DefaultClient.Do(req)
	utils.CheckErr(err)
	defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(res)
	// fmt.Println(string(body))
}

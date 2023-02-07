package sendmail

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Sendmail() {

	from := mail.NewEmail("From_name", "From_mail")
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("client_name", "client_mail")
	plainTextContent := "BHUPP"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("AWS_ACCESS_KEY_ID"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

//Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")

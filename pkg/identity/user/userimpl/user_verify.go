package userimpl

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user"
)

func (s *service) SendVerifyMemberEmail(args *user.UserVerifyEmail) {
	from := "nopromise1999@gmail.com"
	pass := "nitranhngao@81299@"

	to := []string{
		args.Email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	t, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}
	
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: [CoffeeShop] Please verify your email \n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct {
		Email   string
		Code    string
		Message string
	}{
		Email:   args.Email,
		Code:    args.Code,
		Message: "Verification code",
	})

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	fmt.Println(err)
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}

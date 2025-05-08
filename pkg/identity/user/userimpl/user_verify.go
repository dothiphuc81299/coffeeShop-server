package userimpl

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/smtp"
	"text/template"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/code"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) SendVerifyMemberEmail(args *user.UserVerifyEmail) {
	from := ""
	pass := ""

	to := []string{
		args.Email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	t, err := template.ParseFiles("templates/template.html")
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
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}

func (s *service) VerifyEmail(ctx context.Context, args user.VerifyEmailCommand) error {
	// result := redisapp.GetValueByKey(args.Code)
	// var res string
	// if result == "" {
	// 	return fmt.Errorf("code not found")
	// }
	// if err := json.Unmarshal([]byte(result), &res); err != nil {
	// 	return err
	// }

	// if res != args.Email {
	// 	fmt.Println("ok")
	// 	return fmt.Errorf("Email Khong hop le")
	// }

	result, err := s.codeStore.FindOneByCondition(ctx, bson.M{"email": args.Email})
	if err != nil {
		return err
	}

	if result.Code != args.Code {
		return fmt.Errorf(" Khong hop le")
	}

	err = s.store.UpdateByCondition(ctx, bson.M{"email": args.Email}, bson.M{"$set": bson.M{"active": true}})
	if err != nil {
		return err
	}

	// err = redisapp.DelKey(args.Code)
	err = s.codeStore.DeleteOne(ctx, args.Email)
	if err != nil {
		return err
	}
	return nil

}

func (s *service) SendEmail(ctx context.Context, mail user.SendUserEmailCommand) error {
	otp, _ := GenerateOTP(6)
	//err := redisapp.SetKeyValue(code, mail.Email, 24*time.Hour)

	argsCode := code.CodedRegisterRaw{
		Id:    primitive.NewObjectID(),
		Email: mail.Email,
		Code:  otp,
	}
	err := s.codeStore.InsertOne(ctx, argsCode)

	if err != nil {
		return err
	}
	fmt.Println(err)
	mailw := mail.Email
	args := &user.UserVerifyEmail{
		Email: mailw,
		Code:  otp,
	}
	s.SendVerifyMemberEmail(args)

	return nil

}

const otpChars = "1234567890"

func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

package config
import (
	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(email, verificationLink string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "golang-api@gmail.com")
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Verify Your Email")
	mailer.SetBody("text/html", "Click the link below to verify your email:<br><a href='"+verificationLink+"'>"+verificationLink+"</a>")

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "houssamtkd03@gmail.com", "ibxu ibtu sxot snzr")
	return dialer.DialAndSend(mailer)
}

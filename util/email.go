package util

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/thteam47/common-libs/errutil"
)

func SendMail(to []string, data string) error {
	from := "thaianhanhthai4@gmail.com"
	password := "gpiwwgpjmszenwtu"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", from, password, smtpHost)
	msg := []byte(fmt.Sprintf(
		"From:  %s\r\n"+
			"To: "+strings.Join(to, ",")+"\r\n"+
			"Subject: ThteaM\r\n"+
			"\r\n%s\r\n", "ThteaM", data))
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		return errutil.Wrapf(err, "smtp.SendMail")
	}
	return nil
}

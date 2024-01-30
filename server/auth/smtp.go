package auth

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"

	"crypto/rand"
	"encoding/base64"

	"github.com/johnkristanf/VoiceForge/server/utils"
)

type SmtpClientMethod interface{
	SendVerificationEmail(string) (int64, error)
}

type SmtpClient struct{
	smtp *smtp.Client
}

var (

	host = "smtp.gmail.com"
    port = "465"
	password = "bkzd cgbb wywn qmvi"
	from = "johnkristan01@gmail.com"
	tokenLength = 24
)

func SmtpConfig() (*SmtpClient, error) {

	TLSconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName: host,
	}

	serverAddress := fmt.Sprintf("%s:%s", host, port)

	connection, err := tls.Dial("tcp", serverAddress, TLSconfig) 
	if err != nil {
		return nil, err
	}

	smtpClient, err := smtp.NewClient(connection, host)
	if err != nil{
		log.Panic(err)
	}

	return &SmtpClient{
		smtp: smtpClient,
	}, nil
}

func GenerateRandomToken(length int) (string, error) {

    token := make([]byte, length)

    _, err := rand.Read(token)
    if err != nil {
        return "", err
    }

    
    return base64.URLEncoding.EncodeToString(token)[:length], nil
}


func (s *SmtpClient) SendVerificationEmail(to string) (int64, error) {

    verificationCode, err := utils.SecureRandomNumber(5)
    if err != nil {
       return 0, err
    }

	fmt.Println("verificationCode", verificationCode)

	body := fmt.Sprintf("Your Verification Code:%d", verificationCode)

	headers := map[string]string{
		"From": from,
		"To": to,
		"Subject": "VoiceForge Email Verification",
	}

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r", k, v)
	}
	message += "\r" + body

	authenticate := smtp.PlainAuth("", from, password, host)


	if err := s.smtp.Auth(authenticate); err != nil{
		return 0, err
	}

	if err := s.smtp.Mail(from); err != nil{
		return 0, err
	}

	if err := s.smtp.Rcpt(headers["To"]); err != nil{
		return 0, err
	}

	writer, err := s.smtp.Data()
    if err != nil {
        return 0, err
    }

    _, err = writer.Write([]byte(message))
    if err != nil {
        return 0, err
    }

	if err := writer.Close(); err != nil{
		return 0, err
	}

	if err := s.smtp.Quit(); err != nil{
		return 0, err
	}

	return verificationCode, nil
}


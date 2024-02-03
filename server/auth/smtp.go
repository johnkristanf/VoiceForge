package auth

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"net/smtp"
	"github.com/johnkristanf/VoiceForge/server/utils"
)

type SmtpClientMethod interface{
	SendVerificationEmail(string) (int64, error)
}

type SmtpClient struct {
    host      string
    port      string
    password  string
    from      string
    smtp      *smtp.Client
}

var (
	tokenLength = 5
)

func NewSmtpClient() (*SmtpClient, error){

	return &SmtpClient{
		host: os.Getenv("SMTP_HOST"),
		port: os.Getenv("SMTP_PORT"),
		password: os.Getenv("SMTP_PASSWORD"),
		from:  os.Getenv("SMTP_FROM"),
	}, nil
}

func (s *SmtpClient) Connect() error {

	TLSconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName: s.host,
	}

	serverAddress := fmt.Sprintf("%s:%s", s.host, s.port)

	connection, err := tls.Dial("tcp", serverAddress, TLSconfig) 
	if err != nil {
		return err
	}

	smtpClient, err := smtp.NewClient(connection, s.host)
	if err != nil{
		log.Panic(err)
	}

	s.smtp = smtpClient
	return nil
}

func (s *SmtpClient) Disconnect() {
    if s.smtp != nil {
        _ = s.smtp.Quit()
    }
}


func (s *SmtpClient) SendVerificationEmail(to string) (int64, error) {

	if err := s.Connect(); err != nil {
        return 0, err
    }
    defer s.Disconnect()

    verificationCode, err := utils.SecureRandomNumber(tokenLength)
    if err != nil {
       return 0, err
    }

	body := fmt.Sprintf("Your Verification Code:%d", verificationCode)

	headers := map[string]string{
		"From": s.from,
		"To": to,
		"Subject": "VoiceForge Email Verification",
	}

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r", k, v)
	}
	message += "\r" + body

	authenticate := smtp.PlainAuth("", s.from, s.password, s.host)


	if err := s.smtp.Auth(authenticate); err != nil{
		return 0, err
	}

	if err := s.smtp.Mail(s.from); err != nil{
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

	return verificationCode, nil
}


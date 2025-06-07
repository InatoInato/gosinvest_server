package utils

import "log"

func SendEmail(to, subject, body string) {
	log.Printf("[EMAIL] To: %s | Subject: %s | Body: %s\n", to, subject, body)
}

func SendSMS(to, message string){
	log.Printf("[SMS] To: %s | Message: %s\n", to, message)
}
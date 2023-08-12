package helpers

import (
	"fmt"
	"log"
	"time"

	"github.com/Nrhlzh-18/todo-app/models"
	"gopkg.in/gomail.v2"
)

func SendEmailForTomorrowSchedules(schedules []models.MSchedule) {
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "zeee.h66@gmail.com", "Nurhalizah18")

	tomorrow := time.Now().Add(24 * time.Hour).Truncate(24 * time.Hour)
	for _, s := range schedules {
		if s.Date.Year() == tomorrow.Year() && s.Date.Month() == tomorrow.Month() && s.Date.Day() == tomorrow.Day() {
			mail := gomail.NewMessage()
			mail.SetHeader("From", "zeee.h66@gmail.com")
			mail.SetHeader("To", "halizahn241@gmail.com")
			mail.SetHeader("Subject", "Reminder: Schedule for Tomorrow")
			mail.SetBody("text/plain", fmt.Sprintf("Schedule with ID %d is tomorrow.", s.ID))

			if err := dialer.DialAndSend(mail); err != nil {
				log.Println("Error sending email:", err)
			} else {
				log.Println("Email sent for schedule ID:", s.ID)
			}
		}
	}
}

package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

type Reminder struct {
	UserID    string
	Message   string
	TriggerAt time.Time
}
type ReminderManager struct {
	Reminders []Reminder
}

func (rm *ReminderManager) AddReminder(userID, message string, triggerAt time.Time) {
	reminder := Reminder{
		UserID:    userID,
		Message:   message,
		TriggerAt: triggerAt,
	}
	rm.Reminders = append(rm.Reminders, reminder)
}

// CheckReminders checks for triggered reminders and sends notifications.
func (rm *ReminderManager) CheckReminders(session *discordgo.Session) {
	now := time.Now()
	var updatedReminders []Reminder

	for _, reminder := range rm.Reminders {
		if now.After(reminder.TriggerAt) {
			// Send a reminder message
			user, err := session.User(reminder.UserID)
			if err != nil {
				log.Printf("Error getting user info: %v", err)
				continue
			}
			message := fmt.Sprintf("Reminder for %s: %s", user.Username, reminder.Message)
			session.ChannelMessageSend(reminder.UserID, message)
		} else {
			updatedReminders = append(updatedReminders, reminder)
		}
	}

	rm.Reminders = updatedReminders
}

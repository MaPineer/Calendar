package scheduler

import (
	"Calendar/models"
	"Calendar/utils"
	"fmt"
	"log"
	"sync"
	"time"
)

var mu sync.Mutex

func StartReminderChecker() {
	//启动websocket监听
	go utils.Start()

	for {
		time.Sleep(time.Second)
		mu.Lock()
		for _, reminders := range models.GetAll() {
			for index, reminder := range reminders {
				reminderTime, _ := time.ParseInLocation("2006-01-02 15:04:05", reminder.Time, time.Local)
				if time.Now().After(reminderTime) {
					var notifier models.Notifier

					switch reminder.NotificationType {
					case "email":
						notifier = &models.EmailNotifier{}
					case "sms":
						notifier = &models.MessageNotifier{}
					default:
						log.Printf("Unknown notification type: %s", reminder.NotificationType)
						continue
					}
					fmt.Println("已经通过" + reminder.NotificationType + "发送给" + reminder.CreatorID)
					notifier.Send(reminder.CreatorID, reminder.Content)
					utils.Broadcast <- models.Reminder{Content: reminder.Content, CreatorID: reminder.CreatorID}
					models.Delete(reminder.CreatorID, index)
				}
			}
		}
		mu.Unlock()
	}
}

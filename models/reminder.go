package models

type Reminder struct {
	ID               string `json:"id"`                //提醒消息的ID
	Content          string `json:"content"`           //提醒内容
	Time             string `json:"time"`              //提醒时间
	CreatorID        string `json:"creator_id"`        //创建者的ID
	NotificationType string `json:"notification_type"` //通知方式
	Index            int    //每个用户的消息序号
}

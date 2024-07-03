package handlers

import (
	"Calendar/models"
	"Calendar/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/reminder", addReminder)
	router.GET("/reminders/:creatorID", getReminders)
	router.PUT("/reminder/:creatorID/:index", updateReminder)
	router.DELETE("/reminder/:creatorID/:index", deleteReminder)
	router.GET("/ws", func(c *gin.Context) {
		utils.HandleConnections(c.Writer, c.Request)
	})
}

func addReminder(c *gin.Context) {
	var newReminder models.Reminder
	if err := c.ShouldBindJSON(&newReminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Set(newReminder.CreatorID, newReminder)

	c.JSON(http.StatusOK, gin.H{"status": "提醒添加成功"})
}

func getReminders(c *gin.Context) {
	creatorID := c.Param("creatorID")
	var creatorReminders = models.Get(creatorID)

	c.JSON(http.StatusOK, creatorReminders)
}

func updateReminder(c *gin.Context) {
	creatorID := c.Param("creatorID")
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		// 处理转换错误
		c.JSON(400, gin.H{"error": "Invalid index"})
		return
	}
	var updatedReminder models.Reminder
	if err := c.ShouldBindJSON(&updatedReminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reminders := models.Get(creatorID)
	if index >= len(reminders) {
		c.JSON(http.StatusForbidden, gin.H{"error": "要更新的提醒不存在"})
		return
	}

	var reminder = reminders[index]

	if reminder.CreatorID != updatedReminder.CreatorID {
		c.JSON(http.StatusForbidden, gin.H{"error": "您只能修改自己的提醒"})
		return
	}

	models.Update(creatorID, index, updatedReminder)
	c.JSON(http.StatusOK, gin.H{"status": "提醒更新成功"})
}

func deleteReminder(c *gin.Context) {
	creatorID := c.Param("creatorID")
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		// 处理转换错误
		c.JSON(400, gin.H{"error": "Invalid index"})
		return
	}

	models.Delete(creatorID, index)
	c.JSON(http.StatusOK, gin.H{"status": "提醒删除成功"})
}

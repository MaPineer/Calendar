package models

import (
	"fmt"
	"sync"
)

// 定义一个全局的map和互斥锁
var (
	globalMap = make(map[string][]Reminder)
	mutex     sync.Mutex
)

// Set 设置map中的键值对
func Set(key string, value Reminder) {
	mutex.Lock()
	defer mutex.Unlock()
	//设置提醒序号
	value.Index = len(globalMap[key])
	globalMap[key] = append(globalMap[key], value)
}

// Get 获取map中的值
func Get(key string) []Reminder {
	mutex.Lock()
	defer mutex.Unlock()

	var creatorReminders []Reminder

	cnt := 0
	for _, reminder := range globalMap[key] {
		reminder.Index = cnt
		creatorReminders = append(creatorReminders, reminder)
		cnt++
	}
	return creatorReminders
}

func Delete(key string, index int) {
	mutex.Lock()
	defer mutex.Unlock()
	slice := globalMap[key]
	if index >= 0 && index < len(slice) {
		slice = append(slice[:index], slice[index+1:]...)
	}
	globalMap[key] = slice
}

// 获取所有的用户和他们的reminder
func GetAll() [][]Reminder {
	values := [][]Reminder{}
	for _, value := range globalMap {
		values = append(values, value)
	}
	return values
}

func Update(key string, index int, reminder Reminder) {
	mutex.Lock()
	defer mutex.Unlock()
	globalMap[key][index] = reminder
	fmt.Println(globalMap[key][index])
}

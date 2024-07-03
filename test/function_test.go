package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

/*
测试功能:为当前用户添加一个提醒

url := "http://127.0.0.1:8080/reminder"
*/
func TestAddReminder(t *testing.T) {
	url := "http://127.0.0.1:8080/reminder"
	reminderJSON := `{
						"id":"1",
						"creator_id":"user1", 
						"content":"Test Post", 
						"time":"2024-07-04 00:03:05", 
						"notification_type": "email"
				      }`

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(reminderJSON))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建 HTTP 客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// 打印响应状态和响应体
	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Response body: %s\n", string(body))
}

/*
测试功能:获取当前用户的所有提醒

url := "http://127.0.0.1:8080/reminder/:creatorID"
参数解释:在使用这个api的时候 Path会附带上当前创建提醒人的id
creatorID: 创建提醒人的id
*/
func TestGetReminders(t *testing.T) {
	url := "http://127.0.0.1:8080/reminder/user1"

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建 HTTP 客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// 打印响应状态和响应体
	fmt.Printf("Response status: %s\n", resp.Status)
	if body != nil {
		fmt.Printf("Response body: %s\n", string(body))
	}

}

/*
测试功能:更新用户的某个提醒事项

url := "http://127.0.0.1:8080/reminder/:creatorID/:index"
参数解释:在使用这个api的时候 Path会附带上当前创建提醒人的id和需要修改的提醒事项下标
creatorID: 创建提醒人的id
index: 提醒事项的下标
*/
func TestUpdateReminder(t *testing.T) {
	url := "http://127.0.0.1:8080/reminder/user1/0"
	reminderJSON := `{
						"id":"1",
						"creator_id":"user1", 
						"content":"Test Update", 
						"time":"2024-07-03 11:10:35", 
						"notification_type": "email"
				      }`

	// 创建请求
	req, err := http.NewRequest("PUT", url, bytes.NewBufferString(reminderJSON))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建 HTTP 客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// 打印响应状态和响应体
	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Response body: %s\n", string(body))
}

/*
测试功能:删除用户的某个提醒事项

url := "http://127.0.0.1:8080/reminder/:creatorID/:index"
参数解释:在使用这个api的时候 Path会附带上当前创建提醒人的id和需要修改的提醒事项下标
creatorID: 创建提醒人的id
index: 提醒事项的下标
*/
func TestDeleteReminder(t *testing.T) {
	url := "http://127.0.0.1:8080/reminder/user1/0"

	// 创建请求
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建 HTTP 客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// 打印响应状态和响应体
	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Response body: %s\n", string(body))
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// 消息结构体
type Message struct {
	Receiver          string       `json:"receiver"`
	Status            string       `json:"firing"`
	Alerts            []Alerts     `json:"alerts"`
	GroupLabels       GroupLabels  `json:"groupLabels"`
	CommonLabels      CommonLabels `json:"commonLabels"`
	CommonAnnotations string       `json:"commonAnnotations"`
	ExternalURL       string       `json:"externalURL"`
	Version           string       `json:"version"`
	GroupKey          string       `json:"groupKey"`
}

/**
 * 消息结构体Alerts
 */

type Alerts struct {
	Status       string      `json:"status"`
	Labels       Labels      `json:"labels"`
	Annotations  Annotations `json:"annotations"`
	StartsAt     string      `json:"startsAt"`
	EndsAt       string      `json:"endsAt"`
	GeneratorURL string      `json:"generatorURL"`
}

/**
 * 消息结构体Alerts.labels
 */

type Labels struct {
	Alertname string `json:"alertname"`
	Instance  string `json:"instance"`
	Job       string `json:"job"`
	Role      string `json:"role"`
	Service   string `json:"service"`
}

/**
 * 消息结构体Alerts.annotations
 */

type Annotations struct {
	Description string `json:"Description"`
	Summary     string `json:"Summary"`
}

/**
 * 消息结构体Alerts.GroupLabels
 */

type GroupLabels struct {
	Alertname string `json:"alertname"`
}

/**
 * 消息结构体Alerts.CommonLabels
 */

type CommonLabels struct {
	Alertname string `json:"alertname"`
	Job       string `json:"job"`
	Role      string `json:"role"`
	Service   string `json:"service"`
}

/**
 * ResponseMessage返回的消息
 */

type ResponseMessage struct {
	Status      string `json:"status"`
	Alertname   string `json:"alertname"`
	Host        string `json:"host"`
	Role        string `json:"role"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
	StartsAt    string `json:"startsAt"`
	EndsAt      string `json:"endsAt"`
}

/**
 * 发送给telegram的消息格式
 */

type Telegram struct {
	Chat_id              string `json:"chat_id"`
	Text                 string `json:"text"`
	Disable_notification bool   `json:"disable_notification"`
}

func SendMessage(response []ResponseMessage, url string, chat string) {

	for _, v := range response {
		//var responseStr string
		var (
			message Telegram
		)

		message = Telegram{
			Chat_id:              chat,
			Text:                 message.ParuseUri(v),
			Disable_notification: true,
		}

		data, _ := json.Marshal(message)

		message.ReqPost(url, data)
	}

}

func (this *Telegram) ParuseUri(req ResponseMessage) (UriText string) {

	UriText = fmt.Sprintf("告警类型: %s\n事件类型: %s\n告警主机: %s\n主机角色: %s\n告警摘要: %s\n告警描述: %s\n故障开始时间: %s\nEndsAt: %s\n", req.Status, req.Alertname, req.Host, req.Role, req.Summary, req.Description, req.StartsAt, req.EndsAt)

	return
}

func (this *Telegram) ReqPost(url string, data []byte) {
	req := &fasthttp.Request{}

	req.SetRequestURI(url)

	req.SetBody(data)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}

	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	b := resp.Body()

	fmt.Println("result:\r\n", string(b))
}

func (this *Message) FormatBody(msg Message) (response []ResponseMessage) {

	for _, v := range msg.Alerts {
		var (
			mess ResponseMessage
		)

		mess = ResponseMessage{
			Status:      v.Status,
			Alertname:   v.Labels.Alertname,
			Host:        v.Labels.Instance,
			Role:        v.Labels.Role,
			Description: v.Annotations.Description,
			Summary:     v.Annotations.Summary,
			StartsAt:    v.StartsAt,
			EndsAt:      v.EndsAt,
		}

		response = append(response, mess)
	}
	return
}

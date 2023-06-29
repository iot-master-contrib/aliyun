package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
)

func Subscribe() {
	mqtt.Client.Subscribe("notify/+/+", 0, func(client paho.Client, message paho.Message) {
		//topics := strings.Split(message.Topic(), "/")
		//pid := topics[1]
		//id := topics[2]

		//解析数据
		var notify model.Notification
		err := json.Unmarshal(message.Payload(), &notify)
		if err != nil {
			log.Error(err)
			return
		}

		//var alarm model.Alarm 结构体不取 project,device字段
		alarm := make(map[string]any)
		has, err := db.Engine.Table("alarm").Select("alarm.*, product.name as product, device.name as device").
			Join("INNER", "product", "product.id=alarm.product_id").
			Join("INNER", "device", "device.id=alarm.device_id").
			Where("alarm.id=?", notify.AlarmId).
			Get(&alarm)
		if err != nil {
			log.Error(err)
			return
		}
		if !has {
			log.Error("找不到记录")
			return
		}

		var user model.User
		has, err = db.Engine.ID(notify.UserId).Get(&user)
		if err != nil {
			log.Error(err)
			return
		}
		if !has {
			log.Error("找不到用户")
			return
		}

		//只取字符串参数
		m := map[string]string{
			"user": user.Name,
			//"product": alarm.Product,
			//"device":  alarm.Device,
			//"type":    alarm.Type,
			//"title":   alarm.Title,
			//"message": alarm.Message,
		}

		for k, v := range alarm {
			if val, ok := v.(string); ok {
				m[k] = val
			}
		}

		//通知
		for _, t := range notify.Channels {
			switch t {
			case "sms":
				//发送短信
				err = Send([]string{user.Cellphone}, m)
				if err != nil {
					log.Info(err)
				}
			case "voice":
				//打电话
			}
		}

	})
}

package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/iot-master-contrib/aliyun/types"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
)

func Subscribe() {
	mqtt.Client.Subscribe("alarm/+/+", 0, func(client paho.Client, message paho.Message) {
		//topics := strings.Split(message.Topic(), "/")
		//pid := topics[1]
		//id := topics[2]

		//解析数据
		var alarm map[string]interface{} //TODO 定义在payload中
		err := json.Unmarshal(message.Payload(), &alarm)
		if err != nil {
			//log
			return
		}

		m := make(map[string]string)
		for k, v := range alarm {
			if s, ok := v.(string); ok {
				m[k] = s
			}
			//else To string ?
		}

		//查询订阅者
		level := alarm["level"].(uint)
		var subs []*types.Subscriber
		err = db.Engine.Where("level <= ?", level).And("disabled = ?", false).Find(&subs)
		if err != nil && len(subs) > 0 {
			var phones []string
			for _, s := range subs {
				phones = append(phones, s.Cellphone)
			}

			//发送
			err = Send(phones, m)
			if err != nil {
				log.Info(err)
			}
		}

	})
}

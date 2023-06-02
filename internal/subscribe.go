package internal

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
)

func Subscribe() {
	mqtt.Client.Subscribe("alarm/+/+", 0, func(client paho.Client, message paho.Message) {
		//topics := strings.Split(message.Topic(), "/")
		//pid := topics[1]
		//id := topics[2]

		//解析数据
		var alarm map[string]interface{}
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

		err = Send([]string{""}, m)
		if err != nil {
			log.Info(err)
		}
	})
}

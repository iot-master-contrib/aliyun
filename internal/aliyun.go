package internal

import (
	"encoding/json"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sms "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"strings"
)

var client *sms.Client

func Open() (err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(options.Id),
		AccessKeySecret: tea.String(options.Secret),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}
	client, err = sms.NewClient(config)
	return
}

func Send(phone []string, param map[string]string) error {
	p, _ := json.Marshal(param)

	req := &sms.SendSmsRequest{
		PhoneNumbers:  tea.String(strings.Join(phone, ",")),
		SignName:      tea.String(options.Sign),
		TemplateCode:  tea.String(options.Template),
		TemplateParam: tea.String(string(p)),
	}
	runtime := &util.RuntimeOptions{}
	resp, err := client.SendSmsWithOptions(req, runtime)
	if err != nil {
		log.Println(resp)
	}
	//if resp.StatusCode != 200
	if *resp.Body.Code != "OK" {
		return errors.New(*resp.Body.Message)
	}
	return nil
}

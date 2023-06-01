package internal

import (
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sms "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"strings"
)

type Sender struct {
	client *sms.Client
}

func (s *Sender) Connect(id, secret string) error {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: &id,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: &secret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	var err error
	s.client, err = sms.NewClient(config)
	return err
}

func (s *Sender) Send(phone []string, sign, code string, param map[string]string) error {
	p, _ := json.Marshal(param)

	req := &sms.SendSmsRequest{
		PhoneNumbers:  tea.String(strings.Join(phone, ",")),
		SignName:      tea.String(sign),
		TemplateCode:  tea.String(code),
		TemplateParam: tea.String(string(p)),
	}
	runtime := &util.RuntimeOptions{}
	resp, err := s.client.SendSmsWithOptions(req, runtime)
	if err != nil {
		log.Println(resp)
	}
	return err
}

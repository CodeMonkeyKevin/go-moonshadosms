package moonshadosms

import (
  "net/http"
  "net/url"
  "io/ioutil"
)

type SmsClient struct {
  api_key string
}

func (s SmsClient) SendSms(deviceAddress, message string) (body string, err error) {

  resp, err := http.PostForm(("http://" + s.api_key + "@heroku.moonshado.com/sms"), url.Values{"sms[device_address]": {deviceAddress}, "sms[message]": {message}})

  if err != nil {
    return "", err
  }

  defer resp.Body.Close()

  b, _ := ioutil.ReadAll(resp.Body)
  body = string([]byte(b))

  return
}

func NewClient(api_key string) *SmsClient {
  return &SmsClient{api_key}
}

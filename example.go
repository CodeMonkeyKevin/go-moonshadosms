package main

import (
  "fmt"
  "moonshadosms"
  "os"
)

func main() {
  api_key := os.Getenv("MOONSHADO_SMS_URL")

  s := moonshadosms.NewClient(api_key)
  body, err := s.SendSms("1313", "test2")

  if err != nil {
    fmt.Println("Error: ", err)
  } else {
    fmt.Println(body)
  }

}

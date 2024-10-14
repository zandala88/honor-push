# honor-push
honor push go sdk

荣耀消息推送 go sdk

```go
package main

import (
	"fmt"

	honorpush "github.com/zandala88/honor-push"
)

func main() {
	client, err := honorpush.NewClient(1, "your_clientId", "your_clientSecret")
	if err != nil {
		fmt.Println(err)
		return
	}

	clickAction := honorpush.NewClickAction(3)
	androidNotification := honorpush.NewAndroidNotification("睿鹰急救天眼", "附近有呼叫请求").SetClickAction(clickAction)
	android := honorpush.NewAndroidConfig().SetAndroidNotification(androidNotification)
	message := honorpush.NewMessage().SetAndroidConfig(android).SetToken([]string{"your_token"})

	_, err = client.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
}
```
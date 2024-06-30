package mobpush

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	workNo := ""
	appkey, appSecert := "", ""
	client := NewPushClient(appkey, appSecert)
	res, err := client.PushByRids(workNo, "测试推送标题", "测试推送内容", []string{""})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(res))
}

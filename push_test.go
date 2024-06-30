package mobpush

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	appkey, appSecert := "", ""
	client := NewPushClient(appkey, appSecert)

	push := Push{
		Appkey: client.AppKey,
		Source: "webapi",
		PushNotify: PushNotify{
			Plats:          []int{1},
			Type:           1,
			OfflineSeconds: 3600,
			Title:          "测试推送3",
			Content:        "测试推送内容",
			AndroidNotify: &AndroidNotify{
				// Style:            3,
				// Warn:             "123",
				AndroidBadgeType: 1, // 角标固定值不支持
				AndroidBadge:     7,
			},
		},
		PushTarget: PushTarget{
			Target: 4,
			Rids:   []string{""},
		},
	}
	b, err := GetHTTPClient().PostJSON(client, BASE_URL+PUSH_PUSH_URI, push)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}

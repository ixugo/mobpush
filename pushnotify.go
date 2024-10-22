package mobpush

type PushNotify struct {
	TaskCron       int            `json:"taskCron"`
	TaskTime       int64          `json:"taskTime"`
	Plats          []int          `json:"plats"`
	IosProduct     int            `json:"iosProduct"`
	OfflineSeconds int            `json:"offlineSeconds"`
	Content        string         `json:"content"`
	Title          string         `json:"title"`
	Type           int            `json:"type"`
	Url            string         `json:"url"`
	CustomNotify   *CustomNotify  `json:"customNotify"`
	AndroidNotify  *AndroidNotify `json:"androidNotify"`
	IosNotify      *IosNotify     `json:"iosNotify"`
	ExtrasMapList  []PushMap      `json:"extrasMapList"`
	Speed          int            `json:"speed"`
}

type CustomNotify struct {
	CustomType  string `json:"customType"`
	CustomTitle string `json:"customTitle"`
}

type AndroidNotify struct {
	AppName          string   `json:"appName"`
	Title            string   `json:"title"`
	Warn             string   `json:"warn"`
	Style            int      `json:"style"`
	Content          []string `json:"content"`
	Sound            string   `json:"sound"`
	AndroidBadgeType int      `json:"androidBadgeType"` // 1：角标数值取androidBadge值;2：角标数值为androidBadge当前值加1
	AndroidBadge     int      `json:"androidBadge"`     // 角标数值
}

type IosNotify struct {
	Title            string `json:"title"`
	SubTitle         string `json:"subtitle"`
	Sound            string `json:"sound"`
	Badge            string `json:"badge"`
	BadgeType        int    `json:"badgeType"`
	CateGory         string `json:"category"`
	SLIENT           int    `json:"SLIENT"`
	SlientPush       int    `json:"slientPush"`
	ContentAvailable int    `json:"contentAvailable"`
	MutableContent   int    `json:"mutableContent"`
	AttachmentType   int    `json:"attachmentType"`
	Attachment       string `json:"attachment"`
}

type PushMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (push *Push) SetTitle(title string) *Push {
	push.PushNotify.Title = title
	return push
}

func (push *Push) SetContent(content string) *Push {
	push.PushNotify.Content = content
	return push
}

func (push *Push) SetPlats(plats []int) *Push {
	push.PushNotify.Plats = plats
	return push
}

func (push *Push) SetCustomNotify(customNotify CustomNotify) *Push {
	push.PushNotify.CustomNotify = &customNotify
	return push
}

func (push *Push) SetAndroidNotify(androidNotify AndroidNotify) *Push {
	push.PushNotify.AndroidNotify = &androidNotify
	return push
}

func (push *Push) SetIosNotify(iosNotify IosNotify) *Push {
	push.PushNotify.IosNotify = &iosNotify
	return push
}

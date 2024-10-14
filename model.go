package honorpush

type Message struct {
	Data         string         `json:"data,omitempty"`         // 自定义消息负载，通知栏消息支持JSON格式字符串，透传消息支持普通字符串或者JSON格式字符串。样例："your data"，"{'param1':'value1','param2':'value2'}"。 消息体中有data，没有notification和android.notification，消息类型为透传消息。
	Notification *Notification  `json:"notification,omitempty"` // 通知栏消息内容，具体字段请参见Notification的定义。
	Token        []string       `json:"token"`                  // 按照Token向目标用户推消息，样例："token":["token1","token2"]
	Android      *AndroidConfig `json:"android"`                // Android消息推送控制参数，具体字段请参见AndroidConfig的定义
}

type Notification struct {
	Title string `json:"title"` // 通知栏消息的标题。发送通知栏消息时，此处title和android.notification .title两者最少需要设置一个。
	Body  string `json:"body"`  // 通知栏消息的内容。发送通知栏消息时，此处body和android.notification .body两者最少需要设置一个。
	Image string `json:"image"` // 用户自定义的通知栏消息通知小图URL，如果不设置，则不展示通知小图。URL使用的协议必须是HTTPS协议，取值样例：https://example.com/image.png。
}

type AndroidConfig struct {
	TTL                 string               `json:"ttl,omitempty"`            // 消息缓存时间，单位是秒。在用户设备离线时，消息在Push服务器进行缓存，在消息缓存时间内用户设备上线，消息会下发，超过缓存时间后消息会丢弃，默认值为“86400s”（1天），最大值为“1296000s”（15天）。
	BiTag               string               `json:"biTag,omitempty"`          // 批量任务消息标识，消息回执时会返回给应用服务器，应用服务器可以识别biTag对消息的下发情况进行统计分析。
	Data                string               `json:"data,omitempty"`           // 自定义消息负载，此处如果设置了data，则会覆盖外层的data字段。
	AndroidNotification *AndroidNotification `json:"notification,omitempty"`   // Android通知栏消息结构体，具体字段请参见AndroidNotification结构体的定义。
	TargetUserType      int                  `json:"targetUserType,omitempty"` // 0：普通消息（默认值）1：测试消息。每个应用每日可发送该测试消息1000条且不受每日单设备推送数量上限要求。
}

type AndroidNotification struct {
	Title       string             `json:"title"`                // Android通知栏消息标题，如果此处设置了title则会覆盖notification.title字段，且发送通知栏消息时，此处title和notification.title两者最少需要设置一个。
	Body        string             `json:"body"`                 // Android通知栏消息内容，如果此处设置了body则会覆盖notification.body字段，且发送通知栏消息时，此处body和notification.body两者最少需要设置一个。
	ClickAction *ClickAction       `json:"clickAction"`          // 消息点击行为，具体字段请参见ClickAction结构体的定义。 如果是Android通知栏消息时，则该参数必选。
	Image       string             `json:"image,omitempty"`      // 自定义通知小图URL，功能和notification.image字段一样，如果此处设置，则覆盖notification.image中的值。URL使用的协议必须是HTTPS协议，取值样例：https://example.com/image.png。
	Style       int                `json:"style,omitempty"`      // 通知栏样式，取值如下： 0：默认样式 1：大文本样式
	BigTitle    string             `json:"bigTitle,omitempty"`   // Android通知栏消息大文本标题，当style为1时必选，设置bigTitle后通知栏展示时，bigTitle设置后内容要与title一致
	BigBody     string             `json:"bigBody,omitempty"`    // Android通知栏消息大文本内容，当style为1时必选，设置bigBody后通知栏展示时，bigBody设置后内容要与body一致
	Importance  string             `json:"importance,omitempty"` // Android通知消息分类，决定用户设备消息通知行为，取值如下： LOW：资讯营销类消息 NORMAL：服务与通讯类消息
	When        string             `json:"when,omitempty"`       // 设置通知栏消息的到达时间，如果您同时发送多条消息，Android通知栏中的消息根据这个值进行排序，同时将排序后的消息在通知栏上显示。该时间戳为UTC时间戳
	Buttons     *Buttons           `json:"buttons,omitempty"`    // 通知栏消息动作按钮，最多设置3个。具体字段请参见Button结构体的定义。
	Badge       *BadgeNotification `json:"badge,omitempty"`      // Android通知消息角标控制，具体字段请参见BadgeNotification结构体的定义
	NotifyId    int                `json:"notifyId,omitempty"`   // 每条消息在通知显示时的唯一标识。不携带时或者设置-1时，Push NC自动为给每条消息生成一个唯一标识；不同的通知栏消息可以拥有相同的notifyId，实现新的消息覆盖上一条消息功能。
	Tag         string             `json:"tag,omitempty"`        // 消息标签，同一应用下使用同一个消息标签的消息会相互覆盖，只展示最新的一条
	Group       string             `json:"group,omitempty"`      // 消息分组，例如发送10条带有同样group字段的消息，手机上只会展示该组消息中最新的一条和当前该组接收到的消息总数目，不会展示10条消息
}

type ClickAction struct {
	Type   int    `json:"type"`             // 消息点击行为类型，取值如下： 1：打开应用自定义页面 2：点击后打开特定URL 3：点击后打开应用
	Intent string `json:"intent,omitempty"` // 自定义页面中intent的实现，请参见指定intent参数。当type为1时，字段intent和action至少二选一。
	Url    string `json:"url,omitempty"`    // 设置打开特定URL，本字段填写需要打开的URL，URL使用的协议必须是HTTPS协议，
	Action string `json:"action,omitempty"` // 设置通过action打开应用自定义页面时，本字段填写要打开的页面activity对应的action。
}

type Buttons struct {
	Name       string `json:"name"`       // 按钮名称，最大长度40
	ActionType int    `json:"actionType"` // 按钮动作类型： 0：打开应用首页 1：打开应用自定义页面 2：打开指定的网页
	IntentType int    `json:"intentType"` // 打开自定义页面的方式： 0：设置通过intent打开应用自定义页面 1：设置通过action打开应用自定义页面
	Intent     string `json:"intent"`     // 当actionType为1，此字段按照intentType字段设置应用页面的uri或者action，具体设置方式参见打开应用自定义页面。 当actionType为2，此字段设置打开指定网页的URL，URL使用的协议必须是HTTPS协议
	Data       string `json:"data"`       // 最大长度1024。 当字段actionType为0或1时，该字段用于在点击按钮后给应用透传数据，选填
}

type BadgeNotification struct {
	AddNum     int    `json:"addNum"`     // 应用角标累加数字非应用角标实际显示数字
	BadgeClass string `json:"badgeClass"` // 应用入口Activity类全路径。
	SetNum     int    `json:"setNum"`     // 角标设置数字，大于等于0小于100的整数
}

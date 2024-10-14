package honorpush

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) SetData(data string) *Message {
	m.Data = data
	return m
}

func (m *Message) SetToken(token []string) *Message {
	m.Token = token
	return m
}

func (m *Message) SetNotification(config *Notification) *Message {
	m.Notification = config
	return m
}

func (m *Message) SetAndroidConfig(config *AndroidConfig) *Message {
	m.Android = config
	return m
}

func NewNotification(title, content string) *Notification {
	return &Notification{
		Title: title,
		Body:  content,
	}
}

func (n *Notification) SetImage(image string) *Notification {
	n.Image = image
	return n
}

func NewAndroidConfig() *AndroidConfig {
	return &AndroidConfig{}
}

func (a *AndroidConfig) SetTTL(ttl string) *AndroidConfig {
	a.TTL = ttl
	return a
}

func (a *AndroidConfig) SetBiTag(biTag string) *AndroidConfig {
	a.BiTag = biTag
	return a
}

func (a *AndroidConfig) SetData(data string) *AndroidConfig {
	a.Data = data
	return a
}

func (a *AndroidConfig) SetTargetUserType(t int) *AndroidConfig {
	a.TargetUserType = t
	return a
}

func (a *AndroidConfig) SetAndroidNotification(n *AndroidNotification) *AndroidConfig {
	a.AndroidNotification = n
	return a
}

func NewAndroidNotification(title, body string) *AndroidNotification {
	return &AndroidNotification{
		Title: title,
		Body:  body,
	}
}

func (a *AndroidNotification) SetClickAction(c *ClickAction) *AndroidNotification {
	a.ClickAction = c
	return a
}

func (a *AndroidNotification) SetImage(image string) *AndroidNotification {
	a.Image = image
	return a
}

func (a *AndroidNotification) SetStyle(style int) *AndroidNotification {
	a.Style = style
	return a
}

func (a *AndroidNotification) SetBigTitle(title string) *AndroidNotification {
	a.BigTitle = title
	return a
}

func (a *AndroidNotification) SetBigBody(body string) *AndroidNotification {
	a.BigBody = body
	return a
}

func (a *AndroidNotification) SetImportance(importance string) *AndroidNotification {
	a.Importance = importance
	return a
}

func (a *AndroidNotification) SetWhen(when string) *AndroidNotification {
	a.When = when
	return a
}

func (a *AndroidNotification) SetButtons(b *Buttons) *AndroidNotification {
	a.Buttons = b
	return a
}

func (a *AndroidNotification) SetBadge(b *BadgeNotification) *AndroidNotification {
	a.Badge = b
	return a
}

func (a *AndroidNotification) SetNotifyId(id int) *AndroidNotification {
	a.NotifyId = id
	return a
}

func (a *AndroidNotification) SetTag(tag string) *AndroidNotification {
	a.Tag = tag
	return a
}

func (a *AndroidNotification) SetGroup(group string) *AndroidNotification {
	a.Group = group
	return a
}

func NewClickAction(t int) *ClickAction {
	return &ClickAction{
		Type: t,
	}
}

func (c *ClickAction) SetIntent(intent string) *ClickAction {
	c.Intent = intent
	return c
}

func (c *ClickAction) SetUrl(url string) *ClickAction {
	c.Url = url
	return c
}

func (c *ClickAction) SetAction(action string) *ClickAction {
	c.Action = action
	return c
}

func NewButtons(name string, actionType int) *Buttons {
	return &Buttons{
		Name:       name,
		ActionType: actionType,
	}
}

func (b *Buttons) SetIntentType(intentType int) *Buttons {
	b.IntentType = intentType
	return b
}

func (b *Buttons) SetIntent(intent string) *Buttons {
	b.Intent = intent
	return b
}

func (b *Buttons) SetData(data string) *Buttons {
	b.Data = data
	return b
}

func NewBadgeNotification(badgeClass string) *BadgeNotification {
	return &BadgeNotification{
		BadgeClass: badgeClass,
	}
}

func (b *BadgeNotification) SetAddNum(addNum int) *BadgeNotification {
	b.AddNum = addNum
	return b
}

func (b *BadgeNotification) SetSetNum(setNum int) *BadgeNotification {
	b.SetNum = setNum
	return b
}

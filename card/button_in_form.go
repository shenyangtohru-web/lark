package card

var _ Element = (*ButtonInFormBlock)(nil)

// ButtonInFormBlock 表单中按钮元素
type ButtonInFormBlock struct {
	text       *TextBlock
	name       string
	url        string
	multiURL   *URLBlock
	btnType    string
	value      map[string]interface{}
	confirm    *ConfirmBlock
	actionType string
}

type buttonInFormRenderer struct {
	ElementTag
	Name       string                 `json:"name"`
	Text       Renderer               `json:"text"`
	URL        string                 `json:"url,omitempty"`
	MultiURL   Renderer               `json:"multi_url,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Value      map[string]interface{} `json:"value,omitempty"`
	Confirm    Renderer               `json:"confirm,omitempty"`
	ActionType string                 `json:"action_type,omitempty"`
}

// Render 渲染为 Renderer
func (b *ButtonInFormBlock) Render() Renderer {
	ret := buttonInFormRenderer{
		ElementTag: ElementTag{
			Tag: "button",
		},
		Text:       b.text.Render(),
		URL:        b.url,
		Type:       b.btnType,
		Value:      b.value,
		Name:       b.name,
		ActionType: b.actionType,
	}
	if b.multiURL != nil {
		ret.MultiURL = b.multiURL.Render()
	}
	if b.confirm != nil {
		ret.Confirm = b.confirm.Render()
	}
	return ret
}

// ButtonInForm 按钮交互元素
func ButtonInForm(text *TextBlock, name string) *ButtonInFormBlock {
	return (&ButtonInFormBlock{text: text, name: name}).Default()
}

// URL 按钮的跳转链接
func (b *ButtonInFormBlock) URL(u string) *ButtonInFormBlock {
	b.url = u
	return b
}

// MultiURL 按钮的多端差异跳转链接
func (b *ButtonInFormBlock) MultiURL(u *URLBlock) *ButtonInFormBlock {
	b.multiURL = u
	return b
}

// Value 点击后发送给业务方的数据
func (b *ButtonInFormBlock) Value(v map[string]interface{}) *ButtonInFormBlock {
	b.value = v
	return b
}

// Confirm 点击后二次确认的弹框
func (b *ButtonInFormBlock) Confirm(title, text string) *ButtonInFormBlock {
	b.confirm = Confirm(title, text)
	return b
}

// Default 设置按钮样式（次要按钮）
func (b *ButtonInFormBlock) Default() *ButtonInFormBlock {
	b.btnType = "default"
	b.actionType = "link"
	return b
}

// Primary 设置按钮样式（主要按钮）
func (b *ButtonInFormBlock) Primary() *ButtonInFormBlock {
	b.btnType = "primary"
	return b
}

// Danger 设置按钮样式（警示按钮）
func (b *ButtonInFormBlock) Danger() *ButtonInFormBlock {
	b.btnType = "danger"
	return b
}

// Link 仅链接跳转
func (b *ButtonInFormBlock) Link() *ButtonInFormBlock {
	b.actionType = "link"
	return b
}

// Request 仅回传交互
func (b *ButtonInFormBlock) Request() *ButtonInFormBlock {
	b.actionType = "request"
	return b
}

// Multi 链接跳转+回传交互同时生效
func (b *ButtonInFormBlock) Multi() *ButtonInFormBlock {
	b.actionType = "multi"
	return b
}

// FormSubmit 触发表单容器的提交事件，异步提交表单容器中所有用户填写的表单内容
func (b *ButtonInFormBlock) FormSubmit() *ButtonInFormBlock {
	b.actionType = "form_submit"
	return b
}

// FormReset 触发表单容器的取消提交事件，重置所有表单组件的输入值为初始值
func (b *ButtonInFormBlock) FormReset() *ButtonInFormBlock {
	b.actionType = "form_reset"
	return b
}

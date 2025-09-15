package card

var _ Element = (*TextBlock)(nil)
var _ Element = (*MarkdownBlock)(nil)

// TextBlock 文本元素
type TextBlock struct {
	tag     string
	content string
	lines   int
	href    map[string]*URLBlock
}

type textRenderer struct {
	ElementTag
	Content string              `json:"content"`
	Lines   int                 `json:"lines,omitempty"`
	Href    map[string]Renderer `json:"href,omitempty"`
}

// Text 文本模块
func Text(s string) *TextBlock {
	return &TextBlock{content: s, tag: "plain_text"}
}

// Render 渲染为 Renderer
func (t *TextBlock) Render() Renderer {
	ret := textRenderer{
		ElementTag: ElementTag{
			Tag: t.tag,
		},
		Content: t.content,
		Lines:   t.lines,
	}
	if len(t.href) > 0 {
		ret.Href = make(map[string]Renderer, len(t.href))
		for k, v := range t.href {
			ret.Href[k] = v.Render()
		}
	}
	return ret
}

// LarkMd 嵌入使用的 Markdown 展示模式
func (t *TextBlock) LarkMd() *TextBlock {
	t.tag = "lark_md"
	return t
}

// Lines 内容展示的行数
func (t *TextBlock) Lines(l int) *TextBlock {
	t.lines = l
	return t
}

// Href 设置文本中 []($urlVal) 格式的链接值，仅在 LarkMd 和 Markdown 模块中可用
func (t *TextBlock) Href(name string, url *URLBlock) *TextBlock {
	if t.href == nil {
		t.href = make(map[string]*URLBlock)
	}
	t.href[name] = url
	return t
}

// MarkdownBlock Markdown文本元素
type MarkdownBlock struct {
	content   string
	textAlign string
	href      map[string]*URLBlock
	textSize  string
	icon      *MarkdownIcon
}

type markdownRenderer struct {
	ElementTag
	Content   string              `json:"content"`
	TextAlign string              `json:"text_align,omitempty"`
	TextSize  string              `json:"text_size, omitempty"`
	Href      map[string]Renderer `json:"href,omitempty"`
	Icon      *MarkdownIcon       `json:"icon,omitempty"`
}

type MarkdownIcon struct {
	Tag   string `json:"tag,omitempty"`
	Token string `json:"token,omitempty"`
	Color string `json:"color,omitempty"`
}

// Markdown 单独使用的 Markdown 文本模块
func Markdown(s string) *MarkdownBlock {
	return &MarkdownBlock{
		content: s,
	}
}

// AlignCenter .
func (m *MarkdownBlock) AlignCenter() *MarkdownBlock {
	m.textAlign = "center"
	return m
}

// AlignLeft .
func (m *MarkdownBlock) AlignLeft() *MarkdownBlock {
	m.textAlign = "left"
	return m
}

// AlignRight .
func (m *MarkdownBlock) AlignRight() *MarkdownBlock {
	m.textAlign = "right"
	return m
}

func (m *MarkdownBlock) SetTextSizeNormal() *MarkdownBlock {
	m.textSize = "normal"
	return m
}

func (m *MarkdownBlock) SetTextSizeHeading() *MarkdownBlock {
	m.textSize = "heading"
	return m
}

func (m *MarkdownBlock) SetTextSizeNotation() *MarkdownBlock {
	m.textSize = "notation"
	return m
}

// Href 设置文本中 []($urlVal) 格式的链接值，仅在 LarkMd 和 Markdown 模块中可用
func (m *MarkdownBlock) Href(name string, url *URLBlock) *MarkdownBlock {
	if m.href == nil {
		m.href = make(map[string]*URLBlock)
	}
	m.href[name] = url
	return m
}

// Icon 设置 Markdown 图标
func (m *MarkdownBlock) Icon(tkn, color string) *MarkdownBlock {
	m.icon = &MarkdownIcon{
		Tag:   "standard_icon",
		Token: tkn,
		Color: color,
	}
	return m
}

// Render 渲染为 Renderer
func (m *MarkdownBlock) Render() Renderer {
	ret := markdownRenderer{
		ElementTag: ElementTag{
			Tag: "markdown",
		},
		Content:   m.content,
		TextAlign: m.textAlign,
		TextSize:  m.textSize,
		Icon:      m.icon,
	}
	if len(m.href) > 0 {
		ret.Href = make(map[string]Renderer, len(m.href))
		for k, v := range m.href {
			ret.Href[k] = v.Render()
		}
	}
	return ret
}

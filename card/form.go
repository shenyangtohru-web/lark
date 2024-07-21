package card

var _ Element = (*FormBlock)(nil)

// FormBlock 表单容器
type FormBlock struct {
	name     string
	elements []Element
}

type formRenderer struct {
	ElementTag
	Name     string     `json:"name,omitempty"`
	Elements []Renderer `json:"elements,omitempty"`
}

// Render 渲染为 Renderer
func (f *FormBlock) Render() Renderer {
	ret := formRenderer{
		ElementTag: ElementTag{
			Tag: "form",
		},
		Name:     f.name,
		Elements: renderElements(f.elements),
	}

	return ret
}

// Form 新建表单
func Form(name string) *FormBlock {
	return &FormBlock{
		name: name,
	}
}

// AddElements 添加元素
func (f *FormBlock) AddElements(elements ...Element) *FormBlock {
	f.elements = append(f.elements, elements...)
	return f
}

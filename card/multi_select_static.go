package card

var _ Element = (*SelectBlock)(nil)

type SelectOption struct {
	text  *TextBlock
	icon  *IconBlock
	value string
}

type SelectBlock struct {
	typeA1         string
	name           string
	placeholder    string
	width          string
	required       *bool
	disabled       *bool
	selectedValues []string
	options        []SelectOption
}

type selectRenderer struct {
	ElementTag
	Type           string   `json:"type,omitempty"`
	Name           string   `json:"name"`
	Placeholder    Renderer `json:"placeholder,omitempty"`
	Width          string   `json:"width,omitempty"`
	Required       *bool    `json:"required,omitempty"`
	Disabled       *bool    `json:"disabled,omitempty"`
	SelectedValues []string `json:"selected_values,omitempty"`
	Options        []struct {
		Text  Renderer `json:"text,omitempty"`
		Icon  Renderer `json:"icon,omitempty"`
		Value string   `json:"value"`
	} `json:"options"`
}

// Render 渲染为 Renderer
func (s *SelectBlock) Render() Renderer {
	ret := selectRenderer{
		ElementTag: ElementTag{
			Tag: "multi_select_static",
		},
		Name: s.name,
	}
	if s.typeA1 != "" {
		ret.Type = s.typeA1
	}
	if s.placeholder != "" {
		ret.Placeholder = Text(s.placeholder).Render()
	}
	if s.width != "" {
		ret.Width = s.width
	}
	if s.required != nil {
		ret.Required = s.required
	}
	if s.disabled != nil {
		ret.Disabled = s.disabled
	}
	if s.selectedValues != nil {
		ret.SelectedValues = s.selectedValues
	}
	if s.options != nil {
		for _, op := range s.options {
			op1 := struct {
				Text  Renderer `json:"text,omitempty"`
				Icon  Renderer `json:"icon,omitempty"`
				Value string   `json:"value"`
			}{
				Text:  op.text.Render(),
				Icon:  op.icon.Render(),
				Value: op.value,
			}
			ret.Options = append(ret.Options, op1)
		}
	}

	return ret
}

func (s *SelectBlock) SetName(name string) {
	s.name = name
}

func (s *SelectBlock) SetTypeText() {
	s.typeA1 = "text"
}

func (s *SelectBlock) SetPlaceholder(text string) {
	s.placeholder = text
}
func (s *SelectBlock) SetWidthFill() {
	s.width = "fill"
}

func (s *SelectBlock) SetSelectedValues(vals []string) {
	s.selectedValues = vals
}

func (s *SelectBlock) SetDisabled() {
	disabled := true
	s.disabled = &disabled
}

func (s *SelectBlock) SetOptions(ops []SelectOption) {
	s.options = ops
}

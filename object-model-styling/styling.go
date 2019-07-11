package objectmodelstyling

// https://developer.mozilla.org/en-US/docs/Web/CSS/Reference

// ElementStyle - FIXME
type ElementStyle struct {
	Align      *ElementAlign      `json:"align,omitempty"`
	Animation  *ElementAnimation  `json:"animation,omitempty"`
	Background *ElementBackground `json:"background,omitempty"`
	Border     *ElementBorder     `json:"border,omitempty"`
	Font       *ElementFont       `json:"font,omitempty"`
	Inset      *ElementInset      `json:"inset,omitempty"`
	Justify    *ElementJustify    `json:"justify,omitempty"`
	Line       *ElementLine       `json:"line,omitempty"`
	Margin     *ElementMargin     `json:"margin,omitempty"`
	Mask       *ElementMask       `json:"mask,omitempty"`
	Max        *ElementMax        `json:"max,omitempty"`
	Min        *ElementMin        `json:"min,omitempty"`
	Outline    *ElementOutline    `json:"outline,omitempty"`
	Overflow   *ElementOverflow   `json:"overflow,omitempty"`
	Padding    *ElementPadding    `json:"padding,omitempty"`
	Scroll     *ElementScroll     `json:"scroll,omitempty"`
	Spahe      *ElementSpahe      `json:"spahe,omitempty"`
	Text       *ElementText       `json:"text,omitempty"`
	Transform  *ElementTransform  `json:"transform,omitempty"`
	Transition *ElementTransition `json:"transition,omitempty"`
	Word       *ElementWord       `json:"word,omitempty"`
	Element    *Element           `json:"element,omitempty"`
}

// GlobalStyle - FIXME
type GlobalStyle struct {
	Scrollbar *GlobalScrollbar
}

// LayoutGrid

// LayoutList

// LayoutPage

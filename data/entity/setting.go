package entity

type Setting struct {
	SystemContent string
	AiModel       string
	Temperature   float32
	MaxTokens     uint
}

func (Setting) TableName() string {
	return "setting"
}

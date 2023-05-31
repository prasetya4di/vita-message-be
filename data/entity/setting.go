package entity

type Setting struct {
	SystemContent string  `json:"system_content" gorm:"type:text"`
	AiModel       string  `json:"ai_model" gorm:"size:50"`
	Temperature   float32 `json:"temperature" gorm:"type:double"`
	MaxTokens     uint    `json:"max_tokens" gorm:"type:uint"`
}

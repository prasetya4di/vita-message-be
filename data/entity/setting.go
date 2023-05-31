package entity

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	SystemContent string
	AiModel       string
	Temperature   float32
	MaxTokens     uint
}

func (Setting) TableName() string {
	return "setting"
}

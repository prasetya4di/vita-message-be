package entity

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	SystemContext string
	AiModel       string
	Temperature   float32
	MaxTokens     uint
}

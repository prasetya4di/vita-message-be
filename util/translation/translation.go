package translation

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

func LoadTranslation() *i18n.Localizer {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("util/translation/active.en.toml")
	bundle.MustLoadMessageFile("util/translation/active.id.toml")

	localizer := i18n.NewLocalizer(bundle, language.English.String(), language.Indonesian.String())

	return localizer
}

func UnknownImageMessage(localizer *i18n.Localizer) string {
	message := &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "UnknownImage",
		},
	}

	result, _ := localizer.Localize(message)

	return result
}

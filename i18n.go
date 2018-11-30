package glocalize

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Localization is a wrapper around the Localizer struct.
// This is done to add helper functions to ease translation
type Localization struct {
	i18n.Localizer
}

// A is an alias to map[string]interface{} to allow for passing
// translation variables
type A map[string]interface{}

// GetLocalization returns the localization object for a given language
func GetLocalization(language string, bundle *i18n.Bundle) (*Localization, error) {
	loc := i18n.NewLocalizer(bundle, "en-AU")

	return &Localization{
		Localizer: *loc,
	}, nil
}

// MustGetLocalization is the same as GetLocalization, but it will panic on error
func MustGetLocalization(language string, bundle *i18n.Bundle) *Localization {
	loc, err := GetLocalization(language, bundle)
	if err != nil {
		panic(err)
	}

	return loc
}

// T will lookup the localization key in the language file
func (l *Localization) T(key string) string {
	str, err := l.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: l.applyDefaultTranslationArgs(nil),
	})
	if err != nil {
		return fmt.Sprintf("⚠️ Failed to localize. %s %s", key, err)
	}

	return str
}

// TP will lookup the localization key in the language file
// and use the plural version according to the count param
func (l *Localization) TP(key string, count int) string {
	str, err := l.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		PluralCount:  count,
		TemplateData: l.applyDefaultTranslationArgs(nil),
	})
	if err != nil {
		return fmt.Sprintf("⚠️ Failed to localize. %s %s", key, err)
	}

	return str
}

// TA will lookup the localization key in the language file
// while also substituting arguments
func (l *Localization) TA(key string, args A) string {
	args = l.applyDefaultTranslationArgs(args)

	str, err := l.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: args,
	})
	if err != nil {
		return fmt.Sprintf("⚠️ Failed to localize. %s %s", key, err)
	}

	return str
}

// TAP will lookup the localization key in the language file
// while also substituting arguments
// while also providing a pluralization option
func (l *Localization) TAP(key string, args A, count int) string {
	args = l.applyDefaultTranslationArgs(args)

	str, err := l.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: args,
		PluralCount:  count,
	})
	if err != nil {
		return fmt.Sprintf("⚠️ Failed to localize. %s %s", key, err)
	}

	return str
}

func (l *Localization) applyDefaultTranslationArgs(args A) A {
	if args == nil {
		args = make(map[string]interface{})
	}

	args["HolidayTerm"] = l.TPlain("holidayTerm")
	args["HolidayTermPlural"] = l.TPlain("holidayTerm_plural")
	args["OfferTerm"] = l.TPlain("offerTerm")
	args["OfferTermPlural"] = l.TPlain("offerTerm_plural")
	args["Brand"] = l.TPlain("brand")
	args["HolidayTermQR"] = l.TPlain("holidayTerm_qr")
	args["HolidayTermPluralQR"] = l.TPlain("holidayTerm_plural_qr")
	args["OfferTermPluralQR"] = l.TPlain("offerTerm_plural_qr")

	return args
}

// TPlain will translate the key without replacing standard vars
func (l *Localization) TPlain(key string) string {
	return l.MustLocalize(&i18n.LocalizeConfig{
		MessageID: key,
	})
}

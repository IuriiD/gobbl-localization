package glocalize

import (
	"github.com/calebhiebert/gobbl"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Middleware will set a localization object on the context
// this can later be loaded to do translations.
// This middleware will pull the language from the "lang" flag
// so make sure it's set
func Middleware(bundle *i18n.Bundle) gbl.MiddlewareFunction {
	return func(c *gbl.Context) {
		var loc *Localization

		if c.HasFlag("lang") {
			loc = MustGetLocalization(c.GetStringFlag("lang"), bundle)
		} else {
			loc = MustGetLocalization("en-US", bundle)
		}

		c.Flag("__localizer", loc)
		c.Next()
	}
}

// GetCurrentLocalization will retrieve the current localization from the gobbl context
func GetCurrentLocalization(c *gbl.Context) *Localization {
	if c.HasFlag("__localizer") {
		return c.GetFlag("__localizer").(*Localization)
	}

	panic("Localizer not present on context")
}

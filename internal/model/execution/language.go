package model_execution

import errs "github.com/markgredasov/viewcode/internal/errors"

type Language string

const (
	LanguageGo     = Language("go")
	LanguagePython = Language("py")
	LanguageRust   = Language("rs")
)

func NewLanguage(lang string) (Language, error) {
	switch lang {
	case string(LanguageGo):
		return LanguageGo, nil
	case string(LanguagePython):
		return LanguagePython, nil
	case string(LanguageRust):
		return LanguageRust, nil
	}

	return "", errs.ErrInvalidLanguage
}

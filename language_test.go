package nook_test

import (
	"testing"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"golang.org/x/text/language"
)

func TestSupportedLanguageTags(t *testing.T) {
	tags := nook.SupportedLanguageTags()
	if len(tags) != 13 {
		t.Fatalf("len(nook.SupportedLanguageTags()) = %d", len(tags))
	}
	if tags[0] != language.AmericanEnglish {
		t.Fatalf("nook.SupportedLanguageTags()[0] = %s", tags[0])
	}
	if tags[len(tags)-1] != language.TraditionalChinese {
		t.Fatalf("nook.SupportedLanguageTags()[last] = %s", tags[len(tags)-1])
	}

	tags[0] = language.German

	fresh := nook.SupportedLanguageTags()
	if fresh[0] != language.AmericanEnglish {
		t.Fatalf("nook.SupportedLanguageTags()[0] after mutation = %s", fresh[0])
	}
}

func TestLanguagesComplete(t *testing.T) {
	if !animal.Alligator.Name.Complete() {
		t.Fatal("animal.Alligator.Name.Complete() = false")
	}

	values := nook.Languages{
		language.AmericanEnglish: {
			Language: language.AmericanEnglish,
			Value:    "hello",
		},
		language.CanadianFrench: {
			Language: language.CanadianFrench,
		},
	}
	if values.Complete() {
		t.Fatal("Languages.Complete() = true")
	}
}

func TestLanguagesMissingSupportedTags(t *testing.T) {
	values := nook.Languages{
		language.AmericanEnglish: {
			Language: language.AmericanEnglish,
			Value:    "hello",
		},
		language.CanadianFrench: {
			Language: language.CanadianFrench,
		},
		language.Japanese: {
			Language: language.Japanese,
			Value:    "こんにちは",
		},
	}

	missing := values.MissingSupportedTags()
	if len(missing) != 11 {
		t.Fatalf("len(Languages.MissingSupportedTags()) = %d", len(missing))
	}
	if missing[0] != language.CanadianFrench {
		t.Fatalf("Languages.MissingSupportedTags()[0] = %s", missing[0])
	}
	if missing[len(missing)-1] != language.TraditionalChinese {
		t.Fatalf("Languages.MissingSupportedTags()[last] = %s", missing[len(missing)-1])
	}
}

func TestLanguagesMissingTags(t *testing.T) {
	values := nook.Languages{
		language.AmericanEnglish: {
			Language: language.AmericanEnglish,
			Value:    "hello",
		},
		language.Japanese: {
			Language: language.Japanese,
			Value:    "こんにちは",
		},
		language.Spanish: {
			Language: language.Spanish,
		},
	}

	missing := values.MissingTags(
		language.Spanish,
		language.German,
		language.German,
		language.AmericanEnglish,
	)
	if len(missing) != 2 {
		t.Fatalf("len(Languages.MissingTags(...)) = %d", len(missing))
	}
	if missing[0] != language.Spanish {
		t.Fatalf("Languages.MissingTags(...)[0] = %s", missing[0])
	}
	if missing[1] != language.German {
		t.Fatalf("Languages.MissingTags(...)[1] = %s", missing[1])
	}
}

func TestLanguagesTags(t *testing.T) {
	australianEnglish := language.MustParse("en-AU")

	values := nook.Languages{
		language.AmericanEnglish: {
			Language: language.AmericanEnglish,
			Value:    "hello",
		},
		language.Japanese: {
			Language: language.Japanese,
			Value:    "こんにちは",
		},
		language.Spanish: {
			Language: language.Spanish,
			Value:    "hola",
		},
		australianEnglish: {
			Language: australianEnglish,
			Value:    "g'day",
		},
		language.German: {
			Language: language.German,
		},
	}

	tags := values.Tags()
	if len(tags) != 4 {
		t.Fatalf("len(Languages.Tags()) = %d", len(tags))
	}
	if tags[0] != language.AmericanEnglish {
		t.Fatalf("Languages.Tags()[0] = %s", tags[0])
	}
	if tags[1] != language.Japanese {
		t.Fatalf("Languages.Tags()[1] = %s", tags[1])
	}
	if tags[2] != language.Spanish {
		t.Fatalf("Languages.Tags()[2] = %s", tags[2])
	}
	if tags[3] != australianEnglish {
		t.Fatalf("Languages.Tags()[3] = %s", tags[3])
	}
}

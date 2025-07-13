package cow

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	// belleBirthday represents belle birthday.
	belleBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// belleCode represents belle code.
	belleCode = nook.Code{
		Value: ""}
)

var (
	// belleAmericanEnglishName represents belle american english name.
	belleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Belle"}

	// belleCanadianFrenchName represents belle canadian french name.
	belleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// belleDutchName represents belle dutch name.
	belleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// belleFrenchName represents belle french name.
	belleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Belle"}

	// belleGermanName represents belle german name.
	belleGermanName = nook.Name{
		Language: language.German,
		Value:    "Heidi"}

	// belleItalianName represents belle italian name.
	belleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Isabella"}

	// belleJapaneseName represents belle japanese name.
	belleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミルク"}

	// belleKoreanName represents belle korean name.
	belleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// belleLatinAmericanSpanishName represents belle latin american spanish name.
	belleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// belleRussianName represents belle russian name.
	belleRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// belleSimplifiedChineseName represents belle simplified chinese name.
	belleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "洁洁"}

	// belleSpanishName represents belle spanish name.
	belleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bella"}

	// belleTraditionalChineseName represents belle traditional chinese name.
	belleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// belleName represents belle name.
	belleName = nook.Languages{
		language.AmericanEnglish:      belleAmericanEnglishName,
		language.CanadianFrench:       belleCanadianFrenchName,
		language.Dutch:                belleDutchName,
		language.French:               belleFrenchName,
		language.German:               belleGermanName,
		language.Italian:              belleItalianName,
		language.Japanese:             belleJapaneseName,
		language.Korean:               belleKoreanName,
		language.LatinAmericanSpanish: belleLatinAmericanSpanishName,
		language.Russian:              belleRussianName,
		language.SimplifiedChinese:    belleSimplifiedChineseName,
		language.Spanish:              belleSpanishName,
		language.TraditionalChinese:   belleTraditionalChineseName}
)

var (
	// belleCharacter represents belle character.
	belleCharacter = nook.Character{
		Animal:   animal.Cow,
		Birthday: belleBirthday,
		Code:     belleCode,
		Key:      character.Belle,
		Gender:   gender.Female,
		Name:     belleName,
		Special:  false}
)

var (
	// belleAmericanEnglishPhrase represents belle american english phrase.
	belleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cuddles"}

	// belleCanadianFrenchPhrase represents belle canadian french phrase.
	belleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "meuh-euh"}

	// belleDutchPhrase represents belle dutch phrase.
	belleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// belleFrenchPhrase represents belle french phrase.
	belleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meuh-euh"}

	// belleGermanPhrase represents belle german phrase.
	belleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knutsch"}

	// belleItalianPhrase represents belle italian phrase.
	belleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "amuuur"}

	// belleJapanesePhrase represents belle japanese phrase.
	belleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "にゅう"}

	// belleKoreanPhrase represents belle korean phrase.
	belleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// belleLatinAmericanSpanishPhrase represents belle latin american spanish phrase.
	belleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// belleRussianPhrase represents belle russian phrase.
	belleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// belleSimplifiedChinesePhrase represents belle simplified chinese phrase.
	belleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Unknown"}

	// belleSpanishPhrase represents belle spanish phrase.
	belleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cencerrito"}

	// belleTraditionalChinesePhrase represents belle traditional chinese phrase.
	belleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// bellePhrase represents belle phrase.
	bellePhrase = nook.Languages{
		language.AmericanEnglish:      belleAmericanEnglishPhrase,
		language.CanadianFrench:       belleCanadianFrenchPhrase,
		language.Dutch:                belleDutchPhrase,
		language.French:               belleFrenchPhrase,
		language.German:               belleGermanPhrase,
		language.Italian:              belleItalianPhrase,
		language.Japanese:             belleJapanesePhrase,
		language.Korean:               belleKoreanPhrase,
		language.LatinAmericanSpanish: belleLatinAmericanSpanishPhrase,
		language.Russian:              belleRussianPhrase,
		language.SimplifiedChinese:    belleSimplifiedChinesePhrase,
		language.Spanish:              belleSpanishPhrase,
		language.TraditionalChinese:   belleTraditionalChinesePhrase}
)

var (
	// Belle represents belle.
	Belle = nook.Villager{
		Character:   belleCharacter,
		Personality: personality.Peppy,
		Phrase:      bellePhrase}
)

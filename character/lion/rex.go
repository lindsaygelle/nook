package lion

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
	// rexBirthday represents rex birthday.
	rexBirthday = nook.Birthday{
		Day:   24,
		Month: time.July}
)

var (
	// rexCode represents rex code.
	rexCode = nook.Code{
		Value: "lon02"}
)

var (
	// rexAmericanEnglishName represents rex american english name.
	rexAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rex"}

	// rexCanadianFrenchName represents rex canadian french name.
	rexCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Léo"}

	// rexDutchName represents rex dutch name.
	rexDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rex"}

	// rexFrenchName represents rex french name.
	rexFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léo"}

	// rexGermanName represents rex german name.
	rexGermanName = nook.Name{
		Language: language.German,
		Value:    "Rex"}

	// rexItalianName represents rex italian name.
	rexItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rex"}

	// rexJapaneseName represents rex japanese name.
	rexJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サンデー"}

	// rexKoreanName represents rex korean name.
	rexKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "렉스"}

	// rexLatinAmericanSpanishName represents rex latin american spanish name.
	rexLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Leoncio"}

	// rexRussianName represents rex russian name.
	rexRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рекс"}

	// rexSimplifiedChineseName represents rex simplified chinese name.
	rexSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晴天"}

	// rexSpanishName represents rex spanish name.
	rexSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Leoncio"}

	// rexTraditionalChineseName represents rex traditional chinese name.
	rexTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "晴天"}
)

var (
	// rexName represents rex name.
	rexName = nook.Languages{
		language.AmericanEnglish:      rexAmericanEnglishName,
		language.CanadianFrench:       rexCanadianFrenchName,
		language.Dutch:                rexDutchName,
		language.French:               rexFrenchName,
		language.German:               rexGermanName,
		language.Italian:              rexItalianName,
		language.Japanese:             rexJapaneseName,
		language.Korean:               rexKoreanName,
		language.LatinAmericanSpanish: rexLatinAmericanSpanishName,
		language.Russian:              rexRussianName,
		language.SimplifiedChinese:    rexSimplifiedChineseName,
		language.Spanish:              rexSpanishName,
		language.TraditionalChinese:   rexTraditionalChineseName}
)

var (
	// rexCharacter represents rex character.
	rexCharacter = nook.Character{
		Animal:   animal.Lion,
		Birthday: rexBirthday,
		Code:     rexCode,
		Key:      character.Rex,
		Gender:   gender.Male,
		Name:     rexName,
		Special:  false}
)

var (
	// rexAmericanEnglishPhrase represents rex american english phrase.
	rexAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cool cat"}

	// rexCanadianFrenchPhrase represents rex canadian french phrase.
	rexCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lionceau"}

	// rexDutchPhrase represents rex dutch phrase.
	rexDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "coole kat"}

	// rexFrenchPhrase represents rex french phrase.
	rexFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lionceau"}

	// rexGermanPhrase represents rex german phrase.
	rexGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "coolcat"}

	// rexItalianPhrase represents rex italian phrase.
	rexItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrringo"}

	// rexJapanesePhrase represents rex japanese phrase.
	rexJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だオン"}

	// rexKoreanPhrase represents rex korean phrase.
	rexKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "티라이온"}

	// rexLatinAmericanSpanishPhrase represents rex latin american spanish phrase.
	rexLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrreh"}

	// rexRussianPhrase represents rex russian phrase.
	rexRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "спокойно"}

	// rexSimplifiedChinesePhrase represents rex simplified chinese phrase.
	rexSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "狮狮"}

	// rexSpanishPhrase represents rex spanish phrase.
	rexSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "melenas"}

	// rexTraditionalChinesePhrase represents rex traditional chinese phrase.
	rexTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "獅獅"}
)

var (
	// rexPhrase represents rex phrase.
	rexPhrase = nook.Languages{
		language.AmericanEnglish:      rexAmericanEnglishPhrase,
		language.CanadianFrench:       rexCanadianFrenchPhrase,
		language.Dutch:                rexDutchPhrase,
		language.French:               rexFrenchPhrase,
		language.German:               rexGermanPhrase,
		language.Italian:              rexItalianPhrase,
		language.Japanese:             rexJapanesePhrase,
		language.Korean:               rexKoreanPhrase,
		language.LatinAmericanSpanish: rexLatinAmericanSpanishPhrase,
		language.Russian:              rexRussianPhrase,
		language.SimplifiedChinese:    rexSimplifiedChinesePhrase,
		language.Spanish:              rexSpanishPhrase,
		language.TraditionalChinese:   rexTraditionalChinesePhrase}
)

var (
	// Rex represents rex.
	Rex = nook.Villager{
		Character:   rexCharacter,
		Personality: personality.Lazy,
		Phrase:      rexPhrase}
)

package kangaroo

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
	// marcieBirthday represents marcie birthday.
	marcieBirthday = nook.Birthday{
		Day:   31,
		Month: time.May}
)

var (
	// marcieCode represents marcie code.
	marcieCode = nook.Code{
		Value: "kgr10"}
)

var (
	// marcieAmericanEnglishName represents marcie american english name.
	marcieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marcie"}

	// marcieCanadianFrenchName represents marcie canadian french name.
	marcieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Selma"}

	// marcieDutchName represents marcie dutch name.
	marcieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Marcie"}

	// marcieFrenchName represents marcie french name.
	marcieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Selma"}

	// marcieGermanName represents marcie german name.
	marcieGermanName = nook.Name{
		Language: language.German,
		Value:    "Marlies"}

	// marcieItalianName represents marcie italian name.
	marcieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Adelaide"}

	// marcieJapaneseName represents marcie japanese name.
	marcieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マリア"}

	// marcieKoreanName represents marcie korean name.
	marcieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마리아"}

	// marcieLatinAmericanSpanishName represents marcie latin american spanish name.
	marcieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brisa"}

	// marcieRussianName represents marcie russian name.
	marcieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марси"}

	// marcieSimplifiedChineseName represents marcie simplified chinese name.
	marcieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛莉亚"}

	// marcieSpanishName represents marcie spanish name.
	marcieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brisa"}

	// marcieTraditionalChineseName represents marcie traditional chinese name.
	marcieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪莉亞"}
)

var (
	// marcieName represents marcie name.
	marcieName = nook.Languages{
		language.AmericanEnglish:      marcieAmericanEnglishName,
		language.CanadianFrench:       marcieCanadianFrenchName,
		language.Dutch:                marcieDutchName,
		language.French:               marcieFrenchName,
		language.German:               marcieGermanName,
		language.Italian:              marcieItalianName,
		language.Japanese:             marcieJapaneseName,
		language.Korean:               marcieKoreanName,
		language.LatinAmericanSpanish: marcieLatinAmericanSpanishName,
		language.Russian:              marcieRussianName,
		language.SimplifiedChinese:    marcieSimplifiedChineseName,
		language.Spanish:              marcieSpanishName,
		language.TraditionalChinese:   marcieTraditionalChineseName}
)

var (
	// marcieCharacter represents marcie character.
	marcieCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: marcieBirthday,
		Code:     marcieCode,
		Key:      character.Marcie,
		Gender:   gender.Female,
		Name:     marcieName,
		Special:  false}
)

var (
	// marcieAmericanEnglishPhrase represents marcie american english phrase.
	marcieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pouches"}

	// marcieCanadianFrenchPhrase represents marcie canadian french phrase.
	marcieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "wallaby"}

	// marcieDutchPhrase represents marcie dutch phrase.
	marcieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "buidels"}

	// marcieFrenchPhrase represents marcie french phrase.
	marcieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "wallaby"}

	// marcieGermanPhrase represents marcie german phrase.
	marcieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hüpf"}

	// marcieItalianPhrase represents marcie italian phrase.
	marcieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "me oui"}

	// marcieJapanesePhrase represents marcie japanese phrase.
	marcieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "いいのよ"}

	// marcieKoreanPhrase represents marcie korean phrase.
	marcieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "조아요"}

	// marcieLatinAmericanSpanishPhrase represents marcie latin american spanish phrase.
	marcieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nainoná"}

	// marcieRussianPhrase represents marcie russian phrase.
	marcieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кармашек"}

	// marcieSimplifiedChinesePhrase represents marcie simplified chinese phrase.
	marcieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "好唷"}

	// marcieSpanishPhrase represents marcie spanish phrase.
	marcieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "nainoná"}

	// marcieTraditionalChinesePhrase represents marcie traditional chinese phrase.
	marcieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "好唷"}
)

var (
	// marciePhrase represents marcie phrase.
	marciePhrase = nook.Languages{
		language.AmericanEnglish:      marcieAmericanEnglishPhrase,
		language.CanadianFrench:       marcieCanadianFrenchPhrase,
		language.Dutch:                marcieDutchPhrase,
		language.French:               marcieFrenchPhrase,
		language.German:               marcieGermanPhrase,
		language.Italian:              marcieItalianPhrase,
		language.Japanese:             marcieJapanesePhrase,
		language.Korean:               marcieKoreanPhrase,
		language.LatinAmericanSpanish: marcieLatinAmericanSpanishPhrase,
		language.Russian:              marcieRussianPhrase,
		language.SimplifiedChinese:    marcieSimplifiedChinesePhrase,
		language.Spanish:              marcieSpanishPhrase,
		language.TraditionalChinese:   marcieTraditionalChinesePhrase}
)

var (
	// Marcie represents marcie.
	Marcie = nook.Villager{
		Character:   marcieCharacter,
		Personality: personality.Normal,
		Phrase:      marciePhrase}
)

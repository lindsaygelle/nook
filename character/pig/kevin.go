package pig

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
	// kevinBirthday represents kevin birthday.
	kevinBirthday = nook.Birthday{
		Day:   26,
		Month: time.April}
)

var (
	// kevinCode represents kevin code.
	kevinCode = nook.Code{
		Value: "pig15"}
)

var (
	// kevinAmericanEnglishName represents kevin american english name.
	kevinAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kevin"}

	// kevinCanadianFrenchName represents kevin canadian french name.
	kevinCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jean-Bon"}

	// kevinDutchName represents kevin dutch name.
	kevinDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kevin"}

	// kevinFrenchName represents kevin french name.
	kevinFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jean-Bon"}

	// kevinGermanName represents kevin german name.
	kevinGermanName = nook.Name{
		Language: language.German,
		Value:    "Kevin"}

	// kevinItalianName represents kevin italian name.
	kevinItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kotekiño"}

	// kevinJapaneseName represents kevin japanese name.
	kevinJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "イノッチ"}

	// kevinKoreanName represents kevin korean name.
	kevinKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "멧지"}

	// kevinLatinAmericanSpanishName represents kevin latin american spanish name.
	kevinLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Porcinio"}

	// kevinRussianName represents kevin russian name.
	kevinRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кевин"}

	// kevinSimplifiedChineseName represents kevin simplified chinese name.
	kevinSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "伊比利"}

	// kevinSpanishName represents kevin spanish name.
	kevinSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Porcinio"}

	// kevinTraditionalChineseName represents kevin traditional chinese name.
	kevinTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "伊比利"}
)

var (
	// kevinName represents kevin name.
	kevinName = nook.Languages{
		language.AmericanEnglish:      kevinAmericanEnglishName,
		language.CanadianFrench:       kevinCanadianFrenchName,
		language.Dutch:                kevinDutchName,
		language.French:               kevinFrenchName,
		language.German:               kevinGermanName,
		language.Italian:              kevinItalianName,
		language.Japanese:             kevinJapaneseName,
		language.Korean:               kevinKoreanName,
		language.LatinAmericanSpanish: kevinLatinAmericanSpanishName,
		language.Russian:              kevinRussianName,
		language.SimplifiedChinese:    kevinSimplifiedChineseName,
		language.Spanish:              kevinSpanishName,
		language.TraditionalChinese:   kevinTraditionalChineseName}
)

var (
	// kevinCharacter represents kevin character.
	kevinCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: kevinBirthday,
		Code:     kevinCode,
		Key:      character.Kevin,
		Gender:   gender.Male,
		Name:     kevinName,
		Special:  false}
)

var (
	// kevinAmericanEnglishPhrase represents kevin american english phrase.
	kevinAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "weeweewee"}

	// kevinCanadianFrenchPhrase represents kevin canadian french phrase.
	kevinCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "toutébon"}

	// kevinDutchPhrase represents kevin dutch phrase.
	kevinDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ieieieie"}

	// kevinFrenchPhrase represents kevin french phrase.
	kevinFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "toutébon"}

	// kevinGermanPhrase represents kevin german phrase.
	kevinGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quiiiek"}

	// kevinItalianPhrase represents kevin italian phrase.
	kevinItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "boink"}

	// kevinJapanesePhrase represents kevin japanese phrase.
	kevinJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウリウリ"}

	// kevinKoreanPhrase represents kevin korean phrase.
	kevinKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냠냠"}

	// kevinLatinAmericanSpanishPhrase represents kevin latin american spanish phrase.
	kevinLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bidoink"}

	// kevinRussianPhrase represents kevin russian phrase.
	kevinRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "уи-и-и"}

	// kevinSimplifiedChinesePhrase represents kevin simplified chinese phrase.
	kevinSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "比利比利"}

	// kevinSpanishPhrase represents kevin spanish phrase.
	kevinSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bidoink"}

	// kevinTraditionalChinesePhrase represents kevin traditional chinese phrase.
	kevinTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "比利比利"}
)

var (
	// kevinPhrase represents kevin phrase.
	kevinPhrase = nook.Languages{
		language.AmericanEnglish:      kevinAmericanEnglishPhrase,
		language.CanadianFrench:       kevinCanadianFrenchPhrase,
		language.Dutch:                kevinDutchPhrase,
		language.French:               kevinFrenchPhrase,
		language.German:               kevinGermanPhrase,
		language.Italian:              kevinItalianPhrase,
		language.Japanese:             kevinJapanesePhrase,
		language.Korean:               kevinKoreanPhrase,
		language.LatinAmericanSpanish: kevinLatinAmericanSpanishPhrase,
		language.Russian:              kevinRussianPhrase,
		language.SimplifiedChinese:    kevinSimplifiedChinesePhrase,
		language.Spanish:              kevinSpanishPhrase,
		language.TraditionalChinese:   kevinTraditionalChinesePhrase}
)

var (
	// Kevin represents kevin.
	Kevin = nook.Villager{
		Character:   kevinCharacter,
		Personality: personality.Jock,
		Phrase:      kevinPhrase}
)

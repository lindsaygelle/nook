package gorilla

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
	// rocketBirthday represents rocket birthday.
	rocketBirthday = nook.Birthday{
		Day:   14,
		Month: time.April}
)

var (
	// rocketCode represents rocket code.
	rocketCode = nook.Code{
		Value: "gor09"}
)

var (
	// rocketAmericanEnglishName represents rocket american english name.
	rocketAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rocket"}

	// rocketCanadianFrenchName represents rocket canadian french name.
	rocketCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gertrude"}

	// rocketDutchName represents rocket dutch name.
	rocketDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rocket"}

	// rocketFrenchName represents rocket french name.
	rocketFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gertrude"}

	// rocketGermanName represents rocket german name.
	rocketGermanName = nook.Name{
		Language: language.German,
		Value:    "Katrin"}

	// rocketItalianName represents rocket italian name.
	rocketItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kinga"}

	// rocketJapaneseName represents rocket japanese name.
	rocketJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "４ごう"}

	// rocketKoreanName represents rocket korean name.
	rocketKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "4호"}

	// rocketLatinAmericanSpanishName represents rocket latin american spanish name.
	rocketLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gloria"}

	// rocketRussianName represents rocket russian name.
	rocketRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рокет"}

	// rocketSimplifiedChineseName represents rocket simplified chinese name.
	rocketSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿四"}

	// rocketSpanishName represents rocket spanish name.
	rocketSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gloria"}

	// rocketTraditionalChineseName represents rocket traditional chinese name.
	rocketTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿四"}
)

var (
	// rocketName represents rocket name.
	rocketName = nook.Languages{
		language.AmericanEnglish:      rocketAmericanEnglishName,
		language.CanadianFrench:       rocketCanadianFrenchName,
		language.Dutch:                rocketDutchName,
		language.French:               rocketFrenchName,
		language.German:               rocketGermanName,
		language.Italian:              rocketItalianName,
		language.Japanese:             rocketJapaneseName,
		language.Korean:               rocketKoreanName,
		language.LatinAmericanSpanish: rocketLatinAmericanSpanishName,
		language.Russian:              rocketRussianName,
		language.SimplifiedChinese:    rocketSimplifiedChineseName,
		language.Spanish:              rocketSpanishName,
		language.TraditionalChinese:   rocketTraditionalChineseName}
)

var (
	// rocketCharacter represents rocket character.
	rocketCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: rocketBirthday,
		Code:     rocketCode,
		Key:      character.Rocket,
		Gender:   gender.Female,
		Name:     rocketName,
		Special:  false}
)

var (
	// rocketAmericanEnglishPhrase represents rocket american english phrase.
	rocketAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "vroom"}

	// rocketCanadianFrenchPhrase represents rocket canadian french phrase.
	rocketCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "patouche"}

	// rocketDutchPhrase represents rocket dutch phrase.
	rocketDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "vroem"}

	// rocketFrenchPhrase represents rocket french phrase.
	rocketFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "patouche"}

	// rocketGermanPhrase represents rocket german phrase.
	rocketGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brahaha"}

	// rocketItalianPhrase represents rocket italian phrase.
	rocketItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uhatà"}

	// rocketJapanesePhrase represents rocket japanese phrase.
	rocketJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "そりゃっ"}

	// rocketKoreanPhrase represents rocket korean phrase.
	rocketKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "끼얍"}

	// rocketLatinAmericanSpanishPhrase represents rocket latin american spanish phrase.
	rocketLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uga-uga"}

	// rocketRussianPhrase represents rocket russian phrase.
	rocketRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "на взлет"}

	// rocketSimplifiedChinesePhrase represents rocket simplified chinese phrase.
	rocketSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "煞啊"}

	// rocketSpanishPhrase represents rocket spanish phrase.
	rocketSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uga-uga"}

	// rocketTraditionalChinesePhrase represents rocket traditional chinese phrase.
	rocketTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "煞啊"}
)

var (
	// rocketPhrase represents rocket phrase.
	rocketPhrase = nook.Languages{
		language.AmericanEnglish:      rocketAmericanEnglishPhrase,
		language.CanadianFrench:       rocketCanadianFrenchPhrase,
		language.Dutch:                rocketDutchPhrase,
		language.French:               rocketFrenchPhrase,
		language.German:               rocketGermanPhrase,
		language.Italian:              rocketItalianPhrase,
		language.Japanese:             rocketJapanesePhrase,
		language.Korean:               rocketKoreanPhrase,
		language.LatinAmericanSpanish: rocketLatinAmericanSpanishPhrase,
		language.Russian:              rocketRussianPhrase,
		language.SimplifiedChinese:    rocketSimplifiedChinesePhrase,
		language.Spanish:              rocketSpanishPhrase,
		language.TraditionalChinese:   rocketTraditionalChinesePhrase}
)

var (
	// Rocket represents rocket.
	Rocket = nook.Villager{
		Character:   rocketCharacter,
		Personality: personality.BigSister,
		Phrase:      rocketPhrase}
)

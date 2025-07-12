package dog

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
	// portiaBirthday represents portia birthday.
	portiaBirthday = nook.Birthday{
		Day:   25,
		Month: time.October}
)

var (
	// portiaCode represents portia code.
	portiaCode = nook.Code{
		Value: "dog05"}
)

var (
	// portiaAmericanEnglishName represents portia american english name.
	portiaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Portia"}

	// portiaCanadianFrenchName represents portia canadian french name.
	portiaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dalma"}

	// portiaDutchName represents portia dutch name.
	portiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Portia"}

	// portiaFrenchName represents portia french name.
	portiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dalma"}

	// portiaGermanName represents portia german name.
	portiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Isolde"}

	// portiaItalianName represents portia italian name.
	portiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Porzia"}

	// portiaJapaneseName represents portia japanese name.
	portiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブレンダ"}

	// portiaKoreanName represents portia korean name.
	portiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "블랜더"}

	// portiaLatinAmericanSpanishName represents portia latin american spanish name.
	portiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Amanda"}

	// portiaRussianName represents portia russian name.
	portiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Порция"}

	// portiaSimplifiedChineseName represents portia simplified chinese name.
	portiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "傅兰"}

	// portiaSpanishName represents portia spanish name.
	portiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Amanda"}

	// portiaTraditionalChineseName represents portia traditional chinese name.
	portiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "傅蘭"}
)

var (
	// portiaName represents portia name.
	portiaName = nook.Languages{
		language.AmericanEnglish:      portiaAmericanEnglishName,
		language.CanadianFrench:       portiaCanadianFrenchName,
		language.Dutch:                portiaDutchName,
		language.French:               portiaFrenchName,
		language.German:               portiaGermanName,
		language.Italian:              portiaItalianName,
		language.Japanese:             portiaJapaneseName,
		language.Korean:               portiaKoreanName,
		language.LatinAmericanSpanish: portiaLatinAmericanSpanishName,
		language.Russian:              portiaRussianName,
		language.SimplifiedChinese:    portiaSimplifiedChineseName,
		language.Spanish:              portiaSpanishName,
		language.TraditionalChinese:   portiaTraditionalChineseName}
)

var (
	// portiaCharacter represents portia character.
	portiaCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: portiaBirthday,
		Code:     portiaCode,
		Key:      character.Portia,
		Gender:   gender.Female,
		Name:     portiaName,
		Special:  false}
)

var (
	// portiaAmericanEnglishPhrase represents portia american english phrase.
	portiaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ruffian"}

	// portiaCanadianFrenchPhrase represents portia canadian french phrase.
	portiaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "brute"}

	// portiaDutchPhrase represents portia dutch phrase.
	portiaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "keffer"}

	// portiaFrenchPhrase represents portia french phrase.
	portiaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "brute"}

	// portiaGermanPhrase represents portia german phrase.
	portiaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "frechi"}

	// portiaItalianPhrase represents portia italian phrase.
	portiaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ronfietto"}

	// portiaJapanesePhrase represents portia japanese phrase.
	portiaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "フンッ"}

	// portiaKoreanPhrase represents portia korean phrase.
	portiaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흥"}

	// portiaLatinAmericanSpanishPhrase represents portia latin american spanish phrase.
	portiaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ruf-ruf"}

	// portiaRussianPhrase represents portia russian phrase.
	portiaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тяв-ряв"}

	// portiaSimplifiedChinesePhrase represents portia simplified chinese phrase.
	portiaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "屋屋"}

	// portiaSpanishPhrase represents portia spanish phrase.
	portiaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ruf-ruf"}

	// portiaTraditionalChinesePhrase represents portia traditional chinese phrase.
	portiaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "屋屋"}
)

var (
	// portiaPhrase represents portia phrase.
	portiaPhrase = nook.Languages{
		language.AmericanEnglish:      portiaAmericanEnglishPhrase,
		language.CanadianFrench:       portiaCanadianFrenchPhrase,
		language.Dutch:                portiaDutchPhrase,
		language.French:               portiaFrenchPhrase,
		language.German:               portiaGermanPhrase,
		language.Italian:              portiaItalianPhrase,
		language.Japanese:             portiaJapanesePhrase,
		language.Korean:               portiaKoreanPhrase,
		language.LatinAmericanSpanish: portiaLatinAmericanSpanishPhrase,
		language.Russian:              portiaRussianPhrase,
		language.SimplifiedChinese:    portiaSimplifiedChinesePhrase,
		language.Spanish:              portiaSpanishPhrase,
		language.TraditionalChinese:   portiaTraditionalChinesePhrase}
)

var (
	// Portia represents portia.
	Portia = nook.Villager{
		Character:   portiaCharacter,
		Personality: personality.Snooty,
		Phrase:      portiaPhrase}
)

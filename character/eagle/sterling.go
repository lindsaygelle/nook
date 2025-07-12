package eagle

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
	// sterlingBirthday represents sterling birthday.
	sterlingBirthday = nook.Birthday{
		Day:   11,
		Month: time.December}
)

var (
	// sterlingCode represents sterling code.
	sterlingCode = nook.Code{
		Value: "pbr07"}
)

var (
	// sterlingAmericanEnglishName represents sterling american english name.
	sterlingAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sterling"}

	// sterlingCanadianFrenchName represents sterling canadian french name.
	sterlingCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Manfred"}

	// sterlingDutchName represents sterling dutch name.
	sterlingDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sterling"}

	// sterlingFrenchName represents sterling french name.
	sterlingFrenchName = nook.Name{
		Language: language.French,
		Value:    "Manfred"}

	// sterlingGermanName represents sterling german name.
	sterlingGermanName = nook.Name{
		Language: language.German,
		Value:    "Horst"}

	// sterlingItalianName represents sterling italian name.
	sterlingItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Beccodì"}

	// sterlingJapaneseName represents sterling japanese name.
	sterlingJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ギンカク"}

	// sterlingKoreanName represents sterling korean name.
	sterlingKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "은수리"}

	// sterlingLatinAmericanSpanishName represents sterling latin american spanish name.
	sterlingLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Arni"}

	// sterlingRussianName represents sterling russian name.
	sterlingRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Стерлинг"}

	// sterlingSimplifiedChineseName represents sterling simplified chinese name.
	sterlingSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "银阁"}

	// sterlingSpanishName represents sterling spanish name.
	sterlingSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Arni"}

	// sterlingTraditionalChineseName represents sterling traditional chinese name.
	sterlingTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "銀閣"}
)

var (
	// sterlingName represents sterling name.
	sterlingName = nook.Languages{
		language.AmericanEnglish:      sterlingAmericanEnglishName,
		language.CanadianFrench:       sterlingCanadianFrenchName,
		language.Dutch:                sterlingDutchName,
		language.French:               sterlingFrenchName,
		language.German:               sterlingGermanName,
		language.Italian:              sterlingItalianName,
		language.Japanese:             sterlingJapaneseName,
		language.Korean:               sterlingKoreanName,
		language.LatinAmericanSpanish: sterlingLatinAmericanSpanishName,
		language.Russian:              sterlingRussianName,
		language.SimplifiedChinese:    sterlingSimplifiedChineseName,
		language.Spanish:              sterlingSpanishName,
		language.TraditionalChinese:   sterlingTraditionalChineseName}
)

var (
	// sterlingCharacter represents sterling character.
	sterlingCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: sterlingBirthday,
		Code:     sterlingCode,
		Key:      character.Sterling,
		Gender:   gender.Male,
		Name:     sterlingName,
		Special:  false}
)

var (
	// sterlingAmericanEnglishPhrase represents sterling american english phrase.
	sterlingAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "skraaaaw"}

	// sterlingCanadianFrenchPhrase represents sterling canadian french phrase.
	sterlingCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kaputt"}

	// sterlingDutchPhrase represents sterling dutch phrase.
	sterlingDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "havik"}

	// sterlingFrenchPhrase represents sterling french phrase.
	sterlingFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "kaputt"}

	// sterlingGermanPhrase represents sterling german phrase.
	sterlingGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krahkrah"}

	// sterlingItalianPhrase represents sterling italian phrase.
	sterlingItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "chicchirò"}

	// sterlingJapanesePhrase represents sterling japanese phrase.
	sterlingJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やあッ"}

	// sterlingKoreanPhrase represents sterling korean phrase.
	sterlingKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "호옷"}

	// sterlingLatinAmericanSpanishPhrase represents sterling latin american spanish phrase.
	sterlingLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "escróóó"}

	// sterlingRussianPhrase represents sterling russian phrase.
	sterlingRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "клекот"}

	// sterlingSimplifiedChinesePhrase represents sterling simplified chinese phrase.
	sterlingSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呀啊"}

	// sterlingSpanishPhrase represents sterling spanish phrase.
	sterlingSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mancuernas"}

	// sterlingTraditionalChinesePhrase represents sterling traditional chinese phrase.
	sterlingTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "呀啊"}
)

var (
	// sterlingPhrase represents sterling phrase.
	sterlingPhrase = nook.Languages{
		language.AmericanEnglish:      sterlingAmericanEnglishPhrase,
		language.CanadianFrench:       sterlingCanadianFrenchPhrase,
		language.Dutch:                sterlingDutchPhrase,
		language.French:               sterlingFrenchPhrase,
		language.German:               sterlingGermanPhrase,
		language.Italian:              sterlingItalianPhrase,
		language.Japanese:             sterlingJapanesePhrase,
		language.Korean:               sterlingKoreanPhrase,
		language.LatinAmericanSpanish: sterlingLatinAmericanSpanishPhrase,
		language.Russian:              sterlingRussianPhrase,
		language.SimplifiedChinese:    sterlingSimplifiedChinesePhrase,
		language.Spanish:              sterlingSpanishPhrase,
		language.TraditionalChinese:   sterlingTraditionalChinesePhrase}
)

var (
	// Sterling represents sterling.
	Sterling = nook.Villager{
		Character:   sterlingCharacter,
		Personality: personality.Jock,
		Phrase:      sterlingPhrase}
)

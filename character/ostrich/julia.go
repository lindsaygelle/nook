package ostrich

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
	// juliaBirthday represents julia birthday.
	juliaBirthday = nook.Birthday{
		Day:   31,
		Month: time.July}
)

var (
	// juliaCode represents julia code.
	juliaCode = nook.Code{
		Value: "ost05"}
)

var (
	// juliaAmericanEnglishName represents julia american english name.
	juliaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Julia"}

	// juliaCanadianFrenchName represents julia canadian french name.
	juliaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Julie"}

	// juliaDutchName represents julia dutch name.
	juliaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Julia"}

	// juliaFrenchName represents julia french name.
	juliaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Julie"}

	// juliaGermanName represents julia german name.
	juliaGermanName = nook.Name{
		Language: language.German,
		Value:    "Julia"}

	// juliaItalianName represents julia italian name.
	juliaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giulia"}

	// juliaJapaneseName represents julia japanese name.
	juliaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュリア"}

	// juliaKoreanName represents julia korean name.
	juliaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "줄리아"}

	// juliaLatinAmericanSpanishName represents julia latin american spanish name.
	juliaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Julia"}

	// juliaRussianName represents julia russian name.
	juliaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джулия"}

	// juliaSimplifiedChineseName represents julia simplified chinese name.
	juliaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱莉亚"}

	// juliaSpanishName represents julia spanish name.
	juliaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Julia"}

	// juliaTraditionalChineseName represents julia traditional chinese name.
	juliaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱莉亞"}
)

var (
	// juliaName represents julia name.
	juliaName = nook.Languages{
		language.AmericanEnglish:      juliaAmericanEnglishName,
		language.CanadianFrench:       juliaCanadianFrenchName,
		language.Dutch:                juliaDutchName,
		language.French:               juliaFrenchName,
		language.German:               juliaGermanName,
		language.Italian:              juliaItalianName,
		language.Japanese:             juliaJapaneseName,
		language.Korean:               juliaKoreanName,
		language.LatinAmericanSpanish: juliaLatinAmericanSpanishName,
		language.Russian:              juliaRussianName,
		language.SimplifiedChinese:    juliaSimplifiedChineseName,
		language.Spanish:              juliaSpanishName,
		language.TraditionalChinese:   juliaTraditionalChineseName}
)

var (
	// juliaCharacter represents julia character.
	juliaCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: juliaBirthday,
		Code:     juliaCode,
		Key:      character.Julia,
		Gender:   gender.Female,
		Name:     juliaName,
		Special:  false}
)

var (
	// juliaAmericanEnglishPhrase represents julia american english phrase.
	juliaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "dahling"}

	// juliaCanadianFrenchPhrase represents julia canadian french phrase.
	juliaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "trutruche"}

	// juliaDutchPhrase represents julia dutch phrase.
	juliaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schatje"}

	// juliaFrenchPhrase represents julia french phrase.
	juliaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trutruche"}

	// juliaGermanPhrase represents julia german phrase.
	juliaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "püh"}

	// juliaItalianPhrase represents julia italian phrase.
	juliaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blaaah"}

	// juliaJapanesePhrase represents julia japanese phrase.
	juliaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やだわね"}

	// juliaKoreanPhrase represents julia korean phrase.
	juliaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어머몰라"}

	// juliaLatinAmericanSpanishPhrase represents julia latin american spanish phrase.
	juliaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "churri"}

	// juliaRussianPhrase represents julia russian phrase.
	juliaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дарлинг"}

	// juliaSimplifiedChinesePhrase represents julia simplified chinese phrase.
	juliaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吼唷"}

	// juliaSpanishPhrase represents julia spanish phrase.
	juliaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "churri"}

	// juliaTraditionalChinesePhrase represents julia traditional chinese phrase.
	juliaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吼唷"}
)

var (
	// juliaPhrase represents julia phrase.
	juliaPhrase = nook.Languages{
		language.AmericanEnglish:      juliaAmericanEnglishPhrase,
		language.CanadianFrench:       juliaCanadianFrenchPhrase,
		language.Dutch:                juliaDutchPhrase,
		language.French:               juliaFrenchPhrase,
		language.German:               juliaGermanPhrase,
		language.Italian:              juliaItalianPhrase,
		language.Japanese:             juliaJapanesePhrase,
		language.Korean:               juliaKoreanPhrase,
		language.LatinAmericanSpanish: juliaLatinAmericanSpanishPhrase,
		language.Russian:              juliaRussianPhrase,
		language.SimplifiedChinese:    juliaSimplifiedChinesePhrase,
		language.Spanish:              juliaSpanishPhrase,
		language.TraditionalChinese:   juliaTraditionalChinesePhrase}
)

var (
	// Julia represents julia.
	Julia = nook.Villager{
		Character:   juliaCharacter,
		Personality: personality.Snooty,
		Phrase:      juliaPhrase}
)

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
	// ameliaBirthday represents amelia birthday.
	ameliaBirthday = nook.Birthday{
		Day:   19,
		Month: time.November}
)

var (
	// ameliaCode represents amelia code.
	ameliaCode = nook.Code{
		Value: "pbr01"}
)

var (
	// ameliaAmericanEnglishName represents amelia american english name.
	ameliaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Amelia"}

	// ameliaCanadianFrenchName represents amelia canadian french name.
	ameliaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Aurélie"}

	// ameliaDutchName represents amelia dutch name.
	ameliaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Amelia"}

	// ameliaFrenchName represents amelia french name.
	ameliaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Aurélie"}

	// ameliaGermanName represents amelia german name.
	ameliaGermanName = nook.Name{
		Language: language.German,
		Value:    "Adelheid"}

	// ameliaItalianName represents amelia italian name.
	ameliaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Amelia"}

	// ameliaJapaneseName represents amelia japanese name.
	ameliaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アンデス"}

	// ameliaKoreanName represents amelia korean name.
	ameliaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안데스"}

	// ameliaLatinAmericanSpanishName represents amelia latin american spanish name.
	ameliaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Amelia"}

	// ameliaRussianName represents amelia russian name.
	ameliaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Амелия"}

	// ameliaSimplifiedChineseName represents amelia simplified chinese name.
	ameliaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "安地斯"}

	// ameliaSpanishName represents amelia spanish name.
	ameliaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Amelia"}

	// ameliaTraditionalChineseName represents amelia traditional chinese name.
	ameliaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "安地斯"}
)

var (
	// ameliaName represents amelia name.
	ameliaName = nook.Languages{
		language.AmericanEnglish:      ameliaAmericanEnglishName,
		language.CanadianFrench:       ameliaCanadianFrenchName,
		language.Dutch:                ameliaDutchName,
		language.French:               ameliaFrenchName,
		language.German:               ameliaGermanName,
		language.Italian:              ameliaItalianName,
		language.Japanese:             ameliaJapaneseName,
		language.Korean:               ameliaKoreanName,
		language.LatinAmericanSpanish: ameliaLatinAmericanSpanishName,
		language.Russian:              ameliaRussianName,
		language.SimplifiedChinese:    ameliaSimplifiedChineseName,
		language.Spanish:              ameliaSpanishName,
		language.TraditionalChinese:   ameliaTraditionalChineseName}
)

var (
	// ameliaCharacter represents amelia character.
	ameliaCharacter = nook.Character{
		Animal:   animal.Eagle,
		Birthday: ameliaBirthday,
		Code:     ameliaCode,
		Key:      character.Amelia,
		Gender:   gender.Female,
		Name:     ameliaName,
		Special:  false}
)

var (
	// ameliaAmericanEnglishPhrase represents amelia american english phrase.
	ameliaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "eaglet"}

	// ameliaCanadianFrenchPhrase represents amelia canadian french phrase.
	ameliaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "croooa"}

	// ameliaDutchPhrase represents amelia dutch phrase.
	ameliaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "broeder"}

	// ameliaFrenchPhrase represents amelia french phrase.
	ameliaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croooa"}

	// ameliaGermanPhrase represents amelia german phrase.
	ameliaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "adlerauge"}

	// ameliaItalianPhrase represents amelia italian phrase.
	ameliaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mini mini"}

	// ameliaJapanesePhrase represents amelia japanese phrase.
	ameliaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "カラカラ"}

	// ameliaKoreanPhrase represents amelia korean phrase.
	ameliaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "엄멈머"}

	// ameliaLatinAmericanSpanishPhrase represents amelia latin american spanish phrase.
	ameliaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cro-ah"}

	// ameliaRussianPhrase represents amelia russian phrase.
	ameliaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вот"}

	// ameliaSimplifiedChinesePhrase represents amelia simplified chinese phrase.
	ameliaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卡拉卡拉"}

	// ameliaSpanishPhrase represents amelia spanish phrase.
	ameliaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chinchilla"}

	// ameliaTraditionalChinesePhrase represents amelia traditional chinese phrase.
	ameliaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "卡拉卡拉"}
)

var (
	// ameliaPhrase represents amelia phrase.
	ameliaPhrase = nook.Languages{
		language.AmericanEnglish:      ameliaAmericanEnglishPhrase,
		language.CanadianFrench:       ameliaCanadianFrenchPhrase,
		language.Dutch:                ameliaDutchPhrase,
		language.French:               ameliaFrenchPhrase,
		language.German:               ameliaGermanPhrase,
		language.Italian:              ameliaItalianPhrase,
		language.Japanese:             ameliaJapanesePhrase,
		language.Korean:               ameliaKoreanPhrase,
		language.LatinAmericanSpanish: ameliaLatinAmericanSpanishPhrase,
		language.Russian:              ameliaRussianPhrase,
		language.SimplifiedChinese:    ameliaSimplifiedChinesePhrase,
		language.Spanish:              ameliaSpanishPhrase,
		language.TraditionalChinese:   ameliaTraditionalChinesePhrase}
)

var (
	// Amelia represents amelia.
	Amelia = nook.Villager{
		Character:   ameliaCharacter,
		Personality: personality.Snooty,
		Phrase:      ameliaPhrase}
)

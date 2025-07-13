package elephant

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
	// axelBirthday represents axel birthday.
	axelBirthday = nook.Birthday{
		Day:   23,
		Month: time.March}
)

var (
	// axelCode represents axel code.
	axelCode = nook.Code{
		Value: "elp06"}
)

var (
	// axelAmericanEnglishName represents axel american english name.
	axelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Axel"}

	// axelCanadianFrenchName represents axel canadian french name.
	axelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Axel"}

	// axelDutchName represents axel dutch name.
	axelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Axel"}

	// axelFrenchName represents axel french name.
	axelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Axel"}

	// axelGermanName represents axel german name.
	axelGermanName = nook.Name{
		Language: language.German,
		Value:    "Axel"}

	// axelItalianName represents axel italian name.
	axelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sandro"}

	// axelJapaneseName represents axel japanese name.
	axelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エックスエル"}

	// axelKoreanName represents axel korean name.
	axelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "엑스엘리"}

	// axelLatinAmericanSpanishName represents axel latin american spanish name.
	axelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eustakio"}

	// axelRussianName represents axel russian name.
	axelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аксель"}

	// axelSimplifiedChineseName represents axel simplified chinese name.
	axelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大大"}

	// axelSpanishName represents axel spanish name.
	axelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eustakio"}

	// axelTraditionalChineseName represents axel traditional chinese name.
	axelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大大"}
)

var (
	// axelName represents axel name.
	axelName = nook.Languages{
		language.AmericanEnglish:      axelAmericanEnglishName,
		language.CanadianFrench:       axelCanadianFrenchName,
		language.Dutch:                axelDutchName,
		language.French:               axelFrenchName,
		language.German:               axelGermanName,
		language.Italian:              axelItalianName,
		language.Japanese:             axelJapaneseName,
		language.Korean:               axelKoreanName,
		language.LatinAmericanSpanish: axelLatinAmericanSpanishName,
		language.Russian:              axelRussianName,
		language.SimplifiedChinese:    axelSimplifiedChineseName,
		language.Spanish:              axelSpanishName,
		language.TraditionalChinese:   axelTraditionalChineseName}
)

var (
	// axelCharacter represents axel character.
	axelCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: axelBirthday,
		Code:     axelCode,
		Key:      character.Axel,
		Gender:   gender.Male,
		Name:     axelName,
		Special:  false}
)

var (
	// axelAmericanEnglishPhrase represents axel american english phrase.
	axelAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "WHONK"}

	// axelCanadianFrenchPhrase represents axel canadian french phrase.
	axelCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "splaf"}

	// axelDutchPhrase represents axel dutch phrase.
	axelDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "PWAAP"}

	// axelFrenchPhrase represents axel french phrase.
	axelFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "splaf"}

	// axelGermanPhrase represents axel german phrase.
	axelGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "TUUUUT"}

	// axelItalianPhrase represents axel italian phrase.
	axelItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "SBONK"}

	// axelJapanesePhrase represents axel japanese phrase.
	axelJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でゴンス"}

	// axelKoreanPhrase represents axel korean phrase.
	axelKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "히힛"}

	// axelLatinAmericanSpanishPhrase represents axel latin american spanish phrase.
	axelLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ankagua"}

	// axelRussianPhrase represents axel russian phrase.
	axelRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "потрубим"}

	// axelSimplifiedChinesePhrase represents axel simplified chinese phrase.
	axelSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘻嘻"}

	// axelSpanishPhrase represents axel spanish phrase.
	axelSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ankagua"}

	// axelTraditionalChinesePhrase represents axel traditional chinese phrase.
	axelTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘻嘻"}
)

var (
	// axelPhrase represents axel phrase.
	axelPhrase = nook.Languages{
		language.AmericanEnglish:      axelAmericanEnglishPhrase,
		language.CanadianFrench:       axelCanadianFrenchPhrase,
		language.Dutch:                axelDutchPhrase,
		language.French:               axelFrenchPhrase,
		language.German:               axelGermanPhrase,
		language.Italian:              axelItalianPhrase,
		language.Japanese:             axelJapanesePhrase,
		language.Korean:               axelKoreanPhrase,
		language.LatinAmericanSpanish: axelLatinAmericanSpanishPhrase,
		language.Russian:              axelRussianPhrase,
		language.SimplifiedChinese:    axelSimplifiedChinesePhrase,
		language.Spanish:              axelSpanishPhrase,
		language.TraditionalChinese:   axelTraditionalChinesePhrase}
)

var (
	// Axel represents axel.
	Axel = nook.Villager{
		Character:   axelCharacter,
		Personality: personality.Jock,
		Phrase:      axelPhrase}
)

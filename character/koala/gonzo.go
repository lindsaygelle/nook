package koala

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
	// gonzoBirthday represents gonzo birthday.
	gonzoBirthday = nook.Birthday{
		Day:   13,
		Month: time.October}
)

var (
	// gonzoCode represents gonzo code.
	gonzoCode = nook.Code{
		Value: "kal04"}
)

var (
	// gonzoAmericanEnglishName represents gonzo american english name.
	gonzoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gonzo"}

	// gonzoCanadianFrenchName represents gonzo canadian french name.
	gonzoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gonzo"}

	// gonzoDutchName represents gonzo dutch name.
	gonzoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gonzo"}

	// gonzoFrenchName represents gonzo french name.
	gonzoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gonzo"}

	// gonzoGermanName represents gonzo german name.
	gonzoGermanName = nook.Name{
		Language: language.German,
		Value:    "Heribert"}

	// gonzoItalianName represents gonzo italian name.
	gonzoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gonzo"}

	// gonzoJapaneseName represents gonzo japanese name.
	gonzoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゴンゾー"}

	// gonzoKoreanName represents gonzo korean name.
	gonzoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "근성"}

	// gonzoLatinAmericanSpanishName represents gonzo latin american spanish name.
	gonzoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bronco"}

	// gonzoRussianName represents gonzo russian name.
	gonzoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гонзо"}

	// gonzoSimplifiedChineseName represents gonzo simplified chinese name.
	gonzoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "钢锁"}

	// gonzoSpanishName represents gonzo spanish name.
	gonzoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bronco"}

	// gonzoTraditionalChineseName represents gonzo traditional chinese name.
	gonzoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鋼鎖"}
)

var (
	// gonzoName represents gonzo name.
	gonzoName = nook.Languages{
		language.AmericanEnglish:      gonzoAmericanEnglishName,
		language.CanadianFrench:       gonzoCanadianFrenchName,
		language.Dutch:                gonzoDutchName,
		language.French:               gonzoFrenchName,
		language.German:               gonzoGermanName,
		language.Italian:              gonzoItalianName,
		language.Japanese:             gonzoJapaneseName,
		language.Korean:               gonzoKoreanName,
		language.LatinAmericanSpanish: gonzoLatinAmericanSpanishName,
		language.Russian:              gonzoRussianName,
		language.SimplifiedChinese:    gonzoSimplifiedChineseName,
		language.Spanish:              gonzoSpanishName,
		language.TraditionalChinese:   gonzoTraditionalChineseName}
)

var (
	// gonzoCharacter represents gonzo character.
	gonzoCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: gonzoBirthday,
		Code:     gonzoCode,
		Key:      character.Gonzo,
		Gender:   gender.Male,
		Name:     gonzoName,
		Special:  false}
)

var (
	// gonzoAmericanEnglishPhrase represents gonzo american english phrase.
	gonzoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mate"}

	// gonzoCanadianFrenchPhrase represents gonzo canadian french phrase.
	gonzoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "calyptus"}

	// gonzoDutchPhrase represents gonzo dutch phrase.
	gonzoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "partner"}

	// gonzoFrenchPhrase represents gonzo french phrase.
	gonzoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "galopin"}

	// gonzoGermanPhrase represents gonzo german phrase.
	gonzoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "partner"}

	// gonzoItalianPhrase represents gonzo italian phrase.
	gonzoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "dinamite"}

	// gonzoJapanesePhrase represents gonzo japanese phrase.
	gonzoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だがや"}

	// gonzoKoreanPhrase represents gonzo korean phrase.
	gonzoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "버텨"}

	// gonzoLatinAmericanSpanishPhrase represents gonzo latin american spanish phrase.
	gonzoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jorf-jorf"}

	// gonzoRussianPhrase represents gonzo russian phrase.
	gonzoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "напарник"}

	// gonzoSimplifiedChinesePhrase represents gonzo simplified chinese phrase.
	gonzoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "钱钱"}

	// gonzoSpanishPhrase represents gonzo spanish phrase.
	gonzoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jorf-jorf"}

	// gonzoTraditionalChinesePhrase represents gonzo traditional chinese phrase.
	gonzoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "錢錢"}
)

var (
	// gonzoPhrase represents gonzo phrase.
	gonzoPhrase = nook.Languages{
		language.AmericanEnglish:      gonzoAmericanEnglishPhrase,
		language.CanadianFrench:       gonzoCanadianFrenchPhrase,
		language.Dutch:                gonzoDutchPhrase,
		language.French:               gonzoFrenchPhrase,
		language.German:               gonzoGermanPhrase,
		language.Italian:              gonzoItalianPhrase,
		language.Japanese:             gonzoJapanesePhrase,
		language.Korean:               gonzoKoreanPhrase,
		language.LatinAmericanSpanish: gonzoLatinAmericanSpanishPhrase,
		language.Russian:              gonzoRussianPhrase,
		language.SimplifiedChinese:    gonzoSimplifiedChinesePhrase,
		language.Spanish:              gonzoSpanishPhrase,
		language.TraditionalChinese:   gonzoTraditionalChinesePhrase}
)

var (
	// Gonzo represents gonzo.
	Gonzo = nook.Villager{
		Character:   gonzoCharacter,
		Personality: personality.Cranky,
		Phrase:      gonzoPhrase}
)

package deer

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
	// zellBirthday represents zell birthday.
	zellBirthday = nook.Birthday{
		Day:   7,
		Month: time.June}
)

var (
	// zellCode represents zell code.
	zellCode = nook.Code{
		Value: "der02"}
)

var (
	// zellAmericanEnglishName represents zell american english name.
	zellAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Zell"}

	// zellCanadianFrenchName represents zell canadian french name.
	zellCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Régis"}

	// zellDutchName represents zell dutch name.
	zellDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Zell"}

	// zellFrenchName represents zell french name.
	zellFrenchName = nook.Name{
		Language: language.French,
		Value:    "Régis"}

	// zellGermanName represents zell german name.
	zellGermanName = nook.Name{
		Language: language.German,
		Value:    "Walter"}

	// zellItalianName represents zell italian name.
	zellItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Antilio"}

	// zellJapaneseName represents zell japanese name.
	zellJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ネルソン"}

	// zellKoreanName represents zell korean name.
	zellKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "넬슨"}

	// zellLatinAmericanSpanishName represents zell latin american spanish name.
	zellLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Corvilo"}

	// zellRussianName represents zell russian name.
	zellRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Зелл"}

	// zellSimplifiedChineseName represents zell simplified chinese name.
	zellSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "森森"}

	// zellSpanishName represents zell spanish name.
	zellSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Corvilo"}

	// zellTraditionalChineseName represents zell traditional chinese name.
	zellTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "森森"}
)

var (
	// zellName represents zell name.
	zellName = nook.Languages{
		language.AmericanEnglish:      zellAmericanEnglishName,
		language.CanadianFrench:       zellCanadianFrenchName,
		language.Dutch:                zellDutchName,
		language.French:               zellFrenchName,
		language.German:               zellGermanName,
		language.Italian:              zellItalianName,
		language.Japanese:             zellJapaneseName,
		language.Korean:               zellKoreanName,
		language.LatinAmericanSpanish: zellLatinAmericanSpanishName,
		language.Russian:              zellRussianName,
		language.SimplifiedChinese:    zellSimplifiedChineseName,
		language.Spanish:              zellSpanishName,
		language.TraditionalChinese:   zellTraditionalChineseName}
)

var (
	// zellCharacter represents zell character.
	zellCharacter = nook.Character{
		Animal:   animal.Deer,
		Birthday: zellBirthday,
		Code:     zellCode,
		Key:      character.Zell,
		Gender:   gender.Male,
		Name:     zellName,
		Special:  false}
)

var (
	// zellAmericanEnglishPhrase represents zell american english phrase.
	zellAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pronk"}

	// zellCanadianFrenchPhrase represents zell canadian french phrase.
	zellCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "wapiti"}

	// zellDutchPhrase represents zell dutch phrase.
	zellDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gazellig"}

	// zellFrenchPhrase represents zell french phrase.
	zellFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "wapiti"}

	// zellGermanPhrase represents zell german phrase.
	zellGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "sproing"}

	// zellItalianPhrase represents zell italian phrase.
	zellItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arriba"}

	// zellJapanesePhrase represents zell japanese phrase.
	zellJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "たしかに"}

	// zellKoreanPhrase represents zell korean phrase.
	zellKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "슴사사"}

	// zellLatinAmericanSpanishPhrase represents zell latin american spanish phrase.
	zellLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bambín"}

	// zellRussianPhrase represents zell russian phrase.
	zellRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "прыг-прыг"}

	// zellSimplifiedChinesePhrase represents zell simplified chinese phrase.
	zellSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鹿营"}

	// zellSpanishPhrase represents zell spanish phrase.
	zellSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bambín"}

	// zellTraditionalChinesePhrase represents zell traditional chinese phrase.
	zellTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鹿營"}
)

var (
	// zellPhrase represents zell phrase.
	zellPhrase = nook.Languages{
		language.AmericanEnglish:      zellAmericanEnglishPhrase,
		language.CanadianFrench:       zellCanadianFrenchPhrase,
		language.Dutch:                zellDutchPhrase,
		language.French:               zellFrenchPhrase,
		language.German:               zellGermanPhrase,
		language.Italian:              zellItalianPhrase,
		language.Japanese:             zellJapanesePhrase,
		language.Korean:               zellKoreanPhrase,
		language.LatinAmericanSpanish: zellLatinAmericanSpanishPhrase,
		language.Russian:              zellRussianPhrase,
		language.SimplifiedChinese:    zellSimplifiedChinesePhrase,
		language.Spanish:              zellSpanishPhrase,
		language.TraditionalChinese:   zellTraditionalChinesePhrase}
)

var (
	// Zell represents zell.
	Zell = nook.Villager{
		Character:   zellCharacter,
		Personality: personality.Smug,
		Phrase:      zellPhrase}
)

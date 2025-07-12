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
	// faithBirthday represents faith birthday.
	faithBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// faithCode represents faith code.
	faithCode = nook.Code{
		Value: ""}
)

var (
	// faithAmericanEnglishName represents faith american english name.
	faithAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Faith"}

	// faithCanadianFrenchName represents faith canadian french name.
	faithCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// faithDutchName represents faith dutch name.
	faithDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// faithFrenchName represents faith french name.
	faithFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kolette"}

	// faithGermanName represents faith german name.
	faithGermanName = nook.Name{
		Language: language.German,
		Value:    "Finchen"}

	// faithItalianName represents faith italian name.
	faithItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Calipso"}

	// faithJapaneseName represents faith japanese name.
	faithJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マーチ"}

	// faithKoreanName represents faith korean name.
	faithKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// faithLatinAmericanSpanishName represents faith latin american spanish name.
	faithLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// faithRussianName represents faith russian name.
	faithRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// faithSimplifiedChineseName represents faith simplified chinese name.
	faithSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// faithSpanishName represents faith spanish name.
	faithSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Felicia"}

	// faithTraditionalChineseName represents faith traditional chinese name.
	faithTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// faithName represents faith name.
	faithName = nook.Languages{
		language.AmericanEnglish:      faithAmericanEnglishName,
		language.CanadianFrench:       faithCanadianFrenchName,
		language.Dutch:                faithDutchName,
		language.French:               faithFrenchName,
		language.German:               faithGermanName,
		language.Italian:              faithItalianName,
		language.Japanese:             faithJapaneseName,
		language.Korean:               faithKoreanName,
		language.LatinAmericanSpanish: faithLatinAmericanSpanishName,
		language.Russian:              faithRussianName,
		language.SimplifiedChinese:    faithSimplifiedChineseName,
		language.Spanish:              faithSpanishName,
		language.TraditionalChinese:   faithTraditionalChineseName}
)

var (
	// faithCharacter represents faith character.
	faithCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: faithBirthday,
		Code:     faithCode,
		Key:      character.Faith,
		Gender:   gender.Female,
		Name:     faithName,
		Special:  false}
)

var (
	// faithAmericanEnglishPhrase represents faith american english phrase.
	faithAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "aloha"}

	// faithCanadianFrenchPhrase represents faith canadian french phrase.
	faithCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// faithDutchPhrase represents faith dutch phrase.
	faithDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// faithFrenchPhrase represents faith french phrase.
	faithFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "aloha"}

	// faithGermanPhrase represents faith german phrase.
	faithGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "aloha"}

	// faithItalianPhrase represents faith italian phrase.
	faithItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "aloha"}

	// faithJapanesePhrase represents faith japanese phrase.
	faithJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "サクッ"}

	// faithKoreanPhrase represents faith korean phrase.
	faithKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// faithLatinAmericanSpanishPhrase represents faith latin american spanish phrase.
	faithLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// faithRussianPhrase represents faith russian phrase.
	faithRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// faithSimplifiedChinesePhrase represents faith simplified chinese phrase.
	faithSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// faithSpanishPhrase represents faith spanish phrase.
	faithSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "aloha"}

	// faithTraditionalChinesePhrase represents faith traditional chinese phrase.
	faithTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// faithPhrase represents faith phrase.
	faithPhrase = nook.Languages{
		language.AmericanEnglish:      faithAmericanEnglishPhrase,
		language.CanadianFrench:       faithCanadianFrenchPhrase,
		language.Dutch:                faithDutchPhrase,
		language.French:               faithFrenchPhrase,
		language.German:               faithGermanPhrase,
		language.Italian:              faithItalianPhrase,
		language.Japanese:             faithJapanesePhrase,
		language.Korean:               faithKoreanPhrase,
		language.LatinAmericanSpanish: faithLatinAmericanSpanishPhrase,
		language.Russian:              faithRussianPhrase,
		language.SimplifiedChinese:    faithSimplifiedChinesePhrase,
		language.Spanish:              faithSpanishPhrase,
		language.TraditionalChinese:   faithTraditionalChinesePhrase}
)

var (
	// Faith represents faith.
	Faith = nook.Villager{
		Character:   faithCharacter,
		Personality: personality.Normal,
		Phrase:      faithPhrase}
)

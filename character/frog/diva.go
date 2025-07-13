package frog

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
	// divaBirthday represents diva birthday.
	divaBirthday = nook.Birthday{
		Day:   2,
		Month: time.October}
)

var (
	// divaCode represents diva code.
	divaCode = nook.Code{
		Value: "flg18"}
)

var (
	// divaAmericanEnglishName represents diva american english name.
	divaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Diva"}

	// divaCanadianFrenchName represents diva canadian french name.
	divaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Violette"}

	// divaDutchName represents diva dutch name.
	divaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Diva"}

	// divaFrenchName represents diva french name.
	divaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Violette"}

	// divaGermanName represents diva german name.
	divaGermanName = nook.Name{
		Language: language.German,
		Value:    "Dörte"}

	// divaItalianName represents diva italian name.
	divaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Viola"}

	// divaJapaneseName represents diva japanese name.
	divaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイーダ"}

	// divaKoreanName represents diva korean name.
	divaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이다"}

	// divaLatinAmericanSpanishName represents diva latin american spanish name.
	divaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Morania"}

	// divaRussianName represents diva russian name.
	divaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дива"}

	// divaSimplifiedChineseName represents diva simplified chinese name.
	divaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小睫"}

	// divaSpanishName represents diva spanish name.
	divaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Morania"}

	// divaTraditionalChineseName represents diva traditional chinese name.
	divaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小睫"}
)

var (
	// divaName represents diva name.
	divaName = nook.Languages{
		language.AmericanEnglish:      divaAmericanEnglishName,
		language.CanadianFrench:       divaCanadianFrenchName,
		language.Dutch:                divaDutchName,
		language.French:               divaFrenchName,
		language.German:               divaGermanName,
		language.Italian:              divaItalianName,
		language.Japanese:             divaJapaneseName,
		language.Korean:               divaKoreanName,
		language.LatinAmericanSpanish: divaLatinAmericanSpanishName,
		language.Russian:              divaRussianName,
		language.SimplifiedChinese:    divaSimplifiedChineseName,
		language.Spanish:              divaSpanishName,
		language.TraditionalChinese:   divaTraditionalChineseName}
)

var (
	// divaCharacter represents diva character.
	divaCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: divaBirthday,
		Code:     divaCode,
		Key:      character.Diva,
		Gender:   gender.Female,
		Name:     divaName,
		Special:  false}
)

var (
	// divaAmericanEnglishPhrase represents diva american english phrase.
	divaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ya know"}

	// divaCanadianFrenchPhrase represents diva canadian french phrase.
	divaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cuicuisse"}

	// divaDutchPhrase represents diva dutch phrase.
	divaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jeweetwel"}

	// divaFrenchPhrase represents diva french phrase.
	divaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cuicuisse"}

	// divaGermanPhrase represents diva german phrase.
	divaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "platsch"}

	// divaItalianPhrase represents diva italian phrase.
	divaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "crabum"}

	// divaJapanesePhrase represents diva japanese phrase.
	divaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ハーン"}

	// divaKoreanPhrase represents diva korean phrase.
	divaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흐응"}

	// divaLatinAmericanSpanishPhrase represents diva latin american spanish phrase.
	divaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "crocró"}

	// divaRussianPhrase represents diva russian phrase.
	divaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "знаешь"}

	// divaSimplifiedChinesePhrase represents diva simplified chinese phrase.
	divaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蛤"}

	// divaSpanishPhrase represents diva spanish phrase.
	divaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "crocró"}

	// divaTraditionalChinesePhrase represents diva traditional chinese phrase.
	divaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蛤"}
)

var (
	// divaPhrase represents diva phrase.
	divaPhrase = nook.Languages{
		language.AmericanEnglish:      divaAmericanEnglishPhrase,
		language.CanadianFrench:       divaCanadianFrenchPhrase,
		language.Dutch:                divaDutchPhrase,
		language.French:               divaFrenchPhrase,
		language.German:               divaGermanPhrase,
		language.Italian:              divaItalianPhrase,
		language.Japanese:             divaJapanesePhrase,
		language.Korean:               divaKoreanPhrase,
		language.LatinAmericanSpanish: divaLatinAmericanSpanishPhrase,
		language.Russian:              divaRussianPhrase,
		language.SimplifiedChinese:    divaSimplifiedChinesePhrase,
		language.Spanish:              divaSpanishPhrase,
		language.TraditionalChinese:   divaTraditionalChinesePhrase}
)

var (
	// Diva represents diva.
	Diva = nook.Villager{
		Character:   divaCharacter,
		Personality: personality.BigSister,
		Phrase:      divaPhrase}
)

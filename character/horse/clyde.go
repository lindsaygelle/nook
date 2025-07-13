package horse

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
	// clydeBirthday represents clyde birthday.
	clydeBirthday = nook.Birthday{
		Day:   1,
		Month: time.May}
)

var (
	// clydeCode represents clyde code.
	clydeCode = nook.Code{
		Value: "hrs10"}
)

var (
	// clydeAmericanEnglishName represents clyde american english name.
	clydeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Clyde"}

	// clydeCanadianFrenchName represents clyde canadian french name.
	clydeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dorian"}

	// clydeDutchName represents clyde dutch name.
	clydeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Clyde"}

	// clydeFrenchName represents clyde french name.
	clydeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dorian"}

	// clydeGermanName represents clyde german name.
	clydeGermanName = nook.Name{
		Language: language.German,
		Value:    "Tommi"}

	// clydeItalianName represents clyde italian name.
	clydeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lallo"}

	// clydeJapaneseName represents clyde japanese name.
	clydeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "デースケ"}

	// clydeKoreanName represents clyde korean name.
	clydeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마철이"}

	// clydeLatinAmericanSpanishName represents clyde latin american spanish name.
	clydeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Moe"}

	// clydeRussianName represents clyde russian name.
	clydeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клайд"}

	// clydeSimplifiedChineseName represents clyde simplified chinese name.
	clydeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "笛笛"}

	// clydeSpanishName represents clyde spanish name.
	clydeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Moe"}

	// clydeTraditionalChineseName represents clyde traditional chinese name.
	clydeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "笛笛"}
)

var (
	// clydeName represents clyde name.
	clydeName = nook.Languages{
		language.AmericanEnglish:      clydeAmericanEnglishName,
		language.CanadianFrench:       clydeCanadianFrenchName,
		language.Dutch:                clydeDutchName,
		language.French:               clydeFrenchName,
		language.German:               clydeGermanName,
		language.Italian:              clydeItalianName,
		language.Japanese:             clydeJapaneseName,
		language.Korean:               clydeKoreanName,
		language.LatinAmericanSpanish: clydeLatinAmericanSpanishName,
		language.Russian:              clydeRussianName,
		language.SimplifiedChinese:    clydeSimplifiedChineseName,
		language.Spanish:              clydeSpanishName,
		language.TraditionalChinese:   clydeTraditionalChineseName}
)

var (
	// clydeCharacter represents clyde character.
	clydeCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: clydeBirthday,
		Code:     clydeCode,
		Key:      character.Clyde,
		Gender:   gender.Male,
		Name:     clydeName,
		Special:  false}
)

var (
	// clydeAmericanEnglishPhrase represents clyde american english phrase.
	clydeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "clip clawp"}

	// clydeCanadianFrenchPhrase represents clyde canadian french phrase.
	clydeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mouh-ouh"}

	// clydeDutchPhrase represents clyde dutch phrase.
	clydeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "klipklop"}

	// clydeFrenchPhrase represents clyde french phrase.
	clydeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mouh-ouh"}

	// clydeGermanPhrase represents clyde german phrase.
	clydeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schauder"}

	// clydeItalianPhrase represents clyde italian phrase.
	clydeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "clop clop"}

	// clydeJapanesePhrase represents clyde japanese phrase.
	clydeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぴろぴろ"}

	// clydeKoreanPhrase represents clyde korean phrase.
	clydeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "다그닥"}

	// clydeLatinAmericanSpanishPhrase represents clyde latin american spanish phrase.
	clydeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "clip-clop"}

	// clydeRussianPhrase represents clyde russian phrase.
	clydeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "цок-цок"}

	// clydeSimplifiedChinesePhrase represents clyde simplified chinese phrase.
	clydeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吹卷卷"}

	// clydeSpanishPhrase represents clyde spanish phrase.
	clydeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "troteras"}

	// clydeTraditionalChinesePhrase represents clyde traditional chinese phrase.
	clydeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吹捲捲"}
)

var (
	// clydePhrase represents clyde phrase.
	clydePhrase = nook.Languages{
		language.AmericanEnglish:      clydeAmericanEnglishPhrase,
		language.CanadianFrench:       clydeCanadianFrenchPhrase,
		language.Dutch:                clydeDutchPhrase,
		language.French:               clydeFrenchPhrase,
		language.German:               clydeGermanPhrase,
		language.Italian:              clydeItalianPhrase,
		language.Japanese:             clydeJapanesePhrase,
		language.Korean:               clydeKoreanPhrase,
		language.LatinAmericanSpanish: clydeLatinAmericanSpanishPhrase,
		language.Russian:              clydeRussianPhrase,
		language.SimplifiedChinese:    clydeSimplifiedChinesePhrase,
		language.Spanish:              clydeSpanishPhrase,
		language.TraditionalChinese:   clydeTraditionalChinesePhrase}
)

var (
	// Clyde represents clyde.
	Clyde = nook.Villager{
		Character:   clydeCharacter,
		Personality: personality.Lazy,
		Phrase:      clydePhrase}
)

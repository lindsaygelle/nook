package cat

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
	// kabukiBirthday represents kabuki birthday.
	kabukiBirthday = nook.Birthday{
		Day:   29,
		Month: time.November}
)

var (
	// kabukiCode represents kabuki code.
	kabukiCode = nook.Code{
		Value: "cat09"}
)

var (
	// kabukiAmericanEnglishName represents kabuki american english name.
	kabukiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kabuki"}

	// kabukiCanadianFrenchName represents kabuki canadian french name.
	kabukiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kabuki"}

	// kabukiDutchName represents kabuki dutch name.
	kabukiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kabuki"}

	// kabukiFrenchName represents kabuki french name.
	kabukiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kabuki"}

	// kabukiGermanName represents kabuki german name.
	kabukiGermanName = nook.Name{
		Language: language.German,
		Value:    "Kabuki"}

	// kabukiItalianName represents kabuki italian name.
	kabukiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kabuki"}

	// kabukiJapaneseName represents kabuki japanese name.
	kabukiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "かぶきち"}

	// kabukiKoreanName represents kabuki korean name.
	kabukiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "가북희"}

	// kabukiLatinAmericanSpanishName represents kabuki latin american spanish name.
	kabukiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kabuki"}

	// kabukiRussianName represents kabuki russian name.
	kabukiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кабуки"}

	// kabukiSimplifiedChineseName represents kabuki simplified chinese name.
	kabukiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "戈伍纪"}

	// kabukiSpanishName represents kabuki spanish name.
	kabukiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kabuki"}

	// kabukiTraditionalChineseName represents kabuki traditional chinese name.
	kabukiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "戈伍紀"}
)

var (
	// kabukiName represents kabuki name.
	kabukiName = nook.Languages{
		language.AmericanEnglish:      kabukiAmericanEnglishName,
		language.CanadianFrench:       kabukiCanadianFrenchName,
		language.Dutch:                kabukiDutchName,
		language.French:               kabukiFrenchName,
		language.German:               kabukiGermanName,
		language.Italian:              kabukiItalianName,
		language.Japanese:             kabukiJapaneseName,
		language.Korean:               kabukiKoreanName,
		language.LatinAmericanSpanish: kabukiLatinAmericanSpanishName,
		language.Russian:              kabukiRussianName,
		language.SimplifiedChinese:    kabukiSimplifiedChineseName,
		language.Spanish:              kabukiSpanishName,
		language.TraditionalChinese:   kabukiTraditionalChineseName}
)

var (
	// kabukiCharacter represents kabuki character.
	kabukiCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: kabukiBirthday,
		Code:     kabukiCode,
		Key:      character.Kabuki,
		Gender:   gender.Male,
		Name:     kabukiName,
		Special:  false}
)

var (
	// kabukiAmericanEnglishPhrase represents kabuki american english phrase.
	kabukiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "meooo-OH"}

	// kabukiCanadianFrenchPhrase represents kabuki canadian french phrase.
	kabukiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mia-ouh"}

	// kabukiDutchPhrase represents kabuki dutch phrase.
	kabukiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mie-auw"}

	// kabukiFrenchPhrase represents kabuki french phrase.
	kabukiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mia-ouh"}

	// kabukiGermanPhrase represents kabuki german phrase.
	kabukiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "miiiau"}

	// kabukiItalianPhrase represents kabuki italian phrase.
	kabukiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "meooo-OH"}

	// kabukiJapanesePhrase represents kabuki japanese phrase.
	kabukiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぃよぉー"}

	// kabukiKoreanPhrase represents kabuki korean phrase.
	kabukiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꺄"}

	// kabukiLatinAmericanSpanishPhrase represents kabuki latin american spanish phrase.
	kabukiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "imikimono"}

	// kabukiRussianPhrase represents kabuki russian phrase.
	kabukiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мя-я-яу"}

	// kabukiSimplifiedChinesePhrase represents kabuki simplified chinese phrase.
	kabukiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咦唷"}

	// kabukiSpanishPhrase represents kabuki spanish phrase.
	kabukiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "imikimono"}

	// kabukiTraditionalChinesePhrase represents kabuki traditional chinese phrase.
	kabukiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咦唷"}
)

var (
	// kabukiPhrase represents kabuki phrase.
	kabukiPhrase = nook.Languages{
		language.AmericanEnglish:      kabukiAmericanEnglishPhrase,
		language.CanadianFrench:       kabukiCanadianFrenchPhrase,
		language.Dutch:                kabukiDutchPhrase,
		language.French:               kabukiFrenchPhrase,
		language.German:               kabukiGermanPhrase,
		language.Italian:              kabukiItalianPhrase,
		language.Japanese:             kabukiJapanesePhrase,
		language.Korean:               kabukiKoreanPhrase,
		language.LatinAmericanSpanish: kabukiLatinAmericanSpanishPhrase,
		language.Russian:              kabukiRussianPhrase,
		language.SimplifiedChinese:    kabukiSimplifiedChinesePhrase,
		language.Spanish:              kabukiSpanishPhrase,
		language.TraditionalChinese:   kabukiTraditionalChinesePhrase}
)

var (
	// Kabuki represents kabuki.
	Kabuki = nook.Villager{
		Character:   kabukiCharacter,
		Personality: personality.Cranky,
		Phrase:      kabukiPhrase}
)

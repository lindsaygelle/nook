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
	// lilyBirthday represents lily birthday.
	lilyBirthday = nook.Birthday{
		Day:   4,
		Month: time.February}
)

var (
	// lilyCode represents lily code.
	lilyCode = nook.Code{
		Value: "flg00"}
)

var (
	// lilyAmericanEnglishName represents lily american english name.
	lilyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lily"}

	// lilyCanadianFrenchName represents lily canadian french name.
	lilyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Raina"}

	// lilyDutchName represents lily dutch name.
	lilyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lily"}

	// lilyFrenchName represents lily french name.
	lilyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Raina"}

	// lilyGermanName represents lily german name.
	lilyGermanName = nook.Name{
		Language: language.German,
		Value:    "Liliane"}

	// lilyItalianName represents lily italian name.
	lilyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gigliola"}

	// lilyJapaneseName represents lily japanese name.
	lilyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイニー"}

	// lilyKoreanName represents lily korean name.
	lilyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레이니"}

	// lilyLatinAmericanSpanishName represents lily latin american spanish name.
	lilyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lili"}

	// lilyRussianName represents lily russian name.
	lilyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лили"}

	// lilySimplifiedChineseName represents lily simplified chinese name.
	lilySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雷妮"}

	// lilySpanishName represents lily spanish name.
	lilySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lili"}

	// lilyTraditionalChineseName represents lily traditional chinese name.
	lilyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雷妮"}
)

var (
	// lilyName represents lily name.
	lilyName = nook.Languages{
		language.AmericanEnglish:      lilyAmericanEnglishName,
		language.CanadianFrench:       lilyCanadianFrenchName,
		language.Dutch:                lilyDutchName,
		language.French:               lilyFrenchName,
		language.German:               lilyGermanName,
		language.Italian:              lilyItalianName,
		language.Japanese:             lilyJapaneseName,
		language.Korean:               lilyKoreanName,
		language.LatinAmericanSpanish: lilyLatinAmericanSpanishName,
		language.Russian:              lilyRussianName,
		language.SimplifiedChinese:    lilySimplifiedChineseName,
		language.Spanish:              lilySpanishName,
		language.TraditionalChinese:   lilyTraditionalChineseName}
)

var (
	// lilyCharacter represents lily character.
	lilyCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: lilyBirthday,
		Code:     lilyCode,
		Key:      character.Lily,
		Gender:   gender.Female,
		Name:     lilyName,
		Special:  false}
)

var (
	// lilyAmericanEnglishPhrase represents lily american english phrase.
	lilyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zzrrbbit"}

	// lilyCanadianFrenchPhrase represents lily canadian french phrase.
	lilyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zzrrbbitt"}

	// lilyDutchPhrase represents lily dutch phrase.
	lilyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "padje"}

	// lilyFrenchPhrase represents lily french phrase.
	lilyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "watich'"}

	// lilyGermanPhrase represents lily german phrase.
	lilyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kröti"}

	// lilyItalianPhrase represents lily italian phrase.
	lilyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gragragra"}

	// lilyJapanesePhrase represents lily japanese phrase.
	lilyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でちゅ"}

	// lilyKoreanPhrase represents lily korean phrase.
	lilyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그래요"}

	// lilyLatinAmericanSpanishPhrase represents lily latin american spanish phrase.
	lilyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "riiibit"}

	// lilyRussianPhrase represents lily russian phrase.
	lilyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ква"}

	// lilySimplifiedChinesePhrase represents lily simplified chinese phrase.
	lilySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雨雨"}

	// lilySpanishPhrase represents lily spanish phrase.
	lilySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mosquita"}

	// lilyTraditionalChinesePhrase represents lily traditional chinese phrase.
	lilyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雨雨"}
)

var (
	// lilyPhrase represents lily phrase.
	lilyPhrase = nook.Languages{
		language.AmericanEnglish:      lilyAmericanEnglishPhrase,
		language.CanadianFrench:       lilyCanadianFrenchPhrase,
		language.Dutch:                lilyDutchPhrase,
		language.French:               lilyFrenchPhrase,
		language.German:               lilyGermanPhrase,
		language.Italian:              lilyItalianPhrase,
		language.Japanese:             lilyJapanesePhrase,
		language.Korean:               lilyKoreanPhrase,
		language.LatinAmericanSpanish: lilyLatinAmericanSpanishPhrase,
		language.Russian:              lilyRussianPhrase,
		language.SimplifiedChinese:    lilySimplifiedChinesePhrase,
		language.Spanish:              lilySpanishPhrase,
		language.TraditionalChinese:   lilyTraditionalChinesePhrase}
)

var (
	// Lily represents lily.
	Lily = nook.Villager{
		Character:   lilyCharacter,
		Personality: personality.Normal,
		Phrase:      lilyPhrase}
)

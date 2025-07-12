package pig

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
	// ganonBirthday represents ganon birthday.
	ganonBirthday = nook.Birthday{
		Day:   21,
		Month: time.February}
)

var (
	// ganonCode represents ganon code.
	ganonCode = nook.Code{
		Value: "pig18"}
)

var (
	// ganonAmericanEnglishName represents ganon american english name.
	ganonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ganon"}

	// ganonCanadianFrenchName represents ganon canadian french name.
	ganonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ganon"}

	// ganonDutchName represents ganon dutch name.
	ganonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// ganonFrenchName represents ganon french name.
	ganonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ganon"}

	// ganonGermanName represents ganon german name.
	ganonGermanName = nook.Name{
		Language: language.German,
		Value:    "Ganon"}

	// ganonItalianName represents ganon italian name.
	ganonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ganon"}

	// ganonJapaneseName represents ganon japanese name.
	ganonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガノン"}

	// ganonKoreanName represents ganon korean name.
	ganonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "가논"}

	// ganonLatinAmericanSpanishName represents ganon latin american spanish name.
	ganonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ganon"}

	// ganonRussianName represents ganon russian name.
	ganonRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// ganonSimplifiedChineseName represents ganon simplified chinese name.
	ganonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// ganonSpanishName represents ganon spanish name.
	ganonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ganon"}

	// ganonTraditionalChineseName represents ganon traditional chinese name.
	ganonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// ganonName represents ganon name.
	ganonName = nook.Languages{
		language.AmericanEnglish:      ganonAmericanEnglishName,
		language.CanadianFrench:       ganonCanadianFrenchName,
		language.Dutch:                ganonDutchName,
		language.French:               ganonFrenchName,
		language.German:               ganonGermanName,
		language.Italian:              ganonItalianName,
		language.Japanese:             ganonJapaneseName,
		language.Korean:               ganonKoreanName,
		language.LatinAmericanSpanish: ganonLatinAmericanSpanishName,
		language.Russian:              ganonRussianName,
		language.SimplifiedChinese:    ganonSimplifiedChineseName,
		language.Spanish:              ganonSpanishName,
		language.TraditionalChinese:   ganonTraditionalChineseName}
)

var (
	// ganonCharacter represents ganon character.
	ganonCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: ganonBirthday,
		Code:     ganonCode,
		Key:      character.Ganon,
		Gender:   gender.Male,
		Name:     ganonName,
		Special:  false}
)

var (
	// ganonAmericanEnglishPhrase represents ganon american english phrase.
	ganonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "heh heh"}

	// ganonCanadianFrenchPhrase represents ganon canadian french phrase.
	ganonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ganouïrk"}

	// ganonDutchPhrase represents ganon dutch phrase.
	ganonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// ganonFrenchPhrase represents ganon french phrase.
	ganonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ganouïrk"}

	// ganonGermanPhrase represents ganon german phrase.
	ganonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ganoink"}

	// ganonItalianPhrase represents ganon italian phrase.
	ganonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "trisgrunf"}

	// ganonJapanesePhrase represents ganon japanese phrase.
	ganonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "フォース"}

	// ganonKoreanPhrase represents ganon korean phrase.
	ganonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "포스"}

	// ganonLatinAmericanSpanishPhrase represents ganon latin american spanish phrase.
	ganonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ganoink"}

	// ganonRussianPhrase represents ganon russian phrase.
	ganonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// ganonSimplifiedChinesePhrase represents ganon simplified chinese phrase.
	ganonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// ganonSpanishPhrase represents ganon spanish phrase.
	ganonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ganoink"}

	// ganonTraditionalChinesePhrase represents ganon traditional chinese phrase.
	ganonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// ganonPhrase represents ganon phrase.
	ganonPhrase = nook.Languages{
		language.AmericanEnglish:      ganonAmericanEnglishPhrase,
		language.CanadianFrench:       ganonCanadianFrenchPhrase,
		language.Dutch:                ganonDutchPhrase,
		language.French:               ganonFrenchPhrase,
		language.German:               ganonGermanPhrase,
		language.Italian:              ganonItalianPhrase,
		language.Japanese:             ganonJapanesePhrase,
		language.Korean:               ganonKoreanPhrase,
		language.LatinAmericanSpanish: ganonLatinAmericanSpanishPhrase,
		language.Russian:              ganonRussianPhrase,
		language.SimplifiedChinese:    ganonSimplifiedChinesePhrase,
		language.Spanish:              ganonSpanishPhrase,
		language.TraditionalChinese:   ganonTraditionalChinesePhrase}
)

var (
	// Ganon represents ganon.
	Ganon = nook.Villager{
		Character:   ganonCharacter,
		Personality: personality.Cranky,
		Phrase:      ganonPhrase}
)

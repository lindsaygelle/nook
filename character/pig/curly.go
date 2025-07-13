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
	// curlyBirthday represents curly birthday.
	curlyBirthday = nook.Birthday{
		Day:   26,
		Month: time.July}
)

var (
	// curlyCode represents curly code.
	curlyCode = nook.Code{
		Value: "pig00"}
)

var (
	// curlyAmericanEnglishName represents curly american english name.
	curlyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Curly"}

	// curlyCanadianFrenchName represents curly canadian french name.
	curlyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tirbou"}

	// curlyDutchName represents curly dutch name.
	curlyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Curly"}

	// curlyFrenchName represents curly french name.
	curlyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tirbou"}

	// curlyGermanName represents curly german name.
	curlyGermanName = nook.Name{
		Language: language.German,
		Value:    "Oink"}

	// curlyItalianName represents curly italian name.
	curlyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ricciolo"}

	// curlyJapaneseName represents curly japanese name.
	curlyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハムカツ"}

	// curlyKoreanName represents curly korean name.
	curlyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "햄까스"}

	// curlyLatinAmericanSpanishName represents curly latin american spanish name.
	curlyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rufueto"}

	// curlyRussianName represents curly russian name.
	curlyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Керли"}

	// curlySimplifiedChineseName represents curly simplified chinese name.
	curlySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "胜利"}

	// curlySpanishName represents curly spanish name.
	curlySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rufueto"}

	// curlyTraditionalChineseName represents curly traditional chinese name.
	curlyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "勝利"}
)

var (
	// curlyName represents curly name.
	curlyName = nook.Languages{
		language.AmericanEnglish:      curlyAmericanEnglishName,
		language.CanadianFrench:       curlyCanadianFrenchName,
		language.Dutch:                curlyDutchName,
		language.French:               curlyFrenchName,
		language.German:               curlyGermanName,
		language.Italian:              curlyItalianName,
		language.Japanese:             curlyJapaneseName,
		language.Korean:               curlyKoreanName,
		language.LatinAmericanSpanish: curlyLatinAmericanSpanishName,
		language.Russian:              curlyRussianName,
		language.SimplifiedChinese:    curlySimplifiedChineseName,
		language.Spanish:              curlySpanishName,
		language.TraditionalChinese:   curlyTraditionalChineseName}
)

var (
	// curlyCharacter represents curly character.
	curlyCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: curlyBirthday,
		Code:     curlyCode,
		Key:      character.Curly,
		Gender:   gender.Male,
		Name:     curlyName,
		Special:  false}
)

var (
	// curlyAmericanEnglishPhrase represents curly american english phrase.
	curlyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nyoink"}

	// curlyCanadianFrenchPhrase represents curly canadian french phrase.
	curlyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coink"}

	// curlyDutchPhrase represents curly dutch phrase.
	curlyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hupsaknor"}

	// curlyFrenchPhrase represents curly french phrase.
	curlyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coink"}

	// curlyGermanPhrase represents curly german phrase.
	curlyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "oinkoink"}

	// curlyItalianPhrase represents curly italian phrase.
	curlyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "niooink"}

	// curlyJapanesePhrase represents curly japanese phrase.
	curlyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どもども"}

	// curlyKoreanPhrase represents curly korean phrase.
	curlyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꿀꿀"}

	// curlyLatinAmericanSpanishPhrase represents curly latin american spanish phrase.
	curlyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñoink"}

	// curlyRussianPhrase represents curly russian phrase.
	curlyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хряк-бряк"}

	// curlySimplifiedChinesePhrase represents curly simplified chinese phrase.
	curlySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "感谢"}

	// curlySpanishPhrase represents curly spanish phrase.
	curlySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñoink"}

	// curlyTraditionalChinesePhrase represents curly traditional chinese phrase.
	curlyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "感謝"}
)

var (
	// curlyPhrase represents curly phrase.
	curlyPhrase = nook.Languages{
		language.AmericanEnglish:      curlyAmericanEnglishPhrase,
		language.CanadianFrench:       curlyCanadianFrenchPhrase,
		language.Dutch:                curlyDutchPhrase,
		language.French:               curlyFrenchPhrase,
		language.German:               curlyGermanPhrase,
		language.Italian:              curlyItalianPhrase,
		language.Japanese:             curlyJapanesePhrase,
		language.Korean:               curlyKoreanPhrase,
		language.LatinAmericanSpanish: curlyLatinAmericanSpanishPhrase,
		language.Russian:              curlyRussianPhrase,
		language.SimplifiedChinese:    curlySimplifiedChinesePhrase,
		language.Spanish:              curlySpanishPhrase,
		language.TraditionalChinese:   curlyTraditionalChinesePhrase}
)

var (
	// Curly represents curly.
	Curly = nook.Villager{
		Character:   curlyCharacter,
		Personality: personality.Jock,
		Phrase:      curlyPhrase}
)

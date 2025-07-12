package bearcub

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
	// murphyBirthday represents murphy birthday.
	murphyBirthday = nook.Birthday{
		Day:   29,
		Month: time.December}
)

var (
	// murphyCode represents murphy code.
	murphyCode = nook.Code{
		Value: "cbr07"}
)

var (
	// murphyAmericanEnglishName represents murphy american english name.
	murphyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Murphy"}

	// murphyCanadianFrenchName represents murphy canadian french name.
	murphyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Eddie"}

	// murphyDutchName represents murphy dutch name.
	murphyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Murphy"}

	// murphyFrenchName represents murphy french name.
	murphyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Eddie"}

	// murphyGermanName represents murphy german name.
	murphyGermanName = nook.Name{
		Language: language.German,
		Value:    "Michael"}

	// murphyItalianName represents murphy italian name.
	murphyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marcello"}

	// murphyJapaneseName represents murphy japanese name.
	murphyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のりお"}

	// murphyKoreanName represents murphy korean name.
	murphyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "머피"}

	// murphyLatinAmericanSpanishName represents murphy latin american spanish name.
	murphyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Orsón"}

	// murphyRussianName represents murphy russian name.
	murphyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мерфи"}

	// murphySimplifiedChineseName represents murphy simplified chinese name.
	murphySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宪雄"}

	// murphySpanishName represents murphy spanish name.
	murphySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Orsón"}

	// murphyTraditionalChineseName represents murphy traditional chinese name.
	murphyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "憲雄"}
)

var (
	// murphyName represents murphy name.
	murphyName = nook.Languages{
		language.AmericanEnglish:      murphyAmericanEnglishName,
		language.CanadianFrench:       murphyCanadianFrenchName,
		language.Dutch:                murphyDutchName,
		language.French:               murphyFrenchName,
		language.German:               murphyGermanName,
		language.Italian:              murphyItalianName,
		language.Japanese:             murphyJapaneseName,
		language.Korean:               murphyKoreanName,
		language.LatinAmericanSpanish: murphyLatinAmericanSpanishName,
		language.Russian:              murphyRussianName,
		language.SimplifiedChinese:    murphySimplifiedChineseName,
		language.Spanish:              murphySpanishName,
		language.TraditionalChinese:   murphyTraditionalChineseName}
)

var (
	// murphyCharacter represents murphy character.
	murphyCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: murphyBirthday,
		Code:     murphyCode,
		Key:      character.Murphy,
		Gender:   gender.Male,
		Name:     murphyName,
		Special:  false}
)

var (
	// murphyAmericanEnglishPhrase represents murphy american english phrase.
	murphyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "laddie"}

	// murphyCanadianFrenchPhrase represents murphy canadian french phrase.
	murphyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "fait que"}

	// murphyDutchPhrase represents murphy dutch phrase.
	murphyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "spruit"}

	// murphyFrenchPhrase represents murphy french phrase.
	murphyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "canaille"}

	// murphyGermanPhrase represents murphy german phrase.
	murphyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "liebelein"}

	// murphyItalianPhrase represents murphy italian phrase.
	murphyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grizzel"}

	// murphyJapanesePhrase represents murphy japanese phrase.
	murphyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですばい"}

	// murphyKoreanPhrase represents murphy korean phrase.
	murphyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "샐리"}

	// murphyLatinAmericanSpanishPhrase represents murphy latin american spanish phrase.
	murphyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "heee"}

	// murphyRussianPhrase represents murphy russian phrase.
	murphyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "товарищ"}

	// murphySimplifiedChinesePhrase represents murphy simplified chinese phrase.
	murphySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罢了"}

	// murphySpanishPhrase represents murphy spanish phrase.
	murphySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pelusita"}

	// murphyTraditionalChinesePhrase represents murphy traditional chinese phrase.
	murphyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "罷了"}
)

var (
	// murphyPhrase represents murphy phrase.
	murphyPhrase = nook.Languages{
		language.AmericanEnglish:      murphyAmericanEnglishPhrase,
		language.CanadianFrench:       murphyCanadianFrenchPhrase,
		language.Dutch:                murphyDutchPhrase,
		language.French:               murphyFrenchPhrase,
		language.German:               murphyGermanPhrase,
		language.Italian:              murphyItalianPhrase,
		language.Japanese:             murphyJapanesePhrase,
		language.Korean:               murphyKoreanPhrase,
		language.LatinAmericanSpanish: murphyLatinAmericanSpanishPhrase,
		language.Russian:              murphyRussianPhrase,
		language.SimplifiedChinese:    murphySimplifiedChinesePhrase,
		language.Spanish:              murphySpanishPhrase,
		language.TraditionalChinese:   murphyTraditionalChinesePhrase}
)

var (
	// Murphy represents murphy.
	Murphy = nook.Villager{
		Character:   murphyCharacter,
		Personality: personality.Cranky,
		Phrase:      murphyPhrase}
)

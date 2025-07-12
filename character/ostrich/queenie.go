package ostrich

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
	// queenieBirthday represents queenie birthday.
	queenieBirthday = nook.Birthday{
		Day:   13,
		Month: time.November}
)

var (
	// queenieCode represents queenie code.
	queenieCode = nook.Code{
		Value: "ost00"}
)

var (
	// queenieAmericanEnglishName represents queenie american english name.
	queenieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Queenie"}

	// queenieCanadianFrenchName represents queenie canadian french name.
	queenieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Reine"}

	// queenieDutchName represents queenie dutch name.
	queenieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Queenie"}

	// queenieFrenchName represents queenie french name.
	queenieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Reine"}

	// queenieGermanName represents queenie german name.
	queenieGermanName = nook.Name{
		Language: language.German,
		Value:    "Isabella"}

	// queenieItalianName represents queenie italian name.
	queenieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Regina"}

	// queenieJapaneseName represents queenie japanese name.
	queenieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タキュ"}

	// queenieKoreanName represents queenie korean name.
	queenieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "택주"}

	// queenieLatinAmericanSpanishName represents queenie latin american spanish name.
	queenieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marujita"}

	// queenieRussianName represents queenie russian name.
	queenieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Квини"}

	// queenieSimplifiedChineseName represents queenie simplified chinese name.
	queenieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "果果"}

	// queenieSpanishName represents queenie spanish name.
	queenieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marujita"}

	// queenieTraditionalChineseName represents queenie traditional chinese name.
	queenieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "果果"}
)

var (
	// queenieName represents queenie name.
	queenieName = nook.Languages{
		language.AmericanEnglish:      queenieAmericanEnglishName,
		language.CanadianFrench:       queenieCanadianFrenchName,
		language.Dutch:                queenieDutchName,
		language.French:               queenieFrenchName,
		language.German:               queenieGermanName,
		language.Italian:              queenieItalianName,
		language.Japanese:             queenieJapaneseName,
		language.Korean:               queenieKoreanName,
		language.LatinAmericanSpanish: queenieLatinAmericanSpanishName,
		language.Russian:              queenieRussianName,
		language.SimplifiedChinese:    queenieSimplifiedChineseName,
		language.Spanish:              queenieSpanishName,
		language.TraditionalChinese:   queenieTraditionalChineseName}
)

var (
	// queenieCharacter represents queenie character.
	queenieCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: queenieBirthday,
		Code:     queenieCode,
		Key:      character.Queenie,
		Gender:   gender.Female,
		Name:     queenieName,
		Special:  false}
)

var (
	// queenieAmericanEnglishPhrase represents queenie american english phrase.
	queenieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chicken"}

	// queenieCanadianFrenchPhrase represents queenie canadian french phrase.
	queenieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "poupoule"}

	// queenieDutchPhrase represents queenie dutch phrase.
	queenieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kippie"}

	// queenieFrenchPhrase represents queenie french phrase.
	queenieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poupoule"}

	// queenieGermanPhrase represents queenie german phrase.
	queenieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hühnchen"}

	// queenieItalianPhrase represents queenie italian phrase.
	queenieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "polpetta"}

	// queenieJapanesePhrase represents queenie japanese phrase.
	queenieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やっぱし"}

	// queenieKoreanPhrase represents queenie korean phrase.
	queenieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "역시"}

	// queenieLatinAmericanSpanishPhrase represents queenie latin american spanish phrase.
	queenieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "plumifá"}

	// queenieRussianPhrase represents queenie russian phrase.
	queenieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "цыпленок"}

	// queenieSimplifiedChinesePhrase represents queenie simplified chinese phrase.
	queenieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "果然"}

	// queenieSpanishPhrase represents queenie spanish phrase.
	queenieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gallina"}

	// queenieTraditionalChinesePhrase represents queenie traditional chinese phrase.
	queenieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "果然"}
)

var (
	// queeniePhrase represents queenie phrase.
	queeniePhrase = nook.Languages{
		language.AmericanEnglish:      queenieAmericanEnglishPhrase,
		language.CanadianFrench:       queenieCanadianFrenchPhrase,
		language.Dutch:                queenieDutchPhrase,
		language.French:               queenieFrenchPhrase,
		language.German:               queenieGermanPhrase,
		language.Italian:              queenieItalianPhrase,
		language.Japanese:             queenieJapanesePhrase,
		language.Korean:               queenieKoreanPhrase,
		language.LatinAmericanSpanish: queenieLatinAmericanSpanishPhrase,
		language.Russian:              queenieRussianPhrase,
		language.SimplifiedChinese:    queenieSimplifiedChinesePhrase,
		language.Spanish:              queenieSpanishPhrase,
		language.TraditionalChinese:   queenieTraditionalChinesePhrase}
)

var (
	// Queenie represents queenie.
	Queenie = nook.Villager{
		Character:   queenieCharacter,
		Personality: personality.Snooty,
		Phrase:      queeniePhrase}
)

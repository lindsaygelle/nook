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
	// henryBirthday represents henry birthday.
	henryBirthday = nook.Birthday{
		Day:   21,
		Month: time.September}
)

var (
	// henryCode represents henry code.
	henryCode = nook.Code{
		Value: "flg19"}
)

var (
	// henryAmericanEnglishName represents henry american english name.
	henryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Henry"}

	// henryCanadianFrenchName represents henry canadian french name.
	henryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Henri"}

	// henryDutchName represents henry dutch name.
	henryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Henry"}

	// henryFrenchName represents henry french name.
	henryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Henri"}

	// henryGermanName represents henry german name.
	henryGermanName = nook.Name{
		Language: language.German,
		Value:    "Freddy"}

	// henryItalianName represents henry italian name.
	henryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Renato"}

	// henryJapaneseName represents henry japanese name.
	henryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヘンリー"}

	// henryKoreanName represents henry korean name.
	henryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "헨리"}

	// henryLatinAmericanSpanishName represents henry latin american spanish name.
	henryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Saperto"}

	// henryRussianName represents henry russian name.
	henryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Генри"}

	// henrySimplifiedChineseName represents henry simplified chinese name.
	henrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亨利"}

	// henrySpanishName represents henry spanish name.
	henrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Saperto"}

	// henryTraditionalChineseName represents henry traditional chinese name.
	henryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亨利"}
)

var (
	// henryName represents henry name.
	henryName = nook.Languages{
		language.AmericanEnglish:      henryAmericanEnglishName,
		language.CanadianFrench:       henryCanadianFrenchName,
		language.Dutch:                henryDutchName,
		language.French:               henryFrenchName,
		language.German:               henryGermanName,
		language.Italian:              henryItalianName,
		language.Japanese:             henryJapaneseName,
		language.Korean:               henryKoreanName,
		language.LatinAmericanSpanish: henryLatinAmericanSpanishName,
		language.Russian:              henryRussianName,
		language.SimplifiedChinese:    henrySimplifiedChineseName,
		language.Spanish:              henrySpanishName,
		language.TraditionalChinese:   henryTraditionalChineseName}
)

var (
	// henryCharacter represents henry character.
	henryCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: henryBirthday,
		Code:     henryCode,
		Key:      character.Henry,
		Gender:   gender.Male,
		Name:     henryName,
		Special:  false}
)

var (
	// henryAmericanEnglishPhrase represents henry american english phrase.
	henryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snoozit"}

	// henryCanadianFrenchPhrase represents henry canadian french phrase.
	henryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kôwa"}

	// henryDutchPhrase represents henry dutch phrase.
	henryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snurkwaak"}

	// henryFrenchPhrase represents henry french phrase.
	henryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pustule"}

	// henryGermanPhrase represents henry german phrase.
	henryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quappi"}

	// henryItalianPhrase represents henry italian phrase.
	henryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "croac"}

	// henryJapanesePhrase represents henry japanese phrase.
	henryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "むにゃ"}

	// henryKoreanPhrase represents henry korean phrase.
	henryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "음냐음냐"}

	// henryLatinAmericanSpanishPhrase represents henry latin american spanish phrase.
	henryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cruacruá"}

	// henryRussianPhrase represents henry russian phrase.
	henryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ква-храп"}

	// henrySimplifiedChinesePhrase represents henry simplified chinese phrase.
	henrySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "滴咕"}

	// henrySpanishPhrase represents henry spanish phrase.
	henrySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cruacruá"}

	// henryTraditionalChinesePhrase represents henry traditional chinese phrase.
	henryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "滴咕"}
)

var (
	// henryPhrase represents henry phrase.
	henryPhrase = nook.Languages{
		language.AmericanEnglish:      henryAmericanEnglishPhrase,
		language.CanadianFrench:       henryCanadianFrenchPhrase,
		language.Dutch:                henryDutchPhrase,
		language.French:               henryFrenchPhrase,
		language.German:               henryGermanPhrase,
		language.Italian:              henryItalianPhrase,
		language.Japanese:             henryJapanesePhrase,
		language.Korean:               henryKoreanPhrase,
		language.LatinAmericanSpanish: henryLatinAmericanSpanishPhrase,
		language.Russian:              henryRussianPhrase,
		language.SimplifiedChinese:    henrySimplifiedChinesePhrase,
		language.Spanish:              henrySpanishPhrase,
		language.TraditionalChinese:   henryTraditionalChinesePhrase}
)

var (
	// Henry represents henry.
	Henry = nook.Villager{
		Character:   henryCharacter,
		Personality: personality.Smug,
		Phrase:      henryPhrase}
)

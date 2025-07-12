package mouse

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
	// candiBirthday represents candi birthday.
	candiBirthday = nook.Birthday{
		Day:   13,
		Month: time.April}
)

var (
	// candiCode represents candi code.
	candiCode = nook.Code{
		Value: "mus08"}
)

var (
	// candiAmericanEnglishName represents candi american english name.
	candiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Candi"}

	// candiCanadianFrenchName represents candi canadian french name.
	candiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sucrette"}

	// candiDutchName represents candi dutch name.
	candiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Candi"}

	// candiFrenchName represents candi french name.
	candiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sucrette"}

	// candiGermanName represents candi german name.
	candiGermanName = nook.Name{
		Language: language.German,
		Value:    "Renate"}

	// candiItalianName represents candi italian name.
	candiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mella"}

	// candiJapaneseName represents candi japanese name.
	candiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "かんゆ"}

	// candiKoreanName represents candi korean name.
	candiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사탕"}

	// candiLatinAmericanSpanishName represents candi latin american spanish name.
	candiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chuchi"}

	// candiRussianName represents candi russian name.
	candiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэнди"}

	// candiSimplifiedChineseName represents candi simplified chinese name.
	candiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "甘油"}

	// candiSpanishName represents candi spanish name.
	candiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chuchi"}

	// candiTraditionalChineseName represents candi traditional chinese name.
	candiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "甘油"}
)

var (
	// candiName represents candi name.
	candiName = nook.Languages{
		language.AmericanEnglish:      candiAmericanEnglishName,
		language.CanadianFrench:       candiCanadianFrenchName,
		language.Dutch:                candiDutchName,
		language.French:               candiFrenchName,
		language.German:               candiGermanName,
		language.Italian:              candiItalianName,
		language.Japanese:             candiJapaneseName,
		language.Korean:               candiKoreanName,
		language.LatinAmericanSpanish: candiLatinAmericanSpanishName,
		language.Russian:              candiRussianName,
		language.SimplifiedChinese:    candiSimplifiedChineseName,
		language.Spanish:              candiSpanishName,
		language.TraditionalChinese:   candiTraditionalChineseName}
)

var (
	// candiCharacter represents candi character.
	candiCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: candiBirthday,
		Code:     candiCode,
		Key:      character.Candi,
		Gender:   gender.Female,
		Name:     candiName,
		Special:  false}
)

var (
	// candiAmericanEnglishPhrase represents candi american english phrase.
	candiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sweetie"}

	// candiCanadianFrenchPhrase represents candi canadian french phrase.
	candiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "claquos"}

	// candiDutchPhrase represents candi dutch phrase.
	candiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zoetje"}

	// candiFrenchPhrase represents candi french phrase.
	candiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "claquos"}

	// candiGermanPhrase represents candi german phrase.
	candiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schatzi"}

	// candiItalianPhrase represents candi italian phrase.
	candiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squitto"}

	// candiJapanesePhrase represents candi japanese phrase.
	candiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですワ"}

	// candiKoreanPhrase represents candi korean phrase.
	candiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "달짝지근"}

	// candiLatinAmericanSpanishPhrase represents candi latin american spanish phrase.
	candiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "escuic"}

	// candiRussianPhrase represents candi russian phrase.
	candiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сладость"}

	// candiSimplifiedChinesePhrase represents candi simplified chinese phrase.
	candiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "滑滑"}

	// candiSpanishPhrase represents candi spanish phrase.
	candiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "yogurín"}

	// candiTraditionalChinesePhrase represents candi traditional chinese phrase.
	candiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "滑滑"}
)

var (
	// candiPhrase represents candi phrase.
	candiPhrase = nook.Languages{
		language.AmericanEnglish:      candiAmericanEnglishPhrase,
		language.CanadianFrench:       candiCanadianFrenchPhrase,
		language.Dutch:                candiDutchPhrase,
		language.French:               candiFrenchPhrase,
		language.German:               candiGermanPhrase,
		language.Italian:              candiItalianPhrase,
		language.Japanese:             candiJapanesePhrase,
		language.Korean:               candiKoreanPhrase,
		language.LatinAmericanSpanish: candiLatinAmericanSpanishPhrase,
		language.Russian:              candiRussianPhrase,
		language.SimplifiedChinese:    candiSimplifiedChinesePhrase,
		language.Spanish:              candiSpanishPhrase,
		language.TraditionalChinese:   candiTraditionalChinesePhrase}
)

var (
	// Candi represents candi.
	Candi = nook.Villager{
		Character:   candiCharacter,
		Personality: personality.Peppy,
		Phrase:      candiPhrase}
)

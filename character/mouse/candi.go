package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	candiBirthday = nook.Birthday{
		Day:   13,
		Month: time.April}
)

var (
	candiCode = nook.Code{
		Value: "mus08"}
)

var (
	candiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Candi"}

	candiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sucrette"}

	candiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Candi"}

	candiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sucrette"}

	candiGermanName = nook.Name{
		Language: language.German,
		Value:    "Renate"}

	candiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mella"}

	candiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "かんゆ"}

	candiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사탕"}

	candiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chuchi"}

	candiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэнди"}

	candiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "甘油"}

	candiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chuchi"}

	candiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "甘油"}
)

var (
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
	candiCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: candiBirthday,
		Code:     candiCode,
		Gender:   gender.Female,
		Name:     candiName}
)

var (
	candiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sweetie"}

	candiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "claquos"}

	candiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zoetje"}

	candiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "claquos"}

	candiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schatzi"}

	candiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squitto"}

	candiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですワ"}

	candiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "달짝지근"}

	candiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "escuic"}

	candiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "сладость"}

	candiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "滑滑"}

	candiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "yogurín"}

	candiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "滑滑"}
)

var (
	candiPhrase = nook.Languages{
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
	Candi = nook.Villager{
		Character:   candiCharacter,
		Personality: personality.Peppy,
		Phrase:      candiPhrase}
)

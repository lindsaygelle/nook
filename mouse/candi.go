package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Sucretteclaquos"}

	candiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Candizoetje"}

	candiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sucretteclaquos"}

	candiGermanName = nook.Name{
		Language: language.German,
		Value:    "Renateschatzi"}

	candiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Mellasquitto"}

	candiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "かんゆですワ"}

	candiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사탕달짝지근"}

	candiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chuchiescuic"}

	candiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэндисладость"}

	candiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "甘油滑滑"}

	candiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chuchiyogurín"}

	candiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "甘油滑滑"}
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
		Animal:   Mouse,
		Birthday: candiBirthday,
		Code:     candiCode,
		Gender:   nook.Female,
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
		Personality: nook.Peppy,
		Phrase:      candiPhrase}
)

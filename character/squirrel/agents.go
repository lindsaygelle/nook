package squirrel

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
	// agentSBirthday represents Agent S' birthday.
	agentSBirthday = nook.Birthday{
		Day:   2,
		Month: time.July}
)

var (
	// agentSCode represents Agent S' unique code.
	agentSCode = nook.Code{
		Value: "squ05"}
)

var (
	// agentSAmericanEnglishName represents Agent S' name in American English.
	agentSAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Agent S"}

	// agentSCanadianFrenchName represents Agent S' name in Canadian French.
	agentSCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ninjette"}

	// agentSDutchName represents Agent S' name in Dutch.
	agentSDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Agent S"}

	// agentSFrenchName represents Agent S' name in French.
	agentSFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ninjette"}

	// agentSGermanName represents Agent S' name in German.
	agentSGermanName = nook.Name{
		Language: language.German,
		Value:    "Karin"}

	// agentSItalianName represents Agent S' name in Italian.
	agentSItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Agente S"}

	// agentSJapaneseName represents Agent S' name in Japanese.
	agentSJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "２ごう"}

	// agentSKoreanName represents Agent S' name in Korean.
	agentSKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "2호"}

	// agentSLatinAmericanSpanishName represents Agent S' name in Latin American Spanish.
	agentSLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fardilia"}

	// agentSRussianName represents Agent S' name in Russian.
	agentSRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Агент С"}

	// agentSSimplifiedChineseName represents Agent S' name in Simplified Chinese.
	agentSSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿二"}

	// agentSSpanishName represents Agent S' name in Spanish.
	agentSSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fardilia"}

	// agentSTraditionalChineseName represents Agent S' name in Traditional Chinese.
	agentSTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿二"}
)

var (
	// agentSName represents Agent S' name in different languages.
	agentSName = nook.Languages{
		language.AmericanEnglish:      agentSAmericanEnglishName,
		language.CanadianFrench:       agentSCanadianFrenchName,
		language.Dutch:                agentSDutchName,
		language.French:               agentSFrenchName,
		language.German:               agentSGermanName,
		language.Italian:              agentSItalianName,
		language.Japanese:             agentSJapaneseName,
		language.Korean:               agentSKoreanName,
		language.LatinAmericanSpanish: agentSLatinAmericanSpanishName,
		language.Russian:              agentSRussianName,
		language.SimplifiedChinese:    agentSSimplifiedChineseName,
		language.Spanish:              agentSSpanishName,
		language.TraditionalChinese:   agentSTraditionalChineseName}
)

var (
	// agentSCharacter represents Agent S' character information.
	agentSCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: agentSBirthday,
		Code:     agentSCode,
		Key:      character.AgentS,
		Gender:   gender.Female,
		Name:     agentSName,
		Special:  false}
)

var (
	// agentSAmericanEnglishPhrase represents Agent S' phrase in American English.
	agentSAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sidekick"}

	// agentSCanadianFrenchPhrase represents Agent S' phrase in Canadian French.
	agentSCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "haï-ya"}

	// agentSDutchPhrase represents Agent S' phrase in Dutch.
	agentSDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sidekick"}

	// agentSFrenchPhrase represents Agent S' phrase in French.
	agentSFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poto"}

	// agentSGermanPhrase represents Agent S' phrase in German.
	agentSGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "strampl"}

	// agentSItalianPhrase represents Agent S' phrase in Italian.
	agentSItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "spalla"}

	// agentSJapanesePhrase represents Agent S' phrase in Japanese.
	agentSJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "せいやっ"}

	// agentSKoreanPhrase represents Agent S' phrase in Korean.
	agentSKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흠냐"}

	// agentSLatinAmericanSpanishPhrase represents Agent S' phrase in Latin American Spanish.
	agentSLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hai-ááá"}

	// agentSRussianPhrase represents Agent S' phrase in Russian.
	agentSRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "коллега"}

	// agentSSimplifiedChinesePhrase represents Agent S' phrase in Simplified Chinese.
	agentSSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘿唷"}

	// agentSSpanishPhrase represents Agent S' phrase in Spanish.
	agentSSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piñita"}

	// agentSTraditionalChinesePhrase represents Agent S' phrase in Traditional Chinese.
	agentSTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘿唷"}
)

var (
	// agentSPhrase represents Agent S' phrases in different languages.
	agentSPhrase = nook.Languages{
		language.AmericanEnglish:      agentSAmericanEnglishPhrase,
		language.CanadianFrench:       agentSCanadianFrenchPhrase,
		language.Dutch:                agentSDutchPhrase,
		language.French:               agentSFrenchPhrase,
		language.German:               agentSGermanPhrase,
		language.Italian:              agentSItalianPhrase,
		language.Japanese:             agentSJapanesePhrase,
		language.Korean:               agentSKoreanPhrase,
		language.LatinAmericanSpanish: agentSLatinAmericanSpanishPhrase,
		language.Russian:              agentSRussianPhrase,
		language.SimplifiedChinese:    agentSSimplifiedChinesePhrase,
		language.Spanish:              agentSSpanishPhrase,
		language.TraditionalChinese:   agentSTraditionalChinesePhrase}
)

var (
	// AgentS represents the character Agent S with her complete information.
	AgentS = nook.Villager{
		Character:   agentSCharacter,
		Personality: personality.Peppy,
		Phrase:      agentSPhrase}
)

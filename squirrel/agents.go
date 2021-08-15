package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	agentsBirthday = nook.Birthday{
		Day:   2,
		Month: time.July}
)

var (
	agentsCode = nook.Code{
		Value: "squ05"}
)

var (
	agentsAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Agent S"}

	agentsCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ninjette"}

	agentsDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Agent S"}

	agentsFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ninjette"}

	agentsGermanName = nook.Name{
		Language: language.German,
		Value:    "Karin"}

	agentsItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Agente S"}

	agentsJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "２ごう"}

	agentsKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "2호"}

	agentsLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fardilia"}

	agentsRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Агент С"}

	agentsSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿二"}

	agentsSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fardilia"}

	agentsTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿二"}
)

var (
	agentsName = nook.Languages{
		language.AmericanEnglish:      agentsAmericanEnglishName,
		language.CanadianFrench:       agentsCanadianFrenchName,
		language.Dutch:                agentsDutchName,
		language.French:               agentsFrenchName,
		language.German:               agentsGermanName,
		language.Italian:              agentsItalianName,
		language.Japanese:             agentsJapaneseName,
		language.Korean:               agentsKoreanName,
		language.LatinAmericanSpanish: agentsLatinAmericanSpanishName,
		language.Russian:              agentsRussianName,
		language.SimplifiedChinese:    agentsSimplifiedChineseName,
		language.Spanish:              agentsSpanishName,
		language.TraditionalChinese:   agentsTraditionalChineseName}
)

var (
	agentsCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: agentsBirthday,
		Code:     agentsCode,
		Gender:   nook.Female,
		Name:     agentsName}
)

var (
	agentsAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sidekick"}

	agentsCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "haï-ya"}

	agentsDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "sidekick"}

	agentsFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poto"}

	agentsGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "strampl"}

	agentsItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "spalla"}

	agentsJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "せいやっ"}

	agentsKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흠냐"}

	agentsLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hai-ááá"}

	agentsRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "коллега"}

	agentsSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘿唷"}

	agentsSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piñita"}

	agentsTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘿唷"}
)

var (
	agentsPhrase = nook.Languages{
		language.AmericanEnglish:      agentsAmericanEnglishName,
		language.CanadianFrench:       agentsCanadianFrenchName,
		language.Dutch:                agentsDutchName,
		language.French:               agentsFrenchName,
		language.German:               agentsGermanName,
		language.Italian:              agentsItalianName,
		language.Japanese:             agentsJapaneseName,
		language.Korean:               agentsKoreanName,
		language.LatinAmericanSpanish: agentsLatinAmericanSpanishName,
		language.Russian:              agentsRussianName,
		language.SimplifiedChinese:    agentsSimplifiedChineseName,
		language.Spanish:              agentsSpanishName,
		language.TraditionalChinese:   agentsTraditionalChineseName}
)

var (
	AgentS = nook.Villager{
		Character:   agentsCharacter,
		Personality: nook.Peppy,
		Phrase:      agentsPhrase}
)

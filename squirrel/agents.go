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
		Value:    "Ninjettehaï-ya"}

	agentsDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Agent Ssidekick"}

	agentsFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ninjettepoto"}

	agentsGermanName = nook.Name{
		Language: language.German,
		Value:    "Karinstrampl"}

	agentsItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Agente Sspalla"}

	agentsJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "２ごうせいやっ"}

	agentsKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "2호흠냐"}

	agentsLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fardiliahai-ááá"}

	agentsRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Агент Сколлега"}

	agentsSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿二嘿唷"}

	agentsSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fardiliapiñita"}

	agentsTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿二嘿唷"}
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

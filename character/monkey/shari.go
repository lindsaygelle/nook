package monkey

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
	// shariBirthday represents shari birthday.
	shariBirthday = nook.Birthday{
		Day:   10,
		Month: time.April}
)

var (
	// shariCode represents shari code.
	shariCode = nook.Code{
		Value: "mnk07"}
)

var (
	// shariAmericanEnglishName represents shari american english name.
	shariAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Shari"}

	// shariCanadianFrenchName represents shari canadian french name.
	shariCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Luna"}

	// shariDutchName represents shari dutch name.
	shariDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Shari"}

	// shariFrenchName represents shari french name.
	shariFrenchName = nook.Name{
		Language: language.French,
		Value:    "Luna"}

	// shariGermanName represents shari german name.
	shariGermanName = nook.Name{
		Language: language.German,
		Value:    "Uta"}

	// shariItalianName represents shari italian name.
	shariItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ninì"}

	// shariJapaneseName represents shari japanese name.
	shariJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シェリー"}

	// shariKoreanName represents shari korean name.
	shariKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "젤리"}

	// shariLatinAmericanSpanishName represents shari latin american spanish name.
	shariLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Montse"}

	// shariRussianName represents shari russian name.
	shariRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шари"}

	// shariSimplifiedChineseName represents shari simplified chinese name.
	shariSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪莉"}

	// shariSpanishName represents shari spanish name.
	shariSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Montse"}

	// shariTraditionalChineseName represents shari traditional chinese name.
	shariTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪莉"}
)

var (
	// shariName represents shari name.
	shariName = nook.Languages{
		language.AmericanEnglish:      shariAmericanEnglishName,
		language.CanadianFrench:       shariCanadianFrenchName,
		language.Dutch:                shariDutchName,
		language.French:               shariFrenchName,
		language.German:               shariGermanName,
		language.Italian:              shariItalianName,
		language.Japanese:             shariJapaneseName,
		language.Korean:               shariKoreanName,
		language.LatinAmericanSpanish: shariLatinAmericanSpanishName,
		language.Russian:              shariRussianName,
		language.SimplifiedChinese:    shariSimplifiedChineseName,
		language.Spanish:              shariSpanishName,
		language.TraditionalChinese:   shariTraditionalChineseName}
)

var (
	// shariCharacter represents shari character.
	shariCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: shariBirthday,
		Code:     shariCode,
		Key:      character.Shari,
		Gender:   gender.Female,
		Name:     shariName,
		Special:  false}
)

var (
	// shariAmericanEnglishPhrase represents shari american english phrase.
	shariAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cheeky"}

	// shariCanadianFrenchPhrase represents shari canadian french phrase.
	shariCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tavu"}

	// shariDutchPhrase represents shari dutch phrase.
	shariDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brutaaltje"}

	// shariFrenchPhrase represents shari french phrase.
	shariFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tavu"}

	// shariGermanPhrase represents shari german phrase.
	shariGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "makimaki"}

	// shariItalianPhrase represents shari italian phrase.
	shariItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "babbuh"}

	// shariJapanesePhrase represents shari japanese phrase.
	shariJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウキキ"}

	// shariKoreanPhrase represents shari korean phrase.
	shariKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우끼네"}

	// shariLatinAmericanSpanishPhrase represents shari latin american spanish phrase.
	shariLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "babuí"}

	// shariRussianPhrase represents shari russian phrase.
	shariRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "щекасто"}

	// shariSimplifiedChinesePhrase represents shari simplified chinese phrase.
	shariSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呜吱吱"}

	// shariSpanishPhrase represents shari spanish phrase.
	shariSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tunante"}

	// shariTraditionalChinesePhrase represents shari traditional chinese phrase.
	shariTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗚吱吱"}
)

var (
	// shariPhrase represents shari phrase.
	shariPhrase = nook.Languages{
		language.AmericanEnglish:      shariAmericanEnglishPhrase,
		language.CanadianFrench:       shariCanadianFrenchPhrase,
		language.Dutch:                shariDutchPhrase,
		language.French:               shariFrenchPhrase,
		language.German:               shariGermanPhrase,
		language.Italian:              shariItalianPhrase,
		language.Japanese:             shariJapanesePhrase,
		language.Korean:               shariKoreanPhrase,
		language.LatinAmericanSpanish: shariLatinAmericanSpanishPhrase,
		language.Russian:              shariRussianPhrase,
		language.SimplifiedChinese:    shariSimplifiedChinesePhrase,
		language.Spanish:              shariSpanishPhrase,
		language.TraditionalChinese:   shariTraditionalChinesePhrase}
)

var (
	// Shari represents shari.
	Shari = nook.Villager{
		Character:   shariCharacter,
		Personality: personality.BigSister,
		Phrase:      shariPhrase}
)

package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	chaiBirthday = nook.Birthday{
		Day:   6,
		Month: time.March}
)

var (
	chaiCode = nook.Code{
		Value: "elp11"}
)

var (
	chaiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chai"}

	chaiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chaïpachipach"}

	chaiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chaikaneel"}

	chaiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chaïpachipach"}

	chaiGermanName = nook.Name{
		Language: language.German,
		Value:    "Chaizimtroll"}

	chaiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chairuuu ruuu"}

	chaiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フィーカロルロル"}

	chaiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "피카롤롤"}

	chaiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chaivuelalto"}

	chaiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чаймашу ушами"}

	chaiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啡卡大耳"}

	chaiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chaivuelalto"}

	chaiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啡卡大耳"}
)

var (
	chaiName = nook.Languages{
		language.AmericanEnglish:      chaiAmericanEnglishName,
		language.CanadianFrench:       chaiCanadianFrenchName,
		language.Dutch:                chaiDutchName,
		language.French:               chaiFrenchName,
		language.German:               chaiGermanName,
		language.Italian:              chaiItalianName,
		language.Japanese:             chaiJapaneseName,
		language.Korean:               chaiKoreanName,
		language.LatinAmericanSpanish: chaiLatinAmericanSpanishName,
		language.Russian:              chaiRussianName,
		language.SimplifiedChinese:    chaiSimplifiedChineseName,
		language.Spanish:              chaiSpanishName,
		language.TraditionalChinese:   chaiTraditionalChineseName}
)

var (
	chaiCharacter = nook.Character{
		Animal:   Elephant,
		Birthday: chaiBirthday,
		Code:     chaiCode,
		Gender:   nook.Female,
		Name:     chaiName}
)

var (
	chaiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "flap flap"}

	chaiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pachipach"}

	chaiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kaneel"}

	chaiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pachipach"}

	chaiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zimtroll"}

	chaiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ruuu ruuu"}

	chaiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ロルロル"}

	chaiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "롤롤"}

	chaiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "vuelalto"}

	chaiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "машу ушами"}

	chaiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大耳"}

	chaiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "vuelalto"}

	chaiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大耳"}
)

var (
	chaiPhrase = nook.Languages{
		language.AmericanEnglish:      chaiAmericanEnglishName,
		language.CanadianFrench:       chaiCanadianFrenchName,
		language.Dutch:                chaiDutchName,
		language.French:               chaiFrenchName,
		language.German:               chaiGermanName,
		language.Italian:              chaiItalianName,
		language.Japanese:             chaiJapaneseName,
		language.Korean:               chaiKoreanName,
		language.LatinAmericanSpanish: chaiLatinAmericanSpanishName,
		language.Russian:              chaiRussianName,
		language.SimplifiedChinese:    chaiSimplifiedChineseName,
		language.Spanish:              chaiSpanishName,
		language.TraditionalChinese:   chaiTraditionalChineseName}
)

var (
	Chai = nook.Villager{
		Character:   chaiCharacter,
		Personality: nook.Peppy,
		Phrase:      chaiPhrase}
)

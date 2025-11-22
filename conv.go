package youdao

import (
	"fmt"
	"github.com/bytedance/gg/gslice"
)

func convYouDaoResp(youDaoResp *rawYouDaoResponse) (*TranslateResponse, error) {
	if youDaoResp == nil {
		return nil, fmt.Errorf("youdao response is nil")
	}

	wordData := gslice.Get(youDaoResp.Ec.Word, 0).ValueOrZero()

	return &TranslateResponse{
		Ukphone:         wordData.Ukphone,
		Usphone:         wordData.Usphone,
		Translations:    gslice.Map(wordData.Trs, func(tr translation) string { return tr.Tr[0].L.I[0] }),
		WebTranslations: convWebTranslations(youDaoResp.WebTrans.WebTranslation),
		WordForms:       gslice.Map(wordData.Wfs, func(wf wordFrom) string { return wf.Wf.Name + wf.Wf.Value }),
		Etymologies: gslice.Map(youDaoResp.Etym.Etyms.Zh, func(etymology etymologyZh) *Etymology {
			return &Etymology{Value: etymology.Value, Desc: etymology.Desc}
		}),
		EgSentences: gslice.Map(youDaoResp.BlngSentsPart.SentencePair, func(sentencePair egSentencePair) *EGSentence {
			return &EGSentence{Sentence: sentencePair.Sentence, Translation: sentencePair.SentenceTranslation}
		}),
		Discrimination: convDiscrimination(youDaoResp),
	}, nil
}

func convWebTranslations(webTranslations []webTranslation) []string {
	var results []string
	for _, webTranslation := range webTranslations {
		if webTranslation.Same != "true" {
			continue
		}
		for _, trans := range webTranslation.Trans {
			results = append(results, trans.Value)
		}
	}
	return results
}

func convDiscrimination(youDaoResp *rawYouDaoResponse) []DiscriminateUsage {
	var maxHeadwords, maxHeadwordsIndex int32
	for i, data := range youDaoResp.Discriminate.Data {
		if int32(len(data.Headwords)) > maxHeadwords {
			maxHeadwords = int32(len(data.Headwords))
			maxHeadwordsIndex = int32(i)
		}
	}

	var results []DiscriminateUsage
	for _, usage := range youDaoResp.Discriminate.Data[maxHeadwordsIndex].Usages {
		results = append(results, DiscriminateUsage{Headword: usage.Headword, Usage: usage.Usage})
	}
	return results
}

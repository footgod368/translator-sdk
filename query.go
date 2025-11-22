package youdao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bytedance/gg/gslice"
	"github.com/footgod368/translator-sdk/utils"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const apiURL = "http://dict.youdao.com/jsonapi"

func Query(ctx context.Context, text string) (*TranslateResponse, error) {
	response, err := http.Get(utils.AppendURLParams(apiURL, map[string]string{"q": text}))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("youdao api error: %s", response.Status)
	}
	youDaoResponse := &rawYouDaoResponse{}
	if err = json.Unmarshal(body, youDaoResponse); err != nil {
		return nil, err
	}
	logrus.Debugln(utils.MarshalJSONToString(youDaoResponse))
	translateResponse, err := convYouDaoResp(youDaoResponse)
	if err != nil {
		return nil, err
	}
	return translateResponse, nil
}

func convYouDaoResp(youDaoResp *rawYouDaoResponse) (*TranslateResponse, error) {
	if youDaoResp == nil {
		return nil, fmt.Errorf("youdao response is nil")
	}

	wordData := gslice.Get(youDaoResp.Ec.Word, 0).ValueOrZero()

	return &TranslateResponse{
		Ukphone:         wordData.Ukphone,
		Usphone:         wordData.Usphone,
		Translations:    gslice.Map(wordData.Trs, func(tr Translation) string { return tr.Tr[0].L.I[0] }),
		WebTranslations: convWebTranslations(youDaoResp.WebTrans.WebTranslation),
		WordForms:       gslice.Map(wordData.Wfs, func(wf WordFrom) string { return wf.Wf.Name + wf.Wf.Value }),
		Etymologies: gslice.Map(youDaoResp.Etym.Etyms.Zh, func(etymology EtymologyZh) *Etymology {
			return &Etymology{Value: etymology.Value, Desc: etymology.Desc}
		}),
		EgSentences: gslice.Map(youDaoResp.BlngSentsPart.SentencePair, func(sentencePair EgSentencePair) *EGSentence {
			return &EGSentence{Sentence: sentencePair.Sentence, Translation: sentencePair.SentenceTranslation}
		}),
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

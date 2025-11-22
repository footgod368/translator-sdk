package youdao

import (
	"context"
	"encoding/json"
	"fmt"
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

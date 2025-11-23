package youdao

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/gg/gconv"
	"github.com/footgod368/translator-sdk/utils"
	"io"
	"net/http"
)

func Suggest(prefix string, num int32) ([]string, error) {
	response, err := http.Get(utils.AppendURLParams(apiURLSuggest,
		map[string]string{
			"q":       prefix,
			"le":      "eng",
			"num":     gconv.To[string](num),
			"doctype": "json",
		}))
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
	suggestResponse := rawSuggestResponse{}
	if err = json.Unmarshal(body, &suggestResponse); err != nil {
		return nil, err
	}
	var results []string
	for _, entry := range suggestResponse.Data.Entries {
		results = append(results, entry.Entry)
	}
	return results, nil
}

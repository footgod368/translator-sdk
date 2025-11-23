package youdao

type rawSuggestResponse struct {
	Data struct {
		Entries []struct {
			Entry   string `json:"entry"`
			Explain string `json:"explain"`
		} `json:"entries"`
		Language string `json:"language"`
		Query    string `json:"query"`
		Type     string `json:"type"`
	} `json:"data"`
	Result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"result"`
}

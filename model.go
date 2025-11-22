package youdao

type rawYouDaoResponse struct {
	WebTrans struct {
		WebTranslation []webTranslation `json:"web-translation"`
	} `json:"web_trans"`
	OxfordAdvanceHtml struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxfordAdvanceHtml"`
	PicDict struct {
		Pic []struct {
			Image string `json:"image"`
			Host  string `json:"host"`
			Url   string `json:"url"`
		} `json:"pic"`
	} `json:"pic_dict"`
	Simple struct {
		Query string `json:"query"`
		Word  []struct {
			Usphone          string `json:"usphone"`
			Ukphone          string `json:"ukphone"`
			Ukspeech         string `json:"ukspeech"`
			ReturnPhrase     string `json:"return-phrase"`
			Usspeech         string `json:"usspeech"`
			CollegeExamVoice struct {
				SpeechWord string `json:"speechWord"`
			} `json:"collegeExamVoice"`
		} `json:"word"`
	} `json:"simple"`
	Phrs struct {
		Word string `json:"word"`
		Phrs []struct {
			Phr struct {
				Headword struct {
					L struct {
						I string `json:"i"`
					} `json:"l"`
				} `json:"headword"`
				Trs []struct {
					Tr struct {
						L struct {
							I string `json:"i"`
						} `json:"l"`
					} `json:"tr"`
				} `json:"trs"`
				Source string `json:"source"`
			} `json:"phr"`
		} `json:"phrs"`
	} `json:"phrs"`
	Oxford struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxford"`
	WordElaboration struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"wordElaboration"`
	Syno struct {
		Synos []struct {
			Syno struct {
				Pos string `json:"pos"`
				Ws  []struct {
					W string `json:"w"`
				} `json:"ws"`
				Tran string `json:"tran"`
			} `json:"syno"`
		} `json:"synos"`
		Word string `json:"word"`
	} `json:"syno"`
	Collins struct {
		CollinsEntries []struct {
			Entries struct {
				Entry []struct {
					TranEntry []struct {
						PosEntry struct {
							Pos     string `json:"pos"`
							PosTips string `json:"pos_tips"`
						} `json:"pos_entry"`
						ExamSents struct {
							Sent []struct {
								ChnSent string `json:"chn_sent"`
								EngSent string `json:"eng_sent"`
							} `json:"sent"`
						} `json:"exam_sents"`
						Tran string `json:"tran"`
					} `json:"tran_entry"`
				} `json:"entry"`
			} `json:"entries"`
			Phonetic     string `json:"phonetic"`
			BasicEntries struct {
				BasicEntry []struct {
					Cet       string `json:"cet"`
					Headword  string `json:"headword"`
					Wordforms struct {
						Wordform []struct {
							Word string `json:"word"`
						} `json:"wordform"`
					} `json:"wordforms"`
				} `json:"basic_entry"`
			} `json:"basic_entries"`
			Headword string `json:"headword"`
			Star     string `json:"star"`
		} `json:"collins_entries"`
	} `json:"collins"`
	WordVideo struct {
		WordVideos []struct {
			Ad struct {
				Avatar string `json:"avatar"`
				Title  string `json:"title"`
				Url    string `json:"url"`
			} `json:"ad"`
			Video struct {
				Cover string `json:"cover"`
				Image string `json:"image"`
				Title string `json:"title"`
				Url   string `json:"url"`
			} `json:"video"`
		} `json:"word_videos"`
	} `json:"word_video"`
	Webster struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"webster"`
	Discriminate struct {
		Data []struct {
			Source string `json:"source"`
			Usages []struct {
				Headword string `json:"headword"`
				Usage    string `json:"usage"`
			} `json:"usages"`
			Headwords []string `json:"headwords"`
			Tran      string   `json:"tran,omitempty"`
		} `json:"data"`
		ReturnPhrase string `json:"return-phrase"`
	} `json:"discriminate"`
	Lang string `json:"lang"`
	Ec   struct {
		ExamType []string `json:"exam_type"`
		Source   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"source"`
		Word []struct {
			Usphone      string        `json:"usphone"`
			Ukphone      string        `json:"ukphone"`
			Ukspeech     string        `json:"ukspeech"`
			Trs          []Translation `json:"trs"`
			Wfs          []WordFrom    `json:"wfs"`
			ReturnPhrase struct {
				L struct {
					I string `json:"i"`
				} `json:"l"`
			} `json:"return-phrase"`
			Usspeech string `json:"usspeech"`
		} `json:"word"`
	} `json:"ec"`
	Ee struct {
		Source struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"source"`
		Word struct {
			Trs []struct {
				Pos string `json:"pos"`
				Tr  []struct {
					Exam struct {
						I struct {
							F struct {
								L []struct {
									I string `json:"i"`
								} `json:"l"`
							} `json:"f"`
						} `json:"i"`
					} `json:"exam,omitempty"`
					L struct {
						I string `json:"i"`
					} `json:"l"`
					SimilarWords []struct {
						Similar string `json:"similar"`
					} `json:"similar-words"`
				} `json:"tr"`
			} `json:"trs"`
			Phone        string `json:"phone"`
			Speech       string `json:"speech"`
			ReturnPhrase struct {
				L struct {
					I string `json:"i"`
				} `json:"l"`
			} `json:"return-phrase"`
		} `json:"word"`
	} `json:"ee"`
	BlngSentsPart struct {
		SentenceCount int              `json:"sentence-count"`
		SentencePair  []EgSentencePair `json:"sentence-pair"`
		More          string           `json:"more"`
		TrsClassify   []struct {
			Proportion string `json:"proportion"`
			Tr         string `json:"tr"`
		} `json:"trs-classify"`
	} `json:"blng_sents_part"`
	Individual struct {
		Trs []struct {
			Pos  string `json:"pos"`
			Tran string `json:"tran"`
		} `json:"trs"`
		Idiomatic []struct {
			Colloc struct {
				En string `json:"en"`
				Zh string `json:"zh"`
			} `json:"colloc"`
		} `json:"idiomatic"`
		Level    string `json:"level"`
		ExamInfo struct {
			Year             int `json:"year"`
			QuestionTypeInfo []struct {
				Time int    `json:"time"`
				Type string `json:"type"`
			} `json:"questionTypeInfo"`
			RecommendationRate int `json:"recommendationRate"`
			Frequency          int `json:"frequency"`
		} `json:"examInfo"`
		ReturnPhrase  string `json:"return-phrase"`
		PastExamSents []struct {
			En     string `json:"en"`
			Source string `json:"source"`
			Zh     string `json:"zh"`
		} `json:"pastExamSents"`
		Mnemonic struct {
			Method string `json:"method"`
			Word   string `json:"word"`
		} `json:"mnemonic"`
	} `json:"individual"`
	CollinsPrimary struct {
		Words struct {
			Indexforms []string `json:"indexforms"`
			Word       string   `json:"word"`
		} `json:"words"`
		Gramcat []struct {
			Audiourl      string `json:"audiourl"`
			Pronunciation string `json:"pronunciation"`
			Senses        []struct {
				Sensenumber string `json:"sensenumber"`
				Examples    []struct {
					Sense struct {
						Lang string `json:"lang"`
						Word string `json:"word"`
					} `json:"sense"`
					Example string `json:"example"`
				} `json:"examples"`
				Definition string `json:"definition"`
				Lang       string `json:"lang"`
				Word       string `json:"word"`
			} `json:"senses"`
			Partofspeech string `json:"partofspeech"`
			Audio        string `json:"audio"`
			Forms        []struct {
				Form string `json:"form"`
			} `json:"forms"`
		} `json:"gramcat"`
	} `json:"collins_primary"`
	RelWord struct {
		Word string `json:"word"`
		Stem string `json:"stem"`
		Rels []struct {
			Rel struct {
				Pos   string `json:"pos"`
				Words []struct {
					Word string `json:"word"`
					Tran string `json:"tran"`
				} `json:"words"`
			} `json:"rel"`
		} `json:"rels"`
	} `json:"rel_word"`
	AuthSentsPart struct {
		SentenceCount int    `json:"sentence-count"`
		More          string `json:"more"`
		Sent          []struct {
			Score      float64 `json:"score"`
			Speech     string  `json:"speech"`
			SpeechSize string  `json:"speech-size"`
			Source     string  `json:"source"`
			Url        string  `json:"url"`
			Foreign    string  `json:"foreign"`
		} `json:"sent"`
	} `json:"auth_sents_part"`
	MagicWords struct {
		MagicWords []struct {
			Paraphrase string `json:"paraphrase"`
			Word       string `json:"word"`
			Info       []struct {
				Phonetic          string `json:"phonetic"`
				VideoUrl          string `json:"videoUrl"`
				PhoneticType      string `json:"phoneticType"`
				VideoCover        string `json:"videoCover"`
				WatermarkVideoUrl string `json:"watermarkVideoUrl"`
			} `json:"info"`
			PremiumVideoStatus int `json:"premiumVideoStatus"`
		} `json:"magic_words"`
	} `json:"magic_words"`
	MediaSentsPart struct {
		SentenceCount int    `json:"sentence-count"`
		More          string `json:"more"`
		Query         string `json:"query"`
		Sent          []struct {
			Mediatype string `json:"@mediatype"`
			Snippets  struct {
				Snippet []struct {
					StreamUrl string `json:"streamUrl"`
					Duration  string `json:"duration,omitempty"`
					Swf       string `json:"swf"`
					Name      string `json:"name"`
					Source    string `json:"source"`
					Win8      string `json:"win8,omitempty"`
					SourceUrl string `json:"sourceUrl,omitempty"`
					ImageUrl  string `json:"imageUrl,omitempty"`
				} `json:"snippet"`
			} `json:"snippets"`
			SpeechSize string `json:"speech-size,omitempty"`
			Eng        string `json:"eng"`
			Chn        string `json:"chn,omitempty"`
		} `json:"sent"`
	} `json:"media_sents_part"`
	ExpandEc struct {
		ReturnPhrase string `json:"return-phrase"`
		Source       struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"source"`
		Word []struct {
			TransList []struct {
				Content struct {
					DetailPos string `json:"detailPos"`
					ExamType  []struct {
						En string `json:"en"`
						Zh string `json:"zh"`
					} `json:"examType"`
					Sents []struct {
						SentOrig   string `json:"sentOrig"`
						SourceType string `json:"sourceType"`
						SentSpeech string `json:"sentSpeech"`
						SentTrans  string `json:"sentTrans"`
						Source     string `json:"source"`
						Type       string `json:"type,omitempty"`
					} `json:"sents"`
				} `json:"content"`
				Trans string `json:"trans"`
			} `json:"transList"`
			Pos string `json:"pos"`
			Wfs []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"wfs"`
		} `json:"word"`
	} `json:"expand_ec"`
	Etym struct {
		Etyms struct {
			En []struct {
				Source string `json:"source"`
				Word   string `json:"word"`
				Value  string `json:"value"`
				Url    string `json:"url"`
				Desc   string `json:"desc"`
			} `json:"en"`
			Zh []EtymologyZh `json:"zh"`
		} `json:"etyms"`
		Word string `json:"word"`
	} `json:"etym"`
	Special struct {
		Summary struct {
			Sources struct {
				Source struct {
					Site string `json:"site"`
					Url  string `json:"url"`
				} `json:"source"`
			} `json:"sources"`
			Text string `json:"text"`
		} `json:"summary"`
		CoAdd   string `json:"co-add"`
		Total   string `json:"total"`
		Entries []struct {
			Entry struct {
				Major string `json:"major"`
				Trs   []struct {
					Tr struct {
						Nat      string `json:"nat"`
						ChnSent  string `json:"chnSent"`
						Cite     string `json:"cite"`
						DocTitle string `json:"docTitle"`
						EngSent  string `json:"engSent"`
						Url      string `json:"url"`
					} `json:"tr"`
				} `json:"trs"`
				Num int `json:"num"`
			} `json:"entry"`
		} `json:"entries"`
	} `json:"special"`
	Senior struct {
		EncryptedData string `json:"encryptedData"`
		Source        struct {
			Name string `json:"name"`
		} `json:"source"`
	} `json:"senior"`
	Input      string `json:"input"`
	MusicSents struct {
		SentsData []struct {
			SongName         string `json:"songName"`
			LyricTranslation string `json:"lyricTranslation"`
			Singer           string `json:"singer"`
			CoverImg         string `json:"coverImg"`
			SupportCount     int    `json:"supportCount"`
			Lyric            string `json:"lyric"`
			Link             string `json:"link"`
			LyricList        []struct {
				Duration         int    `json:"duration"`
				LyricTranslation string `json:"lyricTranslation"`
				Lyric            string `json:"lyric"`
				Start            int    `json:"start"`
			} `json:"lyricList"`
			Id              string `json:"id"`
			SongId          string `json:"songId"`
			DecryptedSongId string `json:"decryptedSongId"`
			PlayUrl         string `json:"playUrl"`
		} `json:"sents_data"`
		More bool   `json:"more"`
		Word string `json:"word"`
	} `json:"music_sents"`
	Baike struct {
		Summarys []struct {
			Summary string `json:"summary"`
			Key     string `json:"key"`
		} `json:"summarys"`
		Source struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"source"`
	} `json:"baike"`
	Meta struct {
		Input           string   `json:"input"`
		GuessLanguage   string   `json:"guessLanguage"`
		IsHasSimpleDict string   `json:"isHasSimpleDict"`
		Le              string   `json:"le"`
		Lang            string   `json:"lang"`
		Dicts           []string `json:"dicts"`
	} `json:"meta"`
	Le            string `json:"le"`
	OxfordAdvance struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxfordAdvance"`
}

type Translation struct {
	Tr []struct {
		L struct {
			I []string `json:"i"`
		} `json:"l"`
	} `json:"tr"`
}

type webTranslation struct {
	Same      string `json:"@same,omitempty"`
	Key       string `json:"key"`
	KeySpeech string `json:"key-speech"`
	Trans     []struct {
		Summary struct {
			Line []string `json:"line"`
		} `json:"summary,omitempty"`
		Value   string `json:"value"`
		Support int    `json:"support,omitempty"`
		Url     string `json:"url,omitempty"`
	} `json:"trans"`
}

type WordFrom struct {
	Wf struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"wf"`
}

type EtymologyZh struct {
	Source string `json:"source"`
	Word   string `json:"word"`
	Value  string `json:"value"`
	Url    string `json:"url"`
	Desc   string `json:"desc"`
}

type EgSentencePair struct {
	Sentence            string `json:"sentence"`
	SentenceEng         string `json:"sentence-eng"`
	SentenceTranslation string `json:"sentence-translation"`
	SpeechSize          string `json:"speech-size"`
	AlignedWords        struct {
		Src struct {
			Chars []struct {
				S      string `json:"@s"`
				E      string `json:"@e"`
				Aligns struct {
					Sc []struct {
						Id string `json:"@id"`
					} `json:"sc"`
					Tc []struct {
						Id string `json:"@id"`
					} `json:"tc"`
				} `json:"aligns"`
				Id string `json:"@id"`
			} `json:"chars"`
		} `json:"src"`
		Tran struct {
			Chars []struct {
				S      string `json:"@s"`
				E      string `json:"@e"`
				Aligns struct {
					Sc []struct {
						Id string `json:"@id"`
					} `json:"sc"`
					Tc []struct {
						Id string `json:"@id"`
					} `json:"tc"`
				} `json:"aligns"`
				Id string `json:"@id"`
			} `json:"chars"`
		} `json:"tran"`
	} `json:"aligned-words"`
	Source         string `json:"source"`
	Url            string `json:"url"`
	SentenceSpeech string `json:"sentence-speech"`
}

type TranslateResponse struct {
	Ukphone         string        `json:"ukphone"`
	Usphone         string        `json:"usphone"`
	Translations    []string      `json:"translations"`
	WebTranslations []string      `json:"web_translations"`
	WordForms       []string      `json:"word_forms"`
	Etymologies     []*Etymology  `json:"etymologies"`
	EgSentences     []*EGSentence `json:"eg_sentences"`
}

type Etymology struct {
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

type EGSentence struct {
	Sentence    string `json:"sentence"`
	Translation string `json:"translation"`
}

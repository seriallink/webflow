package webflow

type ErrorMessage struct {
	Msg         string        `json:"msg"`
	Code        int           `json:"code"`
	Name        string        `json:"name"`
	Path        string        `json:"path"`
	Err         string        `json:"err"`
	Problems    []string      `json:"problems"`
	ProblemData []ProblemData `json:"problem_data"`
	Extensions  *Extensions   `json:"extensions"`
}

type ProblemData struct {
	Slug string `json:"slug"`
	Msg  string `json:"msg"`
}

type Extensions struct {
	Input *Input `json:"input"`
	Meta  *Meta  `json:"meta"`
}

type Input struct {
	CollectionId     *ObjectId `json:"collection_id"`
	ItemId           *ObjectId `json:"item_id"`
	Target           string    `json:"target"`
	Mode             string    `json:"mode"`
	NeedStaging      bool      `json:"need_staging"`
	NeedLive         bool      `json:"need_live"`
	NeedCollections  bool      `json:"need_collections"`
	NeedStagingDraft bool      `json:"need_staging_draft"`
	IsPatchMode      bool      `json:"isPatchMode"`
	IsSilentMode     bool      `json:"isSilentMode"`
	SkipInvalidFiles bool      `json:"skipInvalidFiles"`
}

type Meta struct {
	AuthType string    `json:"authType"`
	UserId   *ObjectId `json:"userId"`
}

type ObjectId struct {
	Id       map[string]interface{} `json:"id"`
	BSonType string                 `json:"_bsontype"`
}

func (e *ErrorMessage) Error() string {
	if len(e.Problems) > 0 {
		return e.Problems[0]
	}
	return e.Err
}

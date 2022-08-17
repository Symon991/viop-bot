package twitter

type AddRuleMessage struct {
	Add []Rule `json:"add"`
}

type DeleteRuleMessage struct {
	Delete struct {
		IDs []string `json:"ids"`
	} `json:"delete"`
}

type GetRulesMessage struct {
	Data []Rule `json:"data"`
}

type Rule struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
	Tag   string `json:"tag,omitempty"`
}

type StreamMessage struct {
	Data          Data           `json:"data"`
	MatchingRules []MatchingRule `json:"matching_rules"`
}

type MatchingRule struct {
	ID  string `json:"id"`
	Tag string `json:"tag"`
}

type Data struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

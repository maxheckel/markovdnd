package domain

type Chain struct {
	Type string `json:"type"`
	Chain map[string]map[string]int `json:"chain"`
	StartingWords map[string][]string `json:"starting_words"`
	Input string `json:"-"`
}


func NewChain(text string, chainType string) Chain{
	return Chain{
		Type: chainType,
		Input: text,
		Chain: map[string]map[string]int{},
		StartingWords: map[string][]string{},
	}
}



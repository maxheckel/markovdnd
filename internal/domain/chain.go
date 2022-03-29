package domain

type Chain struct {
	Type string `json:"type"`
	Chain map[string]map[string]int `json:"chain"`
	StartingWords map[string][]string `json:"starting_words"`
	Input string `json:"-"`
	ArticleWords []string `json:"-"`
}


func NewChain(text string, articles []string, chainType string) Chain{
	return Chain{
		Type: chainType,
		Input: text,
		ArticleWords: articles,
		Chain: map[string]map[string]int{},
		StartingWords: map[string][]string{},
	}
}

package get

import (
	"encoding/json"
	"github.com/maxheckel/markovdnd/internal/domain"
	"github.com/maxheckel/markovdnd/internal/services/store/chain"
	"net/http"
)

type Crawled struct {
	Store chain.Store
}

func (cr Crawled) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := domain.AvailableBooks{
		Crawled: []string{},
	}
	var err error
	resp.Crawled, err = cr.Store.GetCrawled()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	output, err := json.Marshal(resp)
	w.Write(output)
}
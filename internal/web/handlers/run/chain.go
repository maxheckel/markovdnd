package run

import (
	"github.com/maxheckel/auto-dnd/internal/services/chainer"
	"github.com/maxheckel/auto-dnd/internal/services/store/chain"
	"net/http"
	"strings"
)

type Chain struct {
	Store chain.Store
}

func (rc Chain) ServeHTTP(w http.ResponseWriter, r *http.Request){

	err := rc.Store.LoadChain("cos")
	if err != nil {
		panic(any(err))
	}
	chains, err := rc.Store.GetChains("cos")
	res := []string{}
	for i := 0; i < 3; i++ {
		chainNum := 0
		paragraph, err := chainer.Run(chains[chainNum], 100)
		res = append(res, paragraph)
		if err != nil {
			panic(any(err))
		}
	}
	w.Write([]byte(strings.Join(res, "\n\n")))

}
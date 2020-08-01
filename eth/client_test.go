package eth

import (
	"math/big"

	lpTypes "github.com/livepeer/go-livepeer/eth/types"
)

func copyTranscoders(transcoders []*lpTypes.Transcoder) []*lpTypes.Transcoder {
	cp := make([]*lpTypes.Transcoder, 0)
	for _, tr := range transcoders {
		trCp := new(lpTypes.Transcoder)
		trCp.Address = tr.Address
		trCp.DelegatedStake = new(big.Int)
		*trCp.DelegatedStake = *tr.DelegatedStake
		cp = append(cp, trCp)
	}
	return cp
}

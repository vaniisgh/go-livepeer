package eth

import (
	"math/big"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	lpTypes "github.com/livepeer/go-livepeer/eth/types"
	"github.com/stretchr/testify/assert"
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

func TestFindTranscoderHints(t *testing.T) {
	assert := assert.New(t)

	transcoders := []*lpTypes.Transcoder{
		{
			Address:        ethcommon.HexToAddress("aaa"),
			DelegatedStake: big.NewInt(5),
		},
		{
			Address:        ethcommon.HexToAddress("bbb"),
			DelegatedStake: big.NewInt(4),
		},
		{
			Address:        ethcommon.HexToAddress("ccc"),
			DelegatedStake: big.NewInt(3),
		},
		{
			Address:        ethcommon.HexToAddress("ddd"),
			DelegatedStake: big.NewInt(2),
		},
		{
			Address:        ethcommon.HexToAddress("eee"),
			DelegatedStake: big.NewInt(1),
		},
	}

	// del == 'aaa' == head
	hints := findTranscoderHints(ethcommon.HexToAddress("aaa"), transcoders)
	assert.Equal(hints.PosPrev, ethcommon.Address{})
	assert.Equal(hints.PosNext, ethcommon.HexToAddress("bbb"))

	// del == 'eee' == tail
	hints = findTranscoderHints(ethcommon.HexToAddress("eee"), transcoders)
	assert.Equal(hints.PosPrev, ethcommon.HexToAddress("ddd"))
	assert.Equal(hints.PosNext, ethcommon.Address{})

	// del == 'ccc'
	hints = findTranscoderHints(ethcommon.HexToAddress("ccc"), transcoders)
	assert.Equal(hints.PosPrev, ethcommon.HexToAddress("bbb"))
	assert.Equal(hints.PosNext, ethcommon.HexToAddress("ddd"))
}

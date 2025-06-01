package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	proto_dice_roll "dice_roll__v1_not_provablyfair/gen/go/dice_roll"
	"dice_roll__v1_not_provablyfair/tests/suite"
)

func TestGame100(t *testing.T) {
	ctx, st := suite.New(t)
	play_resp, err := st.GameClient.Play(ctx, &proto_dice_roll.PlayRequest{})
	require.NoError(t, err)
	assert.NotEmpty(t, play_resp)
	for i := 1; i < 100; i++ {
		play_resp, err = st.GameClient.Play(ctx, &proto_dice_roll.PlayRequest{})
		require.NoError(t, err)
		s := play_resp.ServerRoll
		p := play_resp.PlayerRoll
		if s > p {
			assert.Equal(t, play_resp.Winner, "server")
		} else if s < p {
			assert.Equal(t, play_resp.Winner, "player")
		} else {
			assert.Equal(t, play_resp.Winner, "draft")
		}
	}
}

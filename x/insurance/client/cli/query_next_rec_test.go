package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/DenisCherkasskikh/insurance/testutil/network"
	"github.com/DenisCherkasskikh/insurance/testutil/nullify"
	"github.com/DenisCherkasskikh/insurance/x/insurance/client/cli"
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
)

func networkWithNextRecObjects(t *testing.T) (*network.Network, types.NextRec) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	nextRec := &types.NextRec{}
	nullify.Fill(&nextRec)
	state.NextRec = nextRec
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.NextRec
}

func TestShowNextRec(t *testing.T) {
	net, obj := networkWithNextRecObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.NextRec
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowNextRec(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetNextRecResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.NextRec)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.NextRec),
				)
			}
		})
	}
}

package cli

import (
	"strconv"

	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-record [brand] [model] [year] [owner] [license-plate] [vin-number] [odometer] [side] [part] [payout]",
		Short: "Broadcast message addRecord",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBrand := args[0]
			argModel := args[1]
			argYear := args[2]
			argOwner := args[3]
			argLicensePlate := args[4]
			argVinNumber := args[5]
			argOdometer := args[6]
			argSide := args[7]
			argPart := args[8]
			argPayout := args[9]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddRecord(
				clientCtx.GetFromAddress().String(),
				argBrand,
				argModel,
				argYear,
				argOwner,
				argLicensePlate,
				argVinNumber,
				argOdometer,
				argSide,
				argPart,
				argPayout,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

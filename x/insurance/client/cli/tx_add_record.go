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
		Use:   "add-record [brand] [model] [owner] [licence-plate] [vin-number] [mileage] [side] [part] [payout]",
		Short: "Broadcast message addRecord",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBrand := args[0]
			argModel := args[1]
			argOwner := args[2]
			argLicencePlate := args[3]
			argVinNumber := args[4]
			argMileage := args[5]
			argSide := args[6]
			argPart := args[7]
			argPayout := args[8]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddRecord(
				clientCtx.GetFromAddress().String(),
				argBrand,
				argModel,
				argOwner,
				argLicencePlate,
				argVinNumber,
				argMileage,
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

package cli

import (
	"context"

	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListCrashRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-crash-record",
		Short: "list all crashRecord",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllCrashRecordRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CrashRecordAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowCrashRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-crash-record [index]",
		Short: "shows a crashRecord",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetCrashRecordRequest{
				Index: argIndex,
			}

			res, err := queryClient.CrashRecord(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

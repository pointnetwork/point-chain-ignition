package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"point/x/lockenomics/types"
)

func CmdListDelegatedAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-delegated-amount",
		Short: "list all delegatedAmount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDelegatedAmountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DelegatedAmountAll(context.Background(), params)
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

func CmdShowDelegatedAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-delegated-amount [index]",
		Short: "shows a delegatedAmount",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetDelegatedAmountRequest{
				Index: argIndex,
			}

			res, err := queryClient.DelegatedAmount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

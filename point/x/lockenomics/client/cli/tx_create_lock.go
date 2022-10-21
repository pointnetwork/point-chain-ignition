package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"point/x/lockenomics/types"
)

var _ = strconv.Itoa(0)

func CmdCreateLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-lock [lenght] [validator]",
		Short: "lock tokens delegated to specific validator for a specific period length",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argLenght, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argValidator := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateLock(
				clientCtx.GetFromAddress().String(),
				argLenght,
				argValidator,
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

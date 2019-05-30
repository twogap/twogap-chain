package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
	accountcmd "github.com/twogap/twogap-chain/x/account/client/cli"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

//NewModuleClient - new module
func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group nameservice queries under a subcommand
	namesvcQueryCmd := &cobra.Command{
		Use:   "nameservice",
		Short: "Querying commands for the nameservice module",
	}

	namesvcQueryCmd.AddCommand(client.GetCommands(
		accountcmd.GetCmdResolveName(mc.storeKey, mc.cdc),
		accountcmd.GetCmdWhois(mc.storeKey, mc.cdc),
		accountcmd.GetCmdNames(mc.storeKey, mc.cdc),
	)...)

	return namesvcQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	namesvcTxCmd := &cobra.Command{
		Use:   "nameservice",
		Short: "Nameservice transactions subcommands",
	}

	namesvcTxCmd.AddCommand(client.PostCommands(
		accountcmd.GetCmdBuyName(mc.cdc),
		accountcmd.GetCmdSetName(mc.cdc),
	)...)

	return namesvcTxCmd
}

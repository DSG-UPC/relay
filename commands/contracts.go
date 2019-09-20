package commands

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/kvartalo/relay/config"
	"github.com/kvartalo/relay/eth"
	"github.com/kvartalo/relay/utils"
	"github.com/urfave/cli"
)

var ContractsCommands = []cli.Command{
	{
		Name:  "contracts",
		Usage: "interact with contracts",
		Subcommands: []cli.Command{{
			Name:  "token",
			Usage: "for token smart contract",
			Subcommands: []cli.Command{
				{
					Name:   "deploy",
					Usage:  "deploy device smart contract",
					Action: cmdDeviceDeploy,
				},
				{
					Name:   "mint",
					Usage:  "mint tokens that will go to this Relay address",
					Action: cmdTokenMint,
				},
				{
					Name:   "transfer",
					Usage:  "tranfer tokens to a specified address",
					Action: cmdTokenTransfer,
				},
				{
					Name:   "info",
					Usage:  "gets ifo about the token stats",
					Action: cmdTokenInfo,
				},
				{
					Name:   "renunceownership",
					Usage:  "renunce ownership",
					Action: cmdTokenRenunceOwnership,
				},
			},
		}, // in the future here will come more contracts
		},
	},
}

func cmdDeviceDeploy(c *cli.Context) error {
	name := c.Args().Get(1)
	if name == "" {
		os.Exit(0)
	}
	initValueStr := c.Args().Get(1)
	if initValueStr == "" {
		os.Exit(0)
	}
	initValue, err := utils.StringToBigInt(initValueStr)
	if err != nil {
		color.Red("Amount parsing error. Usage: contracts token mint [amount]")
		os.Exit(0)
	}
	ethSrv := loadKeyStore(c, nil)
	err = ethSrv.DeployDeviceContract(name, initValue)
	return err
}

func cmdTokenMint(c *cli.Context) error {
	amountStr := c.Args().Get(0)
	if amountStr == "" {
		color.Red("No amount specified. Usage: contracts token mint amount deviceAddress dstAddress")
		os.Exit(0)
	}
	amount, err := utils.StringToBigInt(amountStr)
	if err != nil {
		color.Red("Amount parsing error. Usage: contracts ttoken mint amount deviceAddress dstAddress")
		os.Exit(0)
	}

	deviceAddressStr := c.Args().Get(1)
	if deviceAddressStr == "" {
		color.Red("No device specified. Usage: contracts token mint amount deviceAddress dstAddress")
		os.Exit(0)
	}
	deviceAddr := common.HexToAddress(deviceAddressStr)

	destinationAddressStr := c.Args().Get(2)
	if destinationAddressStr == "" {
		color.Red("No destination specified. Usage: contracts token mint amount deviceAddress dstAddress")
		os.Exit(0)
	}
	destinationAddress := common.HexToAddress(destinationAddressStr)
	fromAddr := common.HexToAddress(config.C.Keystorage.Address)

	ethSrv := loadRelay(c, nil)

	auth, err := eth.GetAuth()
	if err != nil {
		return err
	}

	r, s, v, err := getRSV(auth, ethSrv, amount, deviceAddr, fromAddr, destinationAddress)
	if err != nil {
		color.Red("Error signing transactions")
		os.Exit(0)
	}

	// tx, err := ethSrv.Token.Mint(auth, addr, amount)
	// if err != nil {
	// 	return err
	// }

	// color.Green("mint success, tx: " + tx.Hash().Hex())

	tokenBalance, err := ethSrv.devices[deviceAddr].Mint(nil, destinationAddress, r, s, v)
	if err != nil {
		return err
	}
	color.Magenta("current token balance: " + tokenBalance.String() + " Tokens\n")

	return nil
}

func cmdTokenInfo(c *cli.Context) error {

	ethSrv := loadRelay(c, nil)
	tokenContractAddr := common.HexToAddress(config.C.Contracts.Token)
	ethSrv.LoadTokenContract(tokenContractAddr)

	owner, err := ethSrv.Token.Owner(nil)
	if err != nil {
		return err
	}

	totalSupply, err := ethSrv.Token.TotalSupply(nil)
	if err != nil {
		return err
	}

	taxPercent, err := ethSrv.Token.TaxPercent(nil)
	if err != nil {
		return err
	}

	taxDestination, err := ethSrv.Token.TaxDestination(nil)
	if err != nil {
		return err
	}

	color.Green("Token Owner     : %v", owner.Hex())
	color.Green("TotalSupply     : %v", totalSupply.String())
	color.Green("TaxPercent      : %v", taxPercent.String())
	color.Green("TaxDestination  : %v", taxDestination.Hex())

	return nil
}

func cmdTokenRenunceOwnership(c *cli.Context) error {

	ethSrv := loadRelay(c, nil)
	tokenContractAddr := common.HexToAddress(config.C.Contracts.Token)
	ethSrv.LoadTokenContract(tokenContractAddr)

	relayerAddr := common.HexToAddress(config.C.Keystorage.Address)

	nonce, err := ethSrv.Client().PendingNonceAt(context.Background(), relayerAddr)
	if err != nil {
		return err
	}

	gasPrice, err := ethSrv.Client().SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := eth.GetAuth()
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// transfer token
	color.Green("Renuncing ownership ")
	tx, err := ethSrv.Token.RenounceOwnership(auth)
	if err != nil {
		return err
	}

	color.Green("Sent tx %v", tx.Hash().Hex())
	return nil
}

func cmdTokenTransfer(c *cli.Context) error {
	amountStr := c.Args().Get(1)
	if amountStr == "" {
		os.Exit(0)
	}
	amount, err := utils.StringToBigInt(amountStr)
	if err != nil {
		color.Red("No amount specified. Usage: contracts token transfer [addr] [amount]")
		os.Exit(0)
	}

	toAddrHex := c.Args().Get(0)
	if toAddrHex == "" {
		color.Red("No address specified. Usage: contracts token transfer [addr] [amount]")
		os.Exit(0)
	}
	toAddr := common.HexToAddress(toAddrHex)

	ethSrv := loadRelay(c, nil)

	tokenContractAddr := common.HexToAddress(config.C.Contracts.Token)
	ethSrv.LoadTokenContract(tokenContractAddr)

	fromAddr := common.HexToAddress(config.C.Keystorage.Address)
	tokenAddr := common.HexToAddress(config.C.Contracts.Token)

	nonce, err := ethSrv.Client().PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		return err
	}

	gasPrice, err := ethSrv.Client().SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := eth.GetAuth()
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tokenNonce, err := ethSrv.Token.NonceOf(nil, fromAddr)
	if err != nil {
		return err
	}

	tokenNonceBytes := utils.Uint64ToEthBytes(tokenNonce.Uint64())
	amountBytes := utils.Uint64ToEthBytes(amount.Uint64())

	// build tx msg
	var msgBytes []byte
	msgBytes = append(msgBytes, byte(0x19))
	msgBytes = append(msgBytes, byte(0))
	msgBytes = append(msgBytes, tokenAddr.Bytes()...)
	msgBytes = append(msgBytes, tokenNonceBytes...)
	msgBytes = append(msgBytes, fromAddr.Bytes()...)
	msgBytes = append(msgBytes, toAddr.Bytes()...)
	msgBytes = append(msgBytes, amountBytes...)
	fmt.Println(common.ToHex(msgBytes))

	// sign msg
	sig, err := ethSrv.SignBytes(msgBytes)
	if err != nil {
		return err
	}
	sig[64] += 27

	// get r, s, v from sig
	var r, s [32]byte
	copy(r[:], sig[:32])
	copy(s[:], sig[32:64])
	v := sig[64]

	// transfer token
	tx, err := ethSrv.Token.Transfer(auth, fromAddr, toAddr, amount, r, s, v)
	if err != nil {
		return err
	}

	color.Green("Token transfer success, tx: " + tx.Hash().Hex())
	return nil
}

func getRSV(auth *bind.TransactOpts, ethSrv *eth.EthService, amount *big.Int, contractAddr, fromAddr, toAddr common.Address) (r, s [32]byte, v byte, err error) {
	var emptyBytes [32]byte
	nonce, err := ethSrv.Client().PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		return emptyBytes, emptyBytes, byte(0), err
	}
	gasPrice, err := ethSrv.Client().SuggestGasPrice(context.Background())
	if err != nil {
		return emptyBytes, emptyBytes, byte(0), err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tokenNonce, err := ethSrv.Token.NonceOf(nil, fromAddr)
	if err != nil {
		return emptyBytes, emptyBytes, byte(0), err
	}

	tokenNonceBytes := utils.Uint64ToEthBytes(tokenNonce.Uint64())
	amountBytes := utils.Uint64ToEthBytes(amount.Uint64())

	// build tx msg
	var msgBytes []byte
	msgBytes = append(msgBytes, byte(0x19))
	msgBytes = append(msgBytes, byte(0))
	msgBytes = append(msgBytes, contractAddr.Bytes()...)
	msgBytes = append(msgBytes, tokenNonceBytes...)
	msgBytes = append(msgBytes, fromAddr.Bytes()...)
	msgBytes = append(msgBytes, toAddr.Bytes()...)
	msgBytes = append(msgBytes, amountBytes...)
	fmt.Println(common.ToHex(msgBytes))

	// sign msg
	sig, err := ethSrv.SignBytes(msgBytes)

	if err != nil {
		return emptyBytes, emptyBytes, byte(0), err
	}
	sig[64] += 27

	// get r, s, v from sig
	//var r, s [32]byte
	copy(r[:], sig[:32])
	copy(s[:], sig[32:64])
	//v := sig[64]
	v = sig[64]
	return r, s, v, nil

}

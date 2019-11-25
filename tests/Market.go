package tests

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/inn4science/perigord-example/bindings"
	"github.com/inn4science/perigord/contract"
	"github.com/inn4science/perigord/network"
	"github.com/inn4science/perigord/testing"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/viper"
	"gopkg.in/check.v1"
	. "gopkg.in/check.v1"
)

type MarketSuite struct {
	network *network.Network
}

func getMnemonic() string {
	viper.SetConfigFile("perigord.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal()
	}
	mnemonic := viper.GetStringMapString("networks.dev")["mnemonic"]
	return mnemonic
}

func getNetworkAddress() string {
	viper.SetConfigFile("perigord.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal()
	}
	networkAddr := viper.GetStringMapString("networks.dev")["url"]
	return networkAddr
}

func sendETH(s *MarketSuite, c *ethclient.Client, sender int, receiver common.Address, value *big.Int) {

	senderAcc := s.network.Accounts()[sender].Address
	nonce, err := c.PendingNonceAt(context.Background(), senderAcc)
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(6721975) // in units

	gasPrice := big.NewInt(3700000000)
	wallet, err := hdwallet.NewFromMnemonic(getMnemonic())
	toAddress := receiver
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := wallet.PrivateKey(s.network.Accounts()[sender])

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	rawTx := hex.EncodeToString(ts.GetRlp(0))

	var trx *types.Transaction

	rawTxBytes, err := hex.DecodeString(rawTx)
	err = rlp.DecodeBytes(rawTxBytes, &trx)

	err = c.SendTransaction(context.Background(), trx)
	if err != nil {
		log.Fatal(err)
	}
}

func ensureAuth(auth bind.TransactOpts) *bind.TransactOpts {
	return &bind.TransactOpts{
		auth.From,
		auth.Nonce,
		auth.Signer,
		auth.Value,
		auth.GasPrice,
		auth.GasLimit,
		auth.Context}
}

func changeAuth(s MarketSuite, account int) bind.TransactOpts {
	return *s.network.NewTransactor(s.network.Accounts()[account])
}

var _ = check.Suite(&MarketSuite{})

func (s *MarketSuite) SetUpTest(c *check.C) {
	nw, err := testing.SetUpTest()
	if err != nil {
		c.Fatal(err)
	}

	s.network = nw
}

func (s *MarketSuite) TearDownTest(c *check.C) {
	testing.TearDownTest()
}

// USER TESTS GO HERE

func (s *MarketSuite) TestMarket(c *check.C) {
	contractSession := contract.Session("Market")
	c.Assert(contractSession, NotNil)
	contractSessionActual, ok := contractSession.(*bindings.MarketSession)
	c.Assert(ok, Equals, true)
	c.Assert(contractSessionActual, NotNil)

	owner, err := contractSessionActual.Owner()
	if err != nil {
		log.Fatal()
	}
	account0 := s.network.Accounts()[0]
	c.Assert(owner.Hex(), Equals, account0.Address.Hex()) //Owner account is account 0

	ownerInd := 0
	sender := 5
	receiver := 6

	senderAcc := s.network.Accounts()[sender].Address
	receiverAcc := s.network.Accounts()[receiver].Address

	//Call contract on behalf of its owner
	auth := changeAuth(*s, ownerInd)
	_, err = contractSessionActual.Contract.SetSenderReceiverPair(ensureAuth(auth),
		senderAcc, receiverAcc)
	if err != nil {
		log.Fatal()
	}

	auth = changeAuth(*s, sender) //Change auth fo senderAcc to make deposit on behalf of the sender

	client, err := ethclient.Dial(getNetworkAddress())
	if err != nil {
		log.Fatal()
	}

	//Let's check the current balance
	balance, err := client.BalanceAt(context.Background(), contract.AddressOf("Market"), nil)
	if err != nil {
		log.Fatal()
	}
	c.Assert(balance.Int64(), Equals, big.NewInt(0).Int64()) //Balance shoud be 0

	//Let's transfer 3 ETH to the contract on behalf of the sender
	value := big.NewInt(3000000000000000000) // in wei (3 eth)
	contractReceiver := contract.AddressOf("Market")
	sendETH(s, client, sender, contractReceiver, value)

	balance2, err := client.BalanceAt(context.Background(), contract.AddressOf("Market"), nil)
	if err != nil {
		log.Fatal()
	}

	c.Assert(balance2.Int64(), Equals, value.Int64()) //Balance shoud be 3 ETH

	deposit, err := contractSessionActual.GetDeposit(receiverAcc)
	if err != nil {
		log.Fatal()
	}
	c.Assert(deposit.Int64(), Equals, value.Int64()) //Contribution should be correct

	dealID := big.NewInt(1)
	dealSender, err := contractSessionActual.GetDealSender(dealID)
	if err != nil {
		log.Fatal()
	}
	c.Assert(dealSender, Equals, senderAcc) //Sender should be correct

	dealReceiver, err := contractSessionActual.GetDealReceiver(dealID)
	if err != nil {
		log.Fatal()
	}
	c.Assert(dealReceiver, Equals, receiverAcc) //Receiver should be correct

	dealDeposit, err := contractSessionActual.GetDealDeposit(dealID)
	if err != nil {
		log.Fatal()
	}
	c.Assert(dealDeposit.Int64(), Equals, value.Int64()) //Deposit should be correct

	//Let's change contract caller to owner again and return the deposit to the receiver
	auth = changeAuth(*s, ownerInd)

	//Let's check receiver balance before returning deposit
	receiverBalanceInit, err := client.BalanceAt(context.Background(), receiverAcc, nil)
	if err != nil {
		log.Fatal()
	}

	//Let's return the deposit
	_, err = contractSessionActual.ReturnDeposit(dealID)
	if err != nil {
		log.Fatal()
	}

	//Let's check receiver balance after returning deposit
	receiverBalanceFinal, err := client.BalanceAt(context.Background(), receiverAcc, nil)
	if err != nil {
		log.Fatal()
	}

	balanceDifference := big.NewInt(0).Sub(receiverBalanceFinal, receiverBalanceInit)
	c.Assert(balanceDifference.Int64(), Equals, value.Int64()) //Balance change should be correct

}

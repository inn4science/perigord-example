package migrations

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/inn4science/perigord-example/bindings"
	"github.com/inn4science/perigord/contract"
	"github.com/inn4science/perigord/migration"
	"github.com/inn4science/perigord/network"
)

type MarketDeployer struct{}

func (d *MarketDeployer) Deploy(ctx context.Context, network *network.Network) (common.Address, *types.Transaction, interface{}, error) {
	account := network.Accounts()[0]
	err := network.Unlock(account)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	auth := network.NewTransactor(account)
	address, transaction, contract, err := bindings.DeployMarket(auth, network.Client())
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	session := &bindings.MarketSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return address, transaction, session, nil
}

func (d *MarketDeployer) Bind(ctx context.Context, network *network.Network, address common.Address) (interface{}, error) {
	account := network.Accounts()[0]
	err := network.Unlock(account)
	if err != nil {
		return common.Address{}, err
	}

	auth := network.NewTransactor(account)
	contract, err := bindings.NewMarket(address, network.Client())
	if err != nil {
		return nil, err
	}

	session := &bindings.MarketSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return session, nil
}

func init() {
	contract.AddContract("Market", &MarketDeployer{})

	migration.AddMigration(&migration.Migration{
		Number: 2,
		F: func(ctx context.Context, network *network.Network) error {
			if err := contract.Deploy(ctx, "Market", network); err != nil {
				return err
			}

			return nil
		},
	})
}

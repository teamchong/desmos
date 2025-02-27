package ibctesting

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	profilestypes "github.com/desmos-labs/desmos/v4/x/profiles/types"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
)

var (
	ChainIDPrefix   = "testchain"
	globalStartTime = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	TimeIncrement   = time.Second * 5
)

// Coordinator is a testing struct which contains N TestChain's. It handles keeping all chains
// in sync with regards to time.
type Coordinator struct {
	t *testing.T

	Chains map[string]*TestChain
}

// NewCoordinator initializes Coordinator with N TestChain's
func NewCoordinator(t *testing.T, n int) *Coordinator {
	chains := make(map[string]*TestChain)

	for i := 0; i < n; i++ {
		chainID := GetChainID(i + 1)
		chains[chainID] = NewTestChain(t, chainID)
	}
	return &Coordinator{
		t:      t,
		Chains: chains,
	}
}

// SetupClients is a helper function to create clients on both chains. It assumes the
// caller does not anticipate any errors.
func (coord *Coordinator) SetupClients(
	chainA, chainB *TestChain,
	clientType string,
) (string, string) {

	clientA, err := coord.CreateClient(chainA, chainB, clientType)
	require.NoError(coord.t, err)

	clientB, err := coord.CreateClient(chainB, chainA, clientType)
	require.NoError(coord.t, err)

	return clientA, clientB
}

// SetupClientConnections is a helper function to create clients and the appropriate
// connections on both the source and counterparty chain. It assumes the caller does not
// anticipate any errors.
func (coord *Coordinator) SetupClientConnections(
	chainA, chainB *TestChain,
	clientType string,
) (string, string, *TestConnection, *TestConnection) {

	clientA, clientB := coord.SetupClients(chainA, chainB, clientType)

	connA, connB := coord.CreateConnection(chainA, chainB, clientA, clientB)

	return clientA, clientB, connA, connB
}

// CreateClient creates a counterparty client on the source chain and returns the clientID.
func (coord *Coordinator) CreateClient(
	source, counterparty *TestChain,
	clientType string,
) (clientID string, err error) {
	coord.CommitBlock(source, counterparty)

	clientID = source.NewClientID(clientType)

	switch clientType {
	case exported.Tendermint:
		err = source.CreateTMClient(counterparty, clientID)

	default:
		err = fmt.Errorf("client type %s is not supported", clientType)
	}

	if err != nil {
		return "", err
	}

	coord.IncrementTime()

	return clientID, nil
}

// UpdateClient updates a counterparty client on the source chain.
func (coord *Coordinator) UpdateClient(
	source, counterparty *TestChain,
	clientID string,
	clientType string,
) (err error) {
	coord.CommitBlock(source, counterparty)

	switch clientType {
	case exported.Tendermint:
		err = source.UpdateTMClient(counterparty, clientID)

	default:
		err = fmt.Errorf("client type %s is not supported", clientType)
	}

	if err != nil {
		return err
	}

	coord.IncrementTime()

	return nil
}

// CreateConnection constructs and executes connection handshake messages in order to create
// OPEN channels on chainA and chainB. The connection information of for chainA and chainB
// are returned within a TestConnection struct. The function expects the connections to be
// successfully opened otherwise testing will fail.
func (coord *Coordinator) CreateConnection(
	chainA, chainB *TestChain,
	clientA, clientB string,
) (*TestConnection, *TestConnection) {

	connA, connB, err := coord.ConnOpenInit(chainA, chainB, clientA, clientB)
	require.NoError(coord.t, err)

	err = coord.ConnOpenTry(chainB, chainA, connB, connA)
	require.NoError(coord.t, err)

	err = coord.ConnOpenAck(chainA, chainB, connA, connB)
	require.NoError(coord.t, err)

	err = coord.ConnOpenConfirm(chainB, chainA, connB, connA)
	require.NoError(coord.t, err)

	return connA, connB
}

// CreateIBCProfilesChannels constructs and executes channel handshake messages to create OPEN
// ibc-profiles channel to profiles channel on chainA and chainB. The function expects the channels to be
// successfully opened otherwise testing will fail.
func (coord *Coordinator) CreateIBCProfilesChannels(
	chainA, chainB *TestChain,
	connA, connB *TestConnection,
	order channeltypes.Order,
) (TestChannel, TestChannel) {
	return coord.CreateChannel(chainA, chainB, connA, connB, profilestypes.IBCPortID, profilestypes.IBCPortID, order)
}

// CreateChannel constructs and executes channel handshake messages in order to create
// OPEN channels on chainA and chainB. The function expects the channels to be successfully
// opened otherwise testing will fail.
func (coord *Coordinator) CreateChannel(
	chainA, chainB *TestChain,
	connA, connB *TestConnection,
	sourcePortID, counterpartyPortID string,
	order channeltypes.Order,
) (TestChannel, TestChannel) {

	channelA, channelB, err := coord.ChanOpenInit(chainA, chainB, connA, connB, sourcePortID, counterpartyPortID, order)
	require.NoError(coord.t, err)

	err = coord.ChanOpenTry(chainB, chainA, channelB, channelA, connB, order)
	require.NoError(coord.t, err)

	err = coord.ChanOpenAck(chainA, chainB, channelA, channelB)
	require.NoError(coord.t, err)

	err = coord.ChanOpenConfirm(chainB, chainA, channelB, channelA)
	require.NoError(coord.t, err)

	return channelA, channelB
}

// IncrementTime iterates through all the TestChain's and increments their current header time
// by 5 seconds.
//
// CONTRACT: this function must be called after every commit on any TestChain.
func (coord *Coordinator) IncrementTime() {
	for _, chain := range coord.Chains {
		chain.CurrentHeader.Time = chain.CurrentHeader.Time.Add(TimeIncrement)
		chain.App.BeginBlock(abci.RequestBeginBlock{Header: chain.CurrentHeader})
	}
}

// SendMsg delivers a single provided message to the chain. The counterparty
// client is update with the new source consensus state.
func (coord *Coordinator) SendMsg(source, counterparty *TestChain, counterpartyClientID string, msg sdk.Msg) error {
	return coord.SendMsgs(source, counterparty, counterpartyClientID, []sdk.Msg{msg})
}

// SendMsgs delivers the provided messages to the chain. The counterparty
// client is updated with the new source consensus state.
func (coord *Coordinator) SendMsgs(source, counterparty *TestChain, counterpartyClientID string, msgs []sdk.Msg) error {
	if err := source.sendMsgs(msgs...); err != nil {
		return err
	}

	coord.IncrementTime()

	// update source client on counterparty connection
	return coord.UpdateClient(
		counterparty, source,
		counterpartyClientID, exported.Tendermint,
	)
}

// GetChain returns the TestChain using the given chainID and returns an error if it does
// not exist.
func (coord *Coordinator) GetChain(chainID string) *TestChain {
	chain, found := coord.Chains[chainID]
	require.True(coord.t, found, fmt.Sprintf("%s chain does not exist", chainID))
	return chain
}

// GetChainID returns the chainID used for the provided index.
func GetChainID(index int) string {
	return ChainIDPrefix + strconv.Itoa(index)
}

// CommitBlock commits a block on the provided indexes and then increments the global time.
//
// CONTRACT: the passed in list of indexes must not contain duplicates
func (coord *Coordinator) CommitBlock(chains ...*TestChain) {
	for _, chain := range chains {
		chain.App.Commit()
		chain.NextBlock()
	}
	coord.IncrementTime()
}

// ConnOpenInit initializes a connection on the source chain with the state INIT
// using the OpenInit handshake call.
//
// NOTE: The counterparty testing connection will be created even if it is not created in the
// application state.
func (coord *Coordinator) ConnOpenInit(
	source, counterparty *TestChain,
	clientID, counterpartyClientID string,
) (*TestConnection, *TestConnection, error) {
	sourceConnection := source.AddTestConnection(clientID, counterpartyClientID)
	counterpartyConnection := counterparty.AddTestConnection(counterpartyClientID, clientID)

	// initialize connection on source
	if err := source.ConnectionOpenInit(counterparty, sourceConnection, counterpartyConnection); err != nil {
		return sourceConnection, counterpartyConnection, err
	}
	coord.IncrementTime()

	// update source client on counterparty connection
	if err := coord.UpdateClient(
		counterparty, source,
		counterpartyClientID, exported.Tendermint,
	); err != nil {
		return sourceConnection, counterpartyConnection, err
	}

	return sourceConnection, counterpartyConnection, nil
}

// ConnOpenTry initializes a connection on the source chain with the state TRYOPEN
// using the OpenTry handshake call.
func (coord *Coordinator) ConnOpenTry(
	source, counterparty *TestChain,
	sourceConnection, counterpartyConnection *TestConnection,
) error {
	// initialize TRYOPEN connection on source
	if err := source.ConnectionOpenTry(counterparty, sourceConnection, counterpartyConnection); err != nil {
		return err
	}
	coord.IncrementTime()

	// update source client on counterparty connection
	return coord.UpdateClient(
		counterparty, source,
		counterpartyConnection.ClientID, exported.Tendermint,
	)
}

// ConnOpenAck initializes a connection on the source chain with the state OPEN
// using the OpenAck handshake call.
func (coord *Coordinator) ConnOpenAck(
	source, counterparty *TestChain,
	sourceConnection, counterpartyConnection *TestConnection,
) error {
	// set OPEN connection on source using OpenAck
	if err := source.ConnectionOpenAck(counterparty, sourceConnection, counterpartyConnection); err != nil {
		return err
	}
	coord.IncrementTime()

	// update source client on counterparty connection
	return coord.UpdateClient(
		counterparty, source,
		counterpartyConnection.ClientID, exported.Tendermint,
	)
}

// ConnOpenConfirm initializes a connection on the source chain with the state OPEN
// using the OpenConfirm handshake call.
func (coord *Coordinator) ConnOpenConfirm(
	source, counterparty *TestChain,
	sourceConnection, counterpartyConnection *TestConnection,
) error {
	if err := source.ConnectionOpenConfirm(counterparty, sourceConnection, counterpartyConnection); err != nil {
		return err
	}
	coord.IncrementTime()

	// update source client on counterparty connection
	return coord.UpdateClient(
		counterparty, source,
		counterpartyConnection.ClientID, exported.Tendermint,
	)
}

// ChanOpenInit initializes a channel on the source chain with the state INIT
// using the OpenInit handshake call.
//
// NOTE: The counterparty testing channel will be created even if it is not created in the
// application state.
func (coord *Coordinator) ChanOpenInit(
	source, counterparty *TestChain,
	connection, counterpartyConnection *TestConnection,
	sourcePortID, counterpartyPortID string,
	order channeltypes.Order,
) (TestChannel, TestChannel, error) {
	sourceChannel := source.AddTestChannel(connection, sourcePortID)
	counterpartyChannel := counterparty.AddTestChannel(counterpartyConnection, counterpartyPortID)

	// NOTE: only creation of a capability for a transfer or mock port is supported
	// Other applications must bind to the port in InitGenesis or modify this code.
	source.CreatePortCapability(sourceChannel.PortID)
	coord.IncrementTime()

	// initialize channel on source
	if err := source.ChanOpenInit(sourceChannel, counterpartyChannel, order, connection.ID); err != nil {
		return sourceChannel, counterpartyChannel, err
	}
	coord.IncrementTime()

	// update source client on counterparty connection
	if err := coord.UpdateClient(
		counterparty, source,
		counterpartyConnection.ClientID, exported.Tendermint,
	); err != nil {
		return sourceChannel, counterpartyChannel, err
	}

	return sourceChannel, counterpartyChannel, nil
}

// ChanOpenTry initializes a channel on the source chain with the state TRYOPEN
// using the OpenTry handshake call.
func (coord *Coordinator) ChanOpenTry(
	source, counterparty *TestChain,
	sourceChannel, counterpartyChannel TestChannel,
	connection *TestConnection,
	order channeltypes.Order,
) error {

	// initialize channel on source
	if err := source.ChanOpenTry(counterparty, sourceChannel, counterpartyChannel, order, connection.ID); err != nil {
		return err
	}
	coord.IncrementTime()

	// update source client on counterparty connection
	return coord.UpdateClient(
		counterparty, source,
		connection.CounterpartyClientID, exported.Tendermint,
	)
}

// ChanOpenAck initializes a channel on the source chain with the state OPEN
// using the OpenAck handshake call.
func (coord *Coordinator) ChanOpenAck(
	source, counterparty *TestChain,
	sourceChannel, counterpartyChannel TestChannel,
) error {

	if err := source.ChanOpenAck(counterparty, sourceChannel, counterpartyChannel); err != nil {
		return err
	}
	coord.IncrementTime()

	// update source client on counterparty connection
	return coord.UpdateClient(
		counterparty, source,
		sourceChannel.CounterpartyClientID, exported.Tendermint,
	)
}

// ChanOpenConfirm initializes a channel on the source chain with the state OPEN
// using the OpenConfirm handshake call.
func (coord *Coordinator) ChanOpenConfirm(
	source, counterparty *TestChain,
	sourceChannel, counterpartyChannel TestChannel,
) error {

	if err := source.ChanOpenConfirm(counterparty, sourceChannel, counterpartyChannel); err != nil {
		return err
	}
	coord.IncrementTime()

	// update source client on counterparty connection
	return coord.UpdateClient(
		counterparty, source,
		sourceChannel.CounterpartyClientID, exported.Tendermint,
	)
}

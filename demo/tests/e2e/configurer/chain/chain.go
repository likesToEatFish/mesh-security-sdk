package chain

import (
	"fmt"
	"testing"
	"time"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/stretchr/testify/require"

	"github.com/osmosis-labs/mesh-security-sdk/demo/tests/e2e/configurer/config"
	"github.com/osmosis-labs/mesh-security-sdk/demo/tests/e2e/containers"
	"github.com/osmosis-labs/mesh-security-sdk/demo/tests/e2e/initialization"
)

type Config struct {
	initialization.ChainMeta

	ValidatorInitConfigs []*initialization.NodeConfig
	// voting period is number of blocks it takes to deposit, 1.2 seconds per validator to vote on the prop, and a buffer.
	VotingPeriod          float32
	ExpeditedVotingPeriod float32
	// upgrade proposal height for chain.
	UpgradePropHeight    int64
	LatestProposalNumber int
	LatestLockNumber     int
	NodeConfigs          []*NodeConfig

	LatestCodeID int

	t                *testing.T
	containerManager *containers.Manager
}

const (
	// defaultNodeIndex to use for querying and executing transactions.
	// It is used when we are indifferent about the node we are working with.
	defaultNodeIndex = 0
	// waitUntilRepeatPauseTime is the time to wait between each check of the node status.
	waitUntilRepeatPauseTime = 2 * time.Second
	// waitUntilrepeatMax is the maximum number of times to repeat the wait until condition.
	waitUntilrepeatMax = 60
)

func New(t *testing.T, containerManager *containers.Manager, id string, initValidatorConfigs []*initialization.NodeConfig) *Config {
	numVal := float32(len(initValidatorConfigs))
	return &Config{
		ChainMeta: initialization.ChainMeta{
			ID: id,
		},
		ValidatorInitConfigs:  initValidatorConfigs,
		VotingPeriod:          config.PropDepositBlocks + numVal*config.PropVoteBlocks + config.PropBufferBlocks,
		ExpeditedVotingPeriod: config.PropDepositBlocks + numVal*config.PropVoteBlocks + config.PropBufferBlocks - 2,
		t:                     t,
		containerManager:      containerManager,
	}
}

// CreateNode returns new initialized NodeConfig.
func (c *Config) CreateNode(initNode *initialization.Node) *NodeConfig {
	nodeConfig := &NodeConfig{
		Node:             *initNode,
		chainID:          c.ID,
		containerManager: c.containerManager,
		t:                c.t,
	}
	c.NodeConfigs = append(c.NodeConfigs, nodeConfig)
	return nodeConfig
}

// RemoveNode removes node and stops it from running.
func (c *Config) RemoveNode(nodeName string) error {
	for i, node := range c.NodeConfigs {
		if node.Name == nodeName {
			c.NodeConfigs = append(c.NodeConfigs[:i], c.NodeConfigs[i+1:]...)
			return node.Stop()
		}
	}
	return fmt.Errorf("node %s not found", nodeName)
}

// WaitUntilHeight waits for all validators to reach the specified height at the minimum.
// returns error, if any.
func (c *Config) WaitUntilHeight(height int64) {
	// Ensure the nodes are making progress.
	doneCondition := func(syncInfo coretypes.SyncInfo) bool {
		curHeight := syncInfo.LatestBlockHeight

		if curHeight < height {
			c.t.Logf("current block height is %d, waiting to reach: %d", curHeight, height)
			return false
		}

		return !syncInfo.CatchingUp
	}

	for _, node := range c.NodeConfigs {
		c.t.Logf("node container: %s, waiting to reach height %d", node.Name, height)
		node.WaitUntil(doneCondition)
	}
}

// WaitForNumHeights waits for all nodes to go through a given number of heights.
func (c *Config) WaitForNumHeights(heightsToWait int64) {
	node, err := c.GetDefaultNode()
	require.NoError(c.t, err)
	currentHeight, err := node.QueryCurrentHeight()
	require.NoError(c.t, err)
	c.WaitUntilHeight(currentHeight + heightsToWait)
}

func (c *Config) GetDefaultNode() (*NodeConfig, error) {
	return c.getNodeAtIndex(defaultNodeIndex)
}

// GetPersistentPeers returns persistent peers from every node
// associated with a chain.
func (c *Config) GetPersistentPeers() []string {
	peers := make([]string, len(c.NodeConfigs))
	for i, node := range c.NodeConfigs {
		peers[i] = node.PeerID
	}
	return peers
}

func (c *Config) getNodeAtIndex(nodeIndex int) (*NodeConfig, error) {
	if nodeIndex > len(c.NodeConfigs) {
		return nil, fmt.Errorf("node index (%d) is greter than the number of nodes available (%d)", nodeIndex, len(c.NodeConfigs))
	}
	return c.NodeConfigs[nodeIndex], nil
}

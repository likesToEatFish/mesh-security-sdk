package containers

// ImageConfig contains all images and their respective tags
// needed for running e2e tests.
type ImageConfig struct {
	InitRepository string
	InitTag        string

	MeshRepository string
	MeshTag        string

	RelayerRepository string
	RelayerTag        string
}

//nolint:deadcode
const (
	// Current Git branch mesh repo/version. It is meant to be built locally.
	// This image should be pre-built with `make docker-build-debug` either in CI or locally.
	CurrentBranchMeshRepository = "mesh"
	CurrentBranchMeshTag        = "debug"
	// Hermes repo/version for relayer
	relayerRepository = "informalsystems/hermes"
	relayerTag        = "1.5.1"
)

// Returns ImageConfig needed for running e2e test.
func NewImageConfig() ImageConfig {
	config := ImageConfig{
		RelayerRepository: relayerRepository,
		RelayerTag:        relayerTag,
	}

	config.MeshRepository = CurrentBranchMeshRepository
	config.MeshTag = CurrentBranchMeshTag
	return config
}

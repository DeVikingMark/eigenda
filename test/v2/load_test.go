package v2

import (
	"github.com/Layr-Labs/eigenda/common/testutils/random"
	"github.com/docker/go-units"
	"os"
	"testing"
)

func TestLightLoad(t *testing.T) {
	rand := random.NewTestRandom(t)
	c := getClient(t)

	config := DefaultLoadGeneratorConfig()
	config.AverageBlobSize = 100 * units.KiB
	config.BlobSizeStdDev = 50 * units.KiB
	config.BytesPerSecond = 100 * units.KiB

	generator := NewLoadGenerator(config, c, rand)

	signals := make(chan os.Signal)
	go func() {
		<-signals
		generator.Stop()
	}()

	generator.Start(true)
}

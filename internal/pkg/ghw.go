package pkg

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
)
import "github.com/jaypipes/ghw"

func GetTotalDiskSize() (string, error) {
	block, err := ghw.Block()
	if err != nil {
		return "", errors.Wrap(err, "could not retrieve disk size")
	}

	return prettyByteSize(block.TotalPhysicalBytes), nil
}

func prettyByteSize(b uint64) string {
	bf := float64(b)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.1fYiB", bf)
}

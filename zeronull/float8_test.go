package zeronull_test

import (
	"testing"

	"github.com/getoutreach/pgtype/testutil"
	"github.com/getoutreach/pgtype/zeronull"
)

func TestFloat8Transcode(t *testing.T) {
	testutil.TestSuccessfulTranscode(t, "float8", []interface{}{
		(zeronull.Float8)(1),
		(zeronull.Float8)(0),
	})
}

func TestFloat8ConvertsGoZeroToNull(t *testing.T) {
	testutil.TestGoZeroToNullConversion(t, "float8", (zeronull.Float8)(0))
}

func TestFloat8ConvertsNullToGoZero(t *testing.T) {
	testutil.TestNullToGoZeroConversion(t, "float8", (zeronull.Float8)(0))
}

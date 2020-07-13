package purescript_quickcheck

import (
	"bytes"

	"github.com/lunixbochs/struc"
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Test.QuickCheck.Gen")

	exports["float32ToInt32"] = func(x_ Any) Any {
		x, _ := x_.(float32)
		var buf bytes.Buffer
		struc.Pack(&buf, x)
		var i int32
		struc.Unpack(&buf, &i)
		return int(i)
	}
}

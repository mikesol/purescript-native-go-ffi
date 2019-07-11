package purescript_console

import (
	"Effect_Console"
	"fmt"
	. "purescript"
)

func init() {
	exports := Effect_Console.Foreign

	exports["log"] = func(s Any) Any {
		return func() Any {
			fmt.Println(s)
			return nil
		}
	}

	exports["warn"] = func(s Any) Any {
		return func() Any {
			fmt.Println(s)
			return nil
		}
	}
}
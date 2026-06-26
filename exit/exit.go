package exit

import (
	c "flatpkg/colorize"
	"fmt"
	"os"
)

func ErrorExit(err string) {
	fmt.Fprintln(os.Stderr, c.Red.Fprint("[Error]"), err)
	os.Exit(1)
}

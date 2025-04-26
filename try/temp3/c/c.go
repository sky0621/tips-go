package c

import (
	"fmt"
	"github.com/sky0621/tips-go/try/temp3/util"
)

func Cc() error {
	return fmt.Errorf("[%s] failed to call Cc", util.FileWithLineNum())
}

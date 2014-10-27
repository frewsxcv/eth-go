// +build !windows !cgo

package ethutil

import (
	"fmt"
	"strings"

	"github.com/obscuren/mutan"
	"github.com/obscuren/mutan/backends"
	"github.com/obscuren/serpent-go"
)

// General compile function
func Compile(script string, silent bool) (ret []byte, err error) {
	if len(script) > 2 {
		line := strings.Split(script, "\n")[0]

		if len(line) > 1 && line[0:2] == "#!" {
			switch line {
			case "#!serpent":
				byteCode, err := serpent.Compile(script)
				if err != nil {
					return nil, err
				}

				return byteCode, nil
            case "#!hex":
		        hex := strings.Split(script, "\n")[1]
                byteCode := make([]byte, len(hex) / 2)
                var top uint8
                var bottom uint8
                for i := 0; i < len(hex) / 2; i++ {
                    if hex[i*2] >= 96 {
                         top = hex[i*2] - 87 
                    } else { top = hex[i*2] - 48 }
                    if hex[i*2+1] >= 96 { 
                         bottom = hex[i*2+1] - 87
                    } else { bottom = hex[i*2+1] - 48 }
                    byteCode[i] = (top * 16 + bottom)
                }
                return byteCode, nil
			}
		} else {

			compiler := mutan.NewCompiler(backend.NewEthereumBackend())
			compiler.Silent = silent
			byteCode, errors := compiler.Compile(strings.NewReader(script))
			if len(errors) > 0 {
				var errs string
				for _, er := range errors {
					if er != nil {
						errs += er.Error()
					}
				}
				return nil, fmt.Errorf("%v", errs)
			}

			return byteCode, nil
		}
	}

	return nil, nil
}

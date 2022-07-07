package versions

import (
	"github.com/quzhi1/fiber-versioning-tool/lib"
)

var Version1_0 = lib.VersionDef{
	Version:              "1.0",
	RequestBodyChange:    lib.ByteNoop,
	ResponseBodyChange:   lib.ByteNoop,
	QueryParamChange:     lib.MapNoop,
	RequestHeaderChange:  lib.MapNoop,
	ResponseHeaderChange: lib.MapNoop,
}

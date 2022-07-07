package lib

type VersionDef struct {
	Version              string
	RequestBodyChange    func([]byte) ([]byte, error)
	ResponseBodyChange   func([]byte) ([]byte, error)
	QueryParamChange     func(map[string]string) (map[string]string, error)
	RequestHeaderChange  func(map[string]string) (map[string]string, error)
	ResponseHeaderChange func(map[string]string) (map[string]string, error)
}

func ByteNoop(input []byte) ([]byte, error) {
	return input, nil
}

func MapNoop(input map[string]string) (map[string]string, error) {
	return input, nil
}

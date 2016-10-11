package proxy

type Proxy interface {
	Generate([]byte) ([]byte, error)
	Revert([]byte) ([]byte, error)
}

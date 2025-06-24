package init_chain

func CreateInitChainFromSpec(specPath string) InitChain {
	return &initChain{
		specPath: specPath,
	}
}

type InitChain interface {
	Start() error
}

type InitChainNode interface {
	RunNode() error
}

type initChain struct {
	specPath string
}

func (ic *initChain) Start() error {
	return nil
}

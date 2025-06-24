package init_chain

func NewInitChain(specPath string) InitChain {
	return &initChain{
		specPath: specPath,
	}
}

type InitChain interface {
	Start() error
}

type initChain struct {
	specPath string
}

func (ic *initChain) Start() error {
	return nil
}

type InitChainNode interface {
	RunNode() error
}

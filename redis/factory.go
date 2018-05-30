package redis

type Factory struct {
	Pool *Pool
}

func (rf *Factory) HStorage(setName string) *HStorage {
	return &HStorage{Pool: rf.Pool, SetName: setName}
}

func NewFactory(p *Pool) *Factory {
	return &Factory{p}
}

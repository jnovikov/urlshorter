package url

import (
	"urlshorter/keystorage"
	"urlshorter/random"
)

type RandomGenerator struct {
	shortStorage *keystorage.StorageWrapper
	urlStorage   *keystorage.StorageWrapper
	urlSize      int
}

func (rg *RandomGenerator) New(url string) string {
	short := rg.urlStorage.GetString(url)
	if short != "" {
		return short
	}

	short = random.String(rg.urlSize)
	for rg.shortStorage.Exists(short) {
		short = random.String(rg.urlSize)
	}
	rg.shortStorage.Set(short, url)
	rg.urlStorage.Set(url, short)
	return short
}

func (rg *RandomGenerator) Find(short string) string {
	url := rg.shortStorage.GetString(short)
	return url
}

func (rg *RandomGenerator) SetShortStorage(st keystorage.KeyStorage) *RandomGenerator {
	rg.shortStorage = keystorage.NewWrapper(st)
	return rg
}

func (rg *RandomGenerator) SetUrlStorage(st keystorage.KeyStorage) *RandomGenerator {
	rg.urlStorage = keystorage.NewWrapper(st)
	return rg
}

func (rg *RandomGenerator) SetUrlSize(n int) *RandomGenerator {
	rg.urlSize = n
	return rg
}

func NewRandomGenerator() *RandomGenerator {
	return new(RandomGenerator)
}

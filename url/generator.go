package url

type Generator interface {
	New(url string) string
	Find(short string) string
}

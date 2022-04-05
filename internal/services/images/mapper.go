package images

type Mapper interface {
	Map(word string) (string, error)
}

type mapper struct {
	seeds map[string][]string
}

func (m mapper) Map(word string) (string, error) {
	panic("implement me")
}

func NewMapper() Mapper{

	return mapper{}
}
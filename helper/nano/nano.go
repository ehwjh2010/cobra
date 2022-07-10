package nano

import gonanoid "github.com/matoous/go-nanoid/v2"

func GetNanoId() (string, error) {
	return gonanoid.New()
}

func MustGetNanoId() string {
	return gonanoid.Must()
}

func GetNonaIdWithCustom(alphabet string, size int) (string, error) {
	return gonanoid.Generate(alphabet, size)
}

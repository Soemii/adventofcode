package adventofcode

type Challenge interface {
	GetYear() int
	GetDay() int
	GetChallenge() int

	Execute(rawFile string) error
}

package adventofcode

type Challenge interface {
	GetYear() int
	GetDay() int

	ExecuteFirst(input string) (string, error)
	ExecuteSecond(input string) (string, error)
}

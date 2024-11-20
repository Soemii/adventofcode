package adventofcode

type Challenge[prepareResultFirst any, prepareResultSecond any] interface {
	GetYear() int
	GetDay() int

	PrepareFirst(rawfile string) prepareResultFirst
	PrepareSecond(rawfile string) prepareResultSecond

	ExecuteFirst(input prepareResultFirst) (string, error)
	ExecuteSecond(input prepareResultSecond) (string, error)
}

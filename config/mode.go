package config

type EncryptMode int

const (
	Row        EncryptMode = 1
	RowAndLine EncryptMode = 2
)

type EncryptMethodType int

const (
	Logistic EncryptMethodType = 1
	//Sigmoid  EncryptMethodType = 2
)

type EncryptMethod func(x float64) float64

var (
	LogisticFunc EncryptMethod = func(x float64) float64 { return 1 - 2*x*x }
	//SigmoidFunc  EncryptMethod = func(x float64) float64 { return 1 / (1 + math.Exp(-x)) }
)

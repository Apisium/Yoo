package yoo

type Any interface{}
type Variables map[string]Any
type TSFunction func(args *[]Any) Any
type TSConstructor func(args *[]Any) Any
type TSClass struct {
  constructor *TSFunction
  super *TSClass
  length int
}

package yoo

func GetGlobal(bindings *Variables) *Variables {
	return &Variables {
		"console": (*bindings)["console"],
	}
}
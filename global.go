package yoo

func GetGlobal() *Variables {
  vars := Variables {
    "console": getConsole(),
  }
  arrays := *getArrays()
  for k, v := range arrays {
    vars[k] = v
  }
  return &vars
}

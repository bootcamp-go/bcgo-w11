package people

func NewHuman(name string, age int) human {
	// default values
	defaultName := "John"
	defaultAge := 42
	if name == "" {
		defaultName = name
	}
	if age == 0 {
		defaultAge = age
	}

	return human{
		Name: defaultName,
		Age:  defaultAge,
	}
}

type human struct {
	Name string `json:"name"` // `key:"value" key2:"value2" keyComplex:"value1,value2,value3"`
	Age  int    `json:"age"`
}

// human struct -> json (human{"name": "John", "age": 42})
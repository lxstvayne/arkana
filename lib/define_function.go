package lib

type DefineFunction struct {
	argsNames []string
	body      Statement
}

func NewDefineFunction(argNames []string, body Statement) *DefineFunction {
	return &DefineFunction{
		argsNames: argNames,
		body:      body,
	}
}

func (df *DefineFunction) Execute(args ...Value) Value {
	// EXECUTE and check return statement
	df.body.Execute()
	return NewNumberValue(0)
}

func (df *DefineFunction) ArgsCount() int {
	return len(df.argsNames)
}

func (df *DefineFunction) ArgName(idx int) string {
	if idx < 0 || idx >= df.ArgsCount() {
		return ""
	}

	return df.argsNames[idx]
}

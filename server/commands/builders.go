package commands

const (
	StringSingleWord = iota
	StringQuotablePhrase
	StringGreedyPhrase
)

const (
	EntitySingle = iota + 1
	EntityPlayerOnly
)

type argumentType int

const (
	Bool argumentType = iota
	Float
	Double
	Integer
	Long
	String
)

func NewArgument(name string, t argumentType) Argument {
	return Argument{
		Name:   name,
		Parser: Parser{ID: int32(t)},
	}
}

func (a Argument) Min(min uint64) Argument {
	//todo add checks, dont allow arguments that arent numbers to access this
	a.Parser.Properties.Flags |= 1
	a.Parser.Properties.Min = min
	return a
}

func (a Argument) Max(max uint64) Argument {
	//todo add checks, dont allow arguments that arent numbers to access this
	a.Parser.Properties.Flags |= 2
	a.Parser.Properties.Max = max
	return a
}

func NewBoolArgument(name string) Argument {
	return Argument{
		Name: name,
		Parser: Parser{
			ID: 0,
		},
	}
}

func NewFloatArgument(name string, properties struct {
	Min *uint64
	Max *uint64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = *properties.Min
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = *properties.Max
	}
	return Argument{
		Name: name,
		Parser: Parser{
			ID:         1,
			Properties: props,
		},
	}
}

func NewFloatArgument2(name string, min, max float32) Argument {
	props := Properties{Flags: 0x02}

	return Argument{
		Name: name,
		Parser: Parser{
			ID:         1,
			Properties: props,
		},
	}
}

func NewDoubleArgument(name string, properties struct {
	Min *uint64
	Max *uint64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = *properties.Min
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = *properties.Max
	}
	return Argument{
		Name: name,
		Parser: Parser{
			ID:         2,
			Properties: props,
		},
	}
}

func NewIntegerArgument(name string, properties struct {
	Min *int64
	Max *int64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = uint64(*properties.Min)
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = uint64(*properties.Max)
	}
	return Argument{
		Name: name,
		Parser: Parser{
			ID:         3,
			Properties: props,
		},
	}
}

func NewLongArgument(name string, properties struct {
	Min *int64
	Max *int64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = uint64(*properties.Min)
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = uint64(*properties.Max)
	}
	return Argument{
		Name: name,
		Parser: Parser{
			ID:         4,
			Properties: props,
		},
	}
}

func NewStringArgument(name string, properties byte) Argument {
	props := Properties{Flags: properties}
	return Argument{
		Name: name,
		Parser: Parser{
			ID:         5,
			Properties: props,
		},
	}
}

func NewEntityArgument(name string, properties byte) Argument {
	props := Properties{Flags: properties}
	return Argument{
		Name: name,
		Parser: Parser{
			ID:         6,
			Properties: props,
		},
	}
}

func NewGamemodeArgument(name string) Argument {
	return Argument{
		Name: name,
		Parser: Parser{
			ID: 39,
		},
	}
}

func NewChatComponentArgument(name string) Argument {
	return Argument{
		Name: name,
		Parser: Parser{
			ID: 17,
		},
	}
}

func NewLiteralArgument(name string) Argument {
	return Argument{
		Name:           name,
		SuggestionType: "",
	}
}

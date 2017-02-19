package datatype

var lengthTypes = []*DataType{
	&DataType{
		Group:       GroupLength,
		Names:       []string{"angstrom"},
		DisplayName: "angstrom",
		Factor:      10000000000,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"centimeter", "cm"},
		DisplayName: "cm",
		Factor:      100,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"chain"},
		DisplayName: "chain",
		Factor:      0.049709695378987,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"decimeter", "dm"},
		DisplayName: "dm",
		Factor:      10,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"foot", "ft"},
		DisplayName: "ft",
		Factor:      0.54680664916885,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"fathom"},
		DisplayName: "fathom",
		Factor:      3.2808398950131,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"furlong"},
		DisplayName: "furlong",
		Factor:      0.0049709695378987,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"inch", "in"},
		DisplayName: "in",
		Factor:      39.370078740157,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"kilometer", "km"},
		DisplayName: "km",
		Factor:      0.001,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"league"},
		DisplayName: "league",
		Factor:      0.00020712373074577,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"meter", "m"},
		DisplayName: "m",
		Factor:      1,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"mile", "mi"},
		DisplayName: "mi",
		Factor:      0.00062137119223733,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"millimeter", "mm"},
		DisplayName: "mm",
		Factor:      1000,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"micrometer", "µm"},
		DisplayName: "µm",
		Factor:      1000000,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"micron", "µ"},
		DisplayName: "µ",
		Factor:      1000000,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"nanometer", "nm"},
		DisplayName: "nm",
		Factor:      1000000000,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"parsec"},
		DisplayName: "parsec",
		Factor:      3.2407792896393E-17,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"rod"},
		DisplayName: "rod",
		Factor:      0.19883878151595,
	},
	&DataType{
		Group:       GroupLength,
		Names:       []string{"yard", "yd"},
		DisplayName: "yd",
		Factor:      1.0936132983377,
	},
}

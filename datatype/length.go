package datatype

var lengthTypes = []DataType{
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"angstrom", "angstroms"},
			DisplayName: "angstrom",
		},
		Factor: 10000000000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"centimeter", "centimetre", "centimeters", "centimetres", "cm"},
			DisplayName: "cm",
		},
		Factor: 100,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"chain", "chains"},
			DisplayName: "chain",
		},
		Factor: 0.049709695378987,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"decimeter", "decimetre", "decimeters", "decimetres", "dm"},
			DisplayName: "dm",
		},
		Factor: 10,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"foot", "feet", "ft"},
			DisplayName: "ft",
		},
		Factor: 0.54680664916885,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"fathom", "fathoms"},
			DisplayName: "fathom",
		},
		Factor: 3.2808398950131,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"furlong", "furlongs"},
			DisplayName: "furlong",
		},
		Factor: 0.0049709695378987,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"inch", "inches", "in"},
			DisplayName: "in",
		},
		Factor: 39.370078740157,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"kilometer", "kilometre", "kilometers", "kilometres", "km"},
			DisplayName: "km",
		},
		Factor: 0.001,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"league", "leagues"},
			DisplayName: "league",
		},
		Factor: 0.00020712373074577,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"light year", "light years"},
			DisplayName: "l.y.",
		},
		Factor: 1.057001218773e-16,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"meter", "metre", "meters", "metres", "m"},
			DisplayName: "m",
		},
		Factor: 1,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"mile", "miles", "mi"},
			DisplayName: "mi",
		},
		Factor: 0.00062137119223733,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"millimeter", "millimetre", "millimeters", "millimetres", "mm"},
			DisplayName: "mm",
		},
		Factor: 1000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"micrometer", "micrometre", "micrometers", "micrometres", "µm"},
			DisplayName: "µm",
		},
		Factor: 1000000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"micron", "microns", "µ"},
			DisplayName: "µ",
		},
		Factor: 1000000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"nanometer", "nanometre", "nanometers", "nanometres", "nm"},
			DisplayName: "nm",
		},
		Factor: 1000000000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"nautical mile", "n.m.", "nautical miles"},
			DisplayName: "n.m.",
		},
		Factor: 0.000539957,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"parsec", "parsecs"},
			DisplayName: "parsec",
		},
		Factor: 3.2407792896393E-17,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"rod", "rods"},
			DisplayName: "rod",
		},
		Factor: 0.19883878151595,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupLength,
			Names:       []string{"yard", "yards", "yd"},
			DisplayName: "yd",
		},
		Factor: 1.0936132983377,
	},
}

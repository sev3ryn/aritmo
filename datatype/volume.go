package datatype

var volumeTypes = []DataType{
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"barrel", "bbl"},
			DisplayName: "bbl",
		},
		Factor: 6.2898107704321,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"cubic centimeter", "cm³"},
			DisplayName: "cm³",
		},
		Factor: 1000000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"cubic decimeter", "cubic decimeters", "cubic decimetre", "cubic decimetres", "dm³"},
			DisplayName: "dm³",
		},
		Factor: 1000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"cubic foot", "cubic feet", "ft³"},
			DisplayName: "ft³",
		},
		Factor: 35.314666721489,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"cubic inch", "cubic inches", "in³"},
			DisplayName: "in³",
		},
		Factor: 61023.744094732,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"cubic meter", "cubic metre", "cubic meters", "cubic metres", "m³"},
			DisplayName: "m³",
		},
		Factor: 1,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"cubic millimeter", "cubic millimeters", "cubic millimetre", "cubic millimetres", "mm³"},
			DisplayName: "mm³",
		},
		Factor: 1000000000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"cubic yard", "cubic yards", "yd³"},
			DisplayName: "yd³",
		},
		Factor: 1.3079506193144,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"centiliter", "centilitre", "centiliters", "centilitres", "cl"},
			DisplayName: "cl",
		},
		Factor: 100000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"dekaliter", "dekalitre", "dekaliters", "dekalitres", "dal"},
			DisplayName: "dal",
		},
		Factor: 100,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"hectoliter", "hectolitre", "hectoliters", "hectolitres", "hl"},
			DisplayName: "hl",
		},
		Factor: 10,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"kiloliter", "kilolitre", "kiloliters", "kilolitres", "kl"},
			DisplayName: "kl",
		},
		Factor: 1,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"liter", "litre", "liters", "litres", "l"},
			DisplayName: "l",
		},
		Factor: 1000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"microliter", "microlitre", "microliters", "microlitres", "µl"},
			DisplayName: "µl",
		},
		Factor: 1000000000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"milliliter", "millilitre", "milliliters", "millilitres", "ml"},
			DisplayName: "ml",
		},
		Factor: 1000000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"fluid dram", "fluid drams", "fl dr"},
			DisplayName: "fl dr",
		},
		Factor: 270512.18161474,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"fluid ounce", "fluid drams", "fl oz"},
			DisplayName: "fl oz",
		},
		Factor: 33814.022701843,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"gallon", "gallons", "gal"},
			DisplayName: "gal",
		},
		Factor: 264.17205235815,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"gill", "gills"},
			DisplayName: "gill",
		},
		Factor: 8453.5056754608,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"minim", "min"},
			DisplayName: "min",
		},
		Factor: 16230730.896885,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"peck", "pecks", "pk"},
			DisplayName: "pk",
		},
		Factor: 113.51037303361,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"pint", "pints", "pt"},
			DisplayName: "pt",
		},
		Factor: 2113.3764188652,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupVolume,
			Names:       []string{"quart", "qt"},
			DisplayName: "qt",
		},
		Factor: 1056.6882094326,
	},
}

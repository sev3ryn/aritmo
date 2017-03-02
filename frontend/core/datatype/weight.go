package datatype

var weightTypes = []DataType{
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"carat", "ct"},
			DisplayName: "ct",
		},
		Factor: 5000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"cental", "centals"},
			DisplayName: "cwt",
		},
		Factor: 0.022046226218488,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"centigram", "centigrams"},
			DisplayName: "cg",
		},
		Factor: 100000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"dekagram", "dekagrams", "dg"},
			DisplayName: "dg",
		},
		Factor: 100,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"dram", "drams", "dr", "ʒ"},
			DisplayName: "dr",
		},
		Factor: 564.38339119329,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"grain", "grains", "gr"},
			DisplayName: "gr",
		},
		Factor: 15432.358352941,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"gram", "grams", "g"},
			DisplayName: "g",
		},
		Factor: 1000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"kilogram", "kilograms", "kg"},
			DisplayName: "kg",
		},
		Factor: 1,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"microgram", "micrograms", "µg"},
			DisplayName: "µg",
		},
		Factor: 1000000000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"milligram", "milligrams", "mg"},
			DisplayName: "mg",
		},
		Factor: 1000000,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"ounce", "ounces", "oz", "℥"},
			DisplayName: "oz",
		},
		Factor: 35.27396194958,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"pennyweight", "pennyweights", "dwt"},
			DisplayName: "dwt",
		},
		Factor: 643.01493137256,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"pound", "pounds", "lb"},
			DisplayName: "lb",
		},
		Factor: 2.2046226218488,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"quarter", "quarters", "qr"},
			DisplayName: "qr",
		},
		Factor: 0.078736522208885,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"stone", "stones", "st"},
			DisplayName: "st",
		},
		Factor: 0.15747304441777,
	},
	&SimpleDataType{
		b: &BaseDataType{
			Group:       GroupWeight,
			Names:       []string{"tonne", "ton", "megagram", "tonnes", "tons", "megagrams", "t"},
			DisplayName: "t",
		},
		Factor: 0.001,
	},
}

package datatype

var weightTypes = []DataType{
	DataType{
		Group:       GroupWeight,
		Names:       []string{"carat", "metric"},
		DisplayName: "metric",
		Factor:      5000,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"cental"},
		DisplayName: "cental",
		Factor:      0.022046226218488,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"centigram"},
		DisplayName: "centigram",
		Factor:      100000,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"dekagram"},
		DisplayName: "dekagram",
		Factor:      100,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"dram", "dr"},
		DisplayName: "dr",
		Factor:      564.38339119329,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"grain", "gr"},
		DisplayName: "gr",
		Factor:      15432.358352941,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"gram", "g"},
		DisplayName: "g",
		Factor:      1000,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"hundredweight", "UK"},
		DisplayName: "UK",
		Factor:      0.019684130552221,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"kilogram", "kg"},
		DisplayName: "kg",
		Factor:      1,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"microgram", "µg"},
		DisplayName: "µg",
		Factor:      1000000000,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"milligram", "mg"},
		DisplayName: "mg",
		Factor:      1000000,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"newton", "Earth"},
		DisplayName: "Earth",
		Factor:      9.80665,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"ounce", "oz"},
		DisplayName: "oz",
		Factor:      35.27396194958,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"pennyweight", "dwt"},
		DisplayName: "dwt",
		Factor:      643.01493137256,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"pound", "lb"},
		DisplayName: "lb",
		Factor:      2.2046226218488,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"quarter"},
		DisplayName: "quarter",
		Factor:      0.078736522208885,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"stone"},
		DisplayName: "stone",
		Factor:      0.15747304441777,
	},
	DataType{
		Group:       GroupWeight,
		Names:       []string{"tonne", "t"},
		DisplayName: "t",
		Factor:      0.001,
	},
}

package scan

import "testing"

type lexTest struct {
	input string
	items []Item
}

func fEOF() Item {
	return Item{ItemEOF, "", 0}
}

func itm(typ itemType, val string) Item {
	return Item{typ, val, 0}
}

var lexTests = []lexTest{
	// just digits
	{"0", []Item{itm(ItemNumber, "0"), fEOF()}},
	{"10", []Item{itm(ItemNumber, "10"), fEOF()}},
	{"0.2", []Item{itm(ItemNumber, "0.2"), fEOF()}},
	{"1.23", []Item{itm(ItemNumber, "1.23"), fEOF()}},
	{"-5", []Item{itm(ItemNumber, "-5"), fEOF()}},
	{"-5.1", []Item{itm(ItemNumber, "-5.1"), fEOF()}},
	{" 1  ", []Item{itm(ItemNumber, "1"), fEOF()}},
	// incorrect
	{"", []Item{itm(ItemError, "Unexpected EOF - at col 0")}},
	{"   ", []Item{itm(ItemError, "Unexpected EOF - at col 3")}},
	{"1.", []Item{itm(ItemError, "Unexpected char - at col 2")}},
	{"1. 0", []Item{itm(ItemError, "Unexpected char - at col 2")}},
	{"12d", []Item{itm(ItemNumber, "12"), itm(ItemError, "Unexpected char - at col 2")}},
	{"- 5", []Item{itm(ItemError, "Unexpected char - at col 1")}},

	// base operations
	{"1+1", []Item{
		itm(ItemNumber, "1"),
		itm(ItemAdd, "+"),
		itm(ItemNumber, "1"),
		fEOF()},
	},
	{"1-1", []Item{
		itm(ItemNumber, "1"),
		itm(ItemSub, "-"),
		itm(ItemNumber, "1"),
		fEOF()},
	},
	{"1*1", []Item{
		itm(ItemNumber, "1"),
		itm(ItemMul, "*"),
		itm(ItemNumber, "1"),
		fEOF()},
	},
	{"1/1", []Item{
		itm(ItemNumber, "1"),
		itm(ItemDiv, "/"),
		itm(ItemNumber, "1"),
		fEOF()},
	},
	{"1+1-2", []Item{
		itm(ItemNumber, "1"),
		itm(ItemAdd, "+"),
		itm(ItemNumber, "1"),
		itm(ItemSub, "-"),
		itm(ItemNumber, "2"),
		fEOF()},
	},

	//incorrect
	{"1++2", []Item{
		itm(ItemNumber, "1"),
		itm(ItemAdd, "+"),
		itm(ItemError, "Unexpected char - at col 2"),
	}},
	{"1+2+", []Item{
		itm(ItemNumber, "1"),
		itm(ItemAdd, "+"),
		itm(ItemNumber, "2"),
		itm(ItemAdd, "+"),
		itm(ItemError, "Unexpected EOF - at col 4"),
	}},

	//complex operations
	{
		"(1+1)/2", []Item{
			itm(ItemLParen, "("),
			itm(ItemNumber, "1"),
			itm(ItemAdd, "+"),
			itm(ItemNumber, "1"),
			itm(ItemRParen, ")"),
			itm(ItemDiv, "/"),
			itm(ItemNumber, "2"),
			fEOF()},
	},
	{
		"1+2*(3+4)-(-5)", []Item{
			itm(ItemNumber, "1"),
			itm(ItemAdd, "+"),
			itm(ItemNumber, "2"),
			itm(ItemMul, "*"),
			itm(ItemLParen, "("),
			itm(ItemNumber, "3"),
			itm(ItemAdd, "+"),
			itm(ItemNumber, "4"),
			itm(ItemRParen, ")"),
			itm(ItemSub, "-"),
			itm(ItemLParen, "("),
			itm(ItemNumber, "-5"),
			itm(ItemRParen, ")"),
			fEOF()},
	},
	{
		"1+((1+1)/(-2-100))", []Item{
			itm(ItemNumber, "1"),
			itm(ItemAdd, "+"),
			itm(ItemLParen, "("),
			itm(ItemLParen, "("),
			itm(ItemNumber, "1"),
			itm(ItemAdd, "+"),
			itm(ItemNumber, "1"),
			itm(ItemRParen, ")"),
			itm(ItemDiv, "/"),
			itm(ItemLParen, "("),
			itm(ItemNumber, "-2"),
			itm(ItemSub, "-"),
			itm(ItemNumber, "100"),
			itm(ItemRParen, ")"),
			itm(ItemRParen, ")"),
			fEOF()},
	},

	//incorrect
	{"1+1)/2", []Item{
		itm(ItemNumber, "1"),
		itm(ItemAdd, "+"),
		itm(ItemNumber, "1"),
		itm(ItemError, "No matching opening parenteses for closing one at col 3"),
	}},
	{"(1+1))/2", []Item{
		itm(ItemLParen, "("),
		itm(ItemNumber, "1"),
		itm(ItemAdd, "+"),
		itm(ItemNumber, "1"),
		itm(ItemRParen, ")"),
		itm(ItemError, "No matching opening parenteses for closing one at col 5"),
	}},
	{"(1+1)(1/2)", []Item{
		itm(ItemLParen, "("),
		itm(ItemNumber, "1"),
		itm(ItemAdd, "+"),
		itm(ItemNumber, "1"),
		itm(ItemRParen, ")"),
		itm(ItemError, "Unexpected char - at col 5"),
	}},
	{"1+()", []Item{
		itm(ItemNumber, "1"),
		itm(ItemAdd, "+"),
		itm(ItemLParen, "("),
		itm(ItemError, "Unexpected char - at col 3"),
	}},
	{"(1+1", []Item{
		itm(ItemLParen, "("),
		itm(ItemNumber, "1"),
		itm(ItemAdd, "+"),
		itm(ItemNumber, "1"),
		itm(ItemError, "Unexpected EOF - at col 4"),
	}},

	// identifiers
	{"t= 2", []Item{
		itm(ItemVariable, "t"),
		itm(ItemEqual, "="),
		itm(ItemNumber, "2"),
		fEOF(),
	}},
	{"t+1", []Item{
		itm(ItemVariable, "t"),
		itm(ItemAdd, "+"),
		itm(ItemNumber, "1"),
		fEOF(),
	}},
	{"2*t+mtp", []Item{
		itm(ItemNumber, "2"),
		itm(ItemMul, "*"),
		itm(ItemVariable, "t"),
		itm(ItemAdd, "+"),
		itm(ItemVariable, "mtp"),
		fEOF(),
	}},
	{"t==", []Item{
		itm(ItemVariable, "t"),
		itm(ItemEqual, "="),
		itm(ItemError, "Unexpected char - at col 2"),
	}},
	{"t=1=1", []Item{
		itm(ItemVariable, "t"),
		itm(ItemEqual, "="),
		itm(ItemNumber, "1"),
		itm(ItemError, "Unexpected char - at col 3"),
	}},
	{"t+1=1", []Item{
		itm(ItemVariable, "t"),
		itm(ItemAdd, "+"),
		itm(ItemNumber, "1"),
		itm(ItemError, "Unexpected char - at col 3"),
	}},
	{"t$s=1", []Item{
		itm(ItemVariable, "t"),
		itm(ItemError, "Unexpected char - at col 1"),
	}},

	// {"(1 + 2)", []item{
	// 	item{itemLParen, "(", nilType},
	// 	item{itemDoubleLiteral, "1", nilType},
	// 	item{itemAdd, "+", nilType},
	// 	item{itemDoubleLiteral, "2", nilType},
	// 	item{itemRParen, ")", nilType},
	// 	fEof()},
	// },
	// {"1 usd", []item{
	// 	item{itemDoubleLiteral, "1", usd},
	// 	fEof()},
	// },
}

func collect(t *lexTest) (items []Item) {
	// buf := bytes.NewBufferString()
	l := New(t.input)
	for {
		item := l.NextItem()
		items = append(items, item)
		if item.Typ == ItemEOF || item.Typ == ItemError {
			break
		}
	}
	return
}

func equal(i1, i2 []Item) bool {
	if len(i1) != len(i2) {
		return false
	}
	for k := range i1 {
		if i1[k].Typ != i2[k].Typ || i1[k].Val != i2[k].Val {
			return false
		}
	}
	return true
}

func TestLex(t *testing.T) {

	for _, test := range lexTests {
		t.Run(test.input, func(t *testing.T) {
			items := collect(&test)
			if !equal(items, test.items) {
				t.Errorf("\ninput:%q\ngot  :%+v\nwant :%+v", test.input, items, test.items)
			}
		})

	}
}

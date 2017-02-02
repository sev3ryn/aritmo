package scan

import "testing"

type lexTest struct {
	input string
	items []Item
}

func fEOF() Item {
	return Item{itemEOF, "", 0}
}

func itm(typ itemType, val string) Item {
	return Item{typ, val, 0}
}

var lexTests = []lexTest{
	// just digits
	{"0", []Item{itm(itemNumber, "0"), fEOF()}},
	{"10", []Item{itm(itemNumber, "10"), fEOF()}},
	{"0.2", []Item{itm(itemNumber, "0.2"), fEOF()}},
	{"1.23", []Item{itm(itemNumber, "1.23"), fEOF()}},
	{"-5", []Item{itm(itemNumber, "-5"), fEOF()}},
	{"-5.1", []Item{itm(itemNumber, "-5.1"), fEOF()}},
	{" 1  ", []Item{itm(itemNumber, "1"), fEOF()}},
	// incorrect
	{"", []Item{itm(itemError, "Unexpected EOF - at col 0")}},
	{"   ", []Item{itm(itemError, "Unexpected EOF - at col 3")}},
	{"1.", []Item{itm(itemError, "Unexpected char - at col 2")}},
	{"1. 0", []Item{itm(itemError, "Unexpected char - at col 2")}},
	{"12d", []Item{itm(itemNumber, "12"), itm(itemError, "Unexpected char - at col 2")}},
	{"- 5", []Item{itm(itemError, "Unexpected char - at col 1")}},

	// base operations
	{"1+1", []Item{
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		fEOF()},
	},
	{"1-1", []Item{
		itm(itemNumber, "1"),
		itm(itemSub, "-"),
		itm(itemNumber, "1"),
		fEOF()},
	},
	{"1*1", []Item{
		itm(itemNumber, "1"),
		itm(itemMul, "*"),
		itm(itemNumber, "1"),
		fEOF()},
	},
	{"1/1", []Item{
		itm(itemNumber, "1"),
		itm(itemDiv, "/"),
		itm(itemNumber, "1"),
		fEOF()},
	},
	{"1+1-2", []Item{
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		itm(itemSub, "-"),
		itm(itemNumber, "2"),
		fEOF()},
	},

	//incorrect
	{"1++2", []Item{
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemError, "Unexpected char - at col 2"),
	}},
	{"1+2+", []Item{
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "2"),
		itm(itemAdd, "+"),
		itm(itemError, "Unexpected EOF - at col 4"),
	}},

	//complex operations
	{"(1+1)/2", []Item{
		itm(itemLParen, "("),
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		itm(itemRParen, ")"),
		itm(itemDiv, "/"),
		itm(itemNumber, "2"),
		fEOF()},
	},
	{"1+((1+1)/(-2-100))", []Item{
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemLParen, "("),
		itm(itemLParen, "("),
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		itm(itemRParen, ")"),
		itm(itemDiv, "/"),
		itm(itemLParen, "("),
		itm(itemNumber, "-2"),
		itm(itemSub, "-"),
		itm(itemNumber, "100"),
		itm(itemRParen, ")"),
		itm(itemRParen, ")"),
		fEOF()},
	},

	//incorrect
	{"1+1)/2", []Item{
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		itm(itemError, "No matching opening parenteses for closing one at col 3"),
	}},
	{"(1+1))/2", []Item{
		itm(itemLParen, "("),
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		itm(itemRParen, ")"),
		itm(itemError, "No matching opening parenteses for closing one at col 5"),
	}},
	{"(1+1)(1/2)", []Item{
		itm(itemLParen, "("),
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		itm(itemRParen, ")"),
		itm(itemError, "Unexpected char - at col 5"),
	}},
	{"(1+1", []Item{
		itm(itemLParen, "("),
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		itm(itemError, "Unexpected EOF - at col 4"),
	}},

	// identifiers
	{"t= 2", []Item{
		itm(itemVariable, "t"),
		itm(itemEqual, "="),
		itm(itemNumber, "2"),
		fEOF(),
	}},
	{"t+1", []Item{
		itm(itemVariable, "t"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		fEOF(),
	}},
	{"2*t+mtp", []Item{
		itm(itemNumber, "2"),
		itm(itemMul, "*"),
		itm(itemVariable, "t"),
		itm(itemAdd, "+"),
		itm(itemVariable, "mtp"),
		fEOF(),
	}},
	{"t==", []Item{
		itm(itemVariable, "t"),
		itm(itemEqual, "="),
		itm(itemError, "Unexpected char - at col 2"),
	}},
	{"t=1=1", []Item{
		itm(itemVariable, "t"),
		itm(itemEqual, "="),
		itm(itemNumber, "1"),
		itm(itemError, "Unexpected char - at col 3"),
	}},
	{"t+1=1", []Item{
		itm(itemVariable, "t"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		itm(itemError, "Unexpected char - at col 3"),
	}},
	{"t$s=1", []Item{
		itm(itemVariable, "t"),
		itm(itemError, "Unexpected char - at col 1"),
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
		item := l.nextItem()
		items = append(items, item)
		if item.typ == itemEOF || item.typ == itemError {
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
		if i1[k].typ != i2[k].typ || i1[k].val != i2[k].val {
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

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
	{"0", []Item{
		itm(ItemNumber, "0"),
		itm(ItemBareDataType, ""),
		fEOF(),
	}},
	{"10", []Item{
		itm(ItemNumber, "10"),
		itm(ItemBareDataType, ""),
		fEOF(),
	}},
	{"0.2", []Item{
		itm(ItemNumber, "0.2"),
		itm(ItemBareDataType, ""),
		fEOF(),
	}},
	{"1.23", []Item{
		itm(ItemNumber, "1.23"),
		itm(ItemBareDataType, ""),
		fEOF(),
	}},
	{"-5", []Item{
		itm(ItemNumber, "-5"),
		itm(ItemBareDataType, ""),
		fEOF(),
	}},
	{"-5.1", []Item{
		itm(ItemNumber, "-5.1"),
		itm(ItemBareDataType, ""),
		fEOF(),
	}},
	{" 1  ", []Item{
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		fEOF(),
	}},
	// incorrect
	{"", []Item{
		itm(ItemError, "Unexpected EOF - at col 0"),
	}},
	{"   ", []Item{
		itm(ItemError, "Unexpected EOF - at col 3"),
	}},
	{"1.", []Item{
		itm(ItemError, "Unexpected char - at col 2"),
	}},
	{"1. 0", []Item{
		itm(ItemError, "Unexpected char - at col 2"),
	}},
	{"- 5", []Item{
		itm(ItemError, "Unexpected char - at col 1"),
	}},

	// base operations
	{"1+1", []Item{
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "+"),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		fEOF()},
	},
	{"1-1", []Item{
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "-"),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		fEOF()},
	},
	{"1*1", []Item{
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "*"),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		fEOF()},
	},
	{"1/1", []Item{
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "/"),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		fEOF()},
	},
	{"1+1-2", []Item{
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "+"),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "-"),
		itm(ItemNumber, "2"),
		itm(ItemBareDataType, ""),
		fEOF()},
	},

	//incorrect
	{"1++2", []Item{
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "+"),
		itm(ItemError, "Unexpected char - at col 2"),
	}},
	{"1+2+", []Item{
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "+"),
		itm(ItemNumber, "2"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "+"),
		itm(ItemError, "Unexpected EOF - at col 4"),
	}},

	//complex operations
	{
		"(1+1)/2", []Item{
			itm(ItemLParen, "("),
			itm(ItemNumber, "1"),
			itm(ItemBareDataType, ""),
			itm(ItemOperation, "+"),
			itm(ItemNumber, "1"),
			itm(ItemBareDataType, ""),
			itm(ItemRParen, ")"),
			itm(ItemOperation, "/"),
			itm(ItemNumber, "2"),
			itm(ItemBareDataType, ""),
			fEOF()},
	},
	{
		"1+2*(3+4)-(-5)", []Item{
			itm(ItemNumber, "1"),
			itm(ItemBareDataType, ""),
			itm(ItemOperation, "+"),
			itm(ItemNumber, "2"),
			itm(ItemBareDataType, ""),
			itm(ItemOperation, "*"),
			itm(ItemLParen, "("),
			itm(ItemNumber, "3"),
			itm(ItemBareDataType, ""),
			itm(ItemOperation, "+"),
			itm(ItemNumber, "4"),
			itm(ItemBareDataType, ""),
			itm(ItemRParen, ")"),
			itm(ItemOperation, "-"),
			itm(ItemLParen, "("),
			itm(ItemNumber, "-5"),
			itm(ItemBareDataType, ""),
			itm(ItemRParen, ")"),
			fEOF()},
	},
	{
		"1+((1+1)/(-2-100))", []Item{
			itm(ItemNumber, "1"),
			itm(ItemBareDataType, ""),
			itm(ItemOperation, "+"),
			itm(ItemLParen, "("),
			itm(ItemLParen, "("),
			itm(ItemNumber, "1"),
			itm(ItemBareDataType, ""),
			itm(ItemOperation, "+"),
			itm(ItemNumber, "1"),
			itm(ItemBareDataType, ""),
			itm(ItemRParen, ")"),
			itm(ItemOperation, "/"),
			itm(ItemLParen, "("),
			itm(ItemNumber, "-2"),
			itm(ItemBareDataType, ""),
			itm(ItemOperation, "-"),
			itm(ItemNumber, "100"),
			itm(ItemBareDataType, ""),
			itm(ItemRParen, ")"),
			itm(ItemRParen, ")"),
			fEOF()},
	},

	//incorrect
	{"1+1)/2", []Item{
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "+"),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemError, "No matching opening parenteses for closing one at col 3"),
	}},
	{"(1+1))/2", []Item{
		itm(ItemLParen, "("),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "+"),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemRParen, ")"),
		itm(ItemError, "No matching opening parenteses for closing one at col 5"),
	}},
	{"(1+1)(1/2)", []Item{
		itm(ItemLParen, "("),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "+"),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemRParen, ")"),
		itm(ItemError, "Unexpected char - at col 5"),
	}},
	{"1+()", []Item{
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "+"),
		itm(ItemLParen, "("),
		itm(ItemError, "Unexpected char - at col 3"),
	}},
	{"(1+1", []Item{
		itm(ItemLParen, "("),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "+"),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemError, "Unexpected EOF - at col 4"),
	}},

	// identifiers
	{"t= 2", []Item{
		itm(ItemVariable, "t"),
		itm(ItemEqual, "="),
		itm(ItemNumber, "2"),
		itm(ItemBareDataType, ""),
		fEOF(),
	}},
	{"t+1", []Item{
		itm(ItemVariable, "t"),
		itm(ItemOperation, "+"),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		fEOF(),
	}},
	{"2*t+mtp", []Item{
		itm(ItemNumber, "2"),
		itm(ItemBareDataType, ""),
		itm(ItemOperation, "*"),
		itm(ItemVariable, "t"),
		itm(ItemOperation, "+"),
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
		itm(ItemBareDataType, ""),
		itm(ItemError, "Unexpected char - at col 3"),
	}},
	{"t+1=1", []Item{
		itm(ItemVariable, "t"),
		itm(ItemOperation, "+"),
		itm(ItemNumber, "1"),
		itm(ItemBareDataType, ""),
		itm(ItemError, "Unexpected char - at col 3"),
	}},
	{"t$s=1", []Item{
		itm(ItemVariable, "t"),
		itm(ItemError, "Unexpected char - at col 1"),
	}},

	// Datatypes
	{"12usd", []Item{
		itm(ItemNumber, "12"),
		itm(ItemDataType, "usd"),
		fEOF(),
	}},
	{"13 eur ", []Item{
		itm(ItemNumber, "13"),
		itm(ItemDataType, "eur"),
		fEOF(),
	}},
	{"12 usd+13eur% ", []Item{
		itm(ItemNumber, "12"),
		itm(ItemDataType, "usd"),
		itm(ItemOperation, "+"),
		itm(ItemNumber, "13"),
		itm(ItemDataType, "eur"),
		itm(ItemError, "Unexpected char - at col 12"),
	}},
	// {"(1 + 2)", []item{
	// 	item{itemLParen, "(", nilType},
	// 	item{itemDoubleLiteral, "1", nilType},
	// 	item{ItemOperation, "+", nilType},
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

func TestLexOperand(t *testing.T) {

	tests := []struct {
		input string
		want  Item
	}{
		{input: "", want: Item{Typ: ItemError, Val: "Unexpected EOF - at col 0"}},
		{input: "   ", want: Item{Typ: ItemError, Val: "Unexpected EOF - at col 3"}},
		{input: "3000", want: Item{Typ: ItemNumber, Val: "3000"}},
		{input: "  -1003", want: Item{Typ: ItemNumber, Val: "-1003"}},
		{input: "  -", want: Item{Typ: ItemError, Val: "Unexpected EOF - at col 3"}},
		{input: "  -a", want: Item{Typ: ItemError, Val: "Unexpected char - at col 3"}},
		{input: "(", want: Item{Typ: ItemLParen, Val: "("}}, // need to check increment working
		{input: "*", want: Item{Typ: ItemError, Val: "Unexpected char - at col 0"}},
		//{input: "  a", want: Item{Typ: ItemVariable, Val: "a"}}, // need to check no output
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			s := &Scanner{input: []rune(tt.input), items: make(chan Item)}
			go func() { _ = lexOperand(s) }()

			if itm := <-s.items; itm.Typ != tt.want.Typ || itm.Val != tt.want.Val {
				t.Errorf("lexOperand() = %v, want %v", itm, tt.want)
			}
		})
	}
}

func TestLexOperator(t *testing.T) {
	tests := []struct {
		input string
		want  Item
	}{
		{input: "", want: Item{Typ: ItemEOF, Val: ""}},
		{input: "   ", want: Item{Typ: ItemEOF, Val: ""}},
		{input: "+", want: Item{Typ: ItemOperation, Val: "+"}},
		{input: "  -", want: Item{Typ: ItemOperation, Val: "-"}},
		{input: "*", want: Item{Typ: ItemOperation, Val: "*"}},
		{input: "/", want: Item{Typ: ItemOperation, Val: "/"}},
		{input: "to ", want: Item{Typ: ItemOperation, Val: "to"}},
		{input: "tp", want: Item{Typ: ItemError, Val: "Unexpected char - at col 1"}},
		{input: "too", want: Item{Typ: ItemError, Val: "Unexpected char - at col 2"}},
		{input: ")", want: Item{Typ: ItemError, Val: "No matching opening parenteses for closing one at col 0"}}, // need to check decrement working
		{input: "a", want: Item{Typ: ItemError, Val: "Unexpected char - at col 0"}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			s := &Scanner{input: []rune(tt.input), items: make(chan Item)}
			go func() { _ = lexOperator(s) }()

			if itm := <-s.items; itm.Typ != tt.want.Typ || itm.Val != tt.want.Val {
				t.Errorf("lexOperator() = %v, want %v", itm, tt.want)
			}
		})
	}
}

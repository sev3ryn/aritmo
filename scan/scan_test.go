package scan

import "testing"

type lexTest struct {
	input string
	items []item
}

func fEof() item {
	return item{itemEOF, "", 0}
}

func itm(typ itemType, val string) item {
	return item{typ, val, 0}
}

var lexTests = []lexTest{
	// just digits
	{"0", []item{itm(itemNumber, "0"), fEof()}},
	{"10", []item{itm(itemNumber, "10"), fEof()}},
	{"0.2", []item{itm(itemNumber, "0.2"), fEof()}},
	{"1.23", []item{itm(itemNumber, "1.23"), fEof()}},
	//{"-5", []item{itm(itemNumber, "-5"), fEof()}}, //not working yet
	//{"-5.1", []item{itm(itemNumber, "-5.1"), fEof()}}, // not working yet
	{" 1  ", []item{itm(itemNumber, "1"), fEof()}},
	// incorrect
	{"", []item{itm(itemError, "Unexpected EOF - at col 0")}},
	{"   ", []item{itm(itemError, "Unexpected EOF - at col 3")}},
	{"1.", []item{itm(itemError, "Unexpected char - at col 2")}},
	{"1. 0", []item{itm(itemError, "Unexpected char - at col 2")}},
	{"12d", []item{itm(itemNumber, "12"), itm(itemError, "Unexpected char - at col 2")}},
	// {"- 5", []item{item{itemError, "unexpected space"}}},

	// base operations
	{"1+1", []item{
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		fEof()},
	},
	{"1-1", []item{
		itm(itemNumber, "1"),
		itm(itemSub, "-"),
		itm(itemNumber, "1"),
		fEof()},
	},
	{"1*1", []item{
		itm(itemNumber, "1"),
		itm(itemMul, "*"),
		itm(itemNumber, "1"),
		fEof()},
	},
	{"1/1", []item{
		itm(itemNumber, "1"),
		itm(itemDiv, "/"),
		itm(itemNumber, "1"),
		fEof()},
	},

	//incorrect
	{"1++2", []item{
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemError, "Unexpected char - at col 2"),
	}},
	{"1+2+", []item{
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "2"),
		itm(itemAdd, "+"),
		itm(itemError, "Unexpected EOF - at col 4"),
	}},

	//complex operations
	{"1+1-2", []item{
		itm(itemNumber, "1"),
		itm(itemAdd, "+"),
		itm(itemNumber, "1"),
		itm(itemSub, "-"),
		itm(itemNumber, "2"),
		fEof()},
	},
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

func collect(t *lexTest) (items []item) {
	// buf := bytes.NewBufferString()
	l := lex(t.input)
	for {
		item := l.nextItem()
		items = append(items, item)
		if item.typ == itemEOF || item.typ == itemError {
			break
		}
	}
	return
}

func equal(i1, i2 []item) bool {
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

type stringTest struct {
	name  string
	input item
	str   string
}

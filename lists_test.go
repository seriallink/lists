package lists

import "testing"

var (
	QuoteStr = "to be or not to be"
	QuoteSlc = []string{"to", "be", "or", "not", "to", "be"}
)

func TestParseArray(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")

	if len(hamlet.Values) != len(QuoteSlc) {
		t.Errorf("Array haven't been parsed correctly.")
	}

	if hamlet.Separator != " " {
		t.Errorf("Separator haven't been assigned correctly.")
	}

	if hamlet.Index != -1 {
		t.Errorf("Index haven't been initialized correctly.")
	}

}

func TestParseString(t *testing.T) {

	hamlet := ParseString(QuoteStr, " ")

	if len(hamlet.Values) != 6 {
		t.Errorf("Array haven't been parsed correctly.")
	}

	if hamlet.Separator != " " {
		t.Errorf("Separator haven't been assigned correctly.")
	}

	if hamlet.Index != -1 {
		t.Errorf("Index haven't been initialized correctly.")
	}

}

func TestToString(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")

	if hamlet.ToString() != QuoteStr {
		t.Errorf("Strings are not matching.")
	}

}

func TestToArray(t *testing.T) {

	hamlet := ParseString(QuoteStr, " ")

	if len(hamlet.Values) != len(QuoteSlc) {
		t.Errorf("Slices are not matching.")
	}

	for i := range hamlet.Values {
		if hamlet.Values[i] != QuoteSlc[i] {
			t.Errorf("Slices are not matching.")
		}
	}

}

func TestLen(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")

	if hamlet.Len() != len(QuoteSlc) {
		t.Errorf("Len failed.")
	}

}

func TestFirst(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")

	if hamlet.First() != QuoteSlc[0] {
		t.Errorf("First failed.")
	}

}

func TestLast(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")

	if hamlet.Last() != QuoteSlc[len(QuoteSlc)-1] {
		t.Errorf("Last failed.")
	}

}

func TestFind(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")

	if hamlet.Find("question") != -1 || hamlet.Find("be") != 1 {
		t.Errorf("Find failed.")
	}

}

func TestExist(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")

	if hamlet.Exist("question") || !hamlet.Exist("be") {
		t.Errorf("Exist failed.")
	}

}

func TestHasNext(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")

	for hamlet.HasNext() {
		hamlet.Next()
	}

	if hamlet.Index != len(QuoteSlc)-1 {
		t.Errorf("HasNext failed.")
	}

}

func TestMoveTo(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")
	hamlet.MoveTo(2)

	if hamlet.Get() != "or" {
		t.Errorf("MoveTo failed.")
	}

}

func TestRewind(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")
	hamlet.MoveTo(2)
	hamlet.Rewind()

	if hamlet.Index != -1 {
		t.Errorf("Rewind failed.")
	}

}

func TestIsEmpty(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")
	macbeth := List{}

	if hamlet.IsEmpty() || !macbeth.IsEmpty() {
		t.Errorf("IsEmpty failed.")
	}

}

func TestIsOutOfBound(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")

	if hamlet.IsOutOfBound(1) || !hamlet.IsOutOfBound(10) {
		t.Errorf("IsOutOfBound failed.")
	}

}

func TestGet(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")
	hamlet.MoveTo(2)

	if hamlet.Get() != "or" {
		t.Errorf("Get failed.")
	}

}

func TestGetAt(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")

	if value, _ := hamlet.GetAt(2); value != "or" {
		t.Errorf("GetAt failed.")
	}

}

func TestDedup(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")
	hamlet.Dedup()

	if hamlet.ToString() != "to be or not" {
		t.Errorf("Dedup failed.")
	}

}

func TestQuote(t *testing.T) {

	hamlet := ParseArray(QuoteSlc, " ")
	hamlet.Quote('\'')

	if hamlet.ToString() != "'to' 'be' 'or' 'not' 'to' 'be'" {
		t.Errorf("Quote has failed.")
	}

}

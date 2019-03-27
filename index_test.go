package rankreq

import (
	"testing"
)

func TestIndex(t *testing.T) {

	tsvFile, reader, err := FileDescribe("test/test.tsv")
	if err != nil {
		t.Error(err)
	}
	root := Moment{}
	root.Index(tsvFile, reader)

	// Root level
	str := "Root level: "
	if len(root.children) != 1 {
		t.Error(str + "wrong number of children")
	}
	if root.children["2015"] == nil {
		t.Error(str + "wrong children")
	}

	// Year level
	str = "Year level: "
	if len(root.children["2015"].children) != 1 {
		t.Error(str + " wrong number of children")
	}
	if root.children["2015"].children["08"] == nil {
		t.Error(str + "wrong children")
	}

	// Month level
	str = "Month level: "
	if len(root.children["2015"].children["08"].children) != 3 {
		t.Error(str + "wrong number of children")
	}
	if root.children["2015"].children["08"].children["01"] == nil {
		t.Error(str + "wrong children")
	}
	if root.children["2015"].children["08"].children["02"] == nil {
		t.Error(str + "wrong children")
	}
	if root.children["2015"].children["08"].children["04"] == nil {
		t.Error(str + "wrong children")
	}
	if root.children["2015"].children["08"].children["05"] != nil {
		t.Error(str + "wrong children")
	}
	if root.children["2015"].children["08"].children["01"].count != 2 {
		t.Error(str + "wrong moment count")
	}

	// Minute level
	str = "Minute level: "
	if len(root.children["2015"].children["08"].children["01"].children["00"].children["03"].children) != 2 {
		t.Error(str + "wrong number of children")
	}

	// Second level
	str = "Second level: "
	if len(root.children["2015"].children["08"].children["01"].children["00"].children["03"].queries) != 1 {
		t.Error(str + "wrong number of queries for 2015-08-01 00:03:43: should be 1")
	}
	if root.children["2015"].children["08"].children["01"].children["00"].children["03"].children["43"].queries["reqA"] == nil {
		t.Error(str + "wrong query for 2015-08-01 00:03:43: should be reqA")
	}
	len := len(root.children["2015"].children["08"].children["02"].children["05"].children["17"].children["11"].queries)
	if len != 1 {
		t.Error(str+"wrong number of queries for 2015-08-02 05:17:11: should be 1, have", len)
	}
	if root.children["2015"].children["08"].children["02"].children["05"].children["17"].children["11"].queries["reqC"].Count != 2 {
		t.Error(str + "wrong number of occurence of the same query for 2015-08-02 05:17:11: should be 2")
	}
	if root.children["2015"].children["08"].children["04"].children["06"].children["52"].children["16"].queries["reqA"] == nil {
		t.Error(str + "wrong query for 2015-08-04 06:52:16: should be reqA")
	}
}

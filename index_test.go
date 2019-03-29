package rankreq

import (
	"testing"
)

func TestIndex(t *testing.T) {

	tsvFile, reader, err := FileDescribe("test/test_index.tsv")
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
	if root.children[2015] == nil {
		t.Error(str + "wrong children")
	}

	// Year level
	str = "Year level: "
	if len(root.children[2015].children) != 1 {
		t.Error(str + " wrong number of children")
	}
	if root.children[2015].children[8] == nil {
		t.Error(str + "wrong children")
	}

	// Month level
	str = "Month level: "
	if len(root.children[2015].children[8].children) != 3 {
		t.Error(str + "wrong number of children")
	}
	if root.children[2015].children[8].children[1] == nil {
		t.Error(str + "wrong children")
	}
	if root.children[2015].children[8].children[2] == nil {
		t.Error(str + "wrong children")
	}
	if root.children[2015].children[8].children[4] == nil {
		t.Error(str + "wrong children")
	}
	if root.children[2015].children[8].children[5] != nil {
		t.Error(str + "wrong children")
	}
	if root.children[2015].children[8].children[1].count != 2 {
		t.Error(str + "wrong moment count")
	}

	// Minute level
	str = "Minute level: "
	if len(root.children[2015].children[8].children[1].children[0].children[3].children) != 2 {
		t.Error(str + "wrong number of children")
	}

	// Second level
	str = "Second level: "
	if len := len(root.children[2015].children[8].children[1].children[0].children[3].queries); len != 1 {
		t.Error(str+"wrong number of queries for 2015-08-01 00:03:43: should be 1, has", len)
	}
	if _, found := root.children[2015].children[8].children[1].children[0].children[3].children[43].queries["reqA"]; !found {
		t.Error(str + "wrong query for 2015-08-01 00:03:43: should be reqA")
	}
	if len := len(root.children[2015].children[8].children[2].children[5].children[17].children[11].queries); len != 1 {
		t.Error(str+"wrong number of queries for 2015-08-02 05:17:11: should be 1, has", len)
	}
	if root.children[2015].children[8].children[2].children[5].children[17].children[11].queries["reqC"] != 2 {
		t.Error(str + "wrong number of occurence of the same query for 2015-08-02 05:17:11: should be 2")
	}
	if _, found := root.children[2015].children[8].children[4].children[6].children[52].children[16].queries["reqA"]; !found {
		t.Error(str + "wrong query for 2015-08-04 06:52:16: should be reqA")
	}
}

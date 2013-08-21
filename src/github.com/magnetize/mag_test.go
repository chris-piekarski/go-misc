package main
import (
	"testing"
)

func TestPendingNew(t *testing.T){
	clearAllPending()
	
	max := 100

	for i := 0; i < max; i++ {
		me := "chris@cpiekarski.com"
		notMe := "bob.smith@yahoo.com"
		storeNewRequest(me, notMe, "")
	}
	
	if countPendingNew() != max {
		t.Error("Wrong number of new entries")
	} else {
		t.Log("Number of new entries match")
	}
}

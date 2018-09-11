package mpunbound

import (
	"testing"
)

func TestGraphDefinition(t *testing.T) {
	var unbound UnboundPlugin
	expect := 6

	graphdef := unbound.GraphDefinition()
	if len(graphdef) != expect {
		t.Errorf("GetTempfilename: %d should be %d", len(graphdef), expect)
	}
}

func TestParseStats(t *testing.T) {
	stats := `thread0.num.queries=74
thread0.num.cachehits=67
thread0.num.cachemiss=7
thread0.num.prefetch=0
thread0.num.recursivereplies=7
thread0.requestlist.avg=0
thread0.requestlist.max=0
thread0.requestlist.overwritten=0
thread0.requestlist.exceeded=0
thread0.requestlist.current.all=0
thread0.requestlist.current.user=0
thread0.recursion.time.avg=0.027743
thread0.recursion.time.median=0.024576
total.num.queries=74
total.num.cachehits=67
total.num.cachemiss=7
total.num.prefetch=0
total.num.recursivereplies=7
total.requestlist.avg=0
total.requestlist.max=0
total.requestlist.overwritten=0
total.requestlist.exceeded=0
total.requestlist.current.all=0
total.requestlist.current.user=0
total.recursion.time.avg=0.027743
total.recursion.time.median=0.024576
time.now=1535624567.513783
time.up=5717068.166485
time.elapsed=403.873205`

	var unbound UnboundPlugin
	result, err := unbound.parseStats(stats)
	if err != nil {
		t.Errorf("Failed to parse: %s", err)
	}

	expect := map[string]float64{
		"thread0.num.queries":              74,
		"thread0.num.cachehits":            67,
		"thread0.num.cachemiss":            7,
		"thread0.num.prefetch":             0,
		"thread0.num.recursivereplies":     7,
		"thread0.requestlist.avg":          0,
		"thread0.requestlist.max":          0,
		"thread0.requestlist.overwritten":  0,
		"thread0.requestlist.exceeded":     0,
		"thread0.requestlist.current.all":  0,
		"thread0.requestlist.current.user": 0,
		"thread0.recursion.time.avg":       0.027743,
		"thread0.recursion.time.median":    0.024576,
		"total.num.queries":                74,
		"total.num.cachehits":              67,
		"total.num.cachemiss":              7,
		"total.num.prefetch":               0,
		"total.num.recursivereplies":       7,
		"total.requestlist.avg":            0,
		"total.requestlist.max":            0,
		"total.requestlist.overwritten":    0,
		"total.requestlist.exceeded":       0,
		"total.requestlist.current.all":    0,
		"total.requestlist.current.user":   0,
		"total.recursion.time.avg":         0.027743,
		"total.recursion.time.median":      0.024576,
		"time.now":                         1535624567.513783,
		"time.up":                          5717068.166485,
		"time.elapsed":                     403.873205,
	}

	for k := range expect {
		if expect[k] != result[k] {
			t.Errorf("%s does not match\nExpect: %v\nResult: %v\nresult: %v\n", k, expect[k], result[k], result)
		}
	}
}

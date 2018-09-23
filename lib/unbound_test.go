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
	stats := `thread0.num.queries=1452
thread0.num.cachehits=1200
thread0.num.cachemiss=252
thread0.num.prefetch=0
thread0.num.recursivereplies=252
thread0.requestlist.avg=0.678571
thread0.requestlist.max=16
thread0.requestlist.overwritten=0
thread0.requestlist.exceeded=0
thread0.requestlist.current.all=0
thread0.requestlist.current.user=0
thread0.recursion.time.avg=1.357106
thread0.recursion.time.median=0.0284402
total.num.queries=1452
total.num.cachehits=1200
total.num.cachemiss=252
total.num.prefetch=0
total.num.recursivereplies=252
total.requestlist.avg=0.678571
total.requestlist.max=16
total.requestlist.overwritten=0
total.requestlist.exceeded=0
total.requestlist.current.all=0
total.requestlist.current.user=0
total.recursion.time.avg=1.357106
total.recursion.time.median=0.0284402
time.now=1536756976.225834
time.up=10413.923319
time.elapsed=10413.923319
mem.total.sbrk=8523776
mem.cache.rrset=409036
mem.cache.message=402097
mem.mod.iterator=16532
mem.mod.validator=130452
histogram.000000.000000.to.000000.000001=0
histogram.000000.000001.to.000000.000002=0
histogram.000000.000002.to.000000.000004=0
histogram.000000.000004.to.000000.000008=0
histogram.000000.000008.to.000000.000016=0
histogram.000000.000016.to.000000.000032=0
histogram.000000.000032.to.000000.000064=0
histogram.000000.000064.to.000000.000128=0
histogram.000000.000128.to.000000.000256=0
histogram.000000.000256.to.000000.000512=0
histogram.000000.000512.to.000000.001024=3
histogram.000000.001024.to.000000.002048=42
histogram.000000.002048.to.000000.004096=24
histogram.000000.004096.to.000000.008192=9
histogram.000000.008192.to.000000.016384=9
histogram.000000.016384.to.000000.032768=53
histogram.000000.032768.to.000000.065536=40
histogram.000000.065536.to.000000.131072=18
histogram.000000.131072.to.000000.262144=31
histogram.000000.262144.to.000000.524288=17
histogram.000000.524288.to.000001.000000=0
histogram.000001.000000.to.000002.000000=1
histogram.000002.000000.to.000004.000000=0
histogram.000004.000000.to.000008.000000=0
histogram.000008.000000.to.000016.000000=0
histogram.000016.000000.to.000032.000000=1
histogram.000032.000000.to.000064.000000=1
histogram.000064.000000.to.000128.000000=3
histogram.000128.000000.to.000256.000000=0
histogram.000256.000000.to.000512.000000=0
histogram.000512.000000.to.001024.000000=0
histogram.001024.000000.to.002048.000000=0
histogram.002048.000000.to.004096.000000=0
histogram.004096.000000.to.008192.000000=0
histogram.008192.000000.to.016384.000000=0
histogram.016384.000000.to.032768.000000=0
histogram.032768.000000.to.065536.000000=0
histogram.065536.000000.to.131072.000000=0
histogram.131072.000000.to.262144.000000=0
histogram.262144.000000.to.524288.000000=0
num.query.type.A=745
num.query.type.MX=2
num.query.type.AAAA=705
num.query.class.IN=1452
num.query.opcode.QUERY=1452
num.query.tcp=0
num.query.ipv6=0
num.query.flags.QR=0
num.query.flags.AA=0
num.query.flags.TC=0
num.query.flags.RD=1452
num.query.flags.RA=0
num.query.flags.Z=0
num.query.flags.AD=0
num.query.flags.CD=0
num.query.edns.present=0
num.query.edns.DO=0
num.answer.rcode.NOERROR=1416
num.answer.rcode.SERVFAIL=6
num.answer.rcode.NXDOMAIN=30
num.answer.rcode.nodata=680
num.answer.secure=22
num.answer.bogus=0
num.rrset.bogus=0
unwanted.queries=0
unwanted.replies=0`

	var unbound UnboundPlugin
	result, err := unbound.parseStats(stats)
	if err != nil {
		t.Errorf("Failed to parse: %s", err)
	}

	expect := map[string]float64{
		"thread0.num.queries":                      1452,
		"thread0.num.cachehits":                    1200,
		"thread0.num.cachemiss":                    252,
		"thread0.num.prefetch":                     0,
		"thread0.num.recursivereplies":             252,
		"thread0.requestlist.avg":                  0.678571,
		"thread0.requestlist.max":                  16,
		"thread0.requestlist.overwritten":          0,
		"thread0.requestlist.exceeded":             0,
		"thread0.requestlist.current.all":          0,
		"thread0.requestlist.current.user":         0,
		"thread0.recursion.time.avg":               1.357106,
		"thread0.recursion.time.median":            0.0284402,
		"total.num.queries":                        1452,
		"total.num.cachehits":                      1200,
		"total.num.cachemiss":                      252,
		"total.num.prefetch":                       0,
		"total.num.recursivereplies":               252,
		"total.requestlist.avg":                    0.678571,
		"total.requestlist.max":                    16,
		"total.requestlist.overwritten":            0,
		"total.requestlist.exceeded":               0,
		"total.requestlist.current.all":            0,
		"total.requestlist.current.user":           0,
		"total.recursion.time.avg":                 1.357106,
		"total.recursion.time.median":              0.0284402,
		"time.now":                                 1536756976.225834,
		"time.up":                                  10413.923319,
		"time.elapsed":                             10413.923319,
		"mem.total.sbrk":                           8523776,
		"mem.cache.rrset":                          409036,
		"mem.cache.message":                        402097,
		"mem.mod.iterator":                         16532,
		"mem.mod.validator":                        130452,
		"histogram.000000.000000.to.000000.000001": 0,
		"histogram.000000.000001.to.000000.000002": 0,
		"histogram.000000.000002.to.000000.000004": 0,
		"histogram.000000.000004.to.000000.000008": 0,
		"histogram.000000.000008.to.000000.000016": 0,
		"histogram.000000.000016.to.000000.000032": 0,
		"histogram.000000.000032.to.000000.000064": 0,
		"histogram.000000.000064.to.000000.000128": 0,
		"histogram.000000.000128.to.000000.000256": 0,
		"histogram.000000.000256.to.000000.000512": 0,
		"histogram.000000.000512.to.000000.001024": 3,
		"histogram.000000.001024.to.000000.002048": 42,
		"histogram.000000.002048.to.000000.004096": 24,
		"histogram.000000.004096.to.000000.008192": 9,
		"histogram.000000.008192.to.000000.016384": 9,
		"histogram.000000.016384.to.000000.032768": 53,
		"histogram.000000.032768.to.000000.065536": 40,
		"histogram.000000.065536.to.000000.131072": 18,
		"histogram.000000.131072.to.000000.262144": 31,
		"histogram.000000.262144.to.000000.524288": 17,
		"histogram.000000.524288.to.000001.000000": 0,
		"histogram.000001.000000.to.000002.000000": 1,
		"histogram.000002.000000.to.000004.000000": 0,
		"histogram.000004.000000.to.000008.000000": 0,
		"histogram.000008.000000.to.000016.000000": 0,
		"histogram.000016.000000.to.000032.000000": 1,
		"histogram.000032.000000.to.000064.000000": 1,
		"histogram.000064.000000.to.000128.000000": 3,
		"histogram.000128.000000.to.000256.000000": 0,
		"histogram.000256.000000.to.000512.000000": 0,
		"histogram.000512.000000.to.001024.000000": 0,
		"histogram.001024.000000.to.002048.000000": 0,
		"histogram.002048.000000.to.004096.000000": 0,
		"histogram.004096.000000.to.008192.000000": 0,
		"histogram.008192.000000.to.016384.000000": 0,
		"histogram.016384.000000.to.032768.000000": 0,
		"histogram.032768.000000.to.065536.000000": 0,
		"histogram.065536.000000.to.131072.000000": 0,
		"histogram.131072.000000.to.262144.000000": 0,
		"histogram.262144.000000.to.524288.000000": 0,
		"num.query.type.A":                         745,
		"num.query.type.MX":                        2,
		"num.query.type.AAAA":                      705,
		"num.query.class.IN":                       1452,
		"num.query.opcode.QUERY":                   1452,
		"num.query.tcp":                            0,
		"num.query.ipv6":                           0,
		"num.query.flags.QR":                       0,
		"num.query.flags.AA":                       0,
		"num.query.flags.TC":                       0,
		"num.query.flags.RD":                       1452,
		"num.query.flags.RA":                       0,
		"num.query.flags.Z":                        0,
		"num.query.flags.AD":                       0,
		"num.query.flags.CD":                       0,
		"num.query.edns.present":                   0,
		"num.query.edns.DO":                        0,
		"num.answer.rcode.NOERROR":                 1416,
		"num.answer.rcode.SERVFAIL":                6,
		"num.answer.rcode.NXDOMAIN":                30,
		"num.answer.rcode.nodata":                  680,
		"num.answer.secure":                        22,
		"num.answer.bogus":                         0,
		"num.rrset.bogus":                          0,
		"unwanted.queries":                         0,
		"unwanted.replies":                         0,
	}

	for k := range expect {
		if expect[k] != result[k] {
			t.Errorf("%s does not match\nExpect: %v\nResult: %v\nresult: %v\n", k, expect[k], result[k], result)
		}
	}
}

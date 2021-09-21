package mpunbound

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

func (m *UnboundPlugin) unboundGraphDef() map[string]mp.Graphs {
	labelPrefix := strings.Title(strings.Replace(m.MetricKeyPrefix(), "unbound", "Unbound", -1))
	return map[string]mp.Graphs{
		"cache": {
			Label: labelPrefix + " Cache",
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "total.num.cachehits", Label: "Cache Hits"},
				{Name: "total.num.cachemiss", Label: "Cache Miss"},
				{Name: "total.num.prefetch", Label: "Prefetch"},
				{Name: "total.num.recursivereplies", Label: "Recursive Replies"},
			},
		},
		"requestlist": {
			Label: labelPrefix + " Request List",
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "total.requestlist.avg", Label: "Average"},
				{Name: "total.requestlist.max", Label: "Max"},
				{Name: "total.requestlist.overwritten", Label: "Overwritten"},
				{Name: "total.requestlist.exceeded", Label: "Exceeded"},
				{Name: "total.requestlist.current.all", Label: "Current All"},
				{Name: "total.requestlist.current.user", Label: "Current User"},
			},
		},
		"total.recursion": {
			Label: labelPrefix + " Recursion",
			Unit:  mp.UnitFloat,
			Metrics: []mp.Metrics{
				{Name: "total.recursion.time.avg", Label: "Time Average"},
				{Name: "total.recursion.time.median", Label: "Time Median"},
			},
		},
		// Enabled extended statistics
		"mem": {
			Label: labelPrefix + " Memory",
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "mem.total.sbrk.avg", Label: "Sbrk Average"},
				{Name: "mem.cache.rrset", Label: "Cache rrset"},
				{Name: "mem.cache.message", Label: "Cache Message"},
				{Name: "mem.mod.iterator", Label: "Mod Iterator"},
				{Name: "mem.mod.validator", Label: "Mod Validator"},
			},
		},
		"num.query.type": {
			Label: labelPrefix + " Query Type",
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "*", Label: "*"},
			},
		},
		"num.query.flags": {
			Label: labelPrefix + " Query Flags",
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "*", Label: "*"},
			},
		},
		"num.answer.rcode": {
			Label: labelPrefix + " Answer RCode",
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "*", Label: "*"},
			},
		},
		"histogram": {
			Label: labelPrefix + " Histogram",
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "*", Label: "*"},
			},
		},
	}
}

// UnboundPlugin mackerel plugin for Unbound
type UnboundPlugin struct {
	unboundControlCommands []string
	prefix                 string
}

// MetricKeyPrefix interface for PluginWithPrefix
func (m *UnboundPlugin) MetricKeyPrefix() string {
	if m.prefix == "" {
		m.prefix = "unbound"
	}
	return m.prefix
}

func (m *UnboundPlugin) parseStats(out string) (map[string]float64, error) {
	stat := make(map[string]float64)
	var err error

	for _, line := range strings.Split(out, "\n") {
		res := strings.Split(line, "=")
		if len(res) != 2 {
			continue
		}

		stat[res[0]], err = strconv.ParseFloat(res[1], 64)
		if err != nil {
			return nil, err
		}
	}
	return stat, nil
}

// FetchMetrics interface for mackerelplugin
func (m *UnboundPlugin) FetchMetrics() (map[string]float64, error) {
	out, err := exec.Command(m.unboundControlCommands[0], m.unboundControlCommands[1:]...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err, out)
	}

	stat, err := m.parseStats(string(out))
	if err != nil {
		return nil, err
	}

	return stat, nil
}

// GraphDefinition interface for mackerelplugin
func (m *UnboundPlugin) GraphDefinition() map[string]mp.Graphs {
	return m.unboundGraphDef()
}

// Do the plugin
func Do() {
	optUnboundControlPath := flag.String("unbound-control", "/usr/sbin/unbound-control", "Path to unbound-control")
	optUnboundControlOptions := flag.String("option", "", "unbound-control option")
	unboundControlNoReset := flag.Bool("no-reset", false, "use sub command stats_noreset to get stats")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	optMetricKeyPrefix := flag.String("metric-key-prefix", "unbound", "metric key prefix")
	flag.Parse()

	var unbound UnboundPlugin

	unbound.unboundControlCommands = []string{
		*optUnboundControlPath,
	}

	if *optUnboundControlOptions != "" {
		unbound.unboundControlCommands = append(unbound.unboundControlCommands, strings.Split(*optUnboundControlOptions, " ")...)
	}
	if *unboundControlNoReset {
		unbound.unboundControlCommands = append(unbound.unboundControlCommands, "stats_noreset")
	} else {
		unbound.unboundControlCommands = append(unbound.unboundControlCommands, "stats")
	}

	unbound.prefix = *optMetricKeyPrefix
	helper := mp.NewMackerelPlugin(&unbound)
	helper.Tempfile = *optTempfile

	helper.Run()
}

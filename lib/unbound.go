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
				{Name: "cachehits", Label: "Cache Hits"},
				{Name: "cachemiss", Label: "Cache Miss"},
				{Name: "prefetch", Label: "Prefetch"},
				{Name: "recursivereplies", Label: "Recursive Replies"},
			},
		},
		"requestlist": {
			Label: labelPrefix + " Request List",
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "avg", Label: "Average"},
				{Name: "max", Label: "Max"},
				{Name: "overwritten", Label: "Overwritten"},
				{Name: "exceeded", Label: "Exceeded"},
				{Name: "current.all", Label: "Current All"},
				{Name: "current.user", Label: "Current User"},
			},
		},
		"recursion": {
			Label: labelPrefix + " Recursion",
			Unit:  mp.UnitFloat,
			Metrics: []mp.Metrics{
				{Name: "time.avg", Label: "Time Average"},
				{Name: "time.median", Label: "Time Median"},
			},
		},
		// Enabled extended statistics
		"mem": {
			Label: labelPrefix + " Memory",
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "total.sbrk.avg", Label: "Sbrk Average"},
				{Name: "cache.rrset", Label: "Cache rrset"},
				{Name: "cache.message", Label: "Cache Message"},
				{Name: "mod.iterator", Label: "Mod Iterator"},
				{Name: "mod.validator", Label: "Mod Validator"},
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
	}
}

// UnboundPlugin mackerel plugin for Unbound
type UnboundPlugin struct {
	UnboundControlPath string
	Tempfile           string
	prefix             string
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
	out, err := exec.Command(m.UnboundControlPath, "stats").CombinedOutput()
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
	optTempfile := flag.String("tempfile", "", "Temp file name")
	optMetricKeyPrefix := flag.String("metric-key-prefix", "unbound", "metric key prefix")
	flag.Parse()

	var unbound UnboundPlugin

	unbound.UnboundControlPath = *optUnboundControlPath
	unbound.prefix = *optMetricKeyPrefix
	helper := mp.NewMackerelPlugin(&unbound)
	helper.Tempfile = *optTempfile

	helper.Run()
}

package main

import (
	"strings"
	"fmt"
)

func main() {
	str := `RRD: "rrd",
    KEYWORD: "keywordMonitor",
    STN: "statisTopn",
    SPMSEC: "spmSecond",
    GLOBAL_APP_CMS: "globalAppCms",
    DIRECT_SQL: "directSql",
    SPM: "spm",
    METRICS_SPM: "metricsSpm",
    AM_SYSTEM: "system",
    AM_SYS_TOP: "appMonitorSystemTop",
    PMSEC: "secondPatternMatch",
    MAP_PERF: "mapPerf",
    MM: "multiMinute",
    MW_TOMCAT: "tomcat",
    DEVICE: "device",
    CMS: "cms",
    MW: "middleware",
    AM: "appMonitor",
    SEC: "second",
    MW_JVM: "jvm",
    AM_ALARM_RULES: "amAlarmRule",
    APP_CORE_SERVICE: "appCoreService",
    AM_SYS_STAT: "appMonitorSystemStat",
    PROM: "prom",
    AM_SYS: "appMonitorSystem",
    DASHBOARD: "dashboard",
    MW_HSF_DETAIL: "hsfDetail",
    MW_NOTIFY: "notify",
    SM: "singleMinute",
    GC: "generalComp",
    PUGSEC: "purgeSecGroupby",
    PUSECSTN: "purgeSecStatisTopn",
    MW_TAIR: "tair",
    GCSEC: "generalCompSec",
    MW_HSF: "hsf",
    MW_STAT: "middlewareStat",
    MSEC: "multiSecond",
    PUG: "purgeGroupby",
    MW_METAQ: "metaq",
    AM_ERR: "appMonitorError",
    PUSTN: "purgeStatisTopn",
    APP_CMS: "appcms",
    METRICS: "metrics",
    RRDS: "rrds",
    AM_WEB: "web",
    METRICS_SEC: "secondMetrics",
    GREP: "grepEvent",
    TN: "topn",
    MW_TDDL: "tddl",
    PM: "patternMatch",
    UNION: "unionComp",
    SCRIPT_DS: "scriptDs",
    AM_MISTAKE: "appMonitorMistake",
    IOT_ALARM: "iotAlarm"`

	tmp := strings.Split(str, ",\n")
	for i := range tmp {
		itemt := strings.Split(tmp[i], ":")
		fmt.Printf("pluginTypeMap.put(%s, \"%s\");\r\n", strings.TrimSpace(itemt[1]), strings.TrimSpace(itemt[0]))
	}
}

package speedtest

import (
	"bytes"
	"fmt"
	"log/slog"
	"text/template"

	"github.com/jcocozza/go_speedtest/sql"
)

func LoadTemplate(speedTestResult SpeedTestResult) string {
	var err error
	tmp1 := template.New("insert_speedtest_template")
	tmp1, err = tmp1.Parse(sql.InsertSpeedTestTemplate)
	if err != nil {
		slog.Error(fmt.Sprint(err))
	}
	var buf bytes.Buffer
	err = tmp1.Execute(&buf, speedTestResult)
	if err != nil {
		slog.Error(fmt.Sprint(err))
	}
	return buf.String()

}
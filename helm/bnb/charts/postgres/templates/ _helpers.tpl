{{/*
Expand the name of the chart.
*/}}
{{- define "postgresql-chart.name" -}}
{{- .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Expand the name of the chart with the release name.
*/}}
{{- define "postgresql-chart.fullname" -}}
{{- printf "%s-%s" (include "postgresql-chart.name" .) .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create chart name and version as used by the label
*/}}
{{- define "postgresql-chart.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | trunc 63 | trimSuffix "-" -}}
{{- end -}}
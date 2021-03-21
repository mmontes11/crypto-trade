{{/*
Expand the name of the chart.
*/}}
{{- define "crypto-trade.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "crypto-trade.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Secret
*/}}
{{- define "crypto-trade.secret" -}}
{{- default .Chart.Name .Values.secretNameOverride  | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "crypto-trade.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "crypto-trade.labels" -}}
helm.sh/chart: {{ include "crypto-trade.chart" . }}
app.kubernetes.io/name: {{ include "crypto-trade.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Full name API
*/}}
{{- define "crypto-trade.fullnameApi" -}}
{{- printf "%s-%s" (include "crypto-trade.fullname" .) "api" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Selector labels API
*/}}
{{- define "crypto-trade.selectorLabelsApi" -}}
app.kubernetes.io/name: {{ include "crypto-trade.fullnameApi" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
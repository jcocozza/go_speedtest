INSERT INTO speedtest_results (
`type`,`timestamp`,
ping_jitter,ping_latency,ping_low,ping_high,
download_bandwidth,download_bytes,download_elapsed,
    download_latency_iqm,download_latency_low,download_latency_high,download_latency_jitter,
upload_bandwidth,upload_bytes,upload_elapsed,
    upload_latency_iqm,upload_latency_low,upload_latency_high,upload_latency_jitter,
packetLoss,isp,
interface_internalIp,interface_name,interface_macAddr,interface_isVpn,interface_externalIp,
server_id,server_host,server_port,server_name,server_location,server_country,server_ip,
result_id,result_url,result_persisted
) VALUES (
    "{{ .Type }}", "{{ .TimeStamp }}",
    {{ .Ping.Jitter }}, {{ .Ping.Latency }}, {{ .Ping.Low }}, {{ .Ping.High }},
    {{ .Download.Bandwidth }}, {{ .Download.Bytes }}, {{ .Download.Elapsed }},
        {{ .Download.Latency.Iqm }}, {{ .Download.Latency.Low }}, {{ .Download.Latency.High }}, {{ .Download.Latency.Jitter }},
    {{ .Upload.Bandwidth }}, {{ .Upload.Bytes }}, {{ .Upload.Elapsed }},
        {{ .Upload.Latency.Iqm }}, {{ .Upload.Latency.Low }}, {{ .Upload.Latency.High }}, {{ .Upload.Latency.Jitter }},
    {{ .PacketLoss }}, "{{ .Isp }}",
    "{{ .Interface.InternalIP }}", "{{ .Interface.Name }}", "{{ .Interface.MacAddr }}", {{ .Interface.IsVpn }}, "{{ .Interface.ExternalIP }}",
    {{ .Server.ID }}, "{{ .Server.Host }}", {{ .Server.Port }}, "{{ .Server.Name }}", "{{ .Server.Location }}", "{{ .Server.Country }}", "{{ .Server.IP }}",
    "{{ .Result.ID }}", "{{ .Result.URL }}", {{ .Result.Persisted }}
);
// ui/decoy_ui.go
func StartDecoyUI(mode string) {
	switch mode {
	case "torrent":
		showTorrentClient()
	case "cdn":
		showCDNStats()
	case "office":
		showFakeWorkDashboard()
	}
}
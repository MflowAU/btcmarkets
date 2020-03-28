package btcmarkets

// CreateNewReportData data struct to POST to Report API
type CreateNewReportData struct {
	Format string `json:"format"`
	Type   string `json:"type"`
}

// ReportServiceOps perform Report API actions
type ReportServiceOps struct {
	client *BTCMClient
}

package main

type FinanceReport struct {
}

type ReportMaker interface {
	MakeReport(*FinanceReport)
}

type EmailReportMaker struct {
}

func (s *EmailReportMaker) MakeReport(r *FinanceReport) {
	//...
}

type ReportSender interface {
	SendReport(*FinanceReport)
}

type EmailReportSender struct {
}

func (s *EmailReportSender) SendReport(r *FinanceReport) {
	//..email
}

type FileReportSender struct {
}

func (s *FileReportSender) SendReport(r *FinanceReport) {
	//..file
}

//확장을 하면서 변경은 안하는
// type NetworkReportSender struct {
// }

// func (s *NetworkReportSender) SendReport() {

// }

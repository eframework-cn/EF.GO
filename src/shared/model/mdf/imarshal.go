package mdf

type IMarshal interface {
	OnEncode()
	OnDecode()
}

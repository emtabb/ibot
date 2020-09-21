package ibot

type IBot interface {
	GetIpAddress() string
	SetIpAddress(string)
	GetPort() string
	SetPort(string)

	Description(string)
	Describe() string

	GetName() string
	SetName(string)

	GetStatus() *bool
	SetStatus(*bool)

	SetWork(Work)
	Working(Work)
}

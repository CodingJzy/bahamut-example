module github.com/CodingJzy/bahamut_example

go 1.14

replace github.com/valyala/tcplisten => github.com/CodingJzy/tcplisten v0.0.0-20200222015545-59d090b7d0b1

require (
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/mitchellh/copystructure v1.0.0
	go.aporeto.io/bahamut v1.120.0
	go.aporeto.io/elemental v1.110.0
)

module {{ .ModuleName }}

go {{ .GoVersion}}

require (
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/jcmturner/restclient v0.0.0-20180513084530-bfd3f693c938
	github.com/stretchr/testify v1.4.0 // indirect
)

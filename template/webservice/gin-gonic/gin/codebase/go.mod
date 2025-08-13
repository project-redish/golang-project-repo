module {{ .ModuleName }}

go {{ .GoVersion}}

require (
	github.com/gin-gonic/gin v1.10.1
	{{ .Logging.ImportPath }} {{ .Logging.Version }}
)

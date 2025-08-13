module {{ .ModuleName }}

go {{ .GoVersion}}

require (
    goji.io v2.0.2+incompatible
    {{ .Logging.ImportPath }} {{ .Logging.Version }}
)
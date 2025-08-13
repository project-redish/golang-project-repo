package consts

const (
	//TemplatePath the root folder for templates
	TemplatePath       = "/template"
	//CliTempatePath path for cli templates
	CliTempatePath     = "/template/cli"
	REST_TEMPLATE_PATH = "/template/rest"
	CLI_IN_REQUEST     = "cli"
	REST_IN_REQUEST    = "rest"
	OUTPUT_FOLDER      = "output"
	OUTPUT_ZIP         = "outputzip"
)

var (
	//SupportedCliLib list of support cli libraries
	SupportedCliLib = []string{"alecthomas", "spf13", "urfave"}

	//SupportedAppType list of support cli libraries
	SupportedAppType = []string{"cli"}

	//CliConfigurableFiles file path which will be edited by template parser
	CliConfigurableFiles map[string][]string
)

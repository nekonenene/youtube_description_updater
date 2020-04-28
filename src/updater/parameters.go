package updater

import (
	"flag"
	"log"
)

type parameters struct {
	CredentialFilePath string
	TokenFilePath      string
	TargetString       string
	ReplacementString  string
}

var params parameters

func ParseParameters() {
	flag.StringVar(&params.TargetString, "target-string", "", "(*Required) target string")
	flag.StringVar(&params.ReplacementString, "replacement-string", "", "(*Required) replacement string")
	flag.StringVar(&params.CredentialFilePath, "credential-file", defaultCredentialFilePath, "(Option) Download client_secret_*.json from Google Developer Console, and specifiled path")
	flag.StringVar(&params.TokenFilePath, "token-file", defaultTokenFilePath, "(Option) If you want to use your token file, specifiled path")
	flag.Parse()

	if params.TargetString == "" {
		log.Fatalln("-target-string is required")
	}
	if params.ReplacementString == "" {
		log.Fatalln("-replacement-string is required")
	}
}

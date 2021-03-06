package auth

import (
	"log"
)

type Credential struct {
	ConsumerKey    string
	ConsumerSecret string
}

var (
	PublicCredentials = map[string]Credential{
		"Twitter for iPhone": Credential{
			ConsumerKey:    "IQKbtAYlXLripLGPWd0HUA",
			ConsumerSecret: "GgDYlkSvaPxGxC4X8liwpUoqKwwr3lCADbz8A7ADU",
		},
		"Twitter for Android": Credential{
			ConsumerKey:    "3nVuSoBZnx6U4vzUxf5w",
			ConsumerSecret: "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys",
		},
		"Twitter for Google TV": Credential{
			ConsumerKey:    "iAtYJ4HpUVfIUoNnif1DA",
			ConsumerSecret: "172fOpzuZoYzNYaU3mMYvE8m8MEyLbztOdbrUolU",
		},
		"Twitter for iPad": Credential{
			ConsumerKey:    "CjulERsDeqhhjSme66ECg",
			ConsumerSecret: "IQWdVyqFxghAtURHGeGiWAsmCAGmdW3WmbEx6Hck",
		},
	}
)

func DefaultCredential() *Credential {
	return CredentialByClientName("Twitter for iPhone")
}

func CredentialByClientName(clientName string) *Credential {
	credential, hasCredential := PublicCredentials[clientName]
	if hasCredential {
		return &credential
	} else {
		log.Fatal("Not registered client name:", clientName)
		return nil
	}
}

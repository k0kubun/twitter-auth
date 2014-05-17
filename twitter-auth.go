package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/k0kubun/twitter-auth/auth"
)

type Options struct {
	ConsumerKey     string `short:"k" long:"consumer-key" description:"use selected consumer key"`
	ConsumerSecret  string `short:"s" long:"consumer-secret" description:"use selected consumer secret"`
	ClientName      string `short:"n" long:"client-name" description:"use public credentials"`
	ShowClientNames bool   `short:"l" long:"list" default:"false" description:"show available public client names"`
}

func authenticate(options *Options) {
	var credential *auth.Credential

	if len(options.ConsumerKey) > 0 && len(options.ConsumerSecret) > 0 {
		credential = &auth.Credential{
			ConsumerKey:    options.ConsumerKey,
			ConsumerSecret: options.ConsumerSecret,
		}
	} else if len(options.ClientName) > 0 {
		credential = auth.CredentialByClientName(options.ClientName)
	} else {
		credential = auth.DefaultCredential()
	}

	accessToken := auth.Authenticate(credential)
	println("CONSUMER_KEY        :", credential.ConsumerKey)
	println("CONSUMER_SECRET     :", credential.ConsumerSecret)
	println("ACCESS_TOKEN        :", accessToken.Token)
	println("ACCESS_TOKEN_SECRET :", accessToken.Secret)
}

func main() {
	options := new(Options)
	if _, err := flags.Parse(options); err != nil {
		return
	}

	if options.ShowClientNames {
		println("Available clients:")
		for clientName, _ := range auth.PublicCredentials {
			println(" ", clientName)
		}
		return
	}
	authenticate(options)
}

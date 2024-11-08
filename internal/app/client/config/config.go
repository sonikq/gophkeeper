package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	SaveCommandFlgSet *flag.FlagSet
	CredentialsFlgSet *flag.FlagSet
	TextFlgSet        *flag.FlagSet
	BinaryFlgSet      *flag.FlagSet
	LoadCommand       *flag.FlagSet
	LoginCommand      *flag.FlagSet

	NewUserLogin    string
	NewUserPassword string
	UserLogin       string
	UserPassword    string

	Cred   CredInfo
	Text   Info
	Binary Info
	Card   CardInfo
}

type CredInfo struct {
	Login    string
	Password string
	Info
}

type CardInfo struct {
	Number  string
	Owner   string
	Expires string
	Secret  string
	PIN     string
	Info
}

type Info struct {
	Data  string
	Alias string
	Meta  string
	Token string
}

func Load() (Config, error) {
	var cfg = Config{}

	saveCommand := flag.NewFlagSet("save", flag.ExitOnError)
	cfg.SaveCommandFlgSet = saveCommand
	registerCommand := flag.NewFlagSet("register", flag.ExitOnError)
	loadCommand := flag.NewFlagSet("load", flag.ExitOnError)
	cfg.LoadCommand = loadCommand
	loginCommand := flag.NewFlagSet("login", flag.ExitOnError)
	cfg.LoginCommand = loginCommand

	credentialsFlags := flag.NewFlagSet("credentials", flag.ExitOnError)
	cfg.CredentialsFlgSet = credentialsFlags
	textFlags := flag.NewFlagSet("text", flag.ExitOnError)
	cfg.TextFlgSet = textFlags
	binaryFlags := flag.NewFlagSet("binary", flag.ExitOnError)
	cfg.BinaryFlgSet = binaryFlags
	bankCardFlags := flag.NewFlagSet("card", flag.ExitOnError)

	newUserLogin := registerCommand.String("login", "", "new user login")
	cfg.NewUserLogin = *newUserLogin
	newUserPassword := registerCommand.String("password", "", "new user password")
	cfg.NewUserPassword = *newUserPassword

	userLogin := loginCommand.String("login", "", "user login")
	cfg.UserLogin = *userLogin
	userPassword := loginCommand.String("password", "", "user password")
	cfg.UserPassword = *userPassword

	credLogin := credentialsFlags.String("login", "", "credential login")
	cfg.Cred.Login = *credLogin
	credPassword := credentialsFlags.String("password", "", "credential password")
	cfg.Cred.Password = *credPassword
	credAlias := credentialsFlags.String("alias", "", "credentials alias")
	cfg.Cred.Alias = *credAlias
	credMeta := credentialsFlags.String("meta", "", "credential meta")
	cfg.Cred.Meta = *credMeta
	credToken := credentialsFlags.String("token", "", "token")
	cfg.Cred.Token = *credToken

	textData := textFlags.String("data", "", "text to save")
	cfg.Text.Data = *textData
	textAlias := textFlags.String("alias", "", "texts alias")
	cfg.Text.Alias = *textAlias
	textMeta := textFlags.String("meta", "", "text meta")
	cfg.Text.Meta = *textMeta
	textToken := textFlags.String("token", "", "token")
	cfg.Text.Token = *textToken

	binData := binaryFlags.String("data", "", "bin data in encoded form to save")
	cfg.Binary.Data = *binData
	binAlias := binaryFlags.String("alias", "", "binary data alias")
	cfg.Binary.Alias = *binAlias
	binMeta := binaryFlags.String("meta", "", "bin meta")
	cfg.Binary.Meta = *binMeta
	binToken := binaryFlags.String("token", "", "token")
	cfg.Binary.Token = *binToken

	cardNumber := bankCardFlags.String("number", "", "bank card number")
	cfg.Card.Number = *cardNumber
	cardOwner := bankCardFlags.String("owner", "", "bank card owner")
	cfg.Card.Owner = *cardOwner
	cardExpires := bankCardFlags.String("expires at", "", "bank card expiration date")
	cfg.Card.Expires = *cardExpires
	cardSecret := bankCardFlags.String("secret key", "", "bank card s ecret key")
	cfg.Card.Secret = *cardSecret
	cardPIN := bankCardFlags.String("pin", "", "bank card PIN code")
	cfg.Card.PIN = *cardPIN
	cardAlias := bankCardFlags.String("alias", "", "cards alias")
	cfg.Card.Alias = *cardAlias
	cardMeta := bankCardFlags.String("meta", "", "bank card meta")
	cfg.Card.Meta = *cardMeta
	cardToken := bankCardFlags.String("token", "", "token")
	cfg.Card.Token = *cardToken

	if len(os.Args) < 2 {
		fmt.Println("save/register/load/login subcommand required")
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		fmt.Println("credentials/text/binary/ subcommand required")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "save":
		err := saveCommand.Parse(os.Args[2:])
		if err != nil {
			return Config{}, err
		}
	case "load":
		err := loadCommand.Parse(os.Args[2:])
		if err != nil {
			return Config{}, err
		}
	case "login":
		err := loginCommand.Parse(os.Args[2:])
		if err != nil {
			return Config{}, err
		}
	case "register":
		err := registerCommand.Parse(os.Args[2:])
		if err != nil {
			return Config{}, err
		}
	default:
		flag.PrintDefaults()
		return Config{}, errors.New("invalid type of requested data")
	}

	switch os.Args[2] {
	case "credentials":
		err := credentialsFlags.Parse(os.Args[3:])
		if err != nil {
			return Config{}, err
		}
	case "text":
		err := textFlags.Parse(os.Args[3:])
		if err != nil {
			return Config{}, err
		}
	case "binary":
		err := binaryFlags.Parse(os.Args[3:])
		if err != nil {
			return Config{}, err
		}
	case "card":
		err := bankCardFlags.Parse(os.Args[3:])
		if err != nil {
			return Config{}, err
		}
	}

	return cfg, nil
}

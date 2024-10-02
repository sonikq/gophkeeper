package client

import (
	"context"
	"fmt"
	"github.com/sonikq/gophkeeper/internal/app/client/config"
	pb "github.com/sonikq/gophkeeper/internal/delivery/grpc/v1"
	"log"
	"time"
)

var buildVersion = "N/A"
var buildDate = "N/A"
var buildCommit = "N/A"

// printBuildInfo prints the build information.
func printBuildInfo() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
}

type GophKeeperClient struct {
	Client pb.GophKeeperHandlerClient
}

func New() GophKeeperClient {
	return GophKeeperClient{}
}

func (c *GophKeeperClient) Run() {
	printBuildInfo()
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to initialize config: %s", err.Error())
	}

	if cfg.SaveCommandFlgSet.Parsed() {
		if cfg.CredentialsFlgSet.Parsed() {
			err = c.SaveCredentials(ctx, cfg.Cred.Login, cfg.Cred.Password, cfg.Cred.Meta,
				cfg.Cred.Alias, cfg.Cred.Token)
			if err != nil {
				log.Println(err)
			}
		} else if cfg.TextFlgSet.Parsed() {
			err = c.SaveText(ctx, cfg.Text.Data, cfg.Text.Meta, cfg.Text.Alias, cfg.Text.Token)
			if err != nil {
				log.Println(err)
			}
		} else if cfg.BinaryFlgSet.Parsed() {
			err = c.SaveBinary(ctx, cfg.Binary.Data, cfg.Binary.Meta, cfg.Binary.Alias, cfg.Binary.Token)
			if err != nil {
				log.Println(err)
			}
		} else {
			err = c.SaveCard(ctx, cfg.Card.Number, cfg.Card.Owner, cfg.Card.Expires, cfg.Card.Secret,
				cfg.Card.PIN, cfg.Card.Meta, cfg.Card.Alias, cfg.Card.Token)
			if err != nil {
				log.Println(err)
			}
		}
	} else if cfg.LoadCommand.Parsed() {
		if cfg.CredentialsFlgSet.Parsed() {
			c.LoadCredentials(ctx, cfg.Cred.Alias, cfg.Cred.Token)
		} else if cfg.TextFlgSet.Parsed() {
			c.LoadText(ctx, cfg.Text.Alias, cfg.Text.Token)
		} else if cfg.BinaryFlgSet.Parsed() {
			c.LoadBinary(ctx, cfg.Binary.Alias, cfg.Binary.Token)
		} else {
			c.LoadCard(ctx, cfg.Card.Alias, cfg.Card.Token)
		}
	} else if cfg.LoginCommand.Parsed() {
		c.Login(ctx, cfg.UserLogin, cfg.UserPassword)
	} else {
		c.Register(ctx, cfg.NewUserLogin, cfg.NewUserPassword)
	}
}

func (c *GophKeeperClient) SaveText(parent context.Context, text, meta, alias, token string) error {
	if token == "" {
		fmt.Println("register user or authorize before saving data")
		return nil
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	_, err := c.Client.SaveText(ctx, &pb.SaveTextDataRequest{
		Token: token,
		Data: &pb.TextData{
			Uuid: alias,
			Data: text,
			Meta: &pb.Meta{
				Content: meta,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *GophKeeperClient) SaveBinary(parent context.Context, binData, meta, alias, token string) error {
	if token == "" {
		fmt.Println("register user or authorize before saving data")
		return nil
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	_, err := c.Client.SaveBinary(ctx, &pb.SaveBinaryDataRequest{
		Token: token,
		Data: &pb.BinaryData{
			Uuid: alias,
			Data: []byte(binData),
			Meta: &pb.Meta{
				Content: meta,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *GophKeeperClient) SaveCard(parent context.Context, cardNumber, cardOwner, cardExpires, cardSecret, cardPIN, meta, alias, token string) error {
	if token == "" {
		fmt.Println("register user or authorize before saving data")
		return nil
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	_, err := c.Client.SaveBankCard(ctx, &pb.SaveBankCardDataRequest{
		Token: token,
		Data: &pb.BankCardData{
			Uuid:       alias,
			Number:     cardNumber,
			Owner:      cardOwner,
			ExpiresAt:  cardExpires,
			SecretCode: cardSecret,
			PinCode:    cardPIN,
			Meta: &pb.Meta{
				Content: meta,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *GophKeeperClient) SaveCredentials(parent context.Context, login, password, meta, alias, token string) error {
	if token == "" {
		fmt.Println("register user or authorize before saving data")
		return nil
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	_, err := c.Client.SaveCredentials(ctx, &pb.SaveCredentialsDataRequest{
		Token: token,
		Data: &pb.CredentialsData{
			Uuid:     alias,
			Login:    login,
			Password: password,
			Meta: &pb.Meta{
				Content: meta,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *GophKeeperClient) LoadCredentials(parent context.Context, alias, token string) string {
	if token == "" {
		return "register user or authorize before saving data"
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	req, err := c.Client.LoadCredentials(ctx, &pb.LoadCredentialsDataRequest{
		Token: token,
		Uuid:  alias,
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return fmt.Sprintln(req.Data)
}

func (c *GophKeeperClient) LoadText(parent context.Context, alias, token string) string {
	if token == "" {
		return "register user or authorize before saving data"
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	req, err := c.Client.LoadText(ctx, &pb.LoadTextDataRequest{
		Token: token,
		Uuid:  alias,
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return fmt.Sprintln(req.Data)
}

func (c *GophKeeperClient) LoadBinary(parent context.Context, alias, token string) string {
	if token == "" {
		return "register user or authorize before saving data"
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	req, err := c.Client.LoadBinary(ctx, &pb.LoadBinaryDataRequest{
		Token: token,
		Uuid:  alias,
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return fmt.Sprintln(req.Data)
}

func (c *GophKeeperClient) LoadCard(parent context.Context, alias, token string) string {
	if token == "" {
		return "register user or authorize before saving data"
	}
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	req, err := c.Client.LoadBankCard(ctx, &pb.LoadBankCardDataRequest{
		Token: token,
		Uuid:  alias,
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return fmt.Sprintln(req.Data)
}

func (c *GophKeeperClient) Login(parent context.Context, login, password string) string {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	req, err := c.Client.LoginUser(ctx, &pb.LoginUserRequest{
		User: &pb.User{
			Login:    login,
			Password: password,
		},
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return fmt.Sprintf("user authorized. token: %s", req.Token)
}

func (c *GophKeeperClient) Register(parent context.Context, login, password string) string {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	req, err := c.Client.RegisterUser(ctx, &pb.RegisterUserRequest{
		User: &pb.User{
			Login:    login,
			Password: password,
		},
	})
	if err != nil {
		return err.Error()
	}
	if req.Error != "" {
		return req.Error
	}
	return "user registered, please login to work further"
}

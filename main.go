package main

import (
	"os"
	"time"

	"github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
)

func main() {
	client, err := api.NewClient(&api.Config{
		Address: os.Getenv("VAULT_ADDR"),
		Timeout: time.Second * 60,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	initialized, err := client.Sys().InitStatus()
	if err != nil {
		logrus.Fatal(err)
	}

	if initialized {
		logrus.Info("Vault is initialized")
		status, err := client.Sys().SealStatus()
		if err != nil {
			logrus.Fatal(err)
		}

		if status.Sealed {
			logrus.Info("Vault is sealed. Unsealing...")
			_, err := client.Sys().Unseal(os.Getenv("UNSEAL_KEY"))
			if err != nil {
				logrus.Fatal(err)
			}
		} else {
			logrus.Info("Vault is already unlocked")
		}

		logrus.Info("Sleeping for 1h now")
		time.Sleep(1 * time.Hour)
		return
	}
	logrus.Info("Not initialized. Initializing!")

	init, err := client.Sys().Init(&api.InitRequest{
		SecretShares:    1,
		SecretThreshold: 1,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Vault initialized successfully.")
	logrus.Infof("Root token: %s", init.RootToken)
	for _, key := range init.Keys {
		logrus.Infof("Please set the following UNSEAL_KEY: %s", key)
	}

	logrus.Info("Sleeping forever now. Please kill this container when configured correctly.")
	i := make(chan int)
	<-i
}

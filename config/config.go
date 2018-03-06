// Package config provides access to run-time configuration.
package config

import (
	"context"

	"google.golang.org/appengine/datastore"
)

type credentials struct {
	SecretKey   []byte `datastore:",noindex"`
	AccessToken string `datastore:",noindex"`
}

var cachedCreds *credentials

func loadCreds(ctx context.Context) error {
	if cachedCreds != nil {
		return nil
	}

	k := datastore.NewKey(ctx, "credentials", "singleton", 0, nil)

	var c credentials
	if err := datastore.Get(ctx, k, &c); err != nil {
		return err
	}

	cachedCreds = &c
	return nil
}

// SecretKey returns the shared secret used to verify the signature on Github events.
func SecretKey(ctx context.Context) ([]byte, error) {
	if err := loadCreds(ctx); err != nil {
		return nil, err
	}

	return cachedCreds.SecretKey, nil
}

// AccessToken returns the access token used to authenticate requests to Github.
func AccessToken(ctx context.Context) (string, error) {
	if err := loadCreds(ctx); err != nil {
		return "", err
	}

	return cachedCreds.AccessToken, nil
}

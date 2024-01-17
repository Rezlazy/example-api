package config

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type PrivateKey rsa.PrivateKey

func (pk *PrivateKey) UnmarshalText(content []byte) error {
	privateKeyPtr, err := jwt.ParseRSAPrivateKeyFromPEM(content)
	if err != nil {
		return fmt.Errorf("parse rsa private key: %w", err)
	}

	privateKey := *privateKeyPtr
	*pk = PrivateKey(privateKey)

	return nil
}

type PublicKey rsa.PublicKey

func (pk *PublicKey) UnmarshalText(content []byte) error {
	publicKeyPtr, err := jwt.ParseRSAPublicKeyFromPEM(content)
	if err != nil {
		return fmt.Errorf("parse rsa public key: %w", err)
	}

	publicKey := *publicKeyPtr
	*pk = PublicKey(publicKey)

	return nil
}

type Auth struct {
	Email      string      `env:"EMAIl,required"`
	Password   string      `env:"PASSWORD,required"`
	PrivateKey *PrivateKey `env:"PRIVATE_KEY,file" envDefault:"config/auth_private_key"`
	PublicKey  *PublicKey  `env:"PUBLIC_KEY,file" envDefault:"config/auth_public_key"`
}

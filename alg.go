package jwt

import (
	"crypto"
	"errors"
)

// The builtin signing available algorithms.
var (
	// None for unsecured JWTs.
	// An unsecured JWT may be fit for client-side use.
	// For instance, if the session ID is a hard-to-guess number, and
	// the rest of the data is only used by the client for constructing a
	// view, the use of a signature is superfluous.
	// This data can be used by a single-page web application
	// to construct a view with the "pretty" name for the user
	// without hitting the backend while he gets
	// redirected to his last visited page. Even if a malicious user
	// were to modify this data he or she would gain nothing.
	// Example payload:
	//  {
	//    "sub": "user123",
	//    "session": "ch72gsb320000udocl363eofy",
	//    "name": "Pretty Name",
	//    "lastpage": "/views/settings"
	//  }
	NONE Alg = &algNONE{}
	// HMAC-SHA signing algorithms.
	// Keys should be type of []byte.
	HS256 Alg = &algHMAC{"HS256", crypto.SHA256}
	HS384 Alg = &algHMAC{"HS384", crypto.SHA384}
	HS512 Alg = &algHMAC{"HS512", crypto.SHA512}
	// RSA signing algorithms.
	// Sign   key: *rsa.PublicKey (or *rsa.PrivateKey with its PublicKey filled)
	// Verify key: *rsa.PrivateKey
	//
	// 	Signing and verifying RS256 signed tokens is just as easy.
	// The only difference lies in the use of a private/public key pair rather than a shared secret.
	// There are many ways to create RSA keys.
	// OpenSSL is one of the most popular libraries for key creation and management:
	// Generate a private key:
	// $ openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048
	// Derive the public key from the private key:
	// $ openssl rsa -pubout -in private_key.pem -out public_key.pem
	RS256 Alg = &algRSA{"RS256", crypto.SHA256}
	RS384 Alg = &algRSA{"RS384", crypto.SHA384}
	RS512 Alg = &algRSA{"RS512", crypto.SHA512}
	/* TODO:
	• ES384, ES512: SHA-384 and SHA-512 variations of the ES256 algorithm.
	• PS256, PS384, PS512: RSASSA-PSS + MGF1 with SHA256/384/512 variants.
	*/
)

var (
	// ErrTokenSignature indicates that the verification failed.
	ErrTokenSignature = errors.New("invalid token signature")
	// ErrInvalidKey indicates that an algorithm required secret key is not a valid type.
	ErrInvalidKey = errors.New("invalid key")
)

// Alg represents a signing and verifying algorithm.
type Alg interface {
	// Name should return the "alg" JWT field.
	Name() string
	// Sign should return the signed data based on the given
	// full header and payload data and a secret key.
	Sign(headerAndPayload []byte, key interface{}) ([]byte, error)
	// Verify should verify the JWT "signature" (base64-decoded) against
	// the header and payload data's one based on the given secret key.
	Verify(headerAndPayload []byte, signature []byte, key interface{}) error
	// Note:
	// some signing algorithms may be asymmetric,
	// so we accept the headerAndPayload as it's, instead of a Sign's result.
}
package encryption

import "errors"

var (
	ErrGeneratingNonce    = errors.New("failed to generate nonce")
	ErrDecodingBase64Data = errors.New("failed to decode base64 data")
	ErrCreatingCipher     = errors.New("failed to create cipher")
	ErrCreatingGCMCipher  = errors.New("failed to create GCM cipher")
	ErrCipherTextTooShort = errors.New("ciphertext too short")
	ErrDecryptingData     = errors.New("failed to decrypt data")
)

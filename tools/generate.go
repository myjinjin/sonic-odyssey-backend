package tools

//go:generate mockery --dir ../internal/domain/repositories --name UserRepository --output ../internal/usecase/mocks
//go:generate mockery --dir ../internal/domain/repositories --name PasswordResetFlowRepository --output ../internal/usecase/mocks
//go:generate mockery --dir ../infrastructure/logging --name Logger --output ../internal/usecase/mocks
//go:generate mockery --dir ../infrastructure/hash --name PasswordHasher --output ../internal/usecase/mocks
//go:generate mockery --dir ../infrastructure/hash --name EmailHasher --output ../internal/usecase/mocks
//go:generate mockery --dir ../infrastructure/encryption --name Encryptor --output ../internal/usecase/mocks
//go:generate mockery --dir ../infrastructure/email --name EmailSender --output ../internal/usecase/mocks
//go:generate mockery --dir ../internal/usecase --name UserUsecase --output ../internal/controller/http/mocks
//go:generate mockery --dir ../internal/usecase --name MusicUsecase --output ../internal/controller/http/mocks
//go:generate mockery --dir ../infrastructure/spotifyclient --name SpotifyClient --output ../internal/usecase/mocks

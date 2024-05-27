package tools

//go:generate mockery --dir ../internal/domain/repositories --name UserRepository --output ../internal/usecase/mocks
//go:generate mockery --dir ../infrastructure/hash --name PasswordHasher --output ../internal/usecase/mocks
//go:generate mockery --dir ../infrastructure/hash --name EmailHasher --output ../internal/usecase/mocks
//go:generate mockery --dir ../infrastructure/encryption --name Encryptor --output ../internal/usecase/mocks
//go:generate mockery --dir ../infrastructure/email --name EmailSender --output ../internal/usecase/mocks
//go:generate mockery --dir ../internal/usecase --name UserUsecase --output ../internal/controller/http/mocks
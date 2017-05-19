package oauth2

import (
	"context"

	"github.com/MatthewHartstonge/fosite"
)

type RefreshTokenGrantStorage interface {
	RefreshTokenStorage
	PersistRefreshTokenGrantSession(ctx context.Context, requestRefreshSignature, accessSignature, refreshSignature string, request fosite.Requester) error
}

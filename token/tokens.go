package token

import (
	"gopkg.in/oauth2.v3"
	"time"
)

func From(info oauth2.TokenInfo) *Info {
	return &Info{
		ClientID:         info.GetClientID(),
		UserID:           info.GetUserID(),
		RedirectURI:      info.GetRedirectURI(),
		Scope:            info.GetScope(),
		Code:             info.GetCode(),
		CodeCreatedAt:    info.GetCodeCreateAt(),
		CodeExpiresIn:    info.GetCodeExpiresIn(),
		Access:           info.GetAccess(),
		AccessCreatedAt:  info.GetAccessCreateAt(),
		AccessExpiresIn:  info.GetAccessExpiresIn(),
		Refresh:          info.GetRefresh(),
		RefreshCreatedAt: info.GetRefreshCreateAt(),
		RefreshExpiresIn: info.GetRefreshExpiresIn(),
	}
}

type Info struct {
	ClientID         string
	UserID           string
	RedirectURI      string
	Scope            string
	Code             string
	CodeCreatedAt    time.Time
	CodeExpiresIn    time.Duration
	Access           string
	AccessCreatedAt  time.Time
	AccessExpiresIn  time.Duration
	Refresh          string
	RefreshCreatedAt time.Time
	RefreshExpiresIn time.Duration
}

func (i *Info) New() oauth2.TokenInfo {
	return From(i)
}

func (i *Info) GetClientID() string {
	return i.ClientID
}

func (i *Info) SetClientID(ID string) {
	i.ClientID = ID
}

func (i *Info) GetUserID() string {
	return i.UserID
}

func (i *Info) SetUserID(ID string) {
	i.UserID = ID
}

func (i *Info) GetRedirectURI() string {
	return i.RedirectURI
}

func (i *Info) SetRedirectURI(URI string) {
	i.RedirectURI = URI
}

func (i *Info) GetScope() string {
	return i.Scope
}

func (i *Info) SetScope(s string) {
	i.Scope = s
}

func (i *Info) GetCode() string {
	return i.Code
}

func (i *Info) SetCode(c string) {
	i.Code = c
}

func (i *Info) GetCodeCreateAt() time.Time {
	return i.CodeCreatedAt
}

func (i *Info) SetCodeCreateAt(t time.Time) {
	i.CodeCreatedAt = t
}

func (i *Info) GetCodeExpiresIn() time.Duration {
	return i.CodeExpiresIn
}

func (i *Info) SetCodeExpiresIn(d time.Duration) {
	i.CodeExpiresIn = d
}

func (i *Info) GetAccess() string {
	return i.Access
}

func (i *Info) SetAccess(access string) {
	i.Access = access
}

func (i *Info) GetAccessCreateAt() time.Time {
	return i.AccessCreatedAt
}

func (i *Info) SetAccessCreateAt(t time.Time) {
	i.AccessCreatedAt = t
}

func (i *Info) GetAccessExpiresIn() time.Duration {
	return i.AccessExpiresIn
}

func (i *Info) SetAccessExpiresIn(d time.Duration) {
	i.AccessExpiresIn = d
}

func (i *Info) GetRefresh() string {
	return i.Refresh
}

func (i *Info) SetRefresh(r string) {
	i.Refresh = r
}

func (i *Info) GetRefreshCreateAt() time.Time {
	return i.RefreshCreatedAt
}

func (i *Info) SetRefreshCreateAt(t time.Time) {
	i.RefreshCreatedAt = t
}

func (i *Info) GetRefreshExpiresIn() time.Duration {
	return i.RefreshExpiresIn
}

func (i *Info) SetRefreshExpiresIn(d time.Duration) {
	i.RefreshExpiresIn = d
}

package main

type Login struct {
	TenantId string `json:"tenantId"`
	ClientId string `json:"clientId"`
	Secret   string `json:"secret"`
}

type Config struct {
	VaultURI string `json:"vaultURI"`
	Login    `json:"login"`
}

var (
	config Config
)

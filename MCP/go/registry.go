package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_certificates "github.com/apimanagementclient/mcp-server/tools/certificates"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_certificates.CreateCertificates_deleteTool(cfg),
		tools_certificates.CreateCertificates_getTool(cfg),
		tools_certificates.CreateCertificates_createorupdateTool(cfg),
		tools_certificates.CreateCertificates_listbyserviceTool(cfg),
	}
}

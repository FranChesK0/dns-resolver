package cli

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	BoldText      = lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("63"))
	resolvedStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("63"))
)

func QueryingMessage(nameServer string, domainName string) string {
	return fmt.Sprintf("Querying %s for %s\n", BoldText.Render(nameServer), BoldText.Render(domainName))
}

func ResolvedMessage(domainName string, ipAddr string) string {
	return resolvedStyle.Render(fmt.Sprintf("%s -> %s", domainName, ipAddr))
}

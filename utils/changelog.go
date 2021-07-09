package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/desmos-labs/changeset/types"
)

// CollectChanges collects the given entries as a types.TypeChanges instance
func CollectChanges(entries []*types.Entry) types.TypeChanges {
	var typesChanges = types.TypeChanges{}
	for _, entry := range entries {
		existingChanges := typesChanges[entry.Type]
		if existingChanges == nil {
			existingChanges = types.ModuleChanges{}
		}

		existingChanges[entry.Module] = append(existingChanges[entry.Module], entry)
		typesChanges[entry.Type] = existingChanges
	}

	return typesChanges
}

// ConvertToMarkdown converts the given changelog to a Markdown representation
func ConvertToMarkdown(config *types.Config, changelog *types.ChangeLog) (string, error) {
	output := fmt.Sprintf("## %s\n", changelog.Version)

	for typeCode, changes := range changelog.Changes {
		changesType, err := config.GetTypeByCode(typeCode)
		if err != nil {
			return "", err
		}

		output += fmt.Sprintf("### %s\n", strings.Title(changesType.Code.String()))
		for moduleID, entries := range changes {
			module, err := config.GetModuleByID(moduleID)
			if err != nil {
				return "", err
			}

			output += fmt.Sprintf("#### %s\n", module.Description)
			for _, entry := range entries {
				output += fmt.Sprintf("* ([%[1]d](%[2]s/pull/%[1]d)) %[3]s\n",
					entry.PullRequestID, config.GitHubRepo, entry.Description)
			}
			output += "\n"
		}
	}

	return output, nil
}

// UpdateChangelog reads the CHANGELOG file located at the given path, updates it by replacing the
// Unreleased section with the provided data, and returns the updated contents.
func UpdateChangelog(data string, path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("error while opening changelog file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	var unreleasedTitleLine, nextVersionLine int

	var index = 0
	for scanner.Scan() {
		line := scanner.Text()

		isVersionTitle := strings.HasPrefix(line, "## ")
		if isVersionTitle && strings.Contains(strings.ToLower(line), "unreleased") {
			unreleasedTitleLine = index
		} else if isVersionTitle && nextVersionLine == 0 {
			nextVersionLine = index
		}

		lines = append(lines, line)
		index++
	}

	err = scanner.Err()
	if err != nil {
		return "", err
	}

	updatedLines := append(lines[:unreleasedTitleLine], lines[nextVersionLine:]...)
	updatedText := strings.Join(updatedLines, "\n")
	return strings.Join([]string{data, updatedText}, "\n"), nil
}

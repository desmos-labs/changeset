package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
func ConvertToMarkdown(config *types.Config, changelog *types.ChangeLog) ([]string, error) {
	output := []string{
		fmt.Sprintf("## Version %s", changelog.Version),
	}

	for _, changeType := range config.Types {
		moduleChanges := changelog.Changes[changeType.Code]
		if len(moduleChanges) == 0 {
			continue
		}

		changesType, err := config.GetTypeByCode(changeType.Code)
		if err != nil {
			return nil, err
		}

		if changesType.Hide {
			continue
		}

		output = append(output, fmt.Sprintf("### %s", strings.Title(changesType.Title)))

		for _, module := range config.Modules {
			// Get the entries for this module
			entries := moduleChanges.GetEntriesByModule(module.Code)
			if len(entries) == 0 {
				continue
			}

			if module.Code != types.ModuleNone.Code {
				output = append(output, fmt.Sprintf("#### %s", module.Description))
			}

			// Sort the entries based on the pull request number
			sort.SliceStable(entries, func(i, j int) bool {
				return entries[i].PullRequestID < entries[j].PullRequestID
			})

			for _, entry := range entries {
				output = append(output,
					fmt.Sprintf("- ([\\#%[1]d](%[2]s/pull/%[1]d)) %[3]s",
						entry.PullRequestID, config.GetRepoURL(), entry.Description),
				)
			}
			output = append(output, "")
		}
	}

	return output, nil
}

// UpdateChangelog reads the CHANGELOG file located at the given path, updates it by replacing the
// Unreleased section with the provided data, and returns the updated contents.
func UpdateChangelog(data []string, path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("error while opening changelog file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	var unreleasedTitleLine, firstVersionLine = -1, -1

	var index = 0
	for scanner.Scan() {
		line := scanner.Text()

		isVersionTitle := strings.HasPrefix(line, "## ")
		if isVersionTitle && strings.Contains(strings.ToLower(line), "unreleased") {
			unreleasedTitleLine = index
		} else if isVersionTitle && firstVersionLine == -1 {
			firstVersionLine = index
		}

		lines = append(lines, line)
		index++
	}

	err = scanner.Err()
	if err != nil {
		return "", err
	}

	// Compute the lines that will go before: from the line 0 up to the unreleased one
	// (this will make sure we don't exclude top file comments)
	var linesBefore []string
	if unreleasedTitleLine > 0 {
		linesBefore = append([]string{}, lines[0:unreleasedTitleLine]...)
	} else if firstVersionLine > 0 {
		linesBefore = append([]string{}, lines[0:firstVersionLine]...)
	}

	// Compute the lines that will go after the changelog: all previous versions (if any)
	var linesAfter []string
	if firstVersionLine > 0 {
		linesAfter = append([]string{}, lines[firstVersionLine:]...)
	}

	updatedChangelog := append(append(linesBefore, data...), linesAfter...)
	return strings.Join(updatedChangelog, "\n"), nil
}

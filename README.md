# Changeset
> This tool is inspired from [Atlassian's changeset CLI](https://github.com/atlassian/changesets).
 
Changeset allows to easily handle change sets and create changelog entries that are compatible with both [semantic versioning](https://semver.org/) and [Keep A Changelog](https://keepachangelog.com/en/1.0.0/).

## Installation
To install `changeset`, run the following command:

```
go get github.com/desmos-labs/changeset/cmd/changeset
```

## Usage
### Initializing the repo
In order to use `changeset`, the first thing you have to do is initialize the folder that contains your code: 

```
changeset init [github-repo] [version]
```

You need to specify two arguments:
- `github-repo` must represent the GitHub repository URL.  
   This will later be used to create pull request links
- `version` must represent the current version of the software.  
   This will later be used to determine the next version if needed.
  
This will generate a file named `config.yaml` inside a new `.changeset` folder in your current directory. The next thing you have to do is edit such file to make sure the `types` and `modules` arrays contain the type of changes and the supported modules you want to have.  
  
### Adding a changeset entry
Once that's done, you can simply add a new changeset entry using 

```
changeset add
```

This will create an interactive prompt that you can use to fill all the data.

### Collecting the changelog
Finally, when you are ready to collect all the changeset entries into a changelog entry you can simply run 

```
changeset collect
```

This command will get all the current changeset entries and create a CHANGELOG entry. Then, it will replace the `Unreleased` section of your file with the newly generated entry. The new version number will be computed based on the type and backward compatibility of the changes that have been made following the [semantic versioning](https://semver.org/) specifications.

This command supports the following flags

| Flag | Default value | Description |
| :--- | :----------- | :---------- |
| `--path` | `$(curdir)/CHANGELOG.md` | Path to the CHANGELOG file to update |
| `--version` | ` ` | Version to be used instead of the computed one. | 
| `--dry-run` | `false` | If `true`, preview the result without applying it to the CHANGELOG file |
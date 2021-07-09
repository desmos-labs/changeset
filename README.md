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
changeset collect [[version]]
```

The `version` parameter is optional. If specified, `changeset` will use its value as the name of the next version. Otherwise, it will compute the next version based on the type and backward compatibility of the changes that have been made following the [semantic versioning](https://semver.org/) specifications.   
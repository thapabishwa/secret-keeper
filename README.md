# vault-differ

## Pre-requisite
- git
- desired vault tools like ansible-vault, sops etc

## Installation
```
go get github.com/everesthack-incubator/vault-differ
```
## Usage 
```
Usage:
  vault-differ [command]

Available Commands:
  clean       Removes unchanged secrets from git repositories
  decrypt     Runs the decrypt command provided in the config file
  encrypt     Runs the encrypt command provided in the config file
  help        Help about any command
```
## Configuration
- sample `.gitattribute` file
  ```
  *.tfvars diff=ansiblevaultdiffer
  ```
- Run the following to add a git config **(Fix the script a/c to your need)**
  ```bash
  git config diff.ansiblevaultdiffer.textconv "ansible-vault view --vault-password-file ~/.vault_password_file "
  ```
- Contents of config.yaml inside the repository you want to work with. 
  ```yaml
  secrets: ## List of strings 
    - "*.tf"
    - "*.password"
  debug: true ## Shows debug log
  vault_tool: "ansible-vault" ## The tool used to encrypt and decrypt secrets
  encrypt_args: ["encrypt", "--vault-password-file", "~/.vault-password-file"] ## arguments used to encrypt (and clean) the secrets
  decrypt_args: ["decrypt", "--vault-password-file", "~/.vault-password-file"] ## arguments used to decrypt the secrets 
  ```
## Contributors ✨

Thanks goes to these wonderful people:

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://www.thapabishwa.de/"><img src="https://avatars1.githubusercontent.com/u/15176360?v=4?s=50" width="50px;" alt=""/><br /><sub><b>Bishwa Thapa</b></sub></a><br /><a href="https://github.com/everesthack-incubator/vault-differ/commits?author=thapabishwa" title="Code">💻</a> <a href="https://github.com/everesthack-incubator/vault-differ/commits?author=thapabishwa" title="Documentation">📖</a> <a href="#example-thapabishwa" title="Examples">💡</a> <a href="#ideas-thapabishwa" title="Ideas, Planning, & Feedback">🤔</a> <a href="#maintenance-thapabishwa" title="Maintenance">🚧</a> <a href="#platform-thapabishwa" title="Packaging/porting to new platform">📦</a> <a href="#research-thapabishwa" title="Research">🔬</a></td>
    <td align="center"><a href="https://github.com/Kripesh4569"><img src="https://avatars2.githubusercontent.com/u/11332619?v=4?s=50" width="50px;" alt=""/><br /><sub><b>Kripesh Dhakal</b></sub></a><br /><a href="https://github.com/everesthack-incubator/vault-differ/issues?q=author%3AKripesh4569" title="Bug reports">🐛</a> <a href="https://github.com/everesthack-incubator/vault-differ/commits?author=Kripesh4569" title="Code">💻</a> <a href="https://github.com/everesthack-incubator/vault-differ/commits?author=Kripesh4569" title="Documentation">📖</a> <a href="#example-Kripesh4569" title="Examples">💡</a> <a href="#ideas-Kripesh4569" title="Ideas, Planning, & Feedback">🤔</a> <a href="#platform-Kripesh4569" title="Packaging/porting to new platform">📦</a> <a href="https://github.com/everesthack-incubator/vault-differ/pulls?q=is%3Apr+reviewed-by%3AKripesh4569" title="Reviewed Pull Requests">👀</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
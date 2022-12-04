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
  *.tfvars diff=vaultdiffer
  ```
- Run the following to add a git config **(Fix the script a/c to your need)**
  ```bash
  git config diff.vaultdiffer.textconv "ansible-vault view --vault-password-file ~/.vault_password_file "
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
## Improvements
- [x] Improved the code to increase the performance by ~3x while decrypting, cleaning, and encrypting secrets
- [x] Git lock causes the restore process to fail. Added a better mechanism to handle this 
### Before
```
bash-3.2$ time go run ./.. encrypt

real    0m3.420s
user    0m2.554s
sys     0m0.736s

bash-3.2$ time go run ./.. decrypt

real    0m3.018s
user    0m2.465s
sys     0m0.568s

bash-3.2$ time go run ./.. clean

real    0m1.399s
user    0m0.392s
sys     0m0.386s
```

### After
```
bash-3.2$ time go run ./.. encrypt

real    0m1.506s
user    0m3.429s
sys     0m0.878s

bash-3.2$ time go run ./.. decrypt

real    0m0.905s
user    0m3.458s
sys     0m0.790s


bash-3.2$ time go run ./.. clean

real    0m0.837s
user    0m0.385s
sys     0m0.362s
```

## Future Improvements 
- [ ] Add support for multiple vault tools
- [ ] Add support for multiple git attributes
- [ ] Add support for multiple git config
- [ ] Add support for multiple config files
- [ ] Add support for multiple repositories
- [ ] Add support for multiple branches
- [ ] Add support for multiple commits
- [ ] Add support for multiple files
- [ ] Add support for multiple secrets
- [ ] Add support for multiple secret types

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
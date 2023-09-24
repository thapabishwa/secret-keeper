# secret-keeper

Secret Keeper is a tool that helps users to manage and review changes to secrets in encrypted repositories. It does this by filtering files with no secret changes from the git worktree. This makes it easier to review changes to secrets and to commit only updated secrets.

## Features
- Filters files with no secret changes from the git worktree
- Supports Ansible Vault and Sops encrypted repositories
- Easy to use and configure
- Generates a report of the filtered changes

## Benefits

Secret Keeper offers a number of benefits, including:

- Makes it easier to review changes to secrets
- Improves the security of secrets by reducing the amount of time that they are exposed in cleartext in the git worktree
- Reduces the number of unnecessary diffs created in the repo

## Installation
  ```bash
  go install github.com/everesthack-incubator/secret-keeper@latest
  ```

## Pre-requisite
- git
- desired vault tools like ansible-vault, sops etc

## Configuration

- Inside your repository, create a new `config.secret-keeper.yaml` file and modify it as needed. The following is an example configuration file for Ansible Vault and Sops.

  <details>
  <summary>Ansible Vault</summary>

  ```yaml
  secret_files_patterns:
    # The list of file patterns to treat as secrets in the repository across all folders
    - "*.tf"
    - "*.password"
  vault_tool: "ansible-vault"
  # The args to encrypt a file in-place using the vault tool
  encrypt_args:
    - "encrypt"
    - "--vault-password-file"
    - "~/.vault-password-file"
  # The args to decrypt a file in-place using the vault tool
  decrypt_args:
    - "decrypt"
    - "--vault-password-file"
    - "~/.vault-password-file"
  # The args to view secret in the file using the vault tool
  view_args:
    - "view"
    - "--vault-password-file"
    - "~/.vault-password-file"
  ```
  </details>


  <details>
  <summary>Mozilla Sops</summary>

  ```yaml
  secret_files_patterns:
    # The list of file patterns to treat as secrets in the repository across all folders
    - "*.tf"
    - "*.password"
  vault_tool: "sops"
  # The args to encrypt a file in-place using the vault tool
  encrypt_args:
    - "--encrypt"
    - "--in-place"
    - "--pgp"
  # The args to decrypt a file in-place using the vault tool
  decrypt_args:
    - "--decrypt"
    - "--in-place"
    - "--pgp"
  # The args to view secret in the file using the vault tool
  view_args:
    - "--decrypt"
    - "--pgp"
  ```
  </details>


  This configuration file controls the behavior of the tool, allowing you to specify which files should be treated as secrets, enable debug mode, and set the encryption and decryption parameters.

- After creating the configuration file, initialize the repository with the tool
  ```
  secret-keeper init
  ```

## Usage

- Start using the tool
  ```
  secret-keeper encrypt # encrypts all the secrets, if not already encrypted. also cleans the secrets from the git worktree

  secret-keeper clean # cleans the secrets from the git worktree
  
  secret-keeper decrypt # decrypts all the secrets, if not already decrypted.
  ```

## Improvements
- [x] Enhance the performance by ~3x while decrypting, cleaning, and encrypting secrets
- [x] Git lock causes the restore process to fail. Added a better mechanism to handle this
- [x] Ensure that the new/untracked files are not discarded on the clean command
- [x] Ensure that adding a new file to the repo does not cause the clean command to fail 
- [x] Improve the onboarding process

## Future Improvements 
- [ ] Add Support for more secret management tools in the same repo 
- [ ] Add Support for different types of repositories.
- [ ] Add the ability to ignore certain files or directories.
- [ ] Add the ability to generate a report of the filtered changes.
- [ ] Add support for continuous integration (CI) and continuous delivery (CD) pipelines

## Contributors ✨

Thanks goes to these wonderful people:

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://www.thapabishwa.de/"><img src="https://avatars1.githubusercontent.com/u/15176360?v=4?s=50" width="50px;" alt=""/><br /><sub><b>Bishwa Thapa</b></sub></a><br /><a href="https://github.com/everesthack-incubator/secret-keeper/commits?author=thapabishwa" title="Code">💻</a> <a href="https://github.com/everesthack-incubator/secret-keeper/commits?author=thapabishwa" title="Documentation">📖</a> <a href="#example-thapabishwa" title="Examples">💡</a> <a href="#ideas-thapabishwa" title="Ideas, Planning, & Feedback">🤔</a> <a href="#maintenance-thapabishwa" title="Maintenance">🚧</a> <a href="#platform-thapabishwa" title="Packaging/porting to new platform">📦</a> <a href="#research-thapabishwa" title="Research">🔬</a></td>
    <td align="center"><a href="https://github.com/Kripesh4569"><img src="https://avatars2.githubusercontent.com/u/11332619?v=4?s=50" width="50px;" alt=""/><br /><sub><b>Kripesh Dhakal</b></sub></a><br /><a href="https://github.com/everesthack-incubator/secret-keeper/issues?q=author%3AKripesh4569" title="Bug reports">🐛</a> <a href="https://github.com/everesthack-incubator/secret-keeper/commits?author=Kripesh4569" title="Code">💻</a> <a href="https://github.com/everesthack-incubator/secret-keeper/commits?author=Kripesh4569" title="Documentation">📖</a> <a href="#example-Kripesh4569" title="Examples">💡</a> <a href="#ideas-Kripesh4569" title="Ideas, Planning, & Feedback">🤔</a> <a href="#platform-Kripesh4569" title="Packaging/porting to new platform">📦</a> <a href="https://github.com/everesthack-incubator/secret-keeper/pulls?q=is%3Apr+reviewed-by%3AKripesh4569" title="Reviewed Pull Requests">👀</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

## License
Secret Keeper is licensed under the MIT License.
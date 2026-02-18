# Package Distribution Guide

This guide details how to set up and maintain package distribution for Clause CLI across various package managers (Homebrew, Scoop, AUR, Winget, etc.) using GoReleaser.

## Prerequisites

Before releasing, ensure you have the following secrets configured in your GitHub repository:

- `HOMEBREW_TAP_TOKEN`: GitHub Personal Access Token (PAT) with `repo` scope for the Homebrew tap repository.
- `SCOOP_BUCKET_TOKEN`: GitHub PAT with `repo` scope for the Scoop bucket repository.
- `AUR_KEY`: SSH private key for publishing to the Arch User Repository (AUR).
- `GPG_FINGERPRINT`: GPG key fingerprint for signing artifacts (optional but recommended).

---

## 1. Homebrew (macOS/Linux)

Homebrew uses "Taps" (repositories) to host formulas.

### Setup

1.  **Create a Repository**: Create a new public GitHub repository named `homebrew-tap` in your organization (e.g., `clause-cli/homebrew-tap`).
2.  **Initialize**: The repository can be empty; GoReleaser will create the formula file.
3.  **Token**:
    -   Generate a [GitHub PAT](https://github.com/settings/tokens) with `repo` (Full control of private repositories) scope.
    -   Add this token as a repository secret named `HOMEBREW_TAP_TOKEN`.

### Check Configuration

In `.goreleaser.yaml`:

```yaml
brews:
  - name: clause
    tap:
      owner: clause-cli
      name: homebrew-tap
```

### Usage

Users install via:

```bash
brew install clause-cli/tap/clause
```

---

## 2. Scoop (Windows)

Scoop uses "Buckets" (repositories) to host app manifests.

### Setup

1.  **Create a Repository**: Create a new public GitHub repository named `scoop-bucket` (e.g., `clause-cli/scoop-bucket`).
2.  **Initialize**: The repository can be empty.
3.  **Token**:
    -   Use the same GitHub PAT as above or create a new one with `repo` scope.
    -   Add this token as a repository secret named `SCOOP_BUCKET_TOKEN`.

### Check Configuration

In `.goreleaser.yaml`:

```yaml
scoops:
  - name: clause
    bucket:
      owner: clause-cli
      name: scoop-bucket
```

### Usage

Users install via:

```powershell
scoop bucket add clause-cli https://github.com/clause-cli/scoop-bucket
scoop install clause
```

---

## 3. AUR (Arch Linux)

The Arch User Repository (AUR) hosts community-maintained build scripts.

### Setup

1.  **Create an Account**: Register an account on [aur.archlinux.org](https://aur.archlinux.org/).
2.  **Generate SSH Keys**:
    ```bash
    ssh-keygen -t ed25519 -C "aur@clause.dev" -f ~/.ssh/aur
    ```
3.  **Add Public Key**: Add the content of `~/.ssh/aur.pub` to your [AUR account settings](https://aur.archlinux.org/account).
4.  **Create Package**:
    -   You typically need to push an initial empty repository to AUR first to reserve the name.
    -   Run: `git clone ssh://aur@aur.archlinux.org/clause-bin.git` (this will fail if it doesn't exist, you might need to create it via the web interface or push to it).
5.  **Add Private Key Secret**:
    -   Copy the content of the private key (`~/.ssh/aur`).
    -   Add it as a repository secret named `AUR_KEY`.

### Check Configuration

In `.goreleaser.yaml`:

```yaml
aurs:
  - name: clause-bin
    private_key: '{{ .Env.AUR_KEY }}'
    git_url: 'ssh://aur@aur.archlinux.org/clause-bin.git'
```

---

## 4. Winget (Windows)

Winget is the official Windows Package Manager. Submissions are made via Pull Requests to the [microsoft/winget-pkgs](https://github.com/microsoft/winget-pkgs) repository.

### Manual Submission (Recommended for Verification)

GoReleaser generates the Winget manifest but does not automatically submit the PR to Microsoft's repository in the basic configuration.

1.  **Wait for Release**: Let the GitHub Action complete and publish the release.
2.  **Locate Manifest**: GoReleaser will include the generated YAML manifest in the release artifacts (e.g., `dist/clause_manifest.yaml` or inside the `dist` folder in artifacts).
3.  **Submit PR**:
    -   Fork [microsoft/winget-pkgs](https://github.com/microsoft/winget-pkgs).
    -   Create a directory structure: `manifests/c/Clause/ClauseCLI/<version>/`.
    -   Place the generated YAML manifest(s) there.
    -   Submit a Pull Request.

### Automating Winget (Advanced)
To automate this, you can configure GoReleaser to push the manifest to a separate repository (like a staging repo), and then use a separate automation tool (like [winget-create](https://github.com/microsoft/winget-create)) or a GitHub Action to validate and submit the PR to Microsoft.

Current `.goreleaser.yaml` configuration creates the artifact but does not push it:

```yaml
winget:
  - name: clause
    publisher: Clause
    package_identifier: Clause.ClauseCLI
    skip_upload: false # Generates the manifest in dist/
```

---

## 5. NPM (Wrapper)

You can distribute the binary via `npm` by using a wrapper package.

### Setup

1.  **Create NPM Package**: A standard `package.json` with a post-install script that downloads the binary.
2.  **GoReleaser Integration**: GoReleaser can build `nfpm` packages (deb, rpm, apk) but for `npm`, you typically maintain a separate repository (e.g., `clause-cli/npm-wrapper`) that fetches the latest release.

The current configuration generates system packages (`deb`, `rpm`, `apk`) via `nfpms`:

```yaml
nfpms:
  - file_name_template: "{{ .ConventionalFileName }}"
    formats:
      - apk
      - deb
      - rpm
    # ...
```

These are uploaded to the GitHub Release. Users can download them directly or you can host your own `apt`/`yum` repository.

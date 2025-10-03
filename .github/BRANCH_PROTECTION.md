# Branch Protection Configuration

To ensure code quality and prevent breaking changes, the following branch protection rules should be enabled for the `main` branch:

## Recommended Settings

### Require Pull Request Reviews
- ✅ Require approvals: **1**
- ✅ Dismiss stale pull request approvals when new commits are pushed
- ✅ Require review from Code Owners (if CODEOWNERS file exists)

### Require Status Checks
All of the following checks must pass before merging:

**Required CI Checks:**
- ✅ `test` - All tests must pass across Go 1.21, 1.22, 1.23
- ✅ `lint` - golangci-lint must pass
- ✅ `security` - gosec security scan must pass with zero vulnerabilities
- ✅ `benchmark` - Performance benchmarks must complete

**Coverage Requirements:**
- ✅ Overall coverage must be ≥70% (enforced in CI)
- 🎯 Target: 90%+ for v1.0 release

### Additional Protections
- ✅ Require branches to be up to date before merging
- ✅ Require conversation resolution before merging
- ✅ Do not allow bypassing the above settings
- ✅ Restrict who can push to matching branches (maintainers only)

### Linear History (Optional)
- ⚠️ Require linear history - Consider enabling to keep clean git history
- Alternative: Use "Squash and merge" as the default merge method

## How to Enable

### Via GitHub Web UI
1. Go to repository **Settings** → **Branches**
2. Click **Add rule** or edit existing rule for `main`
3. Configure settings as listed above
4. Click **Create** or **Save changes**

### Via GitHub CLI
```bash
# Install GitHub CLI if not already installed
# brew install gh

# Enable branch protection for main
gh api repos/:owner/:repo/branches/main/protection \
  --method PUT \
  --field required_status_checks='{"strict":true,"contexts":["test","lint","security","benchmark"]}' \
  --field enforce_admins=true \
  --field required_pull_request_reviews='{"dismiss_stale_reviews":true,"require_code_owner_reviews":false,"required_approving_review_count":1}' \
  --field restrictions=null
```

## For Contributors

When branch protection is enabled:
- All PRs must pass CI checks before merging
- At least one approving review is required
- Branch must be up-to-date with main before merging
- Force pushes to main are blocked

This ensures code quality and prevents accidental breaking changes.

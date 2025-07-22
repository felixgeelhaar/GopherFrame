#!/bin/bash

# GopherFrame Development Setup
# Installs pre-commit hooks and required tools

set -e

echo "🚀 Setting up GopherFrame development environment..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Check if we're in the right directory
if [ ! -f "go.mod" ] || ! grep -q "gopherFrame" go.mod; then
    echo -e "${RED}❌ Please run this script from the GopherFrame project root${NC}"
    exit 1
fi

# Check if Git is available
if ! command -v git &> /dev/null; then
    echo -e "${RED}❌ Git is required but not installed${NC}"
    exit 1
fi

# Check if Go is available
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is required but not installed${NC}"
    echo -e "${BLUE}💡 Install Go from: https://golang.org/dl/${NC}"
    exit 1
fi

echo -e "${BLUE}🔍 Go version: $(go version)${NC}"

# Install required tools
install_tool() {
    local tool=$1
    local package=$2
    
    if command -v "$tool" &> /dev/null; then
        echo -e "${GREEN}✅ $tool already installed${NC}"
    else
        echo -e "${YELLOW}📦 Installing $tool...${NC}"
        if go install "$package"; then
            echo -e "${GREEN}✅ $tool installed successfully${NC}"
        else
            echo -e "${RED}❌ Failed to install $tool${NC}"
            exit 1
        fi
    fi
}

# Install development tools
echo -e "${BLUE}📦 Installing required tools...${NC}"
install_tool "golangci-lint" "github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
install_tool "goimports" "golang.org/x/tools/cmd/goimports@latest"

# Install pre-commit hook
echo -e "${BLUE}🪝 Installing pre-commit hook...${NC}"

# Create hooks directory if it doesn't exist
mkdir -p .git/hooks

# Copy pre-commit hook from template
if [ -f ".git/hooks/pre-commit" ]; then
    echo -e "${YELLOW}⚠️  Pre-commit hook already exists. Creating backup...${NC}"
    mv .git/hooks/pre-commit .git/hooks/pre-commit.backup
fi

# Create the pre-commit hook
cat > .git/hooks/pre-commit << 'EOF'
#!/bin/bash

# GopherFrame Pre-commit Hook
# Automatically formats and lints code before committing

set -e

echo "🔍 Running pre-commit checks..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if we're in a Git repository
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    echo -e "${RED}❌ Not in a Git repository${NC}"
    exit 1
fi

# Get list of staged Go files
STAGED_GO_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$' || true)

if [ -z "$STAGED_GO_FILES" ]; then
    echo -e "${GREEN}✅ No Go files to check${NC}"
    exit 0
fi

echo -e "${YELLOW}📝 Checking Go files: ${STAGED_GO_FILES}${NC}"

# Check if required tools are installed
check_tool() {
    if ! command -v "$1" &> /dev/null; then
        echo -e "${RED}❌ $1 is not installed. Run scripts/setup-hooks.sh to install required tools.${NC}"
        exit 1
    fi
}

# Check required tools
check_tool "gofmt"
check_tool "goimports" 
check_tool "golangci-lint"

# Format Go files
echo -e "${YELLOW}🔧 Formatting Go code...${NC}"
NEEDS_FORMATTING=0

for file in $STAGED_GO_FILES; do
    # Check if file needs formatting
    if [ "$(gofmt -l "$file")" ]; then
        echo -e "  📝 Formatting $file"
        gofmt -w "$file"
        NEEDS_FORMATTING=1
    fi
    
    # Fix imports
    if ! goimports -l "$file" | grep -q "^$"; then
        echo -e "  🔗 Fixing imports in $file"
        goimports -w "$file"
        NEEDS_FORMATTING=1
    fi
done

# If files were formatted, re-stage them
if [ $NEEDS_FORMATTING -eq 1 ]; then
    echo -e "${GREEN}✅ Code formatted. Re-staging files...${NC}"
    git add $STAGED_GO_FILES
fi

# Run go mod tidy to ensure dependencies are clean
echo -e "${YELLOW}📦 Tidying Go modules...${NC}"
if go mod tidy; then
    # Check if go.mod or go.sum changed
    if git diff --name-only | grep -q "go.mod\|go.sum"; then
        echo -e "${GREEN}✅ Go modules updated. Staging changes...${NC}"
        git add go.mod go.sum
    fi
else
    echo -e "${RED}❌ go mod tidy failed${NC}"
    exit 1
fi

# Run linter
echo -e "${YELLOW}🔍 Running golangci-lint...${NC}"
if ! golangci-lint run --config .golangci.yml $STAGED_GO_FILES; then
    echo -e "${RED}❌ Linting failed. Please fix the issues above.${NC}"
    exit 1
fi

# Run tests for modified packages
echo -e "${YELLOW}🧪 Running tests for affected packages...${NC}"
AFFECTED_PACKAGES=$(echo "$STAGED_GO_FILES" | xargs -I {} dirname {} | sort -u | xargs -I {} go list ./{}... 2>/dev/null || echo ".")

if [ -n "$AFFECTED_PACKAGES" ]; then
    if ! go test -short $AFFECTED_PACKAGES; then
        echo -e "${RED}❌ Tests failed. Please fix failing tests.${NC}"
        exit 1
    fi
fi

# Run quick build check
echo -e "${YELLOW}🔨 Verifying build...${NC}"
if ! go build ./...; then
    echo -e "${RED}❌ Build failed. Please fix build errors.${NC}"
    exit 1
fi

echo -e "${GREEN}✅ All pre-commit checks passed!${NC}"
echo -e "${GREEN}🚀 Ready to commit${NC}"

exit 0
EOF

chmod +x .git/hooks/pre-commit

echo -e "${GREEN}✅ Pre-commit hook installed successfully${NC}"

# Verify installation
echo -e "${BLUE}🔍 Verifying installation...${NC}"

# Test the tools
if golangci-lint version > /dev/null 2>&1; then
    echo -e "${GREEN}✅ golangci-lint: $(golangci-lint version --format short)${NC}"
else
    echo -e "${RED}❌ golangci-lint not working properly${NC}"
fi

if goimports -h > /dev/null 2>&1; then
    echo -e "${GREEN}✅ goimports: installed and working${NC}"
else
    echo -e "${RED}❌ goimports not working properly${NC}"
fi

# Run go mod download to ensure dependencies are available
echo -e "${BLUE}📦 Downloading Go dependencies...${NC}"
if go mod download; then
    echo -e "${GREEN}✅ Dependencies downloaded${NC}"
else
    echo -e "${RED}❌ Failed to download dependencies${NC}"
    exit 1
fi

# Final verification - run a quick format check
echo -e "${BLUE}🧪 Testing pre-commit hook...${NC}"
if git status --porcelain | grep -q '^M'; then
    echo -e "${YELLOW}⚠️  You have uncommitted changes. The pre-commit hook will run on your next commit.${NC}"
else
    echo -e "${GREEN}✅ Pre-commit hook ready. It will run automatically on git commit.${NC}"
fi

echo ""
echo -e "${GREEN}🎉 GopherFrame development environment setup complete!${NC}"
echo ""
echo -e "${BLUE}Next steps:${NC}"
echo -e "  • Make code changes"
echo -e "  • Run: git add ."
echo -e "  • Run: git commit -m 'Your message'"
echo -e "  • The pre-commit hook will automatically format, lint, and test your changes"
echo ""
echo -e "${BLUE}Manual commands:${NC}"
echo -e "  • Format code: gofmt -w ."
echo -e "  • Fix imports: goimports -w ."
echo -e "  • Lint code: golangci-lint run"
echo -e "  • Run tests: go test ./..."
echo ""
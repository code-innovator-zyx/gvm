#!/usr/bin/env bash
set -euo pipefail

GVM_RELEASE="1.0.0"
GVM_HOME="${HOME}/.gvm"

# 获取系统架构
get_arch() {
    case "$(uname -m)" in
        x86_64 | amd64) echo "amd64" ;;
        i386 | i486 | i586) echo "386" ;;
        aarch64 | arm64) echo "arm64" ;;
        armv6l | armv7l) echo "arm" ;;
        s390x) echo "s390x" ;;
        riscv64) echo "riscv64" ;;
        *)
            echo "Unsupported architecture: $(uname -m)" >&2
            exit 1
            ;;
    esac
}

function get_os() {
    echo $(uname -s | awk '{print tolower($0)}')
}
# 安装 gvm
install_gvm() {
    local os arch dest_file url
    os=$(get_os)
    arch=$(get_arch)
    dest_file="${GVM_HOME}/gvm${GVM_RELEASE}.${os}-${arch}.tar.gz"
    url="https://github.com/code-innovator-zyx/gvm/releases/download/v${GVM_RELEASE}/gvm${GVM_RELEASE}.${os}-${arch}.tar.gz"

    echo "[1/3] Downloading ${url}"
    mkdir -p "${GVM_HOME}"
    rm -f "${dest_file}"

    if command -v wget >/dev/null 2>&1; then
        wget -q -O "${dest_file}" "${url}"
    else
        curl -sSL -o "${dest_file}" "${url}"
    fi

    echo "[2/3] Installing gvm to ${GVM_HOME}"
    tar -xzf "${dest_file}" -C "${GVM_HOME}"
    chmod +x "${GVM_HOME}/gvm"
    rm -f "${dest_file}"

    echo "[3/3] Configuring shell environment"
    local shell_config
    for shell_config in "${HOME}/.bashrc" "${HOME}/.zshrc"; do
        if [ -f "$shell_config" ] || [ -w "$HOME" ]; then
            cat >>"$shell_config" <<-'EOF'

# gvm shell setup
export GVM_HOME="${HOME}/.gvm"
export GO_ROOT="${GVM_HOME}/go"
[ -z "$GOPATH" ] && export GOPATH="${HOME}/go"
export PATH="${GVM_HOME}:${GO_ROOT}/bin:${GOPATH}/bin:$PATH"

EOF
        fi
    done

    echo -e "\nInstallation completed. Please restart your terminal or source your shell configuration file."
}

main() {
    install_gvm
}

main
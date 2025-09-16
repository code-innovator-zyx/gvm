#!/usr/bin/env bash
set -e

function get_arch() {
    a=$(uname -m)
    case ${a} in
    "x86_64" | "amd64")
        echo "amd64"
        ;;
    "i386" | "i486" | "i586")
        echo "386"
        ;;
    "aarch64" | "arm64")
        echo "arm64"
        ;;
    "armv6l" | "armv7l")
        echo "arm"
        ;;
    "s390x")
        echo "s390x"
        ;;
    "riscv64")
        echo "riscv64"
        ;;
    *)
        echo ${NIL}
        ;;
    esac
}

function get_os() {
    echo $(uname -s | awk '{print tolower($0)}')
}

function main() {
    local release="1.0.0"
    local os=$(get_os)
    local arch=$(get_arch)
    local dest_file="${HOME}/.gvm/gvm${release}.${os}-${arch}.tar.gz"
    local url="https://github.com/code-innovator-zyx/gvm/releases/download/v${release}/gvm${release}.${os}-${arch}.tar.gz"

    echo "[1/3] Downloading ${url}"
    rm -f "${dest_file}"
    if [ -x "$(command -v wget)" ]; then
        mkdir -p "${HOME}/.gvm"
        wget -q -P "${HOME}/.gvm" "${url}"
    else
        curl -s -S -L --create-dirs -o "${dest_file}" "${url}"
    fi

    echo "[2/3] Install gvm to the ${HOME}/.gvm"
    tar -xz -f "${dest_file}" -C "${HOME}/.gvm"
    chmod +x "${HOME}/.gvm/gvm"
    
    # 删除下载的压缩包
    rm -f "${dest_file}"

    echo "[3/3] Set environment variables"
    # Environment variables setup removed as per user request
    
    if [ -x "$(command -v bash)" ]; then
        cat >>${HOME}/.bashrc <<-'EOF'

# gvm shell setup
export GVM_HOME="${HOME}/.gvm"
export GO_ROOT="${GVM_HOME}/go"
[ -z "$GOPATH" ] && export GOPATH="${HOME}/go"
export PATH="${GVM_HOME}:${GO_ROOT}/bin:${GOPATH}/bin:$PATH"

EOF
    fi

    if [ -x "$(command -v zsh)" ]; then
        cat >>${HOME}/.zshrc <<-'EOF'

# gvm shell setup
export GVM_HOME="${HOME}/.gvm"
export GO_ROOT="${GVM_HOME}/go"
[ -z "$GOPATH" ] && export GOPATH="${HOME}/go"
export PATH="${GVM_HOME}:${GO_ROOT}/bin:${GOPATH}/bin:$PATH"

EOF
    fi

    # Fish shell configuration removed as per user request

    echo -e "\nInstallation completed. Please restart your terminal or source your shell configuration file."

    exit 0
}

main
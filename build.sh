#!/usr/bin/env bash
set -euo pipefail

RELEASE="${1:-1.0.0}"

TARGETS=(
    "darwin_amd64" "darwin_arm64"
    "linux_386" "linux_amd64" "linux_arm" "linux_arm64" "linux_s390x" "linux_riscv64"
#    "windows_386" "windows_amd64" "windows_arm" "windows_arm64"
)

OUTPUT_DIR="./dist"
SHA_FILE="${OUTPUT_DIR}/sha256sum.txt"

rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

function package() {
    local release="$1"
    local osarch="$2"
    local os="${osarch%%_*}"
    local arch="${osarch##*_}"

    local tmp_bin="./gvm"
    if [[ "$os" == "windows" ]]; then
        tmp_bin="${tmp_bin}.exe"
    fi

    echo "[→] Building ${os}-${arch}..."

    GOOS="$os" GOARCH="$arch" CGO_ENABLED=0 GO111MODULE=on GOPROXY="https://goproxy.cn,direct" \
        go build -o "$tmp_bin" .

    if [[ "$os" == "windows" ]]; then
        local pkg_name="gvm${release}.${os}-${arch}.zip"
        zip -j "${OUTPUT_DIR}/${pkg_name}" "$tmp_bin" > /dev/null
    else
        local pkg_name="gvm${release}.${os}-${arch}.tar.gz"
        tar -czf "${OUTPUT_DIR}/${pkg_name}" -C "$(dirname "$tmp_bin")" "$(basename "$tmp_bin")"
    fi

    # 每个平台写独立的 sha256 文件
    shasum -a 256 "${OUTPUT_DIR}/${pkg_name}" > "${OUTPUT_DIR}/${pkg_name}.sha256"

    rm -f "$tmp_bin"
    echo "[✓] $pkg_name built"
}

# ============================
# Main
# ============================

echo "Building gvm version ${RELEASE} for multiple platforms..."
echo "Output dir: $OUTPUT_DIR"

# 并行构建
for target in "${TARGETS[@]}"; do
    package "$RELEASE" "$target" &
done

# 等待所有后台任务完成
wait

# 合并所有 sha256 文件
cat ${OUTPUT_DIR}/*.sha256 > "$SHA_FILE"
rm -f ${OUTPUT_DIR}/*.sha256

# 清理临时文件
go clean
echo "All builds done. SHA256 sums written to $SHA_FILE"
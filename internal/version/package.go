package version

/*
* @Author: zouyx
* @Email: zouyx@knowsec.com
* @Date:   2025/9/11 下午5:07
* @Package:
 */

type Kind string

const (
	// SourceKind 表示源码包（如 .tar.gz, .zip, .tgz）
	SourceKind Kind = "Source"

	// ArchiveKind 表示二进制归档包（解压即可用，如 .zip, .tar.gz）
	ArchiveKind Kind = "Archive"

	// InstallerKind 表示安装程序（需要执行安装过程，如 .exe, .msi, .pkg）
	InstallerKind Kind = "Installer"
)

type OS string

const (
	// 桌面/主流 OS
	Windows OS = "windows"
	MacOS   OS = "darwin" // Go 里对应 darwin，而不是 macos
	Linux   OS = "linux"

	// BSD 家族
	FreeBSD   OS = "freebsd"
	OpenBSD   OS = "openbsd"
	NetBSD    OS = "netbsd"
	Dragonfly OS = "dragonfly"

	// 传统/小众 Unix
	AIX     OS = "aix"
	Illumos OS = "illumos"
	Solaris OS = "solaris"
	Plan9   OS = "plan9"

	// 移动/嵌入式
	Android OS = "android"
	IOS     OS = "ios"

	// 其它
	JS   OS = "js"   // GopherJS / WebAssembly
	WASM OS = "wasm" // WebAssembly runtime
)

type ARCH string

const (
	UnknownARCH ARCH = ""

	// x86 家族
	X86   ARCH = "x86"    // 32-bit
	X8664 ARCH = "x86_64" // 64-bit
	AMD64 ARCH = "amd64"  // 同 x86_64，常用别名
	I386  ARCH = "386"    // 32-bit Intel

	// ARM 家族
	ARM   ARCH = "arm"   // 32-bit ARM
	ARM64 ARCH = "arm64" // 64-bit ARM (AArch64)
	ARMV6 ARCH = "armv6"
	ARMV7 ARCH = "armv7"
	ARMV8 ARCH = "armv8"

	// PowerPC 家族
	PPC     ARCH = "ppc"
	PPC64   ARCH = "ppc64"
	PPC64LE ARCH = "ppc64le" // 小端模式

	// MIPS 家族
	MIPS     ARCH = "mips"
	MIPSLE   ARCH = "mipsle"
	MIPS64   ARCH = "mips64"
	MIPS64LE ARCH = "mips64le"

	// RISC-V
	RISCV64 ARCH = "riscv64"

	// SPARC 家族
	SPARC   ARCH = "sparc"
	SPARC64 ARCH = "sparc64"

	// IBM System z
	S390  ARCH = "s390"
	S390X ARCH = "s390x"

	// LoongArch
	LOONGARCH32 ARCH = "loongarch32"
	LOONGARCH64 ARCH = "loongarch64"
)

type ArtifactInfo struct {
	FileName    string `json:"filename"` // 文件名
	URL         string `json:"url"`      // 下载地址
	Kind        Kind   `json:"kind"`     // 包类型 (Source/Archive/Installer)
	OS          OS     `json:"os"`       // 操作系统
	Arch        ARCH   `json:"arch"`     // 架构
	Size        string `json:"size"`     // 文件大小（字节）
	Checksum    string `json:"checksum"`
	ChecksumURL string `json:"-"`
	Algorithm   string `json:"algorithm"`
}

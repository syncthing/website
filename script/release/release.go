package main

import (
	"cmp"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type githubRelease struct {
	HTMLURL     string    `json:"html_url"`
	Body        string    `json:"body"`
	TagName     string    `json:"tag_name"`
	PublishedAt time.Time `json:"published_at"`
	Assets      []struct {
		Name               string `json:"name"`
		Size               int    `json:"size"`
		BrowserDownloadURL string `json:"browser_download_url"`
	}
}

type downloadPage struct {
	Version string
	OSes    []downloadOS
}

type downloadOS struct {
	OS     string
	Assets []downloadAsset
}

type downloadAsset struct {
	Name        string
	Size        int
	URL         string
	Arch        string
	Recommended bool

	os         string
	osWeight   int
	archWeight int
}

func main() {
	// Load the latest release from GitHub
	res, err := http.Get("https://api.github.com/repos/syncthing/syncthing/releases/latest")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	defer res.Body.Close()
	var rel githubRelease
	err = json.NewDecoder(res.Body).Decode(&rel)
	if err != nil {
		log.Fatal(err)
	}

	// Filter the assets to actual binary Syncthing packages
	var filtered []downloadAsset
	for _, a := range rel.Assets {
		parts := strings.SplitN(a.Name, "-", 4)
		if len(parts) != 4 {
			continue
		}
		if parts[0] != "syncthing" {
			continue
		}
		filtered = append(filtered, downloadAsset{
			Name:        a.Name,
			Size:        a.Size,
			URL:         a.BrowserDownloadURL,
			Arch:        humanReadableArch(parts[2]),
			Recommended: isRecommended(parts[1], parts[2]),

			os:         humanReadableOS(parts[1]),
			osWeight:   osWeight(parts[1]),
			archWeight: archWeight(parts[2]),
		})
	}

	// Sort by operating system and architecture
	slices.SortFunc(filtered, func(a, b downloadAsset) int {
		if a.osWeight != b.osWeight {
			return a.osWeight - b.osWeight
		}
		if a.archWeight != b.archWeight {
			return a.archWeight - b.archWeight
		}
		return cmp.Compare(a.Name, b.Name)
	})

	// Group by OS
	p := downloadPage{
		Version: rel.TagName,
	}
	for _, a := range filtered {
		if len(p.OSes) == 0 || p.OSes[len(p.OSes)-1].OS != a.os {
			log.Println("-", a.os)
			p.OSes = append(p.OSes, downloadOS{
				OS: a.os,
			})
		}
		log.Println("  -", a.Arch)
		p.OSes[len(p.OSes)-1].Assets = append(p.OSes[len(p.OSes)-1].Assets, a)
	}

	// Produce the yaml for the release page
	bs, err := yaml.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(bs)
}

func isRecommended(os, arch string) bool {
	if os == "macos" {
		return arch == "universal"
	}
	return arch == "amd64" || arch == "arm64"
}

func osWeight(os string) int {
	// In order of commonality among our users
	switch os {
	case "linux":
		return 1
	case "windows":
		return 2
	case "macos":
		return 3
	case "freebsd":
		return 4
	case "openbsd":
		return 5
	case "netbsd":
		return 6
	}
	return 9
}

func archWeight(arch string) int {
	// First universal, then 64 bit common archs, then 32 bit common archs,
	// then the rest
	switch arch {
	case "universal":
		return 0
	case "amd64":
		return 1
	case "arm64":
		return 2
	case "386":
		return 3
	case "arm":
		return 4
	case "loong64", "mips64", "mips64le", "ppc64", "ppc64le", "riscv64":
		return 5
	}
	return 9
}

func humanReadableOS(os string) string {
	// Special cases
	switch os {
	case "macos":
		return "macOS"
	}

	// Capitalise the first letter
	os = strings.ToUpper(os[:1]) + os[1:]
	// Known initialisms
	os = strings.Replace(os, "bsd", "BSD", -1)
	return os
}

func humanReadableArch(arch string) string {
	switch arch {
	case "386":
		return "Intel/AMD (32-bit)"
	case "amd64":
		return "Intel/AMD (64-bit)"
	case "arm":
		return "ARM (32-bit)"
	case "arm64":
		return "ARM (64-bit)"
	case "universal":
		return "Universal"
	case "loong64":
		return "Loong64"
	case "riscv":
		return "RISC-V (32 bit)"
	case "riscv64":
		return "RISC-V (64-bit)"
	case "ppc64":
		return "PowerPC (64-bit)"
	case "ppc64le":
		return "PowerPC (64-bit LE)"
	case "mips64":
		return "MIPS (64-bit)"
	case "mips64le":
		return "MIPS (64-bit LE)"
	case "mips":
		return "MIPS (32-bit)"
	case "mipsle":
		return "MIPS (32-bit LE)"
	case "s390x":
		return "IBM zSeries (64-bit)"
	}
	return arch
}

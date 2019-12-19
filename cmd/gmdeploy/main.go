package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: gmdeploy [os] [icon] [path/to/progject]")
		fmt.Println("Flags:")
		flag.PrintDefaults()

		os.Exit(0)
	}

	var targetOS string
	flag.StringVar(&targetOS, "os", "", "target deploy os [windows, darwin, linux]")

	var iconPath string
	flag.StringVar(&iconPath, "icon", "", "applicatio icon file path")

	flag.Parse()

	if len(targetOS) == 0 {
		targetOS = runtime.GOOS
	}

	if len(iconPath) == 0 {
		//@todo: set default icon here
	}

	projectPath, _ := os.Getwd()
	appName := filepath.Base(projectPath)

	// Prepare build dir
	outputDir := filepath.Join(projectPath, "build", targetOS)
	os.RemoveAll(outputDir)
	os.MkdirAll(outputDir, 0755)

	switch targetOS {
	case "darwin":
		// Compile
		cmd := exec.Command("go", "build", ".")
		cmd.Dir = projectPath
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to build: %v", err)
		}

		// Bundle
		macOSPath := filepath.Join(outputDir, fmt.Sprintf("%s.app", appName), "Contents", "MacOS")
		err = os.MkdirAll(macOSPath, 0755)
		if err != nil {
			log.Fatalf("Failed to create %s.app folders:%v", appName, err)
		}
		// Copy compiled executable to build folder
		cmd = exec.Command("mv", appName, macOSPath)
		cmd.Dir = projectPath
		err = cmd.Run()
		if err != nil {
			log.Fatalf("Failed to move executable to destination: %v", err)
		}
		// Prepare Info.plist
		contentsPath := filepath.Join(outputDir, fmt.Sprintf("%s.app", appName), "Contents")
		err = Save(filepath.Join(contentsPath, "Info.plist"), darwin_plist(appName))
		if err != nil {
			log.Fatalf("Failed to generate info.plist:%v", err)
		}
		// Prepare PkgInfo
		err = Save(filepath.Join(contentsPath, "PkgInfo"), darwin_pkginfo())
		if err != nil {
			log.Fatalf("Failed to generate PkgInfo:%v", err)
		}
		// Prepare icon
		resourcesPath := filepath.Join(contentsPath, "Resources")
		os.MkdirAll(resourcesPath, 0755)

		// Rename icon file name to [appName].icns
		cmd = exec.Command("cp", iconPath, filepath.Join(resourcesPath, fmt.Sprintf("%s.icns", appName)))
		cmd.Dir = projectPath
		err = cmd.Run()
		if err != nil {
			log.Fatalf("Failed to move icon to destination: %v", err)
		}

		fmt.Printf("%s.app is generated at %s/build/%s/\n", appName, projectPath, targetOS)

	case "windows":
		//@todo: implement windows bundle
		cmd := exec.Command("go", "build", "-ldflags=\"-s -w -H windowsgui\"", ".")
		cmd.Dir = projectPath
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to build: %v", err)
		}
	default:
		break
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	configDir := "./Configs"
	pcsyncDir := "./pcSync"

	files, err := filepath.Glob(filepath.Join(configDir, "*.json"))
	if err != nil {
		log.Fatal(err)
	}

	if len(files) == 0 {
		// Configs 폴더에 json 파일 없음
		log.Fatal("No project json files found in Configs directory.")
	}

	for _, file := range files {
		// 복사 중...
		fmt.Printf("\n==> Copying %s to pcconfig.json...\n", file)

		err := copyFile(file, filepath.Join(pcsyncDir, "pcconfig.json"))
		if err != nil {
			// 파일 복사 실패
			log.Fatalf("Failed to copy file: %v", err)
		}
		// 복사 완료!
		fmt.Println("pcconfig.json copy complete, running pcsync pushAll...")

		cmd := exec.Command("pcsync", "pushAll", "--yes")
		cmd.Dir = pcsyncDir
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("❌ pcsync pushAll error: %v\noutput: %s", err, string(output))
		} else {
			fmt.Println("✅ pcsync pushAll success")
			fmt.Println(string(output))
		}
	}
	// 모든 프로젝트 pcsync pushAll 작업 완료!
	fmt.Println("\nAll Projects pcsync pushAll success!")
	fmt.Scanln()
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

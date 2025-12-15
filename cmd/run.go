package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func detectBuildTool() (tool string, args []string, err error) {
	isWindows := runtime.GOOS == "windows"

	// Maven wrapper (preferred if exists)
	mvnw := "./mvnw"
	if isWindows {
		mvnw = "mvnw.cmd"
	}
	if fileExists(mvnw) {
		return mvnw, []string{"spring-boot:run"}, nil
	}

	// System Maven fall-back
	if _, err := exec.LookPath("mvn"); err == nil {
		return "mvn", []string{"spring-boot:run"}, nil
	}

	// Gradle wrapper
	gradlew := "./gradlew"
	if isWindows {
		gradlew = "gradlew.bat"
	}
	if fileExists(gradlew) {
		return gradlew,
			[]string{"bootRun"}, // Spring Boot plugin for Gradle
			nil
	}

	// System Gradle fall-back
	if _, err := exec.LookPath("gradle"); err == nil {
		return "gradle", []string{"bootRun"}, nil
	}

	return "", nil, fmt.Errorf("no Maven/Gradle build tool found")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a Spring and Quarkus project",
	Run: func(cmd *cobra.Command, args []string) {

		// Quick Spring Boot detection (properties or yml)
		if !(fileExists("src/main/resources/application.properties") ||
			fileExists("src/main/resources/application.yml")) {

			fmt.Println("❌ Not a Spring Boot project: no application.properties/yml found")
			os.Exit(1)
		}

		// Detect and select correct build tool
		tool, toolArgs, err := detectBuildTool()
		if err != nil {
			fmt.Printf("❌ %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("🚀 Starting Spring Boot using: %s %s\n",
			tool, strings.Join(toolArgs, " "))

		cmdExec := exec.Command(tool, toolArgs...)
		cmdExec.Stdout = os.Stdout
		cmdExec.Stderr = os.Stderr
		cmdExec.Stdin = os.Stdin

		if err := cmdExec.Run(); err != nil {
			fmt.Printf("❌ Error running Spring Boot project: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

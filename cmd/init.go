/*
Copyright © 2025 Subroto Kumar <subrotokumar@outlook.in>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/subrotokumar/springx/internal/initializr"
)

// Command line flags
var (
	groupID      string
	artifactID   string
	version      string
	description  string
	packageName  string
	javaVersion  string
	buildTool    string
	dependencies string
	outputDir    string
	listDeps     bool
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new Spring Boot project",
	Long: `Initialize a new Spring Boot project with the specified configuration.
	
You can customize the project by using various flags to specify dependencies,
build tool, Java version, and other project metadata.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initializr.Run()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Add flags
	initCmd.Flags().StringVarP(&groupID, "group", "g", "", "Group ID (e.g., com.example)")
	initCmd.Flags().StringVarP(&artifactID, "artifact", "a", "", "Artifact ID (defaults to project name)")
	initCmd.Flags().StringVarP(&version, "version", "v", "", "Project version (default: 0.0.1-SNAPSHOT)")
	initCmd.Flags().StringVarP(&description, "description", "d", "", "Project description")
	initCmd.Flags().StringVarP(&packageName, "package", "p", "", "Base package name")
	initCmd.Flags().StringVarP(&javaVersion, "java", "j", "", "Java version (default: 17)")
	initCmd.Flags().StringVarP(&buildTool, "build", "b", "", "Build tool: maven or gradle (default: maven)")
	initCmd.Flags().StringVarP(&dependencies, "deps", "", "", "Comma-separated list of dependencies")
	initCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory (defaults to project name)")
	initCmd.Flags().BoolVarP(&listDeps, "list", "l", false, "List available dependencies")
}

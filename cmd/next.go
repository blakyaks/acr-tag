package cmd

import (
	"os"
	"time"

	"github.com/blakyaks/acr_tag/tag"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var nextCmd = &cobra.Command{
	Use:              "next",
	Args:             cobra.NoArgs,
	TraverseChildren: true,
	Short:            "Generate the next tag",
	Long:             `Based on the current tags of the image and the input from the user, generates the next tag`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		tagType, _ := cmd.Flags().GetString("type")
		environment, _ := cmd.Flags().GetString("environment")
		registry, _ := cmd.Flags().GetString("registry")
		repository, _ := cmd.Flags().GetString("repository")
		debug, _ := cmd.Flags().GetBool("debug")
		tag.NextCmd(username, password, tagType, environment, registry, repository, debug)
	},
}

func init() {
	nextCmd.Flags().StringP("username", "u", "", "Username to authenticate to the registry")
	nextCmd.Flags().StringP("password", "p", "", "Password to authenticate to the registry")
	nextCmd.Flags().StringP("type", "t", "", "Tag type [major, minor, patch, rc, dev]")
	nextCmd.Flags().StringP("environment", "e", "", "Environment [dev, staging, prod]")
	nextCmd.Flags().StringP("registry", "r", "", "Azure Container Registry URL")
	nextCmd.Flags().StringP("repository", "", "", "Repository Name")
	nextCmd.Flags().BoolP("debug", "", false, "Debug")
	rootCmd.AddCommand(nextCmd)

	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = false
	formatter.ForceColors = true
	formatter.TimestampFormat = time.RFC1123

	formatter.SetColorScheme(&prefixed.ColorScheme{
		PrefixStyle:    "blue+b",
		TimestampStyle: "white+h",
	})

	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
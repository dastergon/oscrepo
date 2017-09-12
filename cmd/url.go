package cmd

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/dastergon/oscrepo/lib"
	"github.com/spf13/cobra"
)

const defaultRepoURL = "http://download.opensuse.org/repositories/%s/%s/%s.repo"

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Show the .repo URLs of the projects.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln("Please specify a repository to search for.")
		}
		word := args[0]

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		entry, _ := cmd.Flags().GetInt32("entry")

		if username == "" || password == "" {
			username, password = CfgUsername, CfgPassword
		}

		client := lib.NewBasicAuthClient(username, password)
		systemRelease := lib.GetSystemReleaseName()
		repositories, _ := client.GetRepositories()

		count := int32(0)
		for _, project := range repositories.Projects {

			var buffer bytes.Buffer
			if strings.Contains(project.Name, word) {
				parts := strings.Split(project.Name, ":")
				for i := range parts[0 : len(parts)-1] {
					buffer.WriteString(parts[i] + ":/")
				}
				buffer.WriteString(parts[len(parts)-1])
				availableReleases, _ := client.GetMeta(project.Name)
				releaseExists := false
				for _, m := range availableReleases.Names {
					if m.Name == systemRelease {
						releaseExists = true
					}
				}
				if !releaseExists {
					continue
				}
				count++
				repoURL := fmt.Sprintf(defaultRepoURL, buffer.String(), systemRelease, project.Name)
				if count == entry {
					fmt.Println(repoURL)
					break
				}
				fmt.Println(count, repoURL)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(urlCmd)
}

package cmd

import (
	"fmt"
	"github.com/lucasheld/pfbackup/pfsense"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	url      string
	user     string
	pass     string
	noVerify bool
	path     string

	rootCmd = &cobra.Command{
		Use:   "pfbackup",
		Short: "pfbackup backups pfSense configurations",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "", "", "pfSense url (required)")
	rootCmd.Flags().StringVarP(&user, "user", "", "", "pfSense username (required)")
	rootCmd.Flags().StringVarP(&pass, "pass", "", "", "pfSense password (required)")
	rootCmd.Flags().BoolVarP(&noVerify, "no-verify", "", false, "do not verify ssl certificate")
	rootCmd.Flags().StringVarP(&path, "path", "", ".", "path to output directory")

	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("user")
	rootCmd.MarkFlagRequired("pass")
}

func writeConfigFile(path string, content []byte) {
	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatalf("Failed to create config file: %v", err)
	}
}

func run() {
	url := strings.TrimSuffix(url, "/")
	settings := &pfsense.Settings{
		Url:      url,
		User:     user,
		Pass:     pass,
		NoVerify: noVerify,
	}
	client := pfsense.InitClient(settings)
	pf := &pfsense.Pfsense{
		Settings: settings,
		Client:   client,
	}

	pfsense.Login(pf)
	config := pfsense.GetConfig(pf)
	path = strings.TrimSuffix(path, "/")
	outDir := filepath.Join(path, config.Filename)
	writeConfigFile(outDir, config.Content)
}

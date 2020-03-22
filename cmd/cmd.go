package cmd

import (
	"fmt"
	"github.com/lucasheld/pfbackup/pfsense"
	"github.com/lucasheld/pfbackup/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	showVersion bool
	url         string
	user        string
	pass        string
	noVerify    bool
	path        string

	rootCmd = &cobra.Command{
		Use:   "pfbackup",
		Short: "pfbackup backups pfSense configurations",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkRequiredFlags(cmd.Flags())
		},
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
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Print the version number")

	rootCmd.Flags().StringVarP(&url, "url", "", "", "pfSense url (required)")
	rootCmd.Flags().StringVarP(&user, "user", "", "", "pfSense username (required)")
	rootCmd.Flags().StringVarP(&pass, "pass", "", "", "pfSense password (required)")
	rootCmd.Flags().BoolVarP(&noVerify, "no-verify", "", false, "do not verify ssl certificate")
	rootCmd.Flags().StringVarP(&path, "path", "", ".", "path to output directory")
}

func checkRequiredFlags(flags *pflag.FlagSet) error {
	// show version does not require any flags
	if showVersion {
		return nil
	}

	// check for missing flags
	missingFlags := []string{}
	flags.VisitAll(func(flag *pflag.Flag) {
		for _, f := range []string{"url", "user", "pass"} {
			if f == flag.Name && flag.Value.String() == flag.DefValue {
				missingFlags = append(missingFlags, f)
			}
		}
	})

	// return error of missing flags
	if len(missingFlags) != 0 {
		missingFlagsStr := ""
		for _, f := range missingFlags {
			if missingFlagsStr != "" {
				missingFlagsStr += ", "
			}
			missingFlagsStr += fmt.Sprintf("\"%s\"", f)
		}
		return fmt.Errorf("Required flag(s) %s not set", missingFlagsStr)
	}

	return nil
}

func printVersion() {
	fmt.Printf("pfbackup %s\n", version.Version)
	fmt.Printf("- os/arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("- go version: %s\n", runtime.Version())
}

func writeConfigFile(path string, content []byte) {
	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatalf("Failed to create config file: %v", err)
	}
}

func run() {
	if showVersion {
		printVersion()
		return
	}

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

	pf.Login()
	config := pf.GetConfig()
	path = strings.TrimSuffix(path, "/")
	outDir := filepath.Join(path, config.Filename)
	writeConfigFile(outDir, config.Content)
}

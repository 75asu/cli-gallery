/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"
	"log"

	"github.com/spf13/cobra"
)

// tzoneCmd represents the tzone command
var tzoneCmd = &cobra.Command{
	Use:   "tzone",
	Short: "Get the current time in a given timezone",
	Long: `Get the current time in a given timezone.
	This command takes one argument, the timezone you want to get the current time in.
	eg - tzone tzone EST --date 2006-01-02`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timezone := args[0]
		

		location, _ := time.LoadLocation(timezone)
		dateFlag, _ := cmd.Flags().GetString("date")
		var date string

		if dateFlag != "" {
			date = time.Now().In(location).Format(dateFlag)
			fmt.Printf("Current date in %v: %v\n", timezone, date)
		} else {
			currentTime, err := getTimeInTimezone(timezone)
			if err != nil {
				log.Fatalln("The timezone string is invalid")
			}
			fmt.Println(currentTime)
		}
		
		
	},
}

func init() {
	rootCmd.AddCommand(tzoneCmd)
	tzoneCmd.PersistentFlags().String("date", "", "returns the date in a time zone in a specified format")
}

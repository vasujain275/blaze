package main

import (
	"fmt"
	"os"
)

func main() {
	// Ensure a command is provided
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  ./blaze torrent parse /path/to/file.torrent")
		fmt.Println("  ./blaze torrent download /path/to/file.torrent")
		fmt.Println("  ./blaze magnet parse magnetlink")
		os.Exit(1)
	}

	// Switch on the main command
	switch os.Args[1] {
	case "torrent":
		handleTorrentCommand(os.Args[2:])
	case "magnet":
		handleMagnetCommand(os.Args[2:])
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

func handleTorrentCommand(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  ./blaze torrent parse /path/to/file.torrent")
		fmt.Println("  ./blaze torrent download /path/to/file.torrent")
		os.Exit(1)
	}

	switch args[0] {
	case "parse":
		fmt.Println("Parse")
	case "download":
		fmt.Println("Download")
	default:
		fmt.Println("Unknown subcommand for 'torrent':", args[0])
		os.Exit(1)
	}
}

func handleMagnetCommand(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  ./blaze magnet parse magnetlink")
		os.Exit(1)
	}

	switch args[0] {
	case "parse":
		fmt.Println("Parse")
	default:
		fmt.Println("Unknown subcommand for 'magnet':", args[0])
		os.Exit(1)
	}
}

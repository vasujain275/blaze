# Blazetorrent

## Directory structure:

```
blazetorrent/
│
├── cmd/
│   └── blazetorrent/
│       └── main.go          # CLI entry point
│
├── internal/
│   ├── bencode/
│   │   ├── bencode.go       # Bencode parsing and encoding functions
│   │   ├── decode.go        # Bencode decoding logic
│   │   ├── encode.go        # Bencode encoding logic
│   │   └── bencode_test.go  # Unit tests for Bencode functionality
│   │
│   └── torrent/
│       ├── torrent.go       # Torrent-related functions
│       ├── parse.go         # Torrent parsing logic
│       ├── download.go      # Torrent download logic
│       └── torrent_test.go  # Unit tests for torrent functionality
│
├── pkg/
│   └── blazetorrent/
│       ├── blazetorrent.go  # Public API for the blazetorrent library
│       ├── magnet.go        # Magnet link parsing logic
│       └── torrent.go       # Torrent-related functions using internal packages
│
├── go.mod                   # Go module file
└── go.sum                   # Go module dependencies

```

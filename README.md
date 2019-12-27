# player video downloader
 Player.pl command line video downloader

## Build
```
$ go build
```

## Usage:
```
$ ./player-video-downloader --help
Usage of player-video-downloader:
  -dir string
        Directory path to save the (default ".")
  -id string
        Video id
  -url string
        Player.pl video url
```

### Example

#### Download by video id
`$ ./player-video-downloader -id <videoId> [-dir <downloadDir>]`
```
$ ./player-video-downloader -id 157813
Video id provided.
Downloading...  complete 12.16% 19/157mb
```

#### Download by url
`$ ./player-video-downloader -url <url> [-dir <downloadDir>]`
```
`$ ./player-video-downloader -url player.pl/seriale-online/na-wspolnej-odcinki,144/odcinek-2999,S06E2999,157813
Video url provided.
Found video id:157813
Downloading...  complete 100.00% 157/157mb
Video saved to: ./157813.mp4
```

#### Download to selected dir
`$ ./player-video-downloader -id <videoId> -dir <downloadDir>]`
```
$ ./player-video-downloader -id 157813 -dir ../videos
Video id provided.
Downloading...  complete 100.00% 157/157mb
Video saved to: ../videos/157813.mp4
```



[![Go Report Card](https://goreportcard.com/badge/github.com/darkhz/invidtui)](https://goreportcard.com/report/github.com/darkhz/invidtui)
# invidtui

![demo](demo/demo.gif)

invidtui is an invidious client, which fetches data from invidious instances and displays a user interface in the terminal(TUI), and allows for selecting and playing Youtube audio and video.

Currently, it is tested on Linux and Windows, and it should work on MacOS.

## Features
- Play audio or video
- Search with history support
- Control the video resolution
- Ability to open, view, edit and save m3u8 playlists
- Automatically queries the invidious API and selects the best instance

## Requirements
- MPV
- Youtube-dl

## Installation
You can install the package either via the following command:
```go install github.com/darkhz/invidtui@latest ```

or check the Releases page and download the binary that matches your OS and architecture.

## Usage

    invidtui [<flags>]

    Flags:
      --video-res="720p"  Set the default video resolution.
      --close-instances   Close all currently running instances.
      --mpv-path="mpv"          Specify path to the mpv executable.
      --ytdl-path="youtube-dl"  Specify path to youtube-dl executable or its forks (yt-dlp, yt-dtlp_x86)
      --num-retries=100         Set the number of retries for connecting to the socket.

## Keybindings

### Search

> <kbd>/</kbd><br /> Show search input<br />
>
> <kbd>Ctrl</kbd> + <kbd>e</kbd><br /> Switch between search modes
> (video, playlist, channel)<br />

### Playlist Queue

> <kbd>p</kbd><br /> Open playlist queue. This control will work across
> all pages.<br />
>
> <kbd>Ctrl</kbd>+<kbd>o</kbd><br /> Open saved playlist<br />
>
> <kbd>Ctrl</kbd>+<kbd>s</kbd><br /> Save current playlist queue<br />
>
> <kbd>m</kbd><br /> Move an item in playlist queue. To cancel a move,
> just press <kbd>Enter</kbd> in the same position the move operation
> was started.<br />
>
> <kbd>d</kbd><br /> Delete an item in playlist queue<br />

### Player
Note: These controls will work across all pages (search, playlist or channel pages)<br /><br />

> <kbd>Space</kbd><br /> Pause/unpause<br />
>
> <kbd>Right</kbd><br /> Seek forward<br />
>
> <kbd>Left</kbd><br /> Seek backward<br />
>
> <kbd><</kbd><br /> Switch to previous track<br />
>
> <kbd>></kbd><br /> Switch to next track<br />
>
> <kbd>s</kbd><br /> Cycle shuffle mode (shuffle-playlist)<br />
>
> <kbd>m</kbd><br /> Cycle mute mode<br />
>
> <kbd>l</kbd><br /> Cycle repeat modes (repeat-file,
> repeat-playlist)<br />
>
> <kbd>Shift</kbd>+<kbd>s</kbd><br /> Stop player<br />

### Application

> <kbd>Ctrl</kbd>+<kbd>Z</kbd><br /> Suspend<br />
>
> <kbd>q</kbd><br /> Quit<br />


### Page-based Keybindings

> <kbd>i</kbd><br />
> This control works on the search and channel playlist pages.<br />
> Fetches the Youtube playlist contents from the currently selected entry and displays it in a separate playlist page. <br />
> In case you have exited this page, you can come back to it by pressing <kbd>Alt</kbd>+<kbd>i</kbd> instead of reloading the playlist again.<br/>
>
> <kbd>u</kbd><br />
> This control works on the search page.<br />
> Fetches only videos from a Youtube channel (from the currently selected entry) and displays it in a separate channel video page.<br />
> <kbd>Shift</kbd>+<kbd>u</kbd> fetches only playlists from a Youtube channel and displays it in a separate channel playlist page.
> In case you have exited<br /> this page, you can come back to it by pressing <kbd>Alt</kbd>+<kbd>u</kbd> instead of reloading the channel again.<br />
>
> <kbd>Enter</kbd><br />
> This control works on the search, playlist, channel video and channel playlist pages.<br />
> Fetches more results.<br />
>
> <kbd>a</kbd><br />
> This control works on the search, playlist and channel video list pages.<br />
> Fetches audio of the currently selected entry and adds it to the playlist.<br />
> If the selected entry is a playlist, all the playlist contents will be loaded into<br />
> the playlist queue as audio.
> To immediately play after adding to playlist, press <kbd>Shift</kbd>+<kbd>a</kbd>.<br/>
>
> <kbd>v</kbd><br />
> This control works on the search, playlist and channel video pages<br/>
> Fetches video of the currently selected entry and adds it to the playlist.<br />
> If the selected entry is a playlist, all the playlist contents will be loaded into<br />
> the playlist queue as video.
> To immediately play after adding to playlist, press <kbd>Shift</kbd>+<kbd>v</kbd>.<br/>
>
> <kbd>Ctrl</kbd>+<kbd>x</kbd><br />
> Cancel the fetching of playlist or channel contents (in case it takes a long time,<br/>
> due to slow network speeds for example).<br/>
>
> <kbd>Esc</kbd><br />
> Exit the current page.<br/>

## Additional Notes
- Since Youtube video titles may have many unicode characters (emojis for example), it is recommended to install **noto-fonts** and its variants (noto-fonts-emoji for example). Refer to your distro's documentation on how to install them. On Arch Linux for instance, you can install the fonts using pacman:
  `pacman -S noto-fonts noto-fonts-emoji noto-fonts-extra`<br/>

- For the video mode, only MP4 videos will be played, and currently there is no way to modify this behavior. This will change in later versions.

- The close-instances option should mainly be used if another invidtui instance may be using the socket, if there was an application crash, or if an error pops up like this: ``` Error: Socket exists at /home/test/.config/invidtui/socket, is another instance running?```.

- On Windows, using invidtui in Powershell/CMD will work, but use Windows Terminal for best results.

## Bugs
- Video streams from an invidious instance that are other than 720p or 360p can't currently be played properly when loaded from a saved playlist (only video will be played, audio won't), since we need to merge the audio and video streams, and I have yet to find a way to do that via the m3u8 playlist spec.

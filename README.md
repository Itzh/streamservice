# Small streaming service written in Go

`GET /stream?track={trackName}?seek={sec_ofset}`
returns an audio stream in chunks of 1 sec / sec.

**Note**: Hardcoded to use 320kbit
* `trackname` - the name of the audio file
* `sec_ofset` where to start the stream
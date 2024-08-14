<h1 align="center">
  <br>
  <img src="assets/logo.svg" alt="aos" width="400">
  <br>
  <br>
</h1>

<h5 align="center">aos - Alias on Steroids</h5>
<h4 align="center">A simple command line tool that helps you run complex command lines.</h4>


<p align="center">
  <img src="https://github.com/todoesverso/aos/actions/workflows/main.yml/badge.svg">
  <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square">
</p>

<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#usage">Usage</a> •
  <a href="#installation">Installation</a> •
  <a href="#download">Download</a> •
  <a href="#roadmap">Roadmap</a> •
  <a href="#license">License</a>
</p>

## Key Features

* Save complex command lines in a human readable yaml file and run it
* Add positional arguments
* Get detailed information on the command line 
* Render the final command

## Usage

  ```sh
  $ aos
  AliasOnSteroids

  Usage:
        aos <alias.yaml> [positional arguments of the alias]
  Options:
        In order to keep CLI arguments as straightforward as possible,
        options are passed thru the AOS environment variables.

        AOS=h   ./builds/aos <alias.yaml>       # Prints this usage
        AOS=H   ./builds/aos <alias.yaml>       # Prints a helper short description of the alias
        AOS=E   ./builds/aos <alias.yaml>       # Prints a helper long description of the alias
        AOS=R   ./builds/aos <alias.yaml>       # Renders the command and prints it to stdout
        [AOS=X] ./builds/aos <alias.yaml>       # Runs the command in a shell.

  $ AOS=R aos examples/ffmpeg_gif_to_mp4.yaml examples/input.gif /tmp/output.mp4
ffmpeg -i examples/input.gif -movflags faststart -pix_fmt yuv420p -vf scale=trunc(iw/2)*2:trunc(ih/2)*2 /tmp/output.mp4

  $ AOS=E aos examples/ffmpeg_gif_to_mp4.yaml examples/input.gif /tmp/output.mp4
ffmpeg: video converter
├──-i: input file
├──input.gif: The input GIF
├──-movflags: This option is used to optimize the MP4 file for web
│   streaming. It moves the metadata to the beginning of the
│   file, allowing the video to start playing before it is fully
│   downloaded
│
├──-pix_fmt: This sets the pixel format of the output video. yuv420p is a
│   widely compatible pixel format that ensures the video can be
│   played on most devices and platforms.
│
├──-vf: This applies a video filter (-vf) to scale the video. The
│   scale filter is used to resize the video dimensions. The
│   expression trunc(iw/2)*2 ensures that the width (iw) is an
│   even number, and trunc(ih/2)*2 does the same for the height
│   (ih). This is important because some video codecs require
│   even dimensions
│
└──out.mp4: The output MP4 file

  $ AOS=H aos examples/ffmpeg_gif_to_mp4.yaml examples/input.gif /tmp/output.mp4
┌─ examples/ffmpeg_gif_to_mp4.yaml ────────────────────────────┐
| This AOS converts a GIF file into a MP4 using ffmpeg.        |
|                                                              |
| The provided FFmpeg command processes a GIF image into       |
| an MP4 video format. It employs the -i flag to specify       |
| the input file, input.gif.gif, and outputs the processed     |
| video as output.mp4.                                         |
| The -movflags faststart option optimizes the MP4 container   |
| for progressive playback by relocating essential metadata.   |
| The -pix_fmt yuv420p flag mandates a YUV420p pixel format    |
| for the output video, a widely supported format known for    |
| its balance of color information and compression efficiency. |
| Lastly, the -vf scale=trunc(iw/2)*2:trunc(ih/2)*2 filter     |
| rescales the input image to the nearest even dimensions by   |
| truncating the width and height to multiples of two,         |
| potentially improving compatibility with certain video       |
| codecs and hardware.                                         |
|                                                              |
|                                                              |
└──────────────────────────────────────────────────────────────┘


  $  cat examples/ffmpeg_gif_to_mp4.yaml
description: |
  This AOS converts a GIF file into a MP4 using ffmpeg.

  The provided FFmpeg command processes a GIF image into
  an MP4 video format. It employs the -i flag to specify
  the input file, input.gif.gif, and outputs the processed
  video as output.mp4.
  The -movflags faststart option optimizes the MP4 container
  for progressive playback by relocating essential metadata.
  The -pix_fmt yuv420p flag mandates a YUV420p pixel format
  for the output video, a widely supported format known for
  its balance of color information and compression efficiency.
  Lastly, the -vf scale=trunc(iw/2)*2:trunc(ih/2)*2 filter
  rescales the input image to the nearest even dimensions by
  truncating the width and height to multiples of two,
  potentially improving compatibility with certain video
  codecs and hardware.

command:
  exec: ffmpeg
  description: video converter

arguments:
  - option: -i
    description: input file
  - positional:
      name: input.gif
      description: The input GIF
  - option: -movflags
    value: faststart
    description: >
      This option is used to optimize the MP4 file for web streaming.
      It moves the metadata to the beginning of the file, allowing the
      video to start playing before it is fully downloaded
  - option: -pix_fmt
    value: yuv420p
    description: >
      This sets the pixel format of the output video.
      yuv420p is a widely compatible pixel format that ensures the
      video can be played on most devices and platforms.
  - option: -vf
    value: "scale=trunc(iw/2)*2:trunc(ih/2)*2"
    description: >
      This applies a video filter (-vf) to scale the video.
      The scale filter is used to resize the video dimensions.
      The expression trunc(iw/2)*2 ensures that the width (iw) is an even number,
      and trunc(ih/2)*2 does the same for the height (ih).
      This is important because some video codecs require even dimensions
  - positional:
      name: out.mp4
      description: The output MP4 file
  ```

## Installation

## Download

Pre compiled binaries for several platforms can be downloaded from the [release](https://github.com/todoesverso/aos/releases) section.

## Roadmap

This is just a personal project intended to learn some Go but I'll keep adding features that I see useful or requested by the users. 

## License

MIT

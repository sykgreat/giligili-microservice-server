#!/bin/sh

ffmpeg -y -hwaccel videotoolbox \
  -i "$1" \
  -preset fast -g 120 \
  -sc_threshold 0 \
  -map 0:0 -map 0:1 \
  -map 0:0 -map 0:1 \
  -s:v:0 1280*720 -b:v:0 2800k \
  -s:v:2 858*480 -b:v:2 1200k \
  -crf 23.5 \
  -c:v h264_videotoolbox \
  -c:a copy \
  -var_stream_map "v:0,a:0,name:720p v:1,a:1,name:480p" \
  -f hls \
  -hls_time 6 \
  -hls_playlist_type vod \
  -master_pl_name master.m3u8 \
  -hls_segment_filename "$2""/v%v/seg%d.ts" "$2""/v%v/index.m3u8"

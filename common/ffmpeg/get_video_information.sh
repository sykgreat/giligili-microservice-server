#!/bin/sh

ffprobe -i "$1" -v quiet -print_format json -show_format -show_streams -hide_banner
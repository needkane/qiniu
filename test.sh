#!/bin/bash
ffmpeg -y -i Downloads/kane2.mpg -acodec libfaac -b:a 30k -ar 44100 -r 15 -ac 2 -s 480x272 -vcodec libx264 -refs 2 -x264opts keyint=150:min-keyint=15 -vprofile baseline -level 20 -b:v 200k -vf "drawtext=fontfile=/mnt/hgfs/zm/simhei.ttf: text='来源：迅雷':x=100:y=x/dar:draw='if(gt(n,0),lt(n,250))':fontsize=24:fontcolor=yellow@0.5:shadowy=2"  drawtext.mp4   

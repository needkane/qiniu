#!/bin/bash


for vname in `find ./ -name "ktest*"`

 do 	
 		  en=`expr length $vname` 
          echo $vname;file $vname;
          name=${vname:2:en-6}
# superfast
date>x; 
/usr/bin/time -f '\ntime:%S+%U\n' ./ffmpeg -loglevel warning -i $vname -f mp4 -preset superfast -movflags faststart -acodec libfaac -vcodec libx264 -scodec mov_text -r 25 -vf scale=trunc'(in_w/2)*2:trunc(in_h/2)'*2 -level 3.0 -coder 0 -g 30 -qmax 51 -partitions +parti4x4+partp8x8+partb8x8 -rc_eq 'blurCplx^(1-qComp)' -qmin 10 -trellis 1 -pix_fmt yuv420p -r 30 -b:v 300k -codec:v libx264 -pix_fmt yuv420p -ar 22050 -b:a 128k -codec:a libfaac -map_metadata 0:g -vf [in]scale=floor'(min(320/in_w\,240/in_h)*in_w/2)*2:floor(min(320/in_w\,240/in_h)*in_h/2)'*2[scale1] -y $name+00.mp4 2>>x;
date>>x;
/usr/bin/time -f '\ntime:%S+%U\n' ./ffmpeg -loglevel warning -i $vname -f mp4 -preset superfast -movflags faststart -acodec libfaac -vcodec libx264 -scodec mov_text -r 25 -vf scale=trunc'(in_w/2)*2:trunc(in_h/2)'*2 -level 3.0 -coder 0 -g 30 -qmax 51 -partitions +parti4x4+partp8x8+partb8x8 -rc_eq 'blurCplx^(1-qComp)' -qmin 10 -trellis 1 -pix_fmt yuv420p -r 30 -b:v 1152000 -codec:v libx264 -pix_fmt yuv420p -ar 44100 -b:a 160k -codec:a libfaac -map_metadata 0:g -vf [in]scale=floor'(min(1280/in_w\,720/in_h)*in_w/2)*2:floor(min(1280/in_w\,720/in_h)*in_h/2)'*2[scale1] -y $name+01.mp4 2>>x			
date>>x;
#fast 
/usr/bin/time -f '\ntime:%S+%U\n' ./ffmpeg -loglevel warning -i $vname -f mp4 -preset fast -movflags faststart -acodec libfaac -vcodec libx264 -scodec mov_text -r 25 -vf scale=trunc'(in_w/2)*2:trunc(in_h/2)'*2 -level 3.0 -coder 0 -g 30 -qmax 51 -partitions +parti4x4+partp8x8+partb8x8 -rc_eq 'blurCplx^(1-qComp)' -qmin 10 -trellis 1 -pix_fmt yuv420p -r 30 -b:v 300k -codec:v libx264 -pix_fmt yuv420p -ar 22050 -b:a 128k -codec:a libfaac -map_metadata 0:g -vf [in]scale=floor'(min(320/in_w\,240/in_h)*in_w/2)*2:floor(min(320/in_w\,240/in_h)*in_h/2)'*2[scale1] -y $name+10.mp4 2>>x;
date>>x;
/usr/bin/time -f '\ntime:%S+%U\n' ./ffmpeg -loglevel warning -i $vname -f mp4 -preset fast -movflags faststart -acodec libfaac -vcodec libx264 -scodec mov_text -r 25 -vf scale=trunc'(in_w/2)*2:trunc(in_h/2)'*2 -level 3.0 -coder 0 -g 30 -qmax 51 -partitions +parti4x4+partp8x8+partb8x8 -rc_eq 'blurCplx^(1-qComp)' -qmin 10 -trellis 1 -pix_fmt yuv420p -r 30 -b:v 1152000 -codec:v libx264 -pix_fmt yuv420p -ar 44100 -b:a 160k -codec:a libfaac -map_metadata 0:g -vf [in]scale=floor'(min(1280/in_w\,720/in_h)*in_w/2)*2:floor(min(1280/in_w\,720/in_h)*in_h/2)'*2[scale1] -y $name+11.mp4 2>>x  
date>>x;
#medium
/usr/bin/time -f '\ntime:%S+%U\n' ./ffmpeg -loglevel warning -i $vname -f mp4 -preset medium -movflags faststart -acodec libfaac -vcodec libx264 -scodec mov_text -r 25 -vf scale=trunc'(in_w/2)*2:trunc(in_h/2)'*2 -level 3.0 -coder 0 -g 30 -qmax 51 -partitions +parti4x4+partp8x8+partb8x8 -rc_eq 'blurCplx^(1-qComp)' -qmin 10 -trellis 1 -pix_fmt yuv420p -r 30 -b:v 300k -codec:v libx264 -pix_fmt yuv420p -ar 22050 -b:a 128k -codec:a libfaac -map_metadata 0:g -vf [in]scale=floor'(min(320/in_w\,240/in_h)*in_w/2)*2:floor(min(320/in_w\,240/in_h)*in_h/2)'*2[scale1] -y $name+20.mp4 2>>x;
date>>x;
/usr/bin/time -f '\ntime:%S+%U\n' ./ffmpeg -loglevel warning -i $vname -f mp4 -preset medium -movflags faststart -acodec libfaac -vcodec libx264 -scodec mov_text -r 25 -vf scale=trunc'(in_w/2)*2:trunc(in_h/2)'*2 -level 3.0 -coder 0 -g 30 -qmax 51 -partitions +parti4x4+partp8x8+partb8x8 -rc_eq 'blurCplx^(1-qComp)' -qmin 10 -trellis 1 -pix_fmt yuv420p -r 30 -b:v 1152000 -codec:v libx264 -pix_fmt yuv420p -ar 44100 -b:a 160k -codec:a libfaac -map_metadata 0:g -vf [in]scale=floor'(min(1280/in_w\,720/in_h)*in_w/2)*2:floor(min(1280/in_w\,720/in_h)*in_h/2)'*2[scale1] -y $name+21.mp4 2>>x  
date>>x;
#slow
/usr/bin/time -f '\ntime:%S+%U\n' ./ffmpeg -loglevel warning -i $vname -f mp4 -preset slow -movflags faststart -acodec libfaac -vcodec libx264 -scodec mov_text -r 25 -vf scale=trunc'(in_w/2)*2:trunc(in_h/2)'*2 -level 3.0 -coder 0 -g 30 -qmax 51 -partitions +parti4x4+partp8x8+partb8x8 -rc_eq 'blurCplx^(1-qComp)' -qmin 10 -trellis 1 -pix_fmt yuv420p -r 30 -b:v 300k -codec:v libx264 -pix_fmt yuv420p -ar 22050 -b:a 128k -codec:a libfaac -map_metadata 0:g -vf [in]scale=floor'(min(320/in_w\,240/in_h)*in_w/2)*2:floor(min(320/in_w\,240/in_h)*in_h/2)'*2[scale1] -y $name+30.mp4 2>>x;
date>>x;
/usr/bin/time -f '\ntime:%S+%U\n' ./ffmpeg -loglevel warning -i $vname -f mp4 -preset slow -movflags faststart -acodec libfaac -vcodec libx264 -scodec mov_text -r 25 -vf scale=trunc'(in_w/2)*2:trunc(in_h/2)'*2 -level 3.0 -coder 0 -g 30 -qmax 51 -partitions +parti4x4+partp8x8+partb8x8 -rc_eq 'blurCplx^(1-qComp)' -qmin 10 -trellis 1 -pix_fmt yuv420p -r 30 -b:v 1152000 -codec:v libx264 -pix_fmt yuv420p -ar 44100 -b:a 160k -codec:a libfaac -map_metadata 0:g -vf [in]scale=floor'(min(1280/in_w\,720/in_h)*in_w/2)*2:floor(min(1280/in_w\,720/in_h)*in_h/2)'*2[scale1] -y $name+31.mp4 2>>x  
date>>x;

#m3u8
/usr/bin/time -f '\ntime:%S+%U\n' ./ffmpeg -err_detect aggressive -i $vname -g 90 -dcodec copy -scodec mov_text -acodec libfaac -v warning -flags +loop+mv4 -pix_fmt yuv420p -cmp 256 -partitions +parti4x4+partp8x8+partb8x8 -subq 7 -trellis 1 -refs 5 -coder 0 -me_range 16 -keyint_min 25 -sc_threshold 40 -i_qfactor 0.71 -rc_eq 'blurCplx^(1-qComp)' -qcomp 0.6 -qmin 10 -qmax 51 -qdiff 4 -r 25 -r 30 -b:v 1000k -codec:v libx264 -pix_fmt yuv420p -ar 44100 -b:a 128k -codec:a libfaac -map_metadata 0:g -vf [in]scale=trunc'(in_w/2)*2:trunc(in_h/2)'*2[scale2] -f mp4 -y $name+40.mp4 2>>x
date>>x;
/usr/bin/time -f '\ntime:%S+%U\n' ./ffmpeg -i $vname -v warning -sn -acodec copy -vcodec copy -bsf h264_mp4toannexb -map 0 -f ssegment -segment_format mpegts -segment_list_type m3u8 -segment_time 10 -force_key_frames expr:gte'(t,n_forced*10)' -segment_list_entry_prefix http://needkane.qiniudn.com/MBlL71BJzYZ_T66xjn55UV4WIzc=/luD2OYJLES9-j0VA_xxE9EEJa8ED/ -segment_list /home/qboxserver/test/m3u8+$name/file.m3u8 /home/qboxserver/test//m3u8+$name/seg%d 2>>x
date>>x;

 	

 	
done;


	

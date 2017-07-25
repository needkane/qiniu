 1  cd ~/qbox/ffmpeg/ffmpeg-static/
    2  vi build.sh 
    3  cd build
    4  ls
    5  cd ../
    6  vi build.sh 
    7  ./build.sh 
    8  ls
    9  rm build2
   10  rm -rf  build2
   11  cd build/
   12  ls
   13  cd ../
   14  ls
   15  ./build.sh 
   16  vi /home/qboxtest/qbox/ffmpeg/ffmpeg-static/build/libilbc/signal_processing/vector_scaling_operations_mips.c
   17  vi build.sh
   18  cd build/
   19  ls
   20  cd fribidi-0.19.6/
   21  ls
   22  vi README 
   23  vi INSTALL 
   24  cd ../libass-0.12.0/
   25  LS
   26  ls
   27  vi install-sh 
   28  vi Makefile
   29  cd ../fribidi-0.19.6/
   30  ls
   31  vi README 
   32  vi INSTALL 
   33  cd ../libass-0.12.0/
   34  ls
   35  vi Makefile
   36  ls
   37  cd ..
   38  ls
   39  gcc -v
   40  ls
   41  cd libilbc/
   42  ls
   43  vi Makefile
   44  ls
   45  vi README 
   46  make
   47  vi INSTALL 
   48  ls
   49  cd ilibc
   50  cd ilib
   51  cd ilbc
   52  ls
   53  ./configure
   54  cd ..
   55  ls
   56  cd ilbc/
   57  ls
   58  cd ..
   59  ls
   60  vi INSTALL 
   61  vi ../../build.sh 
   62  cmake .
   63  vi ../../build.sh 
   64  make clean
   65  make 
   66  ls
   67  cd CMakeFiles/
   68  ls
   69  rm * -fr
   70  cd ..
   71  ls
   72  cmake .
   73  make
   74  cd ..
   75  ls
   76  cd libilbc/
   77  ls
   78  cd ../
   79  ll fribidi-0.19.6/
   80  ll 
   81  ll |grep fribidi-0.19.
   82  git branch 
   83  cd ~/qbox/ffmpeg/fop
   84  cd src/qbox.us/
   85  ls
   86  grep 'GetStart(' * -r
   87  git status
   88  git add ffmpegutil/argcheck/arg_check.go
   89  git add fop/avconv/avconv.go
   90  git commit -m "check"
   91  git push origin sewise 
   92  cd avconv
   93  cd fop/avconb
   94  cd fop/avconv
   95  go test
   96  git status
   97  git diff ../../ffmpegutil/argcheck/arg_check.go
   98  cd ../m3u8/
   99  go test
  100  git status
  101  go test
  102  git status
  103  git add m3u8.go
  104  git add ../../ffmpegutil/argcheck/arg_check.go
  105  git commit -m "fix"
  106  git push origin sewise 
  107  cd ~/qbox/ffmpeg/
  108  grep 'getStart(' * -r
  109  vi /etc/ssh/
  110  sudo vi /etc/ssh/ssh_host_dsa_key
  111  sudo vi /etc/ssh/ssh_host_ecdsa_key
  112  sudo vi /etc/ssh/ssh_host_rsa_key
  113  cd ~/.ssh/
  114  ls
  115  vi github
  116  ifconfig
  117  nc -l 1234 < github
  118  cd ~/DE
  119  cd ~/Desktop/
  120  scp -P 18022 zhangle@10.8.0.1:a.jpg .
  121  eog a.jpg
  122  for i in `nb{669..727}` ; do echo $i;done
  123  for i in nb{669..727} ; do echo $i,;done
  124  ssh -p 18022 zhangle@10.8.0.1
  125  top
  126  nc 192.168.1.131 12345 > libilbc.tgz
  127  ll libilbc.tgz 
  128  vim young.srt
  129  vi young.ass
  130  enconv
  131  sudo apt-get install enca
  132  enconv -L latin1 -x UTF-8 young.srt
  133  vi young.srt 
  134  enconv -L latin1 -x UTF-8 young.srt
  135  enconv -L none -x UTF-8 young.srt
  136  rm young.*
  137  vi giver.srt
  138  md5sum giver.srt 
  139  vi american.srt 
  140  gedit american.srt 
  141  enconv -L none -x UTF-8 giver.srt
  142  vi giver.srt 
  143  enconv -L none -x UTF-8 giver.srt 
  144  enconv -x UTF-8 giver.srt 
  145  enconv -L latin -x UTF-8 giver.srt 
  146  enconv --list language 
  147  enca giver.srt 
  148  enca giver.ass
  149  vi abc
  150  enca abc
  151  iconv
  152  iconv -h
  153  iconv --help
  154  iconv -f UTF-8 -t latin1 giver.srt -o GIV
  155  iconv -f UTF-8 -t latin1 giver.srt 
  156  vi giver.srt 
  157  vi into1.srt 
  158  ffmpeg -i sea -vf subtitles=into1.srt out.mp4
  159  ./ffmpeg -i sea -vf subtitles=into1.srt out.mp4
  160  vi into1.srt 
  161  ./ffmpeg -i sea -vf subtitles=into1.srt out.mp4
  162  ./ffmpeg -i sea -vf subtitles=into1.srt -sub_charenc utf-8 out.mp4
  163  vi into1.srt 
  164  ./ffmpeg -i sea -vf subtitles=into1.srt out.mp4
  165  vlc out.mp4 
  166  vi giver.srt
  167  ./ffmpeg -i sea -vf subtitles=giver.srt -sub_charenc latin1 out.mp4
  168  enca 
  169  enca --list languages
  170  enca -h
  171  enca -g giver.ass 
  172  enca -g into1.srt 
  173  enca into1.srt 
  174  vi into1.srt 
  175  iconvreader -h
  176  iconv --help
  177  iconv -f giver.srt
  178  enca into1.srt 
  179  sudo apt-get intall enca
  180  enca -h
  181  enca into1.srt 
  182  locale
  183  enca -L zh giver.srt
  184  enca -L zh into1.srt 
  185  enconv -L zh_CN -x UTF-8 giver.srt
  186  vi giver.srt
  187  ./ffmpeg -i sea -vf subtitles=giver.srt  out.mp4
  188  vlc out.mp4


1  cd ~/qbox/ffmpeg/ffmpeg-static/
    2  vi build.sh
    3  cd ~/Desktop/
    4  ls
    5  ifconfig
    6  ls
    7  ll |grep libass-0.12.0.tar.gz 
    8  ll |grep fribidi-0.19.6.tar.bz2 
    9  nc -l  -p 1234 < libass-0.12.0.tar.gz 
   10  nc -l  1234 < libass-0.12.0.tar.gz 
   11  nc -l  1234 < fribidi-0.19.6.tar.bz2 
   12  nc -l  1234 < libass-0.12.0.tar.gz 
   13  nc -l  1234 < fribidi-0.19.6.tar.bz2 
   14  nc -l  1234 < fribidi-0.19.6.tar.bz2 libass-0.12.0.tar.gz 
   15  nc -l  1234 < fribidi-0.19.6.tar.bz2 
   16  nc -l  1234 < libass-0.12.0.tar.gz 
   17  ll fribidi-0.19.6.tar.bz2 
   18  curl -o xxx 192.168.18.64:12345
   19  ll xx
   20  ll xxx
   21  rm xx
   22  sudo vi /etc/nginx/nginx.conf 
   23  sudo /etc/init.d/nginx restart
   24  http_proxy=http://127.0.0.1:82 curl http://www.baidu.com
   25  sudo -s
   26  cd ~/qbox/service/
   27  rm service/bin/qboximage 
   28  cp ~/Desktop/xxx service/bin/qboximage
   29  ls
   30  vi tools/script/startone.py 
   31  vi /home/qboxtest/qbox/image/fop_cgo/src/qbox.us/app/qboximage/qboximage.conf 
   32  ./go.sh 
   33  ./check.sh 
   34  tail run/qboximage.log 
   35  ./stop.sh 
   36  ll service/bin/qboximgave 
   37  ll service/bin/qboximage 
   38  cd ~/qbox/image/
   39  git branch 
   40  git status
   41  git add .
   42  git checkout .
   43  git checkout -b fsize_default  origin/develop 
   44  git checkout .
   45  git stash
   46  git checkout -b fsize_default  origin/develop 
   47  git fetch origin 
   48  git merge origin/develop 
   49  git status
   50  git add .
   51  git commit -m "default"
   52  git push origin fsize_default 
   53  git push mine fsize_default 
   54  git branch -D fsize_supplement 
   55  ls
   56  cp ~/Desktop/sewise/ .
   57  cp -r ~/Desktop/sewise/ .
   58  ls
   59  rm -rf sewise/
   60  git checkout sewise
   61  git branch 
   62  cd ../ffmpeg/
   63  git checkout sewise 
   64  git branch 
   65  cd fop/src/qbox.us/fop/avconv/
   66  go test
   67  git status
   68  git diff ../../app/qboxffmpeg/qboxffmpeg.conf
   69  git add avconv.go
   70  git commit -m "delete json "
   71  git push origin sewise 
   72  cd ~/qbox/ffmpeg/
   73  git reset --soft 421ceb3df331f757d25f482b88b6b50daaa3542f
   74  git commit -m "soft"
   75  git push origin sewise -f
   76  git fetch qbox/ffmpeg 
   77  git merge qbox/ffmpeg/develop 
   78  git log
   79  view fop/src/qbox.us/fop/avconv/avconv_test.go
   80  git status
   81  git add fop/src/qbox.us/fop/avconv/avconv.go
   82  git add fop/src/qbox.us/fop/avconv/avconv_test.go
   83  git add  fop/src/qbox.us/ffmpegutil/argcheck/arg_check.go
   84  cd fop/src/qbox.us/fop/avconv
   85  go test
   86  go test -run=TestCopy
   87  go test -run=TestConvResolution
   88  go test -run=TestRotate
   89  go test -run=TestTimeClip
   90  go test -run=TestAVConv
   91  go test -run=TestAn
   92  go test
   93  go test -run=TestVn
   94  go test -run=TestCreationTime
   95  go test -run=TestConvertTS
   96  go test -run=TestWriteXing
   97  go tesrt
   98  go test
   99  go test -run=TestH264Profile
  100  go test -run=TestH264Level
  101  go test -run=TestCompressLevel
  102  go test -run=TestVSpeed
  103  go test -run=TestAutoMinBitrate
  104  go test -run=TestMixWaterMark
  105  git status
  106  git add avconv.go
  107  git commit -m "fix"
  108  git push origin sewise 
  109  git checkout duration 
  110  go test -run=TestMixWaterMark
  111  git checkout sew
  112  git checkout sewise 
  113  git checkout .
  114  git checkout sewise 
  115  go test
  116  go test -run=TestMixWaterMark
  117  go test -run=TestTextWaterMark
  118  go test -run=TestWatermark
  119  git status
  120  git checkout duration 
  121  go test -run=TestWatermark
  122  go test -run=TestTextWatermark
  123  go test -run=TestTextWaterMark
  124  go test -run=TestMixWaterMark
  125  go test -run=TestWatermark
  126  go test -run=TestMixWaterMark
  127  go test
  128  git status
  129  git add avconv.go
  130  git commit -m "fix"
  131  git push origin sewise q
  132  git push origin sewise 
  133  cd ~/qbox/ffmpeg/
  134  make test
  135  cd fop/src/qbox.us/fop/m3u8/
  136  go test
  137  git status
  138  cd ~/qbox/ffmpeg/
  139  pwd
  140  ls
  141  cd ffmpeg-static/
  142  ls
  143  cp build.sh tmp.sh
  144  vi tmp.sh 
  145  ./tmp.sh -j 10
  146  pwd
  147  ls
  148  cd target/
  149  ls
  150  pwd
  151  ls
  152  cd ..
  153  ls
  154  cd build/
  155  ls
  156  cd ..
  157  ls
  158  vi tmp.sh 
  159  ./tmp.sh -j 10
  160  vi tmp.sh 
  161  ./tmp.sh 
  162  vi tmp.sh 
  163  ./tmp.sh 
  164  cd build/
  165  ls
  166  mv libilbc/ git_libilbc
  167  tar xvf ~/Desktop/libilbc.tgz 
  168  cd ..
  169  ls
  170  vi tmp.sh 
  171  ./tmp.sh -j 0
  172  ./tmp.sh -j 10
  173  vi tmp.sh 
  174  ./tmp.sh -j 10
  175  cd build/libilbc/
  176  ls
  177  rm CMakeFiles/ -fr
  178  cd -
  179  ./tmp.sh -j 10
  180  cd -
  181  ls
  182  rm CMakeCache.txt 
  183  cd -
  184  ./tmp.sh -j 10
  185  ls
  186  cd build/
  187  ls
  188  tar zxvf ../build_src.tgz 
  189  ls
  190  cd ..
  191  ls
  192  vi build.sh 
  193  cd build/
  194  ls
  195  tar xvf ~/Desktop/libass-0.12.0.tar.gz 
  196  tar xvf ~/Desktop/fribidi-0.19.6.tar.bz2 
  197  ls
  198  cd ..
  199  ls
  200  vi tmp.sh 
  201  ./tmp.sh -j 10
  202  cd -
  203  ls
  204  rm libilbc/ -fr
  205  tar xvf ~/Desktop/libilbc.tgz 
  206  cd ..
  207  ls
  208  ./tmp.sh -j 10
  209  cd- 
  210  cd -
  211  ls
  212  cd libilbc/
  213  ls
  214  rm CMakeFiles/ CMakeCache.txt -fr
  215  pwd
  216  cd -
  217  ls
  218  pwd
  219  ls
  220  cd ..
  221  ls
  222  ./tmp.sh -j 10
  223  163546
  224  sudo apt-get install  fontconfig-dev
  225  sudo apt-get install  libfontconfig-dev
  226  sudo apt-get install libass
  227  sudo apt-get install libass-dev
  228  ls
  229  vi tmp.sh 
  230  ./tmp.sh -j 10
  231  pwd
  232  ls
  233  cd build/ffmpeg-2.2.1/
  234  make clean;
  235  cd -
  236  ./tmp.sh 
  237  ls
  238  cd target/
  239  ls
  240  cd lib/
  241  ls
  242  rm libfribidi.*
  243  rm libass.*
  244  ls
  245  cd ..
  246  ls
  247  cd include/
  248  ls
  249  cd ..
  250  ls
  251  cd ..
  252  ls
  253  cd target/
  254  ls
  255  rm * -fr
  256  ls
  257  cd ..
  258  ls
  259  vi tmp.sh 
  260  ./tmp.sh -j 10
  261  vi tmp.sh 
  262  ./tmp.sh -j 10
  263  vi tmp.sh 
  264  ./tmp.sh -j 10
  265* 
  266  ./tmp.sh -j 10
  267  vi tmp.sh 
  268  ./tmp.sh -j 10
  269  ls
  270  vi tmp.sh 
  271  cd build/
  272  ls
  273  cd ffmpeg-2.2.1/
  274  ls
  275  make clean
  276  ls
  277  vi Makefile 
  278  vi config.mak 
  279  rm config.mak Makefile 
  280  cd -
  281  ls
  282  cd ..
  283  ls
  284  ./tmp.sh -j 10
  285  ./tmp.sh
  286  ls
  287  cd build/
  288  ls
  289  cd ffmpeg-2.2.1/
  290  ls
  291  ./configure 
  292  ls
  293  make
  294  vi Makefile 
  295  make
  296  cd ..
  297  ls
  298  cd ..
  299  ls
  300  cd ttt/
  301  ls
  302  tar zxvf ../build_src.tgz 
  303  ls
  304  cd ffmpeg-2.2.1/
  305  ls
  306  vi Makefile 
  307  cd ..
  308  ls
  309  cp ffmpeg-2.2.1/ ../build -r
  310  cd ../build/
  311  ls
  312  cd ffmpeg-2.2.1/
  313  ls
  314  vi Makefile 
  315  vi config.mak 
  316  rm config.mak 
  317  cd ..
  318  ls
  319  cd ..
  320  ls
  321  ./tmp.sh 
  322  ls
  323  cd target/
  324  ls


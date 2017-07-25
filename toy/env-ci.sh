if [ "$QBOXROOT" = "" ]; then
  QBOXROOT=$(cd ../; pwd)
  export QBOXROOT
fi

GOPATH=$QBOXROOT/account/account:$QBOXROOT/base/com:$QBOXROOT/base/biz:$QBOXROOT/base/qiniu:`pwd`/fop
export GOPATH

#!/usr/bin/env bash

CRYPTO_SDK_LIB_DIR="./crypto-sdk/lib"
CRYPTO_SDK_SWIGFILE="lib.swigcxx"
CRYPTO_SDK_BUILD="${HOME}/dev/orbs/go/src/github.com/orbs-network/crypto-sdk-go/crypto-sdk/build"
#LIBGCRYPT_LIB="${CRYPTO_SDK_BUILD}/mac/libgcrypt/lib/libgcrypt.a"
LIBGCRYPT_LIB_DIR="${CRYPTO_SDK_BUILD}/mac/libgcrypt/lib"
#LIBGPG_LIB="${CRYPTO_SDK_BUILD}/mac/libgpg-error/lib/libgpg-error.a"
LIBGPG_LIB_DIR="${CRYPTO_SDK_BUILD}/mac/libgpg-error/lib"
#LIBGCYPT_H="${HOME}/dev/orbs/go/src/github.com/orbs-network/crypto-sdk-go/crypto-sdk/deps/libgcrypt/libgcrypt-1.8.2/src"

LIBGCYPT_H="${CRYPTO_SDK_BUILD}/mac/libgcrypt/include"
LIBGPG_H="${CRYPTO_SDK_BUILD}/mac/libgpg-error/include"

echo "===> SWIG: Building Go wrapper for C++ code <==="

#export GOOS=linux
#export GOARCH=amd64
export CGO_CPPFLAGS="-I${LIBGCYPT_H} -I${LIBGPG_H}"
export CGO_LDFLAGS="-L${LIBGCRYPT_LIB_DIR} -L${LIBGPG_LIB_DIR} -lgcrypt -lgpg-error"
export CGO_ENABLED=1

#swig -c++ -Wall -go -cgo -intgosize 64 ${CRYPTO_SDK_LIB_DIR}/${CRYPTO_SDK_SWIGFILE}
#swig -go -cgo -intgosize 64 ${CRYPTO_SDK_LIB_DIR}/${CRYPTO_SDK_SWIGFILE}
if [[ $? -ne 0 ]] ; then
    echo "SWIG failed"
    exit 1
fi

# Rename the SWIG file so that 'go build' will not run it again
cd swig/cryptosdk_wrapper
echo "===> Building Go code <==="
# Use go build -work to not delete the work directory when exiting - good for debugging!
#go build
echo "===> Build complete, running tests <==="

go test
echo "===> Tests complete <==="
cd -



# TODO Run tests
echo
echo "Done"
echo
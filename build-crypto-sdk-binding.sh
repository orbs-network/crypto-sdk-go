#!/usr/bin/env bash

CRYPTO_SDK_LIB_DIR="crypto-sdk/lib"
CRYPTO_SDK_SWIGFILE="cryptosdk.swigcxx"
CRYPTO_SDK_BUILD="${HOME}/dev/sandbox/go/src/github.com/orbs-network/crypto-sdk-go/crypto-sdk/build"
#LIBGCRYPT_LIB="${CRYPTO_SDK_BUILD}/mac/libgcrypt/lib/libgcrypt.a"
LIBGCRYPT_LIB_DIR="${CRYPTO_SDK_BUILD}/mac/libgcrypt/lib"
#LIBGPG_LIB="${CRYPTO_SDK_BUILD}/mac/libgpg-error/lib/libgpg-error.a"
LIBGPG_LIB_DIR="${CRYPTO_SDK_BUILD}/mac/libgpg-error/lib"
#LIBGCYPT_H="${HOME}/dev/sandbox/go/src/github.com/orbs-network/crypto-sdk-go/crypto-sdk/deps/libgcrypt/libgcrypt-1.8.2/src"

LIBGCYPT_H="${CRYPTO_SDK_BUILD}/mac/libgcrypt/include/gcrypt.h"
LIBGPG_H="${CRYPTO_SDK_BUILD}/mac/libgpg-error/include/gpg-error.h"

echo "SWIG: Building Go wrapper for C++ code"
swig -c++ -Wall -go -cgo -intgosize 64 ${CRYPTO_SDK_LIB_DIR}/${CRYPTO_SDK_SWIGFILE}
if [[ $? -ne 0 ]] ; then
    echo "SWIG failed"
    exit 1
fi
# Rename the SWIG file so that 'go build' will not run it again
# TODO: find a 'go build' flag that disables automatic running of SWIG, then remove the mv commands
#mv ${CRYPTO_SDK_LIB_DIR}/${CRYPTO_SDK_SWIGFILE} ${CRYPTO_SDK_LIB_DIR}/${CRYPTO_SDK_SWIGFILE}.tmp
cd crypto-sdk-swig-wrapper
echo "Building Go code"
#export CGO_ENABLED=0
export CGO_CFLAGS="-I${LIBGCYPT_H} -I${LIBGPG_H}"
export CGO_LDFLAGS="-L${LIBGCRYPT_LIB_DIR} -L${LIBGPG_LIB_DIR}"
# Use go build -work to not delete the work directory when exiting - good for debugging!
go build -x

#mv ${CRYPTO_SDK_LIB_DIR}/${CRYPTO_SDK_SWIGFILE}.tmp ${CRYPTO_SDK_LIB_DIR}/${CRYPTO_SDK_SWIGFILE}
#echo "Running tests"
#go test
cd -



# TODO Run tests

echo "Done"
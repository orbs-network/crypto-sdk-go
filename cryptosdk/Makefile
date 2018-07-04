LDPATH=../native/mac
INCLUDE_PATH=../native/headers
RPATH=../native/mac
CGO_ENABLED=1
CGO_CXXFLAGS="-std=c++11"
CGO_CPPFLAGS="-I../native/headers"
CGO_LDFLAGS="-L../native/mac -lcryptosdk -Wl,-rpath,../native/mac"
#CGO_LDFLAGS="-L../native/mac -rpath -Wl"
TARGET=crypto-sdk-go

.PHONY: clean

test:
	CGO_CPPFLAGS=$(CGO_CPPFLAGS) CGO_LDFLAGS=$(CGO_LDFLAGS) go test -x


$(TARGET):
	CGO_CPPFLAGS=$(CGO_CPPFLAGS) CGO_LDFLAGS=$(CGO_LDFLAGS) go build -x

%.o: %.cpp
	clang++ -O2 -I${INCLUDE_PATH} -L${LDPATH} -lcryptosdk -rpath ${RPATH} -o $@ -c $^

clean:
	rm -f *.o *.so *.a $(TARGET)
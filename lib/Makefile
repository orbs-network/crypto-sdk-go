
.PHONY: clean

TARGET=crypto-sdk-go

$(TARGET): ../native/mac/libcryptosdk.dylib
	go build -x -o main

%.o: %.cpp
	clang++ -O2 -o $@ -c $^

example: example.cpp
	clang++ -O2 -I../native/headers -L../native/mac -lcryptosdk -rpath ../native/mac -o example example.cpp

clean:
	rm -f *.o *.so *.a $(TARGET)
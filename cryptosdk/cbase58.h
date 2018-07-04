#pragma once

#ifdef __cplusplus
extern "C" {
#endif

//#include <string>
//#include <vector>

//std::string _Base58Encode( std::vector<uint8_t> data);
//std::vector<uint8_t> _Base58Decode( std::string encoded);
const char* Base58Encode(const char* plaintext, int len);
const char* Base58Decode(const char* encoded);
#ifdef __cplusplus
}
#endif

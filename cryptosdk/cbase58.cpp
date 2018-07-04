#include "cbase58.h"
#include "base58.h"

using namespace std;

const char* Base58Encode(const char* plaintext, int len) {
    vector<uint8_t> data = vector<uint8_t>(plaintext, plaintext + len);
    string encoded = Orbs::Base58::Encode(data);
    return encoded.c_str();
}

const char* Base58Decode(const char* encoded) {
    vector<uint8_t> decoded = Orbs::Base58::Decode(encoded);
    std::string decoded_str(decoded.begin(), decoded.end());
    return decoded_str.c_str();
}

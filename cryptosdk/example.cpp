#include <stdio.h>
#include <iostream>
#include "base58.h"

using namespace std;


// Demonstrate calling crypto-sdk directly (from C++)

//int main(){
//    string rawData = "Hello World!";
//    vector<uint8_t> data = vector<uint8_t>(rawData.begin(), rawData.end());
//    string encoded = Orbs::Base58::Encode(data);
//
//    vector<uint8_t> decoded = Orbs::Base58::Decode(encoded);
//    std::string decoded_str(decoded.begin(), decoded.end());
//    cout<<encoded<<std::endl;
//    cout<<decoded_str<<std::endl;
//
//    return 0;
//}
#include <stdio.h>
#include <iostream>
#include "base58.h"

using namespace std;


int main(){
    string rawData = "Hello World!";
    vector<uint8_t> data = vector<uint8_t>(rawData.begin(), rawData.end());
    string encoded = Orbs::Base58::Encode(data);

    vector<uint8_t> decoded = Orbs::Base58::Decode(encoded);

    cout<<encoded<<std::endl;

    return 0;
}
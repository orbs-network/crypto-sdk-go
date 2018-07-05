#include "ced25519key.h"
#include "ed25519key.h"

CED25519Key CED25519KeyInit() {
    Orbs::ED25519Key * ret = new Orbs::ED25519Key();
    return (void*)ret;
}

void CED25519KeyFree(CED25519Key instance) {
	Orbs::ED25519Key * ed25519key = (Orbs::ED25519Key*)instance;
	delete ed25519key;
}

bool CED25519KeyHasPrivateKey(CED25519Key instance) {
	Orbs::ED25519Key * ed25519key = (Orbs::ED25519Key*)instance;
	return ed25519key->HasPrivateKey();
}

const char* CED25519KeyGetPublicKey(CED25519Key instance, char* buf, int* len)   {
   Orbs::ED25519Key * ed25519key = (Orbs::ED25519Key*)instance;
   std::vector<uint8_t> v = ed25519key->GetPublicKey();
   if(buf == 0) {
    *len = v.size();
    return 0;
   }
//   char *res = reinterpret_cast<char *>(malloc(sizeof(uint8_t) * *len));
   memcpy(buf, &v[0], sizeof(uint8_t) * *len);

   return buf;
}

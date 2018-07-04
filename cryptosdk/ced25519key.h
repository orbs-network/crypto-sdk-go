#pragma once
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif
	typedef void* CED25519Key;
	CED25519Key CED25519KeyInit(void);
	void CED25519KeyFree(CED25519Key);
	bool CED25519KeyHasPrivateKey(CED25519Key);
	const char* CED25519KeyGetPublicKey(CED25519Key, int *len);

#ifdef __cplusplus
}
#endif

# BLS - Threshold Signature

## Introduction

Algorithm for threshold encrypting plain text using BLS12381.

**Encryption**  
* generate sk_i shares for n participants (t out of n threhsold)
* encryption key = r*PK, r ephemeral private key by the encryptor
* cipher (V) = plain ^ encryption key
* U = r*G1
* W = r*Hash_to_group(U,V)

**Decryption**  
* ∀ sk_i | share = sk_i * U
* reconstructed key = lagrange interpolation for all shares, evaluate for x=0
* decrypted = cipher ^ encryption key

Inspired by [Skale libBLS](https://github.com/skalenetwork/libBLS/tree/develop/threshold_encryption)
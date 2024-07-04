# Encryptor

A simple command-line utility for encrypting and decrypting files using AES-GCM encryption.

## Features

- Encrypts files and generates a secure key.
- Decrypts files with a provided key.
- Uses AES-GCM for encryption, ensuring data confidentiality and integrity.

## Usage

### Encryption

To encrypt a file:

```sh
./encryptor -from <source_file_path> -to <encrypted_file_path>
```

### Decryption

To decrypt a file:

```sh
./encryptor -from <encrypted_file_path> -to <decrypted_file_path> -mode dec -key <encryption_key>
```

### Command-line Flags
- from: Path to the source file to be encrypted or decrypted.
- to: File name where the encrypted or decrypted file will be saved.
- mode: Operation mode (enc for encryption, dec for decryption). Default is enc.
- key: Key used for encryption or decryption (required for decryption).
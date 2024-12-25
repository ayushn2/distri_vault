# distri_vault

## 📖 Overview
distri_vault is a decentralized file storage system designed to provide secure and efficient data management. It enables peer-to-peer communication, encrypted file storage, and seamless recovery through organized server folders.

By combining peer-to-peer networking and encryption, distri_vault ensures data integrity, confidentiality, and availability, making it ideal for distributed systems.

---

## 🚀 Features

### Core Functionalities
- **Distributed File Storage**: Files are stored in a decentralized manner, ensuring redundancy and availability.
- **Encryption**: Files are encrypted for secure storage on distributed servers, with the original version kept on the primary network.
- **Peer-to-Peer Messaging**: Enables real-time communication between nodes.
- **Data Recovery**:
  - Files are organized under folders based on server IDs.
  - Simplifies folder syncing and file recovery during data loss.

### Additional Features
- **Buffering and Broadcasting**: Efficient data transfer across peers.
- **Optimized Codebase**: Modular and reusable functions reduce duplication and enhance maintainability.

---

## 🛠️ Technology Stack

| Component             | Description                                    |
|-----------------------|------------------------------------------------|
| **Programming Language** | Go (Golang)                                  |
| **Encryption**         | Custom encryption algorithms                   |
| **Networking**         | Peer-to-peer communication protocols           |
| **Storage**            | File-based organization with folder sync       |

---

## 🗂️ Project Structure
    ```plaintext
    distri_vault/
    │
    ├── main.go              # Application entry point
    ├── crypto.go            # Handles encryption and decryption
    ├── peer.go              # Manages peer communication
    ├── server.go            # Manages server-side operations
    ├── store.go             # Handles data storage and retrieval logic
    ├── p2p/                 # Handles peer-to-peer communication protocols
    │   ├── encoding.go      # Data encoding for message transmission
    │   ├── handshake.go     # Manages peer handshake process
    │   ├── message.go       # Defines the structure and logic for messages
    │   ├── tcp_transport.go # Implements transport over TCP
    │   ├── transport.go     # General transport layer abstraction
    └── README.md            # Project documentation

## 💻 Getting Started

### Prerequisites
- **Go** (Golang) installed on your system.
- Basic understanding of distributed systems and networking.

### Setup
1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/distri_vault.git
   cd distri_vault
2. **Run the application**:
   ```bash
   go run main.go
3. **Test the peer-to-peer functionality**:
   - Ensure ports (e.g., 3000 and 4000) are available.
   - Use different terminals to simulate multiple peers.

---

## 🔧 Usage

### Adding Files
1. When a server adds a file, it is automatically stored under a folder named after its server ID.
2. Files are synced to distributed servers for redundancy.

### Retrieving Files
1. Request specific files or folders by server ID.
2. Sync folders to recover data after failure.

### Peer-to-Peer Messaging
1. Establish communication between peers using their respective ports.
2. Use the Send method in peer.go to broadcast messages.

---

## 🌟 Future Enhancements

- **Scalability**: Support for larger networks and higher data volumes.
- **Advanced Encryption**: Implementation of zero-knowledge proofs for enhanced security.
- **Dynamic Node Management**: Adding or removing peers dynamically.
- **File Versioning**: Maintain multiple versions of the same file for better recovery options.

---

## 🤝 Contributors

- [ayushn2](https://github.com/ayushn2) - Creator and Maintainer

--- 

## 🙋‍♂️ Support

For any questions or issues, feel free to [open an issue](https://github.com/ayushn2/distri_vault/issues) or [email me](mailto:iayush.n2@gmail.com).

---

## 📝 Acknowledgements

This project uses the following libraries and tools:

- **[crypto/aes](https://pkg.go.dev/crypto/aes)** - Provides AES encryption and decryption functionalities.
- **[crypto/cipher](https://pkg.go.dev/crypto/cipher)** - Implements block cipher modes and stream ciphers for encryption.
- **[crypto/md5](https://pkg.go.dev/crypto/md5)** - Implements the MD5 hash function.
- **[crypto/rand](https://pkg.go.dev/crypto/rand)** - Provides cryptographically secure random number generation.
- **[encoding/hex](https://pkg.go.dev/encoding/hex)** - Used for encoding and decoding data in hexadecimal format.
- **[encoding/gob](https://pkg.go.dev/encoding/gob)** - Implements the GOB encoding/decoding format for Go objects.
- **[crypto/sha1](https://pkg.go.dev/crypto/sha1)** - Provides SHA-1 hashing functionality.
- **[errors](https://pkg.go.dev/errors)** - Implements error handling and custom errors in Go.
- **[fmt](https://pkg.go.dev/fmt)** - Implements formatted I/O with functions like Printf, Sprintf, etc.
- **[io](https://pkg.go.dev/io)** - Provides basic interfaces for I/O operations such as reading and writing.
- **[log](https://pkg.go.dev/log)** - Implements a simple logging package.
- **[net](https://pkg.go.dev/net)** - Provides networking and internet protocols.
- **[os](https://pkg.go.dev/os)** - Provides functions for OS-level operations, such as file manipulation and environment variables.
- **[strings](https://pkg.go.dev/strings)** - Implements functions for string manipulation.

These libraries have been essential in developing the encryption, peer-to-peer communication, and overall functionality of the system. Thank you to the maintainers of these packages!

---

## 🤝 Contribution

We welcome contributions to improve the project! Here’s how you can contribute:

### 1. **Fork the repository**
   - Click the "Fork" button on the top right corner of the repository page to create a copy of this repository on your GitHub account.

### 2. **Clone the repository**
   - Clone your forked repository to your local machine:
     ```bash
     git clone https://github.com/your-username/distri_vault.git
     ```

### 3. **Create a new branch**
   - It’s important to create a new branch for each contribution:
     ```bash
     git checkout -b your-feature-name
     ```

### 4. **Make changes**
   - Make changes or add features to the project. Ensure your code adheres to the existing style and structure of the project.

### 5. **Commit your changes**
   - Commit your changes with a descriptive message:
     ```bash
     git commit -m "Describe your changes here"
     ```

### 6. **Push your changes**
   - Push your changes to your fork:
     ```bash
     git push origin your-feature-name
     ```

### 7. **Open a Pull Request**
   - Go to your fork on GitHub, and click on "Pull Requests".
   - Click the "New Pull Request" button to submit your changes to the original repository.

### 8. **Discuss and Review**
   - Once your pull request is submitted, it will be reviewed, and any necessary discussions will happen. You may be asked to make changes before the pull request is merged.

### Code of Conduct
Please follow our [Code of Conduct](link-to-code-of-conduct) in all interactions, whether on GitHub or elsewhere.

---

We appreciate your contributions and look forward to improving the project together!

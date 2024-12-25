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
├── storage/
│   ├── server_1/        # Folder for files added by server 1
│   ├── server_2/        # Folder for files added by server 2
│   └── ...
└── README.md            # Project documentation
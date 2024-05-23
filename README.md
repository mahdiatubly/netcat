# Netcat Chat Room

## Overview

Welcome to the Netcat Chat Room project! This project is a practical implementation designed to practice concurrent programming using Go. It mimics some of the essential functionalities of Netcat, providing a chat room where users can connect and communicate in real-time.

## Features

- **Real-time Chat**: Users can connect and chat with others who are connected to the server simultaneously.
- **Chat History**: New users receive the chat history upon joining, ensuring they are up-to-date with the ongoing conversation.
- **Logging**: All chat messages are logged to a file on the server, preserving the conversation history for future reference.

## Getting Started

### Prerequisites

To run this project, you need to have the Go programming language installed on your machine. You can download and install Go from the official [Go website](https://golang.org/dl/).

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/netcat-chat-room.git
    cd netcat-chat-room
    ```

2. Build the project:
    ```sh
    go build -o TCPChat
    ```

3. Run the server:
    ```sh
    ./TCPChat
    ```
    By default, the server listens on port `8989`. You can specify a different port if needed:
    ```sh
    ./TCPChat 2525
    ```

### Connecting to the Chat Room

Users can connect to the chat room using Netcat or similar tools. Open a terminal and run:
```sh
nc localhost 8989
Replace `8989` with the port number specified when running the server.
```

## Usage

Once connected, users will be prompted to enter their name. After entering a name, they can start sending and receiving messages in real-time. The server broadcasts messages from each user to all connected clients.

## Example

### Server Output:

```csharp
Listening on port :8989
```

### Client 1:

```sh
$ nc localhost 8989
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Alice
[2024-05-23 16:03:43][Alice]: Hello, everyone!
```

### Client 2:

```sh
$ nc localhost 8989
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Bob
[2024-05-23 16:04:10][Alice]: Hello, everyone!
[2024-05-23 16:04:32][Bob]: Hi, Alice!
```

## Logging

All messages are logged to a file named `chat.log` on the server. This file contains a complete history of the chat, including timestamps and user names.

## Contributions

Contributions to this project are welcome. Feel free to open issues or submit pull requests with improvements and bug fixes.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

Thank you for using the Netcat Chat Room project! We hope it serves as a useful example of concurrent programming in Go.
```

Save this content into a file with a `.md` extension, and you'll have it in Markdown format.

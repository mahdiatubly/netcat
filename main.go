package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const defaultPort = "8989"
const maxConnections = 10

type Message struct {
	Timestamp  time.Time
	ClientName string
	Text       string
}

type Client struct {
	Conn net.Conn
	Name string
}

var (
	clients  = make(map[net.Conn]Client)
	messages []Message
	mu       sync.Mutex
)

func printUsage() {
	fmt.Println("[USAGE]: ./TCPChat $port")
	os.Exit(1)
}

func broadcastMessage(message, clientName string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formattedMessage := fmt.Sprintf("[%s][%s]: %s\n", timestamp, clientName, message)
	if clientName == "" {
		formattedMessage = fmt.Sprintf("[%s]: %s\n", timestamp, message)
	}
	mu.Lock()
	defer mu.Unlock()
	for _, client := range clients {
		client.Conn.Write([]byte(formattedMessage))
	}
}

func sendChatHistory(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()
	for _, msg := range messages {
		formattedMessage := fmt.Sprintf("[%s][%s]: %s\n", msg.Timestamp.Format("2006-01-02 15:04:05"), msg.ClientName, msg.Text)
		conn.Write([]byte(formattedMessage))
	}
}

func getName(conn net.Conn) string {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		name := scanner.Text()
		if name != "" {
			return name
		}
		conn.Write([]byte("[ENTER YOUR NAME]: "))
	}
	return ""
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Welcome to TCP-Chat!\n" +
		"         _nnnn_\n" +
		"        dGGGGMMb\n" +
		"       @p~qp~~qMb\n" +
		"       M|@||@) M|\n" +
		"       @,----.JM|\n" +
		"      JS^\\__/  qKL\n" +
		"     dZP        qKRb\n" +
		"    dZP          qKKb\n" +
		"   fZP            SMMb\n" +
		"   HZM            MMMM\n" +
		"   FqM            MMMM\n" +
		" __| \".        |\\dS\"qML\n" +
		" |    `.       | `' \\Zq\n" +
		"_)      \\.___.,|     .'\n" +
		"\\____   )MMMMMP|   .'\n" +
		"     `-'       `--'\n" +
		"[ENTER YOUR NAME]: "))

	clientName := getName(conn)
	if clientName == "" {
		return
	}

	client := Client{
		Conn: conn,
		Name: clientName,
	}

	mu.Lock()
	clients[conn] = client
	mu.Unlock()

	sendChatHistory(conn)
	info := fmt.Sprintf("%s has joined the chat", clientName)
	mu.Lock()
	messages = append(messages, Message{Timestamp: time.Now(), Text: info})
	mu.Unlock() 
	broadcastMessage(info, "")
	
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		if message == "" {
			continue
		}

		timestamp := time.Now()
		mu.Lock()
		messages = append(messages, Message{Timestamp: timestamp, ClientName: clientName, Text: message})
		mu.Unlock()

		broadcastMessage(message, clientName)
	}

	mu.Lock()
	delete(clients, conn)
	mu.Unlock()

	info1 := fmt.Sprintf("%s has left the chat", clientName)
	mu.Lock()
	messages = append(messages, Message{Timestamp: time.Now(), Text: info1})
	mu.Unlock() 
	broadcastMessage(info1, "")
}

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		printUsage()
	}

	port := defaultPort
	if len(args) == 1 {
		port = args[0]
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	fmt.Printf("Listening on port :%s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		} else if len(clients) >= maxConnections {
			conn.Write([]byte("Sorry! the discussion room is currently full, try again latter."))
			continue
		}
		go handleConnection(conn)
	}
}


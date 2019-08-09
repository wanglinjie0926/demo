package kunpengBattle

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type KunPengBattleClient struct {
	TeamID       int
	TeamName     string
	ServerIP     string
	ServerPort   int
	conn         *battleConnection
	reader       *bufio.Reader
	writer       *bufio.Writer
	Strategy     Strategy
	legStartChan chan KunPengLegStart
	legEndChan   chan KunPengLegEnd
	roundChan    chan KunPengRound
	errChan      chan error
	terminalChan chan int
}

func NewKunPengBattleClient(teamID int, teamName string, strategy Strategy) KunPengBattleClient {
	client := KunPengBattleClient{TeamID: teamID, TeamName: teamName}
	client.conn = new(battleConnection)
	client.reader = bufio.NewReaderSize(client.conn, 1024*10)
	client.writer = bufio.NewWriterSize(client.conn, 1024*10)

	client.legStartChan = make(chan KunPengLegStart)
	client.legEndChan = make(chan KunPengLegEnd)
	client.roundChan = make(chan KunPengRound)
	client.errChan = make(chan error, 10)
	client.terminalChan = make(chan int)
	client.Strategy = strategy

	return client
}

func (c *KunPengBattleClient) Connect(ip string, port int) error {
	c.ServerIP = ip
	c.ServerPort = port
	address := ip + ":" + strconv.Itoa(port)
	err := c.conn.connect(address)

	return err
}

func (c *KunPengBattleClient) Start() error {

	if err := c.registrate(); err != nil {
		return err
	}

	go c.receive()
	// go c.send()
strategyLoop:
	for {
		select {
		case start := <-c.legStartChan:
			if err := c.Strategy.LegStart(start); err != nil {
				c.errChan <- err
			}
		case end := <-c.legEndChan:
			if err := c.Strategy.LegEnd(end); err != nil {
				c.errChan <- err
			}
		case round := <-c.roundChan:
			action, err := c.Strategy.React(round)
			if err != nil {
				c.errChan <- err
				break
			}
			if err := c.action(action); err != nil {
				c.errChan <- err
			}

		// case <-time.After(time.Millisecond * 750):
		// 	log.Printf("Time OUT!")
		case err := <-c.errChan:
			fmt.Printf("ERROR: %v", err)
		case <-c.terminalChan:
			fmt.Println("GAMEOVER!!!")
			break strategyLoop
		}
	}

	return nil
}

func (c *KunPengBattleClient) receive() {
	scanner := bufio.NewScanner(c.reader)
	scanner.Split(splitPackage)
	for scanner.Scan() {
		msgBytes := scanner.Bytes()
		var msgData json.RawMessage
		msg := &KunPengMsg{Data: &msgData}
		err := json.Unmarshal(msgBytes, msg)
		if err != nil {
			fmt.Printf("Unmarshal err: %v", err)
			continue
		}

		fmt.Println("RECV->", string(msgData))

		switch msg.Name {
		case "leg_start":
			legStart := new(KunPengLegStart)
			err := json.Unmarshal(msgData, legStart)
			if err != nil {
				c.errChan <- err
				continue
			}
			c.legStartChan <- *legStart
		case "round":
			round := new(KunPengRound)
			err := json.Unmarshal(msgData, round)
			if err != nil {
				c.errChan <- err
				continue
			}
			c.roundChan <- *round
		case "leg_end":
			legEnd := new(KunPengLegEnd)
			err := json.Unmarshal(msgData, legEnd)
			if err != nil {
				c.errChan <- err
				continue
			}
			c.legEndChan <- *legEnd
		case "game_over":
			c.terminalChan <- 1
		default:
			c.errChan <- fmt.Errorf("Unknown msg Name: %v", msg.Name)
		}
	}
}

func (c *KunPengBattleClient) registrate() error {
	kpMsg := new(KunPengMsg)
	kpr := KunPengRegistration{TeamID: c.TeamID, TeamName: c.TeamName}
	kpMsg.Name = "registration"
	kpMsg.Data = kpr

	if err := c.Strategy.Registrate(kpr); err != nil {
		return nil
	}

	msgBytes, err := json.Marshal(kpMsg)

	if err != nil {
		return err
	}

	if err := c.send(msgBytes); err != nil {
		return err
	}

	fmt.Println("Registration Sent!!!")

	return nil
}

func (c *KunPengBattleClient) action(act KunPengAction) error {
	kpMsg := new(KunPengMsg)
	kpMsg.Name = "action"
	kpMsg.Data = act

	msgBytes, err := json.Marshal(kpMsg)

	if err != nil {
		return err
	}

	if err := c.send(msgBytes); err != nil {
		return err
	}

	return nil
}

func (c *KunPengBattleClient) send(msgBytes []byte) error {
	msgLen := len(msgBytes)

	if msgLen > 99999 {
		return errors.New("msgLen shoule not greater than 99999")
	}

	msgLenStr := strconv.FormatInt(int64(msgLen), 10)
	sizeBytes := make([]byte, 0, msgLen+5)
	for i := 0; i < 5-len(msgLenStr); i++ {
		sizeBytes = append(sizeBytes, '0')
	}
	sizeBytes = append(sizeBytes, []byte(msgLenStr)...)

	sendBytes := append(sizeBytes, msgBytes...)

	fmt.Println("SEND->", string(msgBytes))

	if _, err := c.writer.Write(sendBytes); err != nil {
		return err
	}

	if err := c.writer.Flush(); err != nil {
		return err
	}

	return nil
}

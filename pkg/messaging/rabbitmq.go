package messaging

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Connection *amqp091.Connection
	Channel    *amqp091.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Printf("Erro ao conectar ao RabbitMQ: %s", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{Connection: conn, Channel: ch}, nil
}

func (rabbit *RabbitMQ) Publish(queue string, body []byte) error {
	_, err := rabbit.Channel.QueueDeclare(
		queue, // nome
		false, // durável
		false, // excluir quando não usado
		false, // exclusivo
		false, // sem espera
		nil,   // argumentos
	)
	if err != nil {
		log.Printf("Erro ao declarar a fila: %s", err)
		return err
	}

	err = rabbit.Channel.Publish(
		"",    // exchange
		queue, // chave de roteamento
		false, // obrigatório
		false, // imediato
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Printf("Erro ao publicar a mensagem: %s", err)
		return err
	}

	log.Printf("Mensagem publicada na fila %s: %s", queue, body)
	return nil
}

func (rabbit *RabbitMQ) Close() {
	if err := rabbit.Channel.Close(); err != nil {
		log.Printf("Erro ao fechar o canal: %s", err)
	}
	if err := rabbit.Connection.Close(); err != nil {
		log.Printf("Erro ao fechar a conexão: %s", err)
	}
}

func (rabbit *RabbitMQ) ConsumeMessages(queue string) {
	msgs, err := rabbit.Channel.Consume(
		queue,
		"",    // consumidor
		true,  // auto-ack
		false, // exclusivo
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Falha ao registrar um consumidor: %s", err)
	}

	log.Printf(" [*] Aguardando mensagens na fila %s. Para sair, pressione CTRL+C", queue)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Mensagem recebida: %s", d.Body)
		}
	}()

	<-forever
}

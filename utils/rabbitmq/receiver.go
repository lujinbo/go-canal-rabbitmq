package rabbitmq

import (
	"fmt"
	"sync"

	"github.com/siddontang/go-log/log"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn       *amqp.Connection
	Channel    *amqp.Channel
	QueueName  string //队列名称
	Exchange   string //交换机
	RoutingKey string
	dns        string
	sync.Mutex
}

type QueueExchange struct {
	QueueName    string // 队列名称
	RoutingKey   string // key值
	ExchangeName string // 交换机名称
	ExchangeType string // 交换机类型
	Dns          string //链接地址
}

func NewRabbitMQ(q QueueExchange) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName:  q.QueueName,
		Exchange:   q.ExchangeName,
		RoutingKey: q.RoutingKey,
		dns:        q.Dns,
	}

	var err error
	rabbitmq.Conn, err = amqp.Dial(rabbitmq.dns)
	if err != nil {
		rabbitmq.failOnErr(err, "创建连接错误")
	}

	rabbitmq.Channel, err = rabbitmq.Conn.Channel()
	if err != nil {
		rabbitmq.failOnErr(err, "获取Channel失败")
	}
	return rabbitmq
}

func (r *RabbitMQ) MqConnect() (err error) {
	mqConn, err := amqp.Dial(r.dns)
	r.Conn = mqConn // 赋值给RabbitMQ对象
	if err != nil {
		fmt.Printf("mq链接失败  :%s \n", err)
		return err
	}
	return
}

// 关闭mq链接
func (r *RabbitMQ) CloseMqConnect() (err error) {
	err = r.Conn.Close()
	if err != nil {
		fmt.Printf("关闭mq链接失败  :%s \n", err)
		return err
	}
	return err
}

// 链接rabbitMQ
func (r *RabbitMQ) MqOpenChannel() (err error) {
	mqConn := r.Conn
	r.Channel, err = mqConn.Channel()
	//defer mqChan.Close()
	if err != nil {
		fmt.Printf("MQ打开管道失败:%s \n", err)
	}
	return err
}

// 链接rabbitMQ
func (r *RabbitMQ) CloseMqChannel() (err error) {
	r.Channel.Close()
	if err != nil {
		fmt.Printf("关闭mq链接失败  :%s \n", err)
	}
	return err
}

func Send(q QueueExchange, msg string) (err error) {
	mq := NewRabbitMQ(q)
	err = mq.MqConnect()
	if err != nil {
		return
	}

	defer func() {
		mq.CloseMqConnect()
	}()

	err = mq.sendMsg(msg)

	return
}

func (mq *RabbitMQ) sendMsg(body string) (err error) {

	err = mq.MqOpenChannel()
	ch := mq.Channel
	if err != nil {
		log.Printf("Channel err  :%s \n", err)
	}

	defer mq.Channel.Close()

	// 1.声明队列
	/*
		如果只有一方声明队列，可能会导致下面的情况：
			a)消费者是无法订阅或者获取不存在的MessageQueue中信息
			b)消息被Exchange接受以后，如果没有匹配的Queue，则会被丢弃

		为了避免上面的问题，所以最好选择两方一起声明
		ps:如果客户端尝试建立一个已经存在的消息队列，Rabbit MQ不会做任何事情，并返回客户端建立成功的
	*/
	_, err = ch.QueueDeclare( // 返回的队列对象内部记录了队列的一些信息，这里没什么用
		mq.QueueName, // 队列名
		true,         // 是否持久化
		false,        // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
		false,        // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
		false,        // 是否阻塞
		nil,          //额外属性（我还不会用）
	)

	if err != nil {
		fmt.Println("声明队列失败", err)
		return err
	}

	// 2.声明交换器
	err = ch.ExchangeDeclare(
		mq.Exchange, //交换器名
		"topic",     //exchange type：一般用fanout、direct、topic
		true,        // 是否持久化
		false,       //是否自动删除（自动删除的前提是至少有一个队列或者交换器与这和交换器绑定，之后所有与这个交换器绑定的队列或者交换器都与此解绑）
		false,       //设置是否内置的。true表示是内置的交换器，客户端程序无法直接发送消息到这个交换器中，只能通过交换器路由到交换器这种方式
		false,       // 是否阻塞
		nil,         // 额外属性
	)

	if err != nil {
		fmt.Println("创建交换机异常", err)
		return err
	}

	// 3.建立Binding(可随心所欲建立多个绑定关系)
	err = ch.QueueBind(
		mq.QueueName,  // 绑定的队列名称
		mq.RoutingKey, // bindkey 用于消息路由分发的key
		mq.Exchange,   // 绑定的exchange名
		false,         // 是否阻塞
		nil,           // 额外属性
	)
	if err != nil {
		fmt.Println("绑定队列和交换器失败", err)
	}

	ch.Publish(
		mq.Exchange,
		mq.RoutingKey,
		//如果为true，根据exchange类型和routekey类型，如果无法找到符合条件的队列，name会把发送的信息返回给发送者
		false,
		//如果为true，当exchange发送到消息队列后发现队列上没有绑定的消费者,则会将消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	log.Infoln("发布成功", body)
	return nil
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, msg string) {
	if err != nil {
		log.Fatal("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

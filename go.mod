module github.com/spiral/jobs/v2

go 1.14

replace github.com/spiral/goridge/v2 v2.4.4 => github.com/korol8484/goridge/v2 v2.4.5-0.20200629170527-5aead22a2027

require (
	github.com/aws/aws-sdk-go v1.16.14
	github.com/beanstalkd/go-beanstalk v0.0.0-20180822062812-53ecdaa3bcfb
	github.com/buger/goterm v0.0.0-20181115115552-c206103e1f37
	github.com/cenkalti/backoff/v4 v4.0.0
	github.com/dustin/go-humanize v1.0.0
	github.com/gofrs/uuid v3.1.0+incompatible
	github.com/json-iterator/go v1.1.9
	github.com/kr/beanstalk v0.0.0-20180818045031-cae1762e4858 // indirect
	github.com/olekukonko/tablewriter v0.0.4
	github.com/prometheus/client_golang v1.5.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	github.com/spiral/goridge/v2 v2.4.4
	github.com/spiral/roadrunner v1.8.0
	github.com/streadway/amqp v0.0.0-20181205114330-a314942b2fd9
	github.com/stretchr/testify v1.5.1
)

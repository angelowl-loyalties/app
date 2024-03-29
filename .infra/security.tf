resource "aws_security_group" "angelowl_ssh" {
  name        = "angelowl-ssh-access"
  description = "Allows inbound SSH traffic"
  vpc_id      = aws_vpc.angelowl.id

  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["10.8.0.0/14"]
  }
}

resource "aws_security_group" "angelowl_postgres" {
  name        = "angelowl-postgres-access"
  description = "Allows inbound Postgres traffic"
  vpc_id      = aws_vpc.angelowl.id

  ingress {
    description = "PostgreSQL"
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["10.8.0.0/14"]
  }
}

resource "aws_security_group" "angelowl_kafka" {
  name        = "angelowl-kafka-access"
  description = "Allows inbound Kafka traffic"
  vpc_id      = aws_vpc.angelowl.id

  ingress {
    description = "Kafka"
    from_port   = 9092
    to_port     = 9096
    protocol    = "tcp"
    cidr_blocks = ["10.8.0.0/14"]
  }

  ingress {
    description = "Zookeeper"
    from_port   = 2181
    to_port     = 2182
    protocol    = "tcp"
    cidr_blocks = ["10.8.0.0/14"]
  }
}

resource "aws_security_group" "angelowl_openvpn" {
  name        = "angelowl-openvpn-access"
  description = "Allows inbound OpenVPN traffic"
  vpc_id      = aws_vpc.angelowl.id

  ingress {
    description = "OpenVPN"
    from_port   = 1194
    to_port     = 1194
    protocol    = "udp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "angelowl_kubeservices" {
  name        = "angelowl-kubeservices-access"
  description = "Allows inbound Kubernetes Services traffic"
  vpc_id      = aws_vpc.angelowl.id

  ingress {
    description = "Kubernetes Services"
    from_port   = 30000
    to_port     = 32767
    protocol    = "tcp"
    cidr_blocks = ["10.8.0.0/14"]
  }

  egress {
    description = "Kubernetes Services"
    from_port   = 30000
    to_port     = 32767
    protocol    = "tcp"
    cidr_blocks = ["10.8.0.0/14"]
  }
}

resource "aws_security_group" "angelowl_http_s_ingress" {
  name        = "angelowl-http-s-ingress"
  description = "Allows inbound HTTP and HTTPS traffic"
  vpc_id      = aws_vpc.angelowl.id

  ingress {
    description = "HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["10.8.0.0/14"]
  }

  ingress {
    description = "HTTPS"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["10.8.0.0/14"]
  }
}

resource "aws_security_group" "angelowl_outbound" {
  name        = "angelowl-outbound-access"
  description = "Allows outbound traffic"
  vpc_id      = aws_vpc.angelowl.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "angelowl_wideopen" {
  name        = "angelowl-wideopen-access"
  description = "Allows INSECURE wide-open traffic"
  vpc_id      = aws_vpc.angelowl.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_key_pair" "angelowl_vpn" {
  key_name   = "angelowl-vpn-key"
  public_key = file("ssh/vpn.pub")
}

resource "aws_key_pair" "angelowl_k3s" {
  key_name   = "angelowl-k3s-key"
  public_key = file("ssh/k3s.pub")
}
